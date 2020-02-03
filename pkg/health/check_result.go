package health

import "time"

type checkResult struct {
	healthCheck healthChecker
	duration    time.Duration
	ok          bool
}

func (c *checkResult) isOk() bool {
	return c.ok
}

func NewCheckResult(healthCheck healthChecker, duration time.Duration, ok bool) *checkResult {
	return &checkResult{
		healthCheck: healthCheck,
		duration:    duration,
		ok:          ok,
	}
}
