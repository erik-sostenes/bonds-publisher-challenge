package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// UserID represents the user unique identifier
type UserID string

func (u UserID) Validate() (*UserID, error) {
	_, err := uuid.Parse(string(u))
	if err != nil {
		return nil, fmt.Errorf("%w = %s", InvalidUserID, err.Error())
	}

	return &u, nil
}

func (u UserID) ID() string {
	return string(u)
}
