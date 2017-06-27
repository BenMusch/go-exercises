package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

const questionResponse = "Sure."
const yellResponse = "Whoa, chill out!"
const nothingResponse = "Fine. Be that way!"
const defaultResponse = "Whatever."

func isNothing(message string) bool {
	return len(message) == 0
}

func isQuestion(message string) bool {
	return message[len(message)-1] == '?'
}

func isYelling(message string) bool {
	capsRegex := regexp.MustCompile("[A-Z]+")
	lowerRegex := regexp.MustCompile("[a-z]+")

	return capsRegex.MatchString(message) && !lowerRegex.MatchString(message)
}

func Hey(message string) string {
	message = strings.TrimSpace(message)
	switch {
	case isNothing(message):
		return nothingResponse
	case isYelling(message):
		return yellResponse
	case isQuestion(message):
		return questionResponse
	default:
		return defaultResponse
	}
}
