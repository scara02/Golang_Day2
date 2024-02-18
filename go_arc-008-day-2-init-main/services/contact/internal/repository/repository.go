package repository

import "architecture_go/services/contact/internal/domain"

type Repository struct {
    ContactRepository ContactRepository
}

func NewRepository(contactRepository ContactRepository) *Repository {
    return &Repository{
        ContactRepository: contactRepository,
    }
}
