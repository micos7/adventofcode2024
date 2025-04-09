package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseValuesPart1(
	index int,
	target int64,
	values []int,
	current int64,
) bool {
	if index == len(values) {
		return current == target
	}

	nextNum := int64(values[index])

	sumResult := current + nextNum

	if parseValuesPart1(index+1, target, values, sumResult) {
		return true
	}

	prodResult := current * nextNum

	if parseValuesPart1(index+1, target, values, prodResult) {
		return true
	}

	return false
}

func parseValuesPart2(
	index int,
	target int64,
	values []int,
	current int64,
) bool {
	if index == len(values) {
		return current == target
	}

	nextNum := int64(values[index])

	sumResult := current + nextNum

	if parseValuesPart2(index+1, target, values, sumResult) {
		return true
	}

	prodResult := current * nextNum

	if parseValuesPart2(index+1, target, values, prodResult) {
		return true
	}

	if current >= 0 {
		strCurrent := strconv.FormatInt(current, 10)
		strNext := strconv.Itoa(values[index])
		concatenatedStr := strCurrent + strNext

		concatenated, err := strconv.ParseInt(concatenatedStr, 10, 64)
		if err == nil {
			if parseValuesPart2(index+1, target, values, concatenated) {
				return true
			}
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSumPart1 := int64(0)
	totalSumPart2 := int64(0)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Printf("Warning: Skipping invalid line format (line %d): %s\n", lineNumber, line)
			continue
		}

		targetStr := strings.TrimSpace(parts[0])
		target, err := strconv.ParseInt(targetStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing target number (line %d): %v\n", lineNumber, err)
			continue
		}

		rawValuesStr := strings.Fields(strings.TrimSpace(parts[1]))
		var parsedValues []int

		validNums := true
		for _, partStr := range rawValuesStr {
			num, err := strconv.Atoi(partStr)
			if err != nil {
				fmt.Printf("Error parsing value number '%s' (line %d): %v\n", partStr, lineNumber, err)
				validNums = false
				break
			}
			parsedValues = append(parsedValues, num)
		}

		if !validNums {
			continue
		}

		if len(parsedValues) == 0 {
			continue
		}

		if parseValuesPart1(1, target, parsedValues, int64(parsedValues[0])) {
			totalSumPart1 += target
		}

		if parseValuesPart2(1, target, parsedValues, int64(parsedValues[0])) {
			totalSumPart2 += target
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total calibration result (Part 1 rules +* only):", totalSumPart1)
	fmt.Println("Total calibration result (Part 2 rules +* concatenation):", totalSumPart2)
}
