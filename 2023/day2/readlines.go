package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func ReadLines(path string) []string {
	var lines []string
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := os.Open(filepath.Join(dir, path))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
