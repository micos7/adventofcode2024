package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputFilename := "input.txt"
	grid, err := readGrid(inputFilename)
	if err != nil {
		log.Fatalf("Failed to read grid: %v", err)
	}
	if len(grid) == 0 {
		fmt.Println("Input grid is empty.")
		return
	}

	count := solveGuardPath(grid)
	fmt.Printf("\nDistinct positions visited on full map: %d\n", count)
}

func readGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", filename, err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file '%s': %w", filename, err)
	}
	if len(grid) > 0 {
		cols := len(grid[0])
		for r := 1; r < len(grid); r++ {
			if len(grid[r]) != cols {
				log.Printf("Warning: Grid is not rectangular. Row %d length %d, expected %d.", r, len(grid[r]), cols)
			}
		}
	}
	return grid, nil
}

type point struct {
	r, c int
}

func findStart(grid [][]rune) (int, int, string, bool) {
	for r, row := range grid {
		for c, cell := range row {
			switch cell {
			case '^':
				return r, c, "up", true
			case 'v':
				return r, c, "down", true
			case '<':
				return r, c, "left", true
			case '>':
				return r, c, "right", true
			}
		}
	}
	return -1, -1, "", false // Not found
}

func turnRight(dir string) string {
	switch dir {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	default:
		return "" // Should not happen
	}
}

func isOutOfBounds(r, c, rows, cols int) bool {
	return r < 0 || r >= rows || c < 0 || c >= cols
}

func solveGuardPath(grid [][]rune) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])

	startR, startC, currentDir, found := findStart(grid)
	if !found {
		fmt.Println("No start symbol (^, v, <, >) found.")
		return 0
	}

	currentRow, currentCol := startR, startC
	visited := make(map[point]bool)
	visited[point{r: currentRow, c: currentCol}] = true

	for {
		dr, dc := 0, 0
		switch currentDir {
		case "up":
			dr = -1
		case "down":
			dr = 1
		case "left":
			dc = -1
		case "right":
			dc = 1
		}
		nextR, nextC := currentRow+dr, currentCol+dc

		isBlocked := false
		if !isOutOfBounds(nextR, nextC, rows, cols) {
			if grid[nextR][nextC] == '#' {
				isBlocked = true
			}
		}

		if isBlocked {
			currentDir = turnRight(currentDir)
		} else {
			currentRow = nextR
			currentCol = nextC

			if isOutOfBounds(currentRow, currentCol, rows, cols) {
				break
			}
			visited[point{r: currentRow, c: currentCol}] = true
		}
	}
	return len(visited)
}
