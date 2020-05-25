package gihealth

import "context"

type Checker interface {
	Check(ctx context.Context) error
}
