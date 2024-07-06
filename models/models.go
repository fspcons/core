package models

import (
	"time"

	"github.com/google/uuid"

	"github.com/fspcons/core/errs"
)

// NewID generates a new id
func NewID() uuid.UUID {
	return uuid.New()
}

// IDIsValid checks if an id is valid
func IDIsValid(id uuid.UUID) bool {
	return id != uuid.Nil
}

// ValidateID validates an id
func ValidateID(id uuid.UUID) error {
	if !IDIsValid(id) {
		return errs.ErrInvalidID
	}
	return nil
}

func ValidateTimestamps(createdAt, updatedAt time.Time) (err error) {
	if createdAt.IsZero() {
		err = errs.AppendNewError(err, "invalid created at")
	}
	if updatedAt.IsZero() {
		err = errs.AppendNewError(err, "invalid updated at")
	}
	if createdAt.After(updatedAt) {
		err = errs.AppendNewError(err, "key created at is greater than updated at")
	}
	return err
}
