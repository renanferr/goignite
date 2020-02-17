package health

import "time"

type CheckResult struct {
	HealthCheck *HealthChecker
	Duration    time.Duration
	Ok          bool
}

func (c *CheckResult) IsOk() bool {
	return c.Ok
}

func NewCheckResult(healthCheck *HealthChecker, duration time.Duration, ok bool) *CheckResult {
	return &CheckResult{
		HealthCheck: healthCheck,
		Duration:    duration,
		Ok:          ok,
	}
}
