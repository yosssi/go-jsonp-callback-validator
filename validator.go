package validator

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^[a-zA-Z_$][0-9a-zA-Z_$]*(?:\[(?:"(?:\\.|[^"\\])*"|\'(?:\\.|[^\'\\])*\'|\d+)\])*?$`)

var reservedKeywords = []string{
	"break",
	"do",
	"instanceof",
	"typeof",
	"case",
	"else",
	"new",
	"var",
	"catch",
	"finally",
	"return",
	"void",
	"continue",
	"for",
	"switch",
	"while",
	"debugger",
	"function",
	"this",
	"with",
	"default",
	"if",
	"throw",
	"delete",
	"in",
	"try",
	"class",
	"enum",
	"extends",
	"super",
	"const",
	"export",
	"import",
	"implements",
	"let",
	"private",
	"public",
	"yield",
	"interface",
	"package",
	"protected",
	"static",
	"null",
	"true",
	"false",
}

// Validate validates the callback function.
func Validate(callback string) bool {
	for _, s := range strings.Split(callback, ".") {
		if !re.MatchString(s) {
			return false
		}

		for _, keyword := range reservedKeywords {
			if s == keyword {
				return false
			}
		}
	}

	return true
}
