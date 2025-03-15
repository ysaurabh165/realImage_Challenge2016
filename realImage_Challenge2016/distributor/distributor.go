package distributor

import (
	"encoding/json"
	"os"
)

// LoadDistributors reads distributor data from a JSON file
func LoadDistributors(filePath string) ([]Distributor, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var distributors []Distributor
	err = json.Unmarshal(file, &distributors)
	if err != nil {
		return nil, err
	}

	return distributors, nil
}
