package model

import (
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

func NewCustomer(name string, address string, postcode int, community string, countrycode string) (*Customer, error) {
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
