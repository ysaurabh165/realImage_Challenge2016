package main

import (
	"fmt"
	"log"

	"realImage_Challenge2016/distributor"
)

func main() {
	distributors, err := distributor.LoadDistributors("data/distributors.json")
	if err != nil {
		log.Fatalf("Failed to load distributor data: %v", err)
	}

	testRegions := []string{
		"CHICAGO-ILLINOIS-UNITEDSTATES",
		"CHENNAI-TAMILNADU-INDIA",
		"BANGALORE-KARNATAKA-INDIA",
		"MUMBAI-MAHARASHTRA-INDIA",
		"BERLIN-GERMANY",
		"PARIS-FRANCE",
	}

	for _, dist := range distributors {
		fmt.Printf("\n=== Distributor: %s ===\n", dist.Name)
		for _, region := range testRegions {
			if distributor.HasPermission(dist, region) {
				fmt.Printf("Distribution in %s: YES\n", region)
			} else {
				fmt.Printf("Distribution in %s: NO\n", region)
			}
		}
	}
}
