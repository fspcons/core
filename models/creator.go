package models

import (
	"context"

	"github.com/fspcons/core/datas"
)

type DefaultCreator[V Validator] struct {
	inserter  datas.Inserter[V]
	processor CreateProcessor[V]
}

func (c *DefaultCreator[V]) Create(ctx context.Context, in CreateInput[V]) (v V, err error) {
	if err = in.Validate(); err != nil {
		return v, err
	}
	v = in.Transform()
	if c.processor != nil {
		if err = c.processor.Process(ctx, &v); err != nil {
			return v, err
		}
	}
	if err = v.Validate(); err != nil {
		return v, err
	}
	return v, c.inserter.Insert(ctx, v)
}

func NewDefaultCreator[V Validator](inserter datas.Inserter[V], processor CreateProcessor[V]) *DefaultCreator[V] {
	return &DefaultCreator[V]{inserter: inserter, processor: processor}
}
