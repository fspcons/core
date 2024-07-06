package models

import (
	"context"

	"github.com/google/uuid"

	"github.com/fspcons/core/datas"
	"github.com/fspcons/core/errs"
)

type DefaultVerifier[T any] struct {
	checker datas.Checker[T]
}

func (v *DefaultVerifier[T]) Verify(ctx context.Context, id uuid.UUID) (err error) {
	if err = ValidateID(id); err != nil {
		return err
	}
	exists, err := v.checker.Exists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return errs.ErrNotFound
	}
	return nil
}

func NewDefaultVerifier[T any](checker datas.Checker[T]) *DefaultVerifier[T] {
	return &DefaultVerifier[T]{checker: checker}
}
