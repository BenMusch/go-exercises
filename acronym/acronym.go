package acronym

import (
	"bytes"
	"regexp"
	"unicode"
)

const testVersion = 3

func GetAbbreviationChar(str string) rune {
	return unicode.ToUpper([]rune(str)[0])
}

func Abbreviate(str string) string {
	var buffer bytes.Buffer
	regex := regexp.MustCompile("\\W")

	for _, word := range regex.Split(str, -1) {
		if len(word) > 0 {
			buffer.WriteRune(GetAbbreviationChar(word))
		}
	}

	return buffer.String()
}
