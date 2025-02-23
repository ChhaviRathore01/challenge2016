package services

import (
	"challenge2016/models"
	"fmt"
)

var Distributors = make(map[string]*models.Distributor)

// CreateDistributor creates a new distributor if it doesn't already exist.
func CreateDistributor(name, parentName string) (*models.Distributor, error) {
	// Check if distributor already exists
	if _, exists := Distributors[name]; exists {
		return nil, fmt.Errorf("Distributor '%s' already exists", name)
	}

	var parent *models.Distributor
	if parentName != "" {
		parent = Distributors[parentName]
	}

	d := models.NewDistributor(name, parent)
	Distributors[name] = d
	return d, nil
}
