package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `json:"email"`
}

type Campaign struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
}

func toContacts(emails []string) []Contact {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i] = Contact{Email: email}
	}
	return contacts
}

func New(name, content string, contacts []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("Name is required")
	} else if content == "" {
		return nil, errors.New("Content is required")
	} else if len(contacts) == 0 {
		return nil, errors.New("At least one contact is required")
	}
	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  toContacts(contacts),
	}, nil
}
