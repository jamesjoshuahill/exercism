package raindrops

import "fmt"

const testVersion = 2

func Convert(number int) string {
	return drop{number}.toString()
}

type drop struct {
	number int
}

func (d drop) toString() string {
	var rain string
	if d.multipleOf(3) {
		rain += "Pling"
	}
	if d.multipleOf(5) {
		rain += "Plang"
	}
	if d.multipleOf(7) {
		rain += "Plong"
	}
	if rain == "" {
		return fmt.Sprintf("%d", d.number)
	}
	return rain
}

func (d drop) multipleOf(x int) bool {
	return d.number%x == 0
}
