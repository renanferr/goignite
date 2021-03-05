package giserver

import (
	"context"
	"sync"
)

func Serve(ctx context.Context, srvs ...func(context.Context)) {

	switch len(srvs) {
	case 0:
		panic("no servers configured")
	case 1:
		srvs[0](ctx)
	default:
		wg := new(sync.WaitGroup)
		wg.Add(len(srvs))

		for _, srv := range srvs {
			srv := srv
			go func() {
				srv(ctx)
				wg.Done()
			}()
		}

		wg.Wait()
	}

}
