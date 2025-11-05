package cpfcnpj

import "regexp"

var (
	re = regexp.MustCompile(`\D+`)
)

func Clean(str string) string {
	return re.ReplaceAllString(str, "")
}
