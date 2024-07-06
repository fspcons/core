package models_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/fspcons/core/errs"
	"github.com/fspcons/core/models"
	"github.com/fspcons/core/tests/mocks"
)

func TestDefaultFinder_Get(t *testing.T) {
	type (
		expected struct {
			result       *string
			error        error
			findCallsCnt int
		}
		testCase struct {
			name       string
			id         uuid.UUID
			findResult string
			findError  error
			expected
		}
	)

	var (
		errExpected = errors.New("expected error")
		resExpected = "expected result"
		tcs         = []testCase{
			{
				name:       "invalid id",
				id:         uuid.Nil,
				findResult: "",
				findError:  nil,
				expected: expected{
					result:       nil,
					error:        errs.ErrInvalidID,
					findCallsCnt: 0,
				},
			},
			{
				name:       "finder error",
				id:         uuid.New(),
				findResult: "",
				findError:  errExpected,
				expected: expected{
					result:       nil,
					error:        errExpected,
					findCallsCnt: 1,
				},
			},
			{
				name:       "success",
				id:         uuid.New(),
				findResult: resExpected,
				findError:  nil,
				expected: expected{
					result:       &resExpected,
					error:        nil,
					findCallsCnt: 1,
				},
			},
		}
	)

	for _, tCase := range tcs {
		t.Run(tCase.name, func(t *testing.T) {
			ctx := context.TODO()
			finderMock := &mocks.FinderMock[string]{
				FindByFunc: func(ctx2 context.Context, id uuid.UUID) (string, error) {
					assert.Equal(t, ctx, ctx2)
					assert.Equal(t, tCase.id, id)
					return tCase.findResult, tCase.findError
				},
			}
			g := models.NewDefaultFinder[string](finderMock)
			res, err := g.FindBy(ctx, tCase.id)
			if tCase.expected.result != nil {
				assert.Equal(t, *tCase.expected.result, res)
			}
			if tCase.expected.error == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tCase.expected.error)
			}
			assert.Len(t, finderMock.FindByCalls(), tCase.expected.findCallsCnt)
		})
	}
}
