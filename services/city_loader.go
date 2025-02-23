package services

import (
	"challenge2016/models"
	"encoding/csv"
	"os"
)

// LoadCities parses the cities.csv file.
func LoadCities(filename string) (*models.GeoNode, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	root := &models.GeoNode{Name: "World", Children: make(map[string]*models.GeoNode)}
	for _, record := range records[1:] {
		city, state, country := record[3], record[4], record[5]
		countryNode := models.GetOrCreateChild(root, country)
		stateNode := models.GetOrCreateChild(countryNode, state)
		models.GetOrCreateChild(stateNode, city)
	}
	return root, nil
}
