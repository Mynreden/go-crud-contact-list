package delivery

import (
	"net/http"
)

type ContactDelivery interface {
	// HTTP request handlers for interacting with contacts
	CreateContactHandler(w http.ResponseWriter, r *http.Request)
	ReadContactHandler(w http.ResponseWriter, r *http.Request)
	UpdateContactHandler(w http.ResponseWriter, r *http.Request)
	DeleteContactHandler(w http.ResponseWriter, r *http.Request)

	// HTTP request handlers for interacting with groups
	CreateGroupHandler(w http.ResponseWriter, r *http.Request)
	ReadGroupHandler(w http.ResponseWriter, r *http.Request)

	// HTTP request handler for adding a contact to a group
	AddContactToGroupHandler(w http.ResponseWriter, r *http.Request)
}
