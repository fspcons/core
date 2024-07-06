package models_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fspcons/core/datas"
	"github.com/fspcons/core/errs"
	"github.com/fspcons/core/models"
	"github.com/fspcons/core/tests/mocks"
)

func TestDefaultSqlLister_List(t *testing.T) {
	type (
		expected struct {
			result []string
			error
			businessErrorMessage string
			indexCallsCnt        int
		}
		testCase struct {
			name       string
			pageSize   uint
			pageNumber uint
			findResult []string
			findError  error
			expected
		}
	)

	var (
		maxPageSize = uint(50)
		errExpected = errors.New("expected error")
		resExpected = []string{"foo", "bar", "baz"}
		tcs         = []testCase{
			{
				name:       "page size is too small",
				pageSize:   0,
				pageNumber: 1,
				findResult: nil,
				findError:  nil,
				expected: expected{
					result:               nil,
					error:                nil,
					businessErrorMessage: "invalid page size",
					indexCallsCnt:        0,
				},
			},
			{
				name:       "page size is too large",
				pageSize:   maxPageSize + 1,
				pageNumber: 1,
				findResult: nil,
				findError:  nil,
				expected: expected{
					result:               nil,
					error:                nil,
					businessErrorMessage: "invalid page size",
					indexCallsCnt:        0,
				},
			},
			{
				name:       "page number is too small",
				pageSize:   maxPageSize,
				pageNumber: 0,
				findResult: nil,
				findError:  nil,
				expected: expected{
					result:               nil,
					error:                nil,
					businessErrorMessage: "invalid page number",
					indexCallsCnt:        0,
				},
			},
			{
				name:       "invalid page size and number",
				pageSize:   0,
				pageNumber: 0,
				findResult: nil,
				findError:  nil,
				expected: expected{
					result:               nil,
					error:                nil,
					businessErrorMessage: "many errors: invalid page size, invalid page number",
					indexCallsCnt:        0,
				},
			},
			{
				name:       "indexer error",
				pageSize:   maxPageSize,
				pageNumber: 1,
				findResult: nil,
				findError:  errExpected,
				expected: expected{
					result:               nil,
					error:                errExpected,
					businessErrorMessage: "",
					indexCallsCnt:        1,
				},
			},
			{
				name:       "indexer error",
				pageSize:   maxPageSize,
				pageNumber: 1,
				findResult: resExpected,
				findError:  nil,
				expected: expected{
					result:               resExpected,
					error:                nil,
					businessErrorMessage: "",
					indexCallsCnt:        1,
				},
			},
		}
	)

	for _, tCase := range tcs {
		t.Run(tCase.name, func(innerT *testing.T) {
			ctx := context.TODO()
			querierMock := &mocks.SqlQuerierMock[mocks.SqlFilterable, string]{
				QueryFunc: func(ctx2 context.Context, f datas.SqlPaginatorFilter[mocks.SqlFilterable]) ([]string, error) {
					assert.Equal(innerT, ctx, ctx2)
					assert.Equal(innerT, tCase.pageSize, f.PageSize)
					assert.Equal(innerT, tCase.pageSize*(tCase.pageNumber-1), f.PageSize*(f.PageNumber-1))
					return tCase.findResult, tCase.findError
				},
			}
			l := models.NewDefaultSqlLister[mocks.SqlFilterable, string](querierMock, maxPageSize)
			res, err := l.List(ctx, datas.SqlPaginatorFilter[mocks.SqlFilterable]{
				PageSize: tCase.pageSize, PageNumber: tCase.pageNumber})
			if tCase.expected.result != nil {
				assert.Equal(innerT, tCase.expected.result, res)
			}
			if tCase.expected.error != nil {
				assert.ErrorIs(innerT, err, tCase.expected.error)
			} else if tCase.expected.businessErrorMessage != "" {
				assert.EqualError(innerT, err, tCase.expected.businessErrorMessage)
				assert.ErrorAs(innerT, err, new(*errs.Error))
			} else {
				assert.NoError(innerT, err)
			}
			assert.Len(innerT, querierMock.QueryCalls(), tCase.expected.indexCallsCnt)
		})
	}
}
