package models

import (
	"fmt"
	"strings"
)

type Distributor struct {
	Name     string
	Parent   *Distributor
	Includes map[string]bool
	Excludes map[string]bool
}

// creates a new distributor.
func NewDistributor(name string, parent *Distributor) *Distributor {
	return &Distributor{
		Name:     name,
		Parent:   parent,
		Includes: make(map[string]bool),
		Excludes: make(map[string]bool),
	}
}

// AddPermission adds an include or exclude permission.
func (d *Distributor) AddPermission(permissionType, region string) error {
	region = strings.ToUpper(region)
	if permissionType == "INCLUDE" {
		if d.Parent != nil && !d.Parent.Includes[region] {
			return fmt.Errorf("parent distributor does not have rights for %s", region)
		}
		d.Includes[region] = true
	} else if permissionType == "EXCLUDE" {
		d.Excludes[region] = true
	} else {
		return fmt.Errorf("invalid permission type: %s", permissionType)
	}
	return nil
}

func (d *Distributor) CanDistribute(city, state, country string) bool {
	city = strings.ToUpper(strings.TrimSpace(city))
	state = strings.ToUpper(strings.TrimSpace(state))
	country = strings.ToUpper(strings.TrimSpace(country))

	loc := fmt.Sprintf("%s-%s-%s", city, state, country)
	stateLoc := fmt.Sprintf("%s-%s", state, country)
	countryLoc := country

	// If explicitly excluded at any level, return false
	if d.isExcluded(loc, stateLoc, countryLoc) {
		return false
	}

	// If included at any level, return true
	if d.isIncluded(loc, stateLoc, countryLoc) {
		return true
	}

	// If the parent distributor has permissions, check those as well
	if d.Parent != nil {
		return d.Parent.CanDistribute(city, state, country)
	}

	return false
}

// Helper method to check exclusion
func (d *Distributor) isExcluded(loc, stateLoc, countryLoc string) bool {
	if d.Excludes[loc] || d.Excludes[stateLoc] || d.Excludes[countryLoc] {
		return true
	}
	if d.Parent != nil {
		return d.Parent.isExcluded(loc, stateLoc, countryLoc)
	}
	return false
}

// Helper method to check inclusion
func (d *Distributor) isIncluded(loc, stateLoc, countryLoc string) bool {
	if d.Includes[loc] || d.Includes[stateLoc] || d.Includes[countryLoc] {
		return true
	}
	if d.Parent != nil {
		return d.Parent.isIncluded(loc, stateLoc, countryLoc)
	}
	return false
}
