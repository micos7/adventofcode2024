package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var leftNumbers, rightNumbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		if len(columns) != 2 {
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		// Convert columns to integers
		left, errL := strconv.Atoi(columns[0])
		right, errR := strconv.Atoi(columns[1])

		if errL != nil || errR != nil {
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		// Append numbers to respective lists
		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Sort both columns
	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	// Calculate the sum of absolute differences
	totalDifference := 0
	for i := range leftNumbers {
		totalDifference += abs(leftNumbers[i] - rightNumbers[i])
	}

	countOccurrences := 0
	for _, leftValue := range leftNumbers {
		countOccurrences += countFrequency(leftValue, rightNumbers)
	}

	// Output results
	fmt.Println("Total Absolute Difference:", totalDifference)
	fmt.Println("Sum of Left-Number Appearances in Right Column:", countOccurrences)
}

// Helper function to compute absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countFrequency(value int, sortedList []int) int {
	count := 0
	for _, num := range sortedList {
		if num == value {
			count++
		}
	}
	return value * count
}
