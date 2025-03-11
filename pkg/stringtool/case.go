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

func UpperSnakeCase(s string) string {
	return strings.ToUpper(SnakeCase(SpaceCase(s)))
}

func UpperDashCase(s string) string {
	return strings.ToUpper(DashCase(SpaceCase(s)))
}

/*
SpaceCase converts a string to a space-separated string.
e.g FooBar -> foo bar
*/
func SpaceCase(s string) string {
	str := ""
	for idx, r := range s {
		if r >= 'A' && r <= 'Z' && idx > 0 && s[idx-1] != ' ' {
			str += " "
		}
		str += string(r)
	}

	charToSpace := []string{"-", "_"}
	for _, c := range charToSpace {
		str = strings.ReplaceAll(str, c, " ")
	}

	return strings.ToLower(str)
}
