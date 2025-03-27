package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	legitReports := 0
	legitReportsDampner := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		numbers := []int{}
		for _, v := range parts {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Skipping invalid input:", v)
				continue
			}
			numbers = append(numbers, num)
		}

		if len(numbers) < 2 {
			continue
		}

		if isValidSequence(numbers) {
			legitReports++
		} else if checkSequenceWithOneRemoval(numbers) {
			legitReportsDampner++
		}
	}

	// Part 1: legitReports
	fmt.Println("Part 1 - Total safe reports:", legitReports)
	// Part 2: legitReports + legitReportsDampner
	fmt.Println("Part 2 - Total safe reports with dampener:", legitReports+legitReportsDampner)
}

func checkSequenceWithOneRemoval(numbers []int) bool {
	for i := range numbers {
		modified := make([]int, len(numbers)-1)
		copy(modified, numbers[:i])
		copy(modified[i:], numbers[i+1:])

		if len(modified) < 2 {
			continue
		}

		if isValidSequence(modified) {
			return true
		}
	}
	return false
}

func isValidSequence(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		absDiff := abs(diff)
		if absDiff < 1 || absDiff > 3 {
			return false
		}
		if diff <= 0 {
			increasing = false
		}
		if diff >= 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
