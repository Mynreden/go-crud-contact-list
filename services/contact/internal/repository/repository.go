package repository

import "ass3/services/contact/internal/domain"

type ContactRepository interface {
	CreateContact(contact domain.Contact) error
	ReadContact(contactID int) (domain.Contact, error)
	UpdateContact(contactID int, newContact domain.Contact) error
	DeleteContact(contactID int) error

	CreateGroup(group domain.Group) error
	ReadGroup(groupID int) (domain.Group, error)

	AddContactToGroup(contactID, groupID int) error
}
