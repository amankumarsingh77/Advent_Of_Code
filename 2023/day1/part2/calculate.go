package part2

import (
	"strings"
	"unicode"
)

func CalculateSum(lines []string) int {
	var totsum int
	numbers := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0,
	}

	for _, line := range lines {
		var first, last int

		for i := 0; i < len(line); i++ {
			if digit := getDigit(line, i, numbers); digit != -1 {
				first = digit
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if digit := getDigit(line, i, numbers); digit != -1 {
				last = digit
				break
			}
		}

		sum := first*10 + last
		//log.Println(first, last, sum)
		totsum += sum
	}

	return totsum
}

func getDigit(line string, index int, numbers map[string]int) int {
	char := rune(line[index])
	if unicode.IsDigit(char) {
		return int(char - '0')
	}

	for key, val := range numbers {
		if strings.HasPrefix(line[index:], key) {
			return val
		}
	}

	return -1
}
