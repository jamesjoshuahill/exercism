package twelve

import (
	"fmt"
	"strings"
)

const testVersion = 1

func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += fmt.Sprintf("%s\n", Verse(i))
	}
	return song
}

func Verse(number int) string {
	return onDay(number) + gaveToMe(number)
}

var ordinals = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func onDay(day int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, ", ordinals[day])
}

var giftList = []string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func gaveToMe(day int) string {
	if day == 1 {
		return giftList[0]
	}
	var gave []string
	for i := day - 1; i > 0; i-- {
		gave = append(gave, giftList[i])
	}
	gave = append(gave, "and "+giftList[0])

	return strings.Join(gave, ", ")
}
