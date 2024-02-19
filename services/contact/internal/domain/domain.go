package domain

import (
	"errors"
	"strings"
)

type Contact struct {
	ID         string
	firstName  string
	lastName   string
	middleName string
	phone      string
}

type Group struct {
	ID       string
	Name     string
	Contacts []Contact
}

func NewContact(id, firstName, lastName, middleName, phone string) (*Contact, error) {
	if len(phone) != len(strings.Trim(phone, "0123456789")) {
		return nil, errors.New("phone number should contain only digits")
	}

	return &Contact{
		ID:         id,
		firstName:  firstName,
		lastName:   lastName,
		middleName: middleName,
		Phone:      phone,
	}, nil
}

func (c *Contact) FullName() string {
	return c.firstName + " " + c.middleName + " " + c.lastName
}

func NewGroup(id, name string, contacts []Contact) (*Group, error) {
	if len(name) > 250 {
		return nil, errors.New("group name should not exceed 250 characters")
	}

	return &Group{
		ID:       id,
		Name:     name,
		Contacts: contacts,
	}, nil
}