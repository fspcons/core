package models

import (
	"context"

	"github.com/fspcons/core/datas"
	"github.com/fspcons/core/errs"
)

type DefaultSqlLister[Q datas.Filterable, T any] struct {
	maxPageSize uint
	querier     datas.SqlQuerier[Q, T]
}

func (ref *DefaultSqlLister[Q, T]) List(ctx context.Context, f datas.SqlPaginatorFilter[Q]) ([]T, error) {
	var err error

	if f.PageSize < 1 || f.PageSize > ref.maxPageSize {
		err = errs.AppendNewError(err, "invalid page size")
	}
	if f.PageNumber < 1 {
		err = errs.AppendNewError(err, "invalid page number")
	}
	if err != nil {
		return nil, err
	}

	return ref.querier.Query(ctx, f)
}

func NewDefaultSqlLister[Q datas.Filterable, T any](querier datas.SqlQuerier[Q, T], maxPageSize uint) *DefaultSqlLister[Q, T] {
	return &DefaultSqlLister[Q, T]{maxPageSize: maxPageSize, querier: querier}
}
