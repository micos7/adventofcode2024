package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	enabledSum := 0
	totalSum := 0
	mulEnabled := true

	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	for scanner.Scan() {
		line := scanner.Text()
		parts := regexp.MustCompile(`(do\(\)|don't\(\))`).Split(line, -1)
		matches := regexp.MustCompile(`(do\(\)|don't\(\))`).FindAllString(line, -1)

		// Process each part sequentially
		for i, part := range parts {
			mulMatches := mulRegex.FindAllStringSubmatch(part, -1)
			for _, match := range mulMatches {
				first, _ := strconv.Atoi(match[1])
				second, _ := strconv.Atoi(match[2])
				product := first * second
				totalSum += product
				if mulEnabled {
					enabledSum += product
				}
			}

			if i < len(matches) {
				if doRegex.MatchString(matches[i]) {
					mulEnabled = true
				} else if dontRegex.MatchString(matches[i]) {
					mulEnabled = false
				}
			}
		}
	}

	fmt.Println("Part 1 - Total sum (all mul calls):", totalSum)
	fmt.Println("Part 2 - Total sum (enabled mul calls only):", enabledSum)
}
