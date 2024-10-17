package stringtool

import (
	"strings"

	"github.com/segmentio/go-camelcase"
	"github.com/segmentio/go-snakecase"
)

func SnakeCase(s string) string {
	return snakecase.Snakecase(s)
}

func CamelCase(s string) string {
	return camelcase.Camelcase(s)
}

func PascalCase(s string) string {
	return UpperFirstLetter(CamelCase(s))
}

func DashCase(s string) string {
	return strings.ReplaceAll(SnakeCase(s), "_", "-")
}
