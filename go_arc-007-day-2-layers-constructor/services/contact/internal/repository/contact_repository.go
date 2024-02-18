package repository

import "architecture_go/services/contact/internal/domain"

type ContactRepository interface {
	CreateContact(id int, fullName string, phoneNumber string) error
    GetContact(id int) (*domain.Contact, error)
    UpdateContact(id int, fullName string, phoneNumber string) error
    DeleteContact(id int) error

	CreateGroup(name string) error

    GetGroup(name string) (*domain.Group, error)

    AddContactToGroup(contactID int, groupName string) error
}
