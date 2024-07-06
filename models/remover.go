package models

import (
	"context"

	"github.com/google/uuid"

	"github.com/fspcons/core/datas"
)

type DefaultRemover[T any] struct {
	deleter datas.Deleter[T]
}

func (r *DefaultRemover[T]) Delete(ctx context.Context, id uuid.UUID) error {
	if err := ValidateID(id); err != nil {
		return err
	}
	return r.deleter.Delete(ctx, id)
}

func NewDefaultDeleter[T any](deleter datas.Deleter[T]) *DefaultRemover[T] {
	return &DefaultRemover[T]{deleter: deleter}
}
