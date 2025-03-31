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
	var topRows [][]int
	var bottomRows [][]int
	readingTop := true

	// Parse input file
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingTop = false
			continue
		}

		parts := strings.Split(line, "|")
		if !readingTop {
			parts = strings.Split(line, ",")
		}

		var row []int
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return
			}
			row = append(row, num)
		}

		if readingTop {
			topRows = append(topRows, row)
		} else {
			bottomRows = append(bottomRows, row)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	correctSum := 0
	incorrectSum := 0

	for _, bottomRow := range bottomRows {
		correct := true

		// Check if the row violates any ordering constraints
		for _, topRow := range topRows {
			if len(topRow) != 2 {
				continue
			}
			first := topRow[0]
			second := topRow[1]
			firstIndex := -1
			secondIndex := -1

			for i, num := range bottomRow {
				if num == first {
					firstIndex = i
				}
				if num == second {
					secondIndex = i
				}
			}

			if firstIndex != -1 && secondIndex != -1 && firstIndex > secondIndex {
				correct = false
				break
			}
		}

		middleNumber := 0

		if correct {
			if len(bottomRow)%2 == 1 {
				middleNumber = bottomRow[len(bottomRow)/2]
			} else if len(bottomRow) > 0 {
				middleNumber = (bottomRow[len(bottomRow)/2-1] + bottomRow[len(bottomRow)/2]) / 2
			}
			correctSum += middleNumber
		} else {
			orderedRow := customSort(bottomRow, topRows)

			if len(orderedRow)%2 == 1 {
				middleNumber = orderedRow[len(orderedRow)/2]
			} else if len(orderedRow) > 0 {
				middleNumber = (orderedRow[len(orderedRow)/2-1] + orderedRow[len(orderedRow)/2]) / 2
			}
			incorrectSum += middleNumber
		}
	}

	fmt.Println("Sum of middle numbers of correct rows:", correctSum)
	fmt.Println("Sum of middle numbers of incorrect rows (ordered):", incorrectSum)
}

func customSort(row []int, constraints [][]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, num := range row {
		graph[num] = []int{}
		inDegree[num] = 0
	}

	for _, constraint := range constraints {
		if len(constraint) != 2 {
			continue
		}

		first := constraint[0]
		second := constraint[1]

		// Check if both elements exist in the row
		firstExists := false
		secondExists := false
		for _, num := range row {
			if num == first {
				firstExists = true
			}
			if num == second {
				secondExists = true
			}
		}

		if firstExists && secondExists {
			graph[first] = append(graph[first], second)
			inDegree[second]++
		}
	}

	var result []int
	var queue []int

	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--

			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) < len(row) {
		for _, num := range row {
			found := false
			for _, res := range result {
				if res == num {
					found = true
					break
				}
			}
			if !found {
				result = append(result, num)
			}
		}
	}

	return result
}
