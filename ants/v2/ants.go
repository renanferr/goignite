package giants

import (
	"context"
	"sync"

	"github.com/panjf2000/ants/v2"
)

type Middleware interface {
	Before(ctx context.Context) context.Context
	After(ctx context.Context)
}

type Task func(ctx context.Context) context.Context

type Wrapper struct {
	pool        *ants.Pool
	middlewares []Middleware
}

func NewWithPool(pool *ants.Pool, middlewares ...Middleware) *Wrapper {
	return &Wrapper{pool: pool, middlewares: middlewares}
}

func (a *Wrapper) Submit(ctx context.Context, wg *sync.WaitGroup, task Task) error {

	wg.Add(1)

	err := ants.Submit(func() {

		for _, m := range a.middlewares {
			ctx = m.Before(ctx)
		}

		ctx = task(ctx)

		for _, m := range a.middlewares {
			m.After(ctx)
		}

		wg.Done()

	})

	return err
}

type WrapperWithFunc struct {
	pool        *ants.PoolWithFunc
	middlewares []Middleware
}

func WithPoolWithFunc(pool *ants.PoolWithFunc, middlewares ...Middleware) *WrapperWithFunc {
	return &WrapperWithFunc{pool: pool, middlewares: middlewares}
}
