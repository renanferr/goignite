package health

import (
	"context"
	"time"
)

var healthChecks []*HealthChecker

func Add(checker *HealthChecker) {
	healthChecks = append(healthChecks, checker)
}

func CheckAll(ctx context.Context) []*CheckResult {

	var results []*CheckResult

	for _, v := range healthChecks {

		start := time.Now()

		err := v.Checker.Check(ctx)

		elapsed := time.Since(start)

		ok := true

		if err != nil {

			ok = false
		}

		results = append(results, NewCheckResult(v, elapsed, ok))
	}

	return results
}