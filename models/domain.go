package models

import (
	"errors"
	"regexp"
)

type Domain struct {
	value string
}

const dnsRegex string = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`

func DomainParse(value string) (*Domain, error) {
	if !isValidDomain(value) {
		return nil, errors.New("error: Given argument is not a valid domain")
	}
	return &Domain{value: value}, nil
}

func isValidDomain(domain string) bool {
	match, _ := regexp.MatchString(dnsRegex, domain)
	return match
}

func (d *Domain) String() string {
	return d.value
}
