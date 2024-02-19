package contact

import (
	"cleancode/services/contact/internal/domain"
	"cleancode/services/entity/internal/repository"
)

type ContactUseCase interface {
	CreateContact(contact *entity.Contact) error
	ReadContact(id int) (*entity.Contact, error)
	UpdateContact(contact *entity.Contact) error
	DeleteContact(id int) error
}
type ContactUseCaseImpl struct {
	repo repository.ContactRepository
}

func (c ContactUseCaseImpl) CreateContact(contact *entity.Contact) error {
	return c.repo.CreateContact(contact)
}

func (c ContactUseCaseImpl) ReadContact(id int) (*entity.Contact, error) {
	return c.repo.ReadContact(id)
}

func (c ContactUseCaseImpl) UpdateContact(contact *entity.Contact) error {
	return c.repo.UpdateContact(contact)
}

func (c ContactUseCaseImpl) DeleteContact(id int) error {
	return c.repo.DeleteContact(id)
}

func NewContactUseCase(repo repository.ContactRepository) *ContactUseCaseImpl {
	return &ContactUseCaseImpl{
		repo: repo,
	}
}