package twelve

import (
	"bytes"
	"fmt"
)

const testVersion = 1
const numDays = 12

type Day struct {
	numeric, gift string
}

var days map[int]Day = map[int]Day{
	1:  Day{"first", "a Partridge in a Pear Tree"},
	2:  Day{"second", "two Turtle Doves"},
	3:  Day{"third", "three French Hens"},
	4:  Day{"fourth", "four Calling Birds"},
	5:  Day{"fifth", "five Gold Rings"},
	6:  Day{"sixth", "six Geese-a-Laying"},
	7:  Day{"seventh", "seven Swans-a-Swimming"},
	8:  Day{"eighth", "eight Maids-a-Milking"},
	9:  Day{"ninth", "nine Ladies Dancing"},
	10: Day{"tenth", "ten Lords-a-Leaping"},
	11: Day{"eleventh", "eleven Pipers Piping"},
	12: Day{"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	var buffer bytes.Buffer

	for i := 1; i <= numDays; i++ {
		buffer.WriteString(Verse(i))
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func Verse(verseNum int) string {
	preface, ok := prefaceForDay(verseNum)

	if !ok {
		return ""
	}

	ending := endingForDay(verseNum)
	return fmt.Sprintf("%s, %s.", preface, ending)
}

func prefaceForDay(dayNum int) (string, bool) {
	day, ok := days[dayNum]

	if !ok {
		return "", ok
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me", day.numeric), ok
}

func endingForDay(dayNum int) string {
	day, ok := days[dayNum]

	if !ok {
		return ""
	}

	if dayNum == 1 {
		return day.gift
	} else if dayNum == 2 {
		return fmt.Sprintf("%s, and %s", day.gift, endingForDay(dayNum-1))
	}
	return fmt.Sprintf("%s, %s", day.gift, endingForDay(dayNum-1))
}
