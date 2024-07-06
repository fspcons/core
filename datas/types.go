package datas

import (
	"context"

	"github.com/google/uuid"
)

type (

	//Filterable generic interface
	Filterable interface {
		GetFilter() map[string]any
	}

	//SqlPaginatorFilter generic sql pagination query filter
	SqlPaginatorFilter[Q Filterable] struct {
		Filter     Q
		PageSize   uint
		PageNumber uint
	}

	//NoSqlPaginatorFilter generic nosql pagination query filter
	NoSqlPaginatorFilter[Q Filterable] struct {
		Filter       Q
		NextPageHash string
	}

	//NoSqlPaginatedList generic nosql pagination response
	NoSqlPaginatedList[T any] struct {
		List         []T
		NextPageHash string
	}

	//go:generate moq -out ../tests/mocks/inserter.go -pkg mocks . Inserter:InserterMock
	// Inserter can insert a new T record into the database
	Inserter[T any] interface {
		Insert(ctx context.Context, model T) error
	}

	//go:generate moq -out ../tests/mocks/deleter.go -pkg mocks . Deleter:DeleterMock
	// Deleter can delete an existing T record from the database
	Deleter[T any] interface {
		Delete(ctx context.Context, id uuid.UUID) error
	}

	//go:generate moq -out ../tests/mocks/finder.go -pkg mocks . Finder:FinderMock
	// Finder can find an existing T record in the database
	Finder[T any] interface {
		FindBy(ctx context.Context, id uuid.UUID) (T, error)
	}

	//go:generate moq -out ../tests/mocks/updater.go -pkg mocks . Updater:UpdaterMock
	// Updater can update an existing T record in the database
	Updater[T any] interface {
		// Update updates an existing T record in the database
		Update(ctx context.Context, t T) error
	}

	//go:generate moq -out ../tests/mocks/sql_querier.go -pkg mocks . SqlQuerier:SqlQuerierMock
	// SqlQuerier can query with Q filter the existing T records in the sql database
	SqlQuerier[Q Filterable, T any] interface {
		Query(ctx context.Context, p SqlPaginatorFilter[Q]) ([]T, error)
	}

	//go:generate moq -out ../tests/mocks/nosql_querier.go -pkg mocks . NoSqlQuerier:NoSqlQuerierMock
	// NoSqlQuerier can query with Q filter the existing T records in the nosql database
	NoSqlQuerier[Q Filterable, T any] interface {
		Query(ctx context.Context, p NoSqlPaginatorFilter[Q]) (*NoSqlPaginatedList[T], error)
	}

	//go:generate moq -out ../tests/mocks/checker.go -pkg mocks . Checker:CheckerMock
	// Checker can check whether a T record exists in the database
	Checker[T any] interface {
		// Exists checks whether a T record exists in the database
		Exists(ctx context.Context, id uuid.UUID) (bool, error)
	}
)
