package utils

import (
	"regexp"
)

// SplitCamelCase places token before each uppercase letter except the first.
func SplitCamelCase(str string, token string) string {
	return regexp.MustCompile(`\B([A-Z]+)`).ReplaceAllString(str, token+`$1`)
}
