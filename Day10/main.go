package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Coordinate struct {
	Row int
	Col int
}

type Path []Coordinate

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
	grid := make([][]int, len(lines))

	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}

	answer := 0

	answer += calculateTotalTrailheadScore(grid)

	fmt.Println(answer)
}

func calculateTotalTrailheadScore(grid [][]int) int {
	totalScore := 0
	var zeroPositions []Coordinate // Store coordinates of all cells containing 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				zeroPositions = append(zeroPositions, Coordinate{r, c})
			}
		}
	}

	if len(zeroPositions) == 0 {
		fmt.Println("No starting points (0) found in the grid.")
		return 0
	}

	trailheadScores := make(map[Coordinate]int)
	startCoordsInOrder := []Coordinate{}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				startPos := Coordinate{r, c}
				startCoordsInOrder = append(startCoordsInOrder, startPos)

				reachableNinesFromThisZero := make(map[Coordinate]bool)

				visitedInPath := make(map[Coordinate]bool)

				findReachableNines_dfs(grid, startPos.Row, startPos.Col, -1, visitedInPath, reachableNinesFromThisZero)

				scoreForThisZero := len(reachableNinesFromThisZero)
				trailheadScores[startPos] = scoreForThisZero

				totalScore += scoreForThisZero
			}
		}
	}

	fmt.Println("Trailhead scores in reading order:")
	for _, startPos := range startCoordsInOrder {
		fmt.Printf("  (%d,%d): %d\n", startPos.Row, startPos.Col, trailheadScores[startPos])
	}

	return totalScore
}

func findReachableNines_dfs(grid [][]int, r, c int, prevValue int, visitedInPath map[Coordinate]bool, reachableNines map[Coordinate]bool) {

	rows := len(grid)
	cols := len(grid[0])

	if r < 0 || r >= rows || c < 0 || c >= cols {
		return
	}

	currentValue := grid[r][c]

	if prevValue == -1 {
		if currentValue != 0 {
			return
		}
	} else {
		if currentValue != prevValue+1 {
			return
		}
	}

	currentCoord := Coordinate{r, c}
	if visitedInPath[currentCoord] {
		return
	}

	visitedInPath[currentCoord] = true

	if currentValue == 9 {
		reachableNines[currentCoord] = true

		delete(visitedInPath, currentCoord)
		return
	}

	if currentValue < 9 {
		moves := []struct{ dr, dc int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // Up, Down, Left, Right

		for _, move := range moves {
			nextR, nextC := r+move.dr, c+move.dc

			findReachableNines_dfs(grid, nextR, nextC, grid[r][c], visitedInPath, reachableNines)
		}
	}
	delete(visitedInPath, currentCoord)
}
