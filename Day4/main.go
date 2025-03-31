package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", filename, err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	firstRowLen := -1

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		row := []rune(line)

		if firstRowLen == -1 {
			firstRowLen = len(row)
		} else if len(row) != firstRowLen {
			log.Printf("Warning: Row has length %d, expected %d. Grid might not be rectangular.", len(row), firstRowLen)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file '%s': %w", filename, err)
	}
	return grid, nil
}

func findWordCount(grid [][]rune, targetWord string) int {
	if len(targetWord) == 0 || len(grid) == 0 {
		return 0
	}
	targetRunes := []rune(targetWord)
	numRows := len(grid)
	wordFoundCount := 0

	dr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dc := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == targetRunes[0] {
				for d := range 8 {
					currentDr, currentDc := dr[d], dc[d]
					foundWordInDirection := true
					for k := 1; k < len(targetRunes); k++ {
						nextR, nextC := r+k*currentDr, c+k*currentDc
						if nextR < 0 || nextR >= numRows || nextC < 0 || nextC >= len(grid[nextR]) || grid[nextR][nextC] != targetRunes[k] {
							foundWordInDirection = false
							break
						}
					}
					if foundWordInDirection {
						wordFoundCount++
					}
				}
			}
		}
	}
	fmt.Printf("Finished searching for word '%s'. Found %d.\n", targetWord, wordFoundCount)
	return wordFoundCount
}

func findPatternCount(grid [][]rune) int {
	patternFoundCount := 0
	rows := len(grid)
	if rows < 3 {
		return 0
	}

	for r := 0; r <= rows-3; r++ {
		if r+2 >= rows {
			break
		}

		cols := len(grid[r])
		if cols < 3 || len(grid[r+1]) < 3 || len(grid[r+2]) < 3 {
			continue
		}

		for c := 0; c <= cols-3; c++ {
			if (grid[r][c] == 'M' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'S') &&
				(grid[r][c+2] == 'M' && grid[r+1][c+1] == 'A' && grid[r+2][c] == 'S') {
				patternFoundCount++
			} else if (grid[r][c] == 'M' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'S') &&
				(grid[r][c+2] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c] == 'M') {
				patternFoundCount++
			} else if (grid[r][c] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'M') &&
				(grid[r][c+2] == 'M' && grid[r+1][c+1] == 'A' && grid[r+2][c] == 'S') {
				patternFoundCount++
			} else if (grid[r][c] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'M') &&
				(grid[r][c+2] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c] == 'M') {
				patternFoundCount++
			}
		}
	}

	fmt.Printf("Finished searching. Found %d X-MAS patterns.\n", patternFoundCount)
	return patternFoundCount
}

func main() {
	inputFilename := "input.txt"
	targetWord := "XMAS"

	grid, err := readGrid(inputFilename)
	if err != nil {
		log.Fatalf("Failed to read grid: %v", err)
	}
	if len(grid) == 0 {
		fmt.Println("Input grid is empty. Exiting.")
		return
	}

	wordCount := findWordCount(grid, targetWord)
	patternCount := findPatternCount(grid) // Calls the MODIFIED function
	fmt.Printf(" Target Word '%s' occurrences: %d\n", targetWord, wordCount)
	fmt.Printf(" Target Pattern (M.M/.A./S.S) occurrences: %d\n", patternCount) // Updated summary label

}
