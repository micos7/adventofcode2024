package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	inputStr := strings.TrimSpace(string(content))
	lines := strings.Split(inputStr, "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// Initialize visited array
	visitedArea := make([][]bool, len(grid))
	visitedPerimeter := make([][]bool, len(grid))
	for i := range visitedArea {
		visitedArea[i] = make([]bool, len(grid[0]))
		visitedPerimeter[i] = make([]bool, len(grid[0]))
	}
	part1 := 0
	for i, row := range grid {
		for j, _ := range row {
			if visitedArea[i][j] {
				continue
			}
			if visitedPerimeter[i][j] {
				continue
			}
			area := area(grid, i, j, visitedArea, grid[i][j])
			perimeter := perimeter(grid, i, j, visitedPerimeter, grid[i][j])
			part1 += area * perimeter
		}
	}

	fmt.Println("Part 1:", part1)

}

func area(grid [][]rune, row int, col int, visited [][]bool, char rune) int {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || visited[row][col] || grid[row][col] != char {
		return 0
	}

	visited[row][col] = true
	return area(grid, row+1, col, visited, char) +
		area(grid, row-1, col, visited, char) +
		area(grid, row, col+1, visited, char) +
		area(grid, row, col-1, visited, char) + 1
}

func perimeter(grid [][]rune, row int, col int, visited [][]bool, char rune) int {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != char {
		return 1
	}
	if visited[row][col] {
		return 0
	}

	visited[row][col] = true
	return perimeter(grid, row+1, col, visited, char) +
		perimeter(grid, row-1, col, visited, char) +
		perimeter(grid, row, col+1, visited, char) +
		perimeter(grid, row, col-1, visited, char)
}
