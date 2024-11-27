package main

import (
	"github.com/amankumarsingh77/aoc/2023/part1"
	"github.com/amankumarsingh77/aoc/2023/part2"
	"log"
)

func main() {
	lines := ReadLines("input.txt")

	part1Ans := part1.CalculateSum(lines)
	log.Println(part1Ans)

	part2Ans := part2.CalculateSum(lines)
	log.Println(part2Ans)
}
