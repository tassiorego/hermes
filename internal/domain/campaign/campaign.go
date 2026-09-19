package campaign

import (
	internalerrors "hermes/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"required,email" json:"email"`
}

type Campaign struct {
	ID        string    `validate:"required" json:"id"`
	Name      string    `validate:"required,min=5,max=100" json:"name"`
	CreatedAt time.Time `validate:"required" json:"created_at"`
	Content   string    `validate:"required,min=10,max=1024" json:"content"`
	Contacts  []Contact `validate:"required,min=1,dive" json:"contacts"`
}

func toContacts(emails []string) []Contact {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i] = Contact{Email: email}
	}
	return contacts
}

func New(name, content string, contacts []string) (*Campaign, error) {
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  toContacts(contacts),
	}

	err := internalerrors.ValidateStruct(campaign)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}
