package part1

import (
	"log"
	"strconv"
	"strings"
)

func CalculateSum(lines []string) int {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var totalSum int
	for _, line := range lines {
		gameID, individualGames := parseLine(line)

		valid := true
		for _, game := range individualGames {
			for color, count := range game {
				if count > limits[color] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			totalSum += gameID
		}
	}

	return totalSum

}

func parseLine(line string) (int, []map[string]int) {
	var individualGames []map[string]int

	parts := strings.Split(line, ":")
	idStr := strings.TrimSpace(strings.Split(parts[0], " ")[1])
	gameID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("Invalid game ID: %s", idStr)
	}

	gameData := strings.Split(parts[1], ";")
	for _, segment := range gameData {
		colorCounts := make(map[string]int)
		individualEntries := strings.Split(segment, ",")
		for _, entry := range individualEntries {
			fields := strings.Fields(strings.TrimSpace(entry))
			if len(fields) < 2 {
				log.Printf("Skipping invalid entry: %s", entry)
				continue
			}

			count, err := strconv.Atoi(fields[0])
			if err != nil {
				log.Printf("Invalid count for entry: %s", fields[0])
				continue
			}

			color := fields[1]
			colorCounts[color] += count
		}
		individualGames = append(individualGames, colorCounts)
	}

	return gameID, individualGames
}
