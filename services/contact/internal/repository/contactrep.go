package repository

import (
	"cleancode/pkg/store/postgres"
	"cleancode/services/entity/internal/domain"
	"fmt"
)

type ContactRepositoryImpl struct {
	conn *postgres.Conn
}

func (c ContactRepositoryImpl) CreateContact(contact *entity.Contact) error {
	fmt.Println("creating a entity")
	return nil
}

func (c ContactRepositoryImpl) ReadContact(id int) (*entity.Contact, error) {
	fmt.Println("reading a entity")
	return nil, nil
}

func (c ContactRepositoryImpl) UpdateContact(contact *entity.Contact) error {
	fmt.Println("updating a entity")
	return nil
}

func (c ContactRepositoryImpl) DeleteContact(id int) error {
	fmt.Println("deleting a entity")
	return nil
}

type ContactRepository interface {
	CreateContact(contact *entity.Contact) error
	ReadContact(id int) (*entity.Contact, error)
	UpdateContact(contact *entity.Contact) error
	DeleteContact(id int) error
}

func NewContactRepository(conn *postgres.Conn) *ContactRepositoryImpl {
	return &ContactRepositoryImpl{
		conn: conn,
	}
}