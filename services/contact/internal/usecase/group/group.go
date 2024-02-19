package group

import (
	"cleancode/services/entity/internal/domain"
	"cleancode/services/entity/internal/repository"
)

type GroupUseCase interface {
	CreateGroup(group *entity.Group) error
	ReadGroup(id int) (*entity.Group, error)
	AddContactToGroup(contactID, groupID int) error
}

type GroupUseCaseImpl struct {
	repo repository.GroupRepository
}

func (g GroupUseCaseImpl) CreateGroup(group *entity.Group) error {
	return g.repo.CreateGroup(group)
}

func (g GroupUseCaseImpl) ReadGroup(id int) (*entity.Group, error) {
	return g.repo.ReadGroup(id)
}

func (g GroupUseCaseImpl) AddContactToGroup(contactID, groupID int) error {
	return g.repo.AddContactToGroup(contactID, groupID)
}

func NewGroupUseCase(repo repository.GroupRepository) *GroupUseCaseImpl {
	return &GroupUseCaseImpl{
		repo: repo,
	}
}