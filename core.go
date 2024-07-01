package core

import "go.uber.org/dig"

// MustProvide helper to use dig's Provide
func MustProvide(di *dig.Container, constructor any, opts ...dig.ProvideOption) {
	if err := di.Provide(constructor, opts...); err != nil {
		panic(err)
	}
}
