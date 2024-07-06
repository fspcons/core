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

func TestDefaultVerifier_Verify(t *testing.T) {
	type (
		expected struct {
			error
			existsCallsCnt int
		}
		testCase struct {
			name         string
			id           uuid.UUID
			existsError  error
			existsReturn bool
			expected
		}
	)

	var (
		id          = uuid.New()
		errExpected = errors.New("expected error")
		testCases   = []testCase{
			{
				name:         "invalid id",
				id:           uuid.Nil,
				existsError:  nil,
				existsReturn: false,
				expected: expected{
					error:          errs.ErrInvalidID,
					existsCallsCnt: 0,
				},
			},
			{
				name:         "exists error",
				id:           id,
				existsError:  errExpected,
				existsReturn: false,
				expected: expected{
					error:          errExpected,
					existsCallsCnt: 1,
				},
			},
			{
				name:         "not found",
				id:           id,
				existsError:  nil,
				existsReturn: false,
				expected: expected{
					error:          errs.ErrNotFound,
					existsCallsCnt: 1,
				},
			},
			{
				name:         "success",
				id:           id,
				existsError:  nil,
				existsReturn: true,
				expected: expected{
					error:          nil,
					existsCallsCnt: 1,
				},
			},
		}
	)

	for _, tCase := range testCases {
		t.Run(tCase.name, func(innerT *testing.T) {
			ctx := context.TODO()
			checkerMock := &mocks.CheckerMock[string]{
				ExistsFunc: func(ctx2 context.Context, id uuid.UUID) (bool, error) {
					assert.Equal(innerT, ctx, ctx2)
					assert.Equal(innerT, tCase.id, id)
					return tCase.existsReturn, tCase.existsError
				},
			}
			v := models.NewDefaultVerifier[string](checkerMock)
			err := v.Verify(ctx, tCase.id)
			if tCase.expected.error != nil {
				assert.ErrorIs(innerT, err, tCase.expected.error)
			} else {
				assert.NoError(innerT, err)
			}
			assert.Len(innerT, checkerMock.ExistsCalls(), tCase.expected.existsCallsCnt)
		})
	}
}
