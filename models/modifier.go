package models

import (
	"context"

	"github.com/google/uuid"

	"github.com/fspcons/core/datas"
	"github.com/fspcons/core/errs"
)

type DefaultModifier[R Refresher[R]] struct {
	getter    Getter[R]
	processor UpdateProcessor[R]
	updater   datas.Updater[R]
}

func (ref *DefaultModifier[R]) Modify(ctx context.Context, id uuid.UUID, in UpdateInput[R]) (res R, err error) {
	if err = ValidateID(id); err != nil {
		return res, err
	}
	if err = in.Validate(); err != nil {
		return res, err
	}
	res, err = ref.getter.Get(ctx, id)
	if err != nil {
		return res, err
	}
	new, hasChanges := in.ApplyChanges(res)
	if !hasChanges {
		return res, errs.ErrNoChanges
	}
	if ref.processor != nil {
		if err = ref.processor.Process(ctx, &new, res); err != nil {
			return res, err
		}
	}
	new = new.RefreshTimestamp()
	if err = new.Validate(); err != nil {
		return res, err
	}
	return new, ref.updater.Update(ctx, new)
}

// NewDefaultModifier godoc
func NewDefaultModifier[R Refresher[R]](g Getter[R], p UpdateProcessor[R], u datas.Updater[R]) *DefaultModifier[R] {
	return &DefaultModifier[R]{getter: g, processor: p, updater: u}
}
