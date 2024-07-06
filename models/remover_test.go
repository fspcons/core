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

func TestDefaultDeleter_Delete(t *testing.T) {
	type (
		expected struct {
			error
			removeCallsCnt int
		}
		testCase struct {
			name        string
			id          uuid.UUID
			removeError error
			expected
		}
	)

	var (
		errExpected = errors.New("expected error")
		tcs         = []testCase{
			{
				name:        "invalid id",
				id:          uuid.Nil,
				removeError: nil,
				expected: expected{
					error:          errs.ErrInvalidID,
					removeCallsCnt: 0,
				},
			},
			{
				name:        "remover error",
				id:          uuid.New(),
				removeError: errExpected,
				expected: expected{
					error:          errExpected,
					removeCallsCnt: 1,
				},
			},
			{
				name: "success",
				id:   uuid.New(),
				expected: expected{
					error:          nil,
					removeCallsCnt: 1,
				},
			},
		}
	)

	for _, tCase := range tcs {
		t.Run(tCase.name, func(t *testing.T) {
			ctx := context.TODO()
			deleterMock := &mocks.DeleterMock[string]{
				DeleteFunc: func(ctx2 context.Context, id uuid.UUID) error {
					assert.Equal(t, ctx, ctx2)
					assert.Equal(t, tCase.id, id)
					return tCase.removeError
				},
			}
			d := models.NewDefaultDeleter[string](deleterMock)
			err := d.Delete(ctx, tCase.id)
			if tCase.expected.error == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tCase.expected.error)
			}
			assert.Len(t, deleterMock.DeleteCalls(), tCase.expected.removeCallsCnt)
		})
	}
}
