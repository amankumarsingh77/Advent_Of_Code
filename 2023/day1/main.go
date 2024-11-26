package main

import (
	"github.com/amankumarsingh77/aoc/2023/part1"
	"github.com/amankumarsingh77/aoc/2023/part2"
	"log"
)

var lines []string

func init() {
	lines = ReadLines("input.txt")
}

func main() {
	//part1 -> without considering the word form
	part1Ans := part1.CalculateSum(lines)
	log.Println(part1Ans)
	//part2 -> with considering the word form
	part2Ans := part2.CalculateSum(lines)
	log.Println(part2Ans)
}
