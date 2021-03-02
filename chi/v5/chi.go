package gichi

import (
	"context"

	"github.com/go-chi/chi/v5"
)

var (
	instance *chi.Mux
)

type ext func(context.Context, *chi.Mux) error

func NewMux(ctx context.Context, exts ...ext) *chi.Mux {

	instance = chi.NewRouter()

	for _, ext := range exts {
		if err := ext(ctx, instance); err != nil {
			panic(err)
		}
	}

	return instance
}
