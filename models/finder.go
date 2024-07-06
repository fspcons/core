package models

import (
	"context"

	"github.com/google/uuid"

	"github.com/fspcons/core/datas"
)

type DefaultFinder[T any] struct {
	finder datas.Finder[T]
}

func (g *DefaultFinder[T]) FindBy(ctx context.Context, id uuid.UUID) (t T, err error) {
	if err = ValidateID(id); err != nil {
		return t, err
	}
	return g.finder.FindBy(ctx, id)
}

func NewDefaultFinder[T any](finder datas.Finder[T]) *DefaultFinder[T] {
	return &DefaultFinder[T]{finder: finder}
}
