package health

type HealthChecker struct {
	Name        string
	Description string
	Checker     Checker
	Required    bool
}

func (c *HealthChecker) IsRequired() bool {
	return c.Required
}


func NewHealthChecker(name string, description string, checker Checker, required bool) *HealthChecker {
	return &HealthChecker{
		Name:        name,
		Description: description,
		Checker:     checker,
		Required:    required,
	}
}
