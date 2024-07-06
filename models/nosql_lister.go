package models

import (
	"context"

	"github.com/fspcons/core/datas"
)

type DefaultNoSqlLister[Q datas.Filterable, T any] struct {
	querier datas.NoSqlQuerier[Q, T]
}

func (ref *DefaultNoSqlLister[Q, T]) List(ctx context.Context, pf datas.NoSqlPaginatorFilter[Q]) (*datas.NoSqlPaginatedList[T], error) {
	return ref.querier.Query(ctx, pf)
}

func NewDefaultNoSqlLister[Q datas.Filterable, T any](querier datas.NoSqlQuerier[Q, T]) *DefaultNoSqlLister[Q, T] {
	return &DefaultNoSqlLister[Q, T]{querier: querier}
}
