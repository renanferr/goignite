package health

type healthChecker struct {
	name        string
	description string
	checker     Checker
	required    bool
}

func NewHealthChecker(name string, description string, checker Checker, required bool) *healthChecker {
	return &healthChecker{
		name:        name,
		description: description,
		checker:     checker,
		required:    required,
	}
}
