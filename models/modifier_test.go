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

func TestDefaultUpdater_Update(t *testing.T) {
	type (
		expected struct {
			error
			inputValidateCallsCnt    int
			getCallsCnt              int
			applyChangesCallsCnt     int
			processCallsCnt          int
			refreshTimestampCallsCnt int
			entityValidateCallsCnt   int
			saveCallsCnt             int
		}
		testCase struct {
			name                string
			omitProcessor       bool
			id                  uuid.UUID
			inputValidateError  error
			getError            error
			applyChangesResult  bool
			processError        error
			entityValidateError error
			saveError           error
			checkResult         bool
			expected
		}
	)

	var (
		id          = uuid.New()
		errExpected = errors.New("expected error")
		testCases   = []testCase{
			{
				name:                "invalid id",
				omitProcessor:       false,
				id:                  uuid.Nil,
				inputValidateError:  nil,
				getError:            nil,
				applyChangesResult:  true,
				processError:        nil,
				entityValidateError: nil,
				saveError:           nil,
				checkResult:         false,
				expected: expected{
					error:                    errs.ErrInvalidID,
					inputValidateCallsCnt:    0,
					getCallsCnt:              0,
					applyChangesCallsCnt:     0,
					processCallsCnt:          0,
					refreshTimestampCallsCnt: 0,
					entityValidateCallsCnt:   0,
					saveCallsCnt:             0,
				},
			},
			{
				name:                "invalid input",
				omitProcessor:       false,
				id:                  id,
				inputValidateError:  errExpected,
				getError:            nil,
				applyChangesResult:  true,
				processError:        nil,
				entityValidateError: nil,
				saveError:           nil,
				checkResult:         false,
				expected: expected{
					error:                    errExpected,
					inputValidateCallsCnt:    1,
					getCallsCnt:              0,
					applyChangesCallsCnt:     0,
					processCallsCnt:          0,
					refreshTimestampCallsCnt: 0,
					entityValidateCallsCnt:   0,
					saveCallsCnt:             0,
				},
			},
			{
				name:                "get error",
				omitProcessor:       false,
				id:                  id,
				inputValidateError:  nil,
				getError:            errExpected,
				applyChangesResult:  true,
				processError:        nil,
				entityValidateError: nil,
				saveError:           nil,
				checkResult:         false,
				expected: expected{
					error:                    errExpected,
					inputValidateCallsCnt:    1,
					getCallsCnt:              1,
					applyChangesCallsCnt:     0,
					processCallsCnt:          0,
					refreshTimestampCallsCnt: 0,
					entityValidateCallsCnt:   0,
					saveCallsCnt:             0,
				},
			},
			{
				name:                "no changes",
				omitProcessor:       false,
				id:                  id,
				inputValidateError:  nil,
				getError:            nil,
				applyChangesResult:  false,
				processError:        nil,
				entityValidateError: nil,
				saveError:           nil,
				checkResult:         false,
				expected: expected{
					error:                    errs.ErrNoChanges,
					inputValidateCallsCnt:    1,
					getCallsCnt:              1,
					applyChangesCallsCnt:     1,
					processCallsCnt:          0,
					refreshTimestampCallsCnt: 0,
					entityValidateCallsCnt:   0,
					saveCallsCnt:             0,
				},
			},
			{
				name:                "process error",
				omitProcessor:       false,
				id:                  id,
				inputValidateError:  nil,
				getError:            nil,
				applyChangesResult:  true,
				processError:        errExpected,
				entityValidateError: nil,
				saveError:           nil,
				checkResult:         false,
				expected: expected{
					error:                    errExpected,
					inputValidateCallsCnt:    1,
					getCallsCnt:              1,
					applyChangesCallsCnt:     1,
					processCallsCnt:          1,
					refreshTimestampCallsCnt: 0,
					entityValidateCallsCnt:   0,
					saveCallsCnt:             0,
				},
			},
			{
				name:                "invalid entity",
				omitProcessor:       false,
				id:                  id,
				inputValidateError:  nil,
				getError:            nil,
				applyChangesResult:  true,
				processError:        nil,
				entityValidateError: errExpected,
				saveError:           nil,
				checkResult:         false,
				expected: expected{
					error:                    errExpected,
					inputValidateCallsCnt:    1,
					getCallsCnt:              1,
					applyChangesCallsCnt:     1,
					processCallsCnt:          1,
					refreshTimestampCallsCnt: 1,
					entityValidateCallsCnt:   1,
					saveCallsCnt:             0,
				},
			},
			{
				name:                "save error",
				omitProcessor:       false,
				id:                  id,
				inputValidateError:  nil,
				getError:            nil,
				applyChangesResult:  true,
				processError:        nil,
				entityValidateError: nil,
				saveError:           errExpected,
				checkResult:         false,
				expected: expected{
					error:                    errExpected,
					inputValidateCallsCnt:    1,
					getCallsCnt:              1,
					applyChangesCallsCnt:     1,
					processCallsCnt:          1,
					refreshTimestampCallsCnt: 1,
					entityValidateCallsCnt:   1,
					saveCallsCnt:             1,
				},
			},
			{
				name:                "success",
				omitProcessor:       false,
				id:                  id,
				inputValidateError:  nil,
				getError:            nil,
				applyChangesResult:  true,
				processError:        nil,
				entityValidateError: nil,
				saveError:           nil,
				checkResult:         true,
				expected: expected{
					error:                    nil,
					inputValidateCallsCnt:    1,
					getCallsCnt:              1,
					applyChangesCallsCnt:     1,
					processCallsCnt:          1,
					refreshTimestampCallsCnt: 1,
					entityValidateCallsCnt:   1,
					saveCallsCnt:             1,
				},
			},
			{
				name:                "success without processor",
				omitProcessor:       true,
				id:                  id,
				inputValidateError:  nil,
				getError:            nil,
				applyChangesResult:  true,
				processError:        nil,
				entityValidateError: nil,
				saveError:           nil,
				checkResult:         false,
				expected: expected{
					error:                    nil,
					inputValidateCallsCnt:    1,
					getCallsCnt:              1,
					applyChangesCallsCnt:     1,
					processCallsCnt:          1,
					refreshTimestampCallsCnt: 1,
					entityValidateCallsCnt:   1,
					saveCallsCnt:             1,
				},
			},
		}
	)

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			ctx := context.TODO()
			refresherMock := &mocks.RefresherMock{
				ValidateFunc: func() error {
					return tCase.entityValidateError
				},
			}
			oldRefresherMock := &mocks.RefresherMock{}
			refresherMock.RefreshTimestampFunc = func() *mocks.RefresherMock {
				return refresherMock
			}
			inputMock := &mocks.UpdateInputMock[*mocks.RefresherMock]{
				ValidateFunc: func() error {
					return tCase.inputValidateError
				},
				ApplyChangesFunc: func(r *mocks.RefresherMock) (*mocks.RefresherMock, bool) {
					assert.Equal(t, oldRefresherMock, r)
					return refresherMock, tCase.applyChangesResult
				},
			}
			getterMock := &mocks.GetterMock[*mocks.RefresherMock]{
				GetFunc: func(ctx2 context.Context, id uuid.UUID) (*mocks.RefresherMock, error) {
					assert.Equal(t, ctx, ctx2)
					assert.Equal(t, tCase.id, id)
					return oldRefresherMock, tCase.getError
				},
			}
			processorMock := &mocks.UpdateProcessorMock[*mocks.RefresherMock]{
				ProcessFunc: func(ctx2 context.Context, upd **mocks.RefresherMock, old *mocks.RefresherMock) error {
					assert.Equal(t, ctx, ctx2)
					assert.Equal(t, refresherMock, *upd)
					assert.Equal(t, oldRefresherMock, old)
					return tCase.processError
				},
			}
			updaterMock := &mocks.UpdaterMock[*mocks.RefresherMock]{
				UpdateFunc: func(ctx2 context.Context, r *mocks.RefresherMock) error {
					assert.Equal(t, ctx, ctx2)
					assert.Equal(t, refresherMock, r)
					return tCase.saveError
				},
			}
			var processor models.UpdateProcessor[*mocks.RefresherMock]
			if !tCase.omitProcessor {
				processor = processorMock
			}
			u := models.NewDefaultModifier[*mocks.RefresherMock](getterMock, processor, updaterMock)
			got, err := u.Modify(ctx, tCase.id, inputMock)
			if tCase.expected.error != nil {
				assert.ErrorIs(t, err, tCase.expected.error)
			} else {
				assert.NoError(t, err)
			}
			if tCase.checkResult {
				assert.Equal(t, refresherMock, got)
			}
			assert.Len(t, inputMock.ValidateCalls(), tCase.expected.inputValidateCallsCnt)
			assert.Len(t, getterMock.GetCalls(), tCase.expected.getCallsCnt)
			assert.Len(t, inputMock.ApplyChangesCalls(), tCase.expected.applyChangesCallsCnt)
			if !tCase.omitProcessor {
				assert.Len(t, processorMock.ProcessCalls(), tCase.expected.processCallsCnt)
			}
			assert.Len(t, refresherMock.RefreshTimestampCalls(), tCase.expected.refreshTimestampCallsCnt)
			assert.Len(t, refresherMock.ValidateCalls(), tCase.expected.entityValidateCallsCnt)
			assert.Len(t, updaterMock.UpdateCalls(), tCase.expected.saveCallsCnt)
		})
	}
}
