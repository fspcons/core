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

func TestDefaultNoSqlLister_List(t *testing.T) {
	type (
		expected struct {
			result *datas.NoSqlPaginatedList[string]
			error
			businessErrorMessage string
			indexCallsCnt        int
		}
		testCase struct {
			name         string
			nextPageHash string
			findResult   *datas.NoSqlPaginatedList[string]
			findError    error
			expected
		}
	)

	var (
		errExpected = errors.New("expected error")
		resExpected = &datas.NoSqlPaginatedList[string]{List: []string{"foo", "bar", "baz"}}
		tcs         = []testCase{
			{
				name:         "querier error",
				nextPageHash: "1",
				findResult:   nil,
				findError:    errExpected,
				expected: expected{
					result:               nil,
					error:                errExpected,
					businessErrorMessage: "",
					indexCallsCnt:        1,
				},
			},
			{
				name:         "querier success",
				nextPageHash: "1",
				findResult:   resExpected,
				findError:    nil,
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
			querierMock := &mocks.NoSqlQuerierMock[mocks.NoSqlFilterable, string]{
				QueryFunc: func(ctx2 context.Context, f datas.NoSqlPaginatorFilter[mocks.NoSqlFilterable]) (*datas.NoSqlPaginatedList[string], error) {
					assert.Equal(innerT, ctx, ctx2)
					assert.Equal(innerT, tCase.nextPageHash, f.NextPageHash)
					return tCase.findResult, tCase.findError
				},
			}
			l := models.NewDefaultNoSqlLister[mocks.NoSqlFilterable, string](querierMock)
			res, err := l.List(ctx, datas.NoSqlPaginatorFilter[mocks.NoSqlFilterable]{
				NextPageHash: tCase.nextPageHash})
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
