package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 3

func Convert(num int) string {
	var buffer bytes.Buffer

	if num%3 == 0 {
		buffer.WriteString("Pling")
	}

	if num%5 == 0 {
		buffer.WriteString("Plang")
	}

	if num%7 == 0 {
		buffer.WriteString("Plong")
	}

	finalString := buffer.String()

	if finalString == "" {
		return strconv.Itoa(num)
	}

	return finalString
}
