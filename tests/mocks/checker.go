// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/fspcons/core/datas"
	"github.com/google/uuid"
	"sync"
)

// Ensure, that CheckerMock does implement datas.Checker.
// If this is not the case, regenerate this file with moq.
var _ datas.Checker[any] = &CheckerMock[any]{}

// CheckerMock is a mock implementation of datas.Checker.
//
//	func TestSomethingThatUsesChecker(t *testing.T) {
//
//		// make and configure a mocked datas.Checker
//		mockedChecker := &CheckerMock{
//			ExistsFunc: func(ctx context.Context, id uuid.UUID) (bool, error) {
//				panic("mock out the Exists method")
//			},
//		}
//
//		// use mockedChecker in code that requires datas.Checker
//		// and then make assertions.
//
//	}
type CheckerMock[T any] struct {
	// ExistsFunc mocks the Exists method.
	ExistsFunc func(ctx context.Context, id uuid.UUID) (bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// Exists holds details about calls to the Exists method.
		Exists []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
	}
	lockExists sync.RWMutex
}

// Exists calls ExistsFunc.
func (mock *CheckerMock[T]) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	if mock.ExistsFunc == nil {
		panic("CheckerMock.ExistsFunc: method is nil but Checker.Exists was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockExists.Lock()
	mock.calls.Exists = append(mock.calls.Exists, callInfo)
	mock.lockExists.Unlock()
	return mock.ExistsFunc(ctx, id)
}

// ExistsCalls gets all the calls that were made to Exists.
// Check the length with:
//
//	len(mockedChecker.ExistsCalls())
func (mock *CheckerMock[T]) ExistsCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockExists.RLock()
	calls = mock.calls.Exists
	mock.lockExists.RUnlock()
	return calls
}
