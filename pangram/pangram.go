package pangram

import (
	"unicode"
)

const testVersion = 1
const unicodeRangeMin = uint16('A')
const unicodeRangeMax = uint16('z')

func IsPangram(str string) bool {
	if len(str) < 26 {
		return false
	}

	charsInString := make(map[rune]bool)
	for _, char := range []rune(str) {
		if unicodeRangeMin <= uint16(char) && unicodeRangeMax >= uint16(char) {
			charsInString[unicode.ToLower(char)] = true
		}
	}

	return len(charsInString) == 26
}
