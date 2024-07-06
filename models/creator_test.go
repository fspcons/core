package models_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fspcons/core/models"
	"github.com/fspcons/core/tests/mocks"
)

func TestDefaultCreator_Create(t *testing.T) {
	type (
		expected struct {
			error
			inputValidateCallsCnt  int
			inputTransformCallsCnt int
			processCallsCnt        int
			entityValidateCallsCnt int
			insertCallsCnt         int
		}
		testCase struct {
			name                string
			omitProcessor       bool
			inputValidateError  error
			processError        error
			entityValidateError error
			insertError         error
			checkResult         bool
			expected
		}
	)

	var (
		errExpected = errors.New("expected error")
		tcs         = []testCase{
			{
				name:                "invalid input",
				omitProcessor:       false,
				inputValidateError:  errExpected,
				processError:        nil,
				entityValidateError: nil,
				insertError:         nil,
				checkResult:         false,
				expected: expected{
					error:                  errExpected,
					inputValidateCallsCnt:  1,
					inputTransformCallsCnt: 0,
					processCallsCnt:        0,
					entityValidateCallsCnt: 0,
					insertCallsCnt:         0,
				},
			},
			{
				name:                "processor error",
				omitProcessor:       false,
				inputValidateError:  nil,
				processError:        errExpected,
				entityValidateError: nil,
				insertError:         nil,
				checkResult:         false,
				expected: expected{
					error:                  errExpected,
					inputValidateCallsCnt:  1,
					inputTransformCallsCnt: 1,
					processCallsCnt:        1,
					entityValidateCallsCnt: 0,
					insertCallsCnt:         0,
				},
			},
			{
				name:                "invalid entity",
				omitProcessor:       false,
				inputValidateError:  nil,
				processError:        nil,
				entityValidateError: errExpected,
				insertError:         nil,
				checkResult:         false,
				expected: expected{
					error:                  errExpected,
					inputValidateCallsCnt:  1,
					inputTransformCallsCnt: 1,
					processCallsCnt:        1,
					entityValidateCallsCnt: 1,
					insertCallsCnt:         0,
				},
			},
			{
				name:                "inserter error",
				omitProcessor:       false,
				inputValidateError:  nil,
				processError:        nil,
				entityValidateError: nil,
				insertError:         errExpected,
				checkResult:         false,
				expected: expected{
					error:                  errExpected,
					inputValidateCallsCnt:  1,
					inputTransformCallsCnt: 1,
					processCallsCnt:        1,
					entityValidateCallsCnt: 1,
					insertCallsCnt:         1,
				},
			},
			{
				name:                "success",
				omitProcessor:       false,
				inputValidateError:  nil,
				processError:        nil,
				entityValidateError: nil,
				insertError:         nil,
				checkResult:         true,
				expected: expected{
					error:                  nil,
					inputValidateCallsCnt:  1,
					inputTransformCallsCnt: 1,
					processCallsCnt:        1,
					entityValidateCallsCnt: 1,
					insertCallsCnt:         1,
				},
			},
			{
				name:                "success without processor",
				omitProcessor:       true,
				inputValidateError:  nil,
				processError:        nil,
				entityValidateError: nil,
				insertError:         nil,
				checkResult:         true,
				expected: expected{
					error:                  nil,
					inputValidateCallsCnt:  1,
					inputTransformCallsCnt: 1,
					processCallsCnt:        1,
					entityValidateCallsCnt: 1,
					insertCallsCnt:         1,
				},
			},
		}
	)

	for _, tCase := range tcs {
		t.Run(tCase.name, func(t *testing.T) {
			ctx := context.TODO()
			validatorMock := &mocks.ValidatorMock{
				ValidateFunc: func() error {
					return tCase.entityValidateError
				},
			}
			inputMock := &mocks.CreateInputMock[*mocks.ValidatorMock]{
				ValidateFunc: func() error {
					return tCase.inputValidateError
				},
				TransformFunc: func() *mocks.ValidatorMock {
					return validatorMock
				},
			}
			processorMock := &mocks.CreateProcessorMock[*mocks.ValidatorMock]{
				ProcessFunc: func(ctx2 context.Context, v **mocks.ValidatorMock) error {
					assert.Equal(t, ctx, ctx2)
					assert.Equal(t, validatorMock, *v)
					return tCase.processError
				},
			}
			inserterMock := &mocks.InserterMock[*mocks.ValidatorMock]{
				InsertFunc: func(ctx2 context.Context, v *mocks.ValidatorMock) error {
					assert.Equal(t, ctx, ctx2)
					assert.Equal(t, validatorMock, v)
					return tCase.insertError
				},
			}
			var processor models.CreateProcessor[*mocks.ValidatorMock]
			if !tCase.omitProcessor {
				processor = processorMock
			}
			c := models.NewDefaultCreator[*mocks.ValidatorMock](inserterMock, processor)
			v, err := c.Create(ctx, inputMock)
			if tCase.expected.error != nil {
				assert.ErrorIs(t, err, tCase.expected.error)
			} else {
				assert.NoError(t, err)
			}
			if tCase.checkResult {
				assert.Equal(t, validatorMock, v)
			}
			assert.Len(t, inputMock.ValidateCalls(), tCase.expected.inputValidateCallsCnt)
			assert.Len(t, inputMock.TransformCalls(), tCase.expected.inputTransformCallsCnt)
			if !tCase.omitProcessor {
				assert.Len(t, processorMock.ProcessCalls(), tCase.expected.processCallsCnt)
			}
			assert.Len(t, validatorMock.ValidateCalls(), tCase.expected.entityValidateCallsCnt)
			assert.Len(t, inserterMock.InsertCalls(), tCase.expected.insertCallsCnt)
		})
	}
}
