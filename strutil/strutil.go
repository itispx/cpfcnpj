package strutil

import "regexp"

var (
	re = regexp.MustCompile("[^0-9]+")
)

func ExtractNumbers(str string) string {
	return re.ReplaceAllString(str, "")
}
