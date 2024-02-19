package repository

import (
	"cleancode/pkg/store/postgres"
	"cleancode/services/entity/internal/domain"
)

type GroupRepositoryImpl struct {
	conn *postgres.Conn
}

func (g GroupRepositoryImpl) CreateGroup(group *entity.Group) error {
	return nil
}

func (g GroupRepositoryImpl) ReadGroup(id int) (*entity.Group, error) {
	return nil, nil
}

func (g GroupRepositoryImpl) AddContactToGroup(contactID, groupID int) error {
	return nil
}

type GroupRepository interface {
	CreateGroup(group *entity.Group) error
	ReadGroup(id int) (*entity.Group, error)
	AddContactToGroup(contactID, groupID int) error
}

func NewGroupRepository(conn *postgres.Conn) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{
		conn: conn,
	}
}
func (c ContactRepositoryImpl) CreateGroup(group *entity.Group) error {
	return nil
}

func (c ContactRepositoryImpl) ReadGroup(id int) (*entity.Group, error) {
	return nil, nil
}

func (c ContactRepositoryImpl) AddContactToGroup(contactID, groupID int) error {
	return nil
}