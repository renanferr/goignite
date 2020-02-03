package health

var checks []Checker

func Add(checker Checker) {
	checks = append(checks, checker)
}

func CheckAll() []checkResult {

	var results []checkResult

	return results
}