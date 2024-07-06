// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/fspcons/core/datas"
	"sync"
)

// Ensure, that NoSqlQuerierMock does implement datas.NoSqlQuerier.
// If this is not the case, regenerate this file with moq.
var _ datas.NoSqlQuerier[datas.Filterable, any] = &NoSqlQuerierMock[datas.Filterable, any]{}

// NoSqlQuerierMock is a mock implementation of datas.NoSqlQuerier.
//
//	func TestSomethingThatUsesNoSqlQuerier(t *testing.T) {
//
//		// make and configure a mocked datas.NoSqlQuerier
//		mockedNoSqlQuerier := &NoSqlQuerierMock{
//			QueryFunc: func(ctx context.Context, p datas.NoSqlPaginatorFilter[Q]) (*datas.NoSqlPaginatedList[T], error) {
//				panic("mock out the Query method")
//			},
//		}
//
//		// use mockedNoSqlQuerier in code that requires datas.NoSqlQuerier
//		// and then make assertions.
//
//	}
type NoSqlQuerierMock[Q datas.Filterable, T any] struct {
	// QueryFunc mocks the Query method.
	QueryFunc func(ctx context.Context, p datas.NoSqlPaginatorFilter[Q]) (*datas.NoSqlPaginatedList[T], error)

	// calls tracks calls to the methods.
	calls struct {
		// Query holds details about calls to the Query method.
		Query []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P datas.NoSqlPaginatorFilter[Q]
		}
	}
	lockQuery sync.RWMutex
}

// Query calls QueryFunc.
func (mock *NoSqlQuerierMock[Q, T]) Query(ctx context.Context, p datas.NoSqlPaginatorFilter[Q]) (*datas.NoSqlPaginatedList[T], error) {
	if mock.QueryFunc == nil {
		panic("NoSqlQuerierMock.QueryFunc: method is nil but NoSqlQuerier.Query was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   datas.NoSqlPaginatorFilter[Q]
	}{
		Ctx: ctx,
		P:   p,
	}
	mock.lockQuery.Lock()
	mock.calls.Query = append(mock.calls.Query, callInfo)
	mock.lockQuery.Unlock()
	return mock.QueryFunc(ctx, p)
}

// QueryCalls gets all the calls that were made to Query.
// Check the length with:
//
//	len(mockedNoSqlQuerier.QueryCalls())
func (mock *NoSqlQuerierMock[Q, T]) QueryCalls() []struct {
	Ctx context.Context
	P   datas.NoSqlPaginatorFilter[Q]
} {
	var calls []struct {
		Ctx context.Context
		P   datas.NoSqlPaginatorFilter[Q]
	}
	mock.lockQuery.RLock()
	calls = mock.calls.Query
	mock.lockQuery.RUnlock()
	return calls
}