package distributor

import "strings"

// normalizeRegion ensures consistent case for comparison
func normalizeRegion(region string) string {
	return strings.ToUpper(strings.TrimSpace(region))
}

// hasPermission checks if the given distributor can distribute in a region
func HasPermission(distributor Distributor, region string) bool {
	normalizedRegion := normalizeRegion(region)

	// Check Exclusions First
	for _, excl := range distributor.Exclude {
		if strings.HasSuffix(normalizedRegion, normalizeRegion(excl)) {
			return false
		}
	}

	// Check Inclusions
	for _, incl := range distributor.Include {
		if strings.HasSuffix(normalizedRegion, normalizeRegion(incl)) {
			return true
		}
	}

	return false
}
