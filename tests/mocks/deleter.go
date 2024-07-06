// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/fspcons/core/datas"
	"github.com/google/uuid"
	"sync"
)

// Ensure, that DeleterMock does implement datas.Deleter.
// If this is not the case, regenerate this file with moq.
var _ datas.Deleter[any] = &DeleterMock[any]{}

// DeleterMock is a mock implementation of datas.Deleter.
//
//	func TestSomethingThatUsesDeleter(t *testing.T) {
//
//		// make and configure a mocked datas.Deleter
//		mockedDeleter := &DeleterMock{
//			DeleteFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the Delete method")
//			},
//		}
//
//		// use mockedDeleter in code that requires datas.Deleter
//		// and then make assertions.
//
//	}
type DeleterMock[T any] struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, id uuid.UUID) error

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
	}
	lockDelete sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *DeleterMock[T]) Delete(ctx context.Context, id uuid.UUID) error {
	if mock.DeleteFunc == nil {
		panic("DeleterMock.DeleteFunc: method is nil but Deleter.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedDeleter.DeleteCalls())
func (mock *DeleterMock[T]) DeleteCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}