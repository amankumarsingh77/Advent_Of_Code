package part1

import (
	"unicode"
)

func CalculateSum(lines []string) int {
	var totsum int
	for _, line := range lines {
		first := -1
		var last int
		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == -1 {
					first = int(char - '0')
					last = first
				} else {
					last = int(char - '0')
				}
			}
		}
		sum := first*10 + last
		totsum += sum
	}
	return totsum
}
