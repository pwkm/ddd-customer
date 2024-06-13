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

func NewCustomer(name string, address string, postcode int, community string, countrycode string) (*Customer, error) {
	// --- verification of the customer
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errorNoName
	}

	address = strings.TrimSpace(address)

	// --- Creation of the customer
	customer := Customer{
		ID:          uuid.New().string(),
		Name:        name,
		Address:     address,
		Postcode:    postcode,
		Community:   community,
		CountryCode: countrycode,
	}
	return &customer, nil
}
