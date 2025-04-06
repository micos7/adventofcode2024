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

	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {

			continue
		}

		parts := strings.Split(line, ":")

		sum, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}

		rawValues := strings.Split(parts[1], " ")
		var parsedValues []int

		for _, part := range rawValues {
			if part == "" {
				continue
			}
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return
			}
			parsedValues = append(parsedValues, num)
		}
		if len(parsedValues) == 0 {
			continue
		}
		if parseValues(1, sum, parsedValues, parsedValues[0]) {
			totalSum += sum
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Sum of middle numbers of correct rows:", totalSum)
}

func parseValues(index int, target int, values []int, current int) bool {
	if index == len(values) {
		return current == target
	}
	return parseValues(index+1, target, values, current+values[index]) ||
		parseValues(index+1, target, values, current*values[index])

}
