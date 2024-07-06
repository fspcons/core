package models

import (
	"context"

	"github.com/google/uuid"
)

// Validator validates the data it holds
//
//go:generate moq -out ./tests/mocks/validator.go -pkg mocks . Validator:ValidatorMock
type Validator interface {
	Validate() error
}

//go:generate moq -out ./tests/mocks/create_input.go -pkg mocks . CreateInput:CreateInputMock
type CreateInput[V Validator] interface {
	Transform() V
	Validator
}

//go:generate moq -out ./tests/mocks/create_processor.go -pkg mocks . CreateProcessor:CreateProcessorMock
type CreateProcessor[V Validator] interface {
	Process(ctx context.Context, v *V) error
}

type Creator[V Validator] interface {
	Create(ctx context.Context, in CreateInput[V]) (V, error)
}

type Lister[T any] interface {
	List(ctx context.Context, pageSize, pageNumber uint) ([]T, error)
}

//go:generate moq -out ./tests/mocks/getter.go -pkg mocks . Getter:GetterMock
type Getter[T any] interface {
	Get(ctx context.Context, id uuid.UUID) (T, error)
}

type Verifier[T any] interface {
	Verify(ctx context.Context, id uuid.UUID) error
}

//go:generate moq -out ./tests/mocks/update_processor.go -pkg mocks . UpdateProcessor:UpdateProcessorMock
type UpdateProcessor[V Validator] interface {
	Process(ctx context.Context, upd *V, old V) error
}

//go:generate moq -out ./tests/mocks/refresher.go -pkg mocks . Refresher:RefresherMock
type Refresher[T any] interface {
	RefreshTimestamp() T
	Validator
}

//go:generate moq -out ./tests/mocks/update_input.go -pkg mocks . UpdateInput:UpdateInputMock
type UpdateInput[R Refresher[R]] interface {
	ApplyChanges(old R) (upd R, hasChanges bool)
	Validator
}

type Updater[R Refresher[R]] interface {
	Update(ctx context.Context, id uuid.UUID, in UpdateInput[R]) (R, error)
}

type Deleter[T any] interface {
	Delete(ctx context.Context, id uuid.UUID) error
}
