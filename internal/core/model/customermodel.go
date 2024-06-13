package model

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

// ---------------------------------------
// Entity: CUSTOMER
// ---------------------------------------

type Customer struct {
	ID          string
	Name        string
	Address     string
	Postcode    int
	Community   string
	CountryCode string
}

var (
	errorNoName           = errors.New("Name field is empty")
	errorNoAddress        = errors.New("Address field is empty")
	errorWrongPostcode    = errors.New("Incorrect postcode")
	errorNoCommunity      = errors.New("Community field is empty")
	errorWrongCountryCode = errors.New("Country code is not correct")
)

func NewCustomer(name string, address string, postcode int, community string, countryCode string) (*Customer, error) {
	// --- verification of the customer
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errorNoName
	}

	address = strings.TrimSpace(address)
	if address == "" {
		return nil, errorNoAddress
	}

	if postcode < 1000 || postcode > 9999 {
		return nil, errorWrongPostcode
	}

	community = strings.TrimSpace(community)
	if len(community) == 0 {
		return nil, errorNoCommunity
	}

	countryCode = strings.TrimSpace(countryCode)
	if len(countryCode) != 2 {
		return nil, errorWrongCountryCode
	}

	// --- Creation of the customer
	customer := Customer{
		ID:          uuid.New().String(),
		Name:        name,
		Address:     address,
		Postcode:    postcode,
		Community:   community,
		CountryCode: countryCode,
	}
	return &customer, nil
}
