package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Coordinate struct {
	Row int
	Col int
}

type Path []Coordinate

func pathToString(path []Coordinate) string {
	parts := make([]string, len(path))
	for i, coord := range path {
		parts[i] = fmt.Sprintf("%d,%d", coord.Row, coord.Col)
	}
	return strings.Join(parts, "-")
}

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

	// Calculate Part 1: Sum of unique paths from each 0 to ANY 9
	part1Answer := calculateTotalReachableNines(grid)
	fmt.Println("\nPart 1: Total sum of unique paths from each 0 to ANY 9:", part1Answer)

	// Calculate Part 2: Sum of unique paths from each 0 to a SPECIFIC 9
	part2Answer := calculateTotalUniquePathsSpecificNine(grid)
	fmt.Println("Part 2: Total sum of unique paths from each 0 to a SPECIFIC 9:", part2Answer)
}

func calculateTotalReachableNines(grid [][]int) int {
	totalScore := 0
	var zeroPositions []Coordinate

	for r := 0; r < len(grid); r++ {
		if len(grid[r]) == 0 {
			continue
		}
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				zeroPositions = append(zeroPositions, Coordinate{r, c})
			}
		}
	}

	if len(zeroPositions) == 0 {
		fmt.Println("Part 1: No starting points (0) found in the grid.")
		return 0
	}

	trailheadScores := make(map[Coordinate]int)

	startCoordsInOrder := make([]Coordinate, len(zeroPositions))
	copy(startCoordsInOrder, zeroPositions)

	for _, startPos := range startCoordsInOrder {
		reachableNinesFromThisZero := make(map[Coordinate]bool)

		visitedInPath := make(map[Coordinate]bool)

		findReachableNines_dfs(grid, startPos.Row, startPos.Col, -1, visitedInPath, reachableNinesFromThisZero)

		scoreForThisZero := len(reachableNinesFromThisZero)
		trailheadScores[startPos] = scoreForThisZero

		totalScore += scoreForThisZero
	}

	fmt.Println("Part 1: Trailhead scores in reading order:")
	sort.Slice(startCoordsInOrder, func(i, j int) bool {
		if startCoordsInOrder[i].Row != startCoordsInOrder[j].Row {
			return startCoordsInOrder[i].Row < startCoordsInOrder[j].Row
		}
		return startCoordsInOrder[i].Col < startCoordsInOrder[j].Col
	})

	for _, startPos := range startCoordsInOrder {
		fmt.Printf(" (%d,%d): %d\n", startPos.Row, startPos.Col, trailheadScores[startPos])
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
		moves := []struct{ dr, dc int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for _, move := range moves {
			nextR, nextC := r+move.dr, c+move.dc

			findReachableNines_dfs(grid, nextR, nextC, grid[r][c], visitedInPath, reachableNines)
		}
	}
	delete(visitedInPath, currentCoord)
}

func calculateTotalUniquePathsSpecificNine(grid [][]int) int {
	totalUniquePathCount := 0
	var zeroPositions []Coordinate
	var ninePositions []Coordinate

	for r := 0; r < len(grid); r++ {
		if len(grid[r]) == 0 {
			continue
		}
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				zeroPositions = append(zeroPositions, Coordinate{r, c})
			}
			if grid[r][c] == 9 {
				ninePositions = append(ninePositions, Coordinate{r, c})
			}
		}
	}

	if len(zeroPositions) == 0 {
		fmt.Println("Part 2: No starting points (0) found in the grid.")
		return 0
	}
	if len(ninePositions) == 0 {
		fmt.Println("Part 2: No ending points (9) found in the grid.")
		return 0
	}

	for _, startPos := range zeroPositions {
		for _, endPos := range ninePositions {
			uniquePathsBetweenZeroAndNine := make(map[string]bool)

			currentPath := []Coordinate{}
			visitedInPath := make(map[Coordinate]bool)

			findUniquePaths_dfs(grid, startPos.Row, startPos.Col, -1, endPos, &currentPath, visitedInPath, uniquePathsBetweenZeroAndNine)

			countForThisPair := len(uniquePathsBetweenZeroAndNine)
			totalUniquePathCount += countForThisPair
		}
	}

	return totalUniquePathCount
}

func findUniquePaths_dfs(grid [][]int, r, c int, prevValue int, endPos Coordinate, currentPath *[]Coordinate, visitedInPath map[Coordinate]bool, uniquePaths map[string]bool) {

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

	*currentPath = append(*currentPath, currentCoord)
	visitedInPath[currentCoord] = true

	targetReached := false
	if currentCoord == endPos {
		if currentValue == 9 {
			targetReached = true
		} else {
			delete(visitedInPath, currentCoord)
			*currentPath = (*currentPath)[:len(*currentPath)-1]
			return
		}
	}

	if targetReached {
		pathToStore := make([]Coordinate, len(*currentPath))
		copy(pathToStore, *currentPath)

		pathKey := pathToString(pathToStore)
		uniquePaths[pathKey] = true

		delete(visitedInPath, currentCoord)
		*currentPath = (*currentPath)[:len(*currentPath)-1]
		return
	}

	if currentValue == 9 {
		delete(visitedInPath, currentCoord)
		*currentPath = (*currentPath)[:len(*currentPath)-1]
		return
	}

	moves := []struct{ dr, dc int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, move := range moves {
		nextR, nextC := r+move.dr, c+move.dc

		findUniquePaths_dfs(grid, nextR, nextC, grid[r][c], endPos, currentPath, visitedInPath, uniquePaths)
	}

	delete(visitedInPath, currentCoord)
	*currentPath = (*currentPath)[:len(*currentPath)-1]
}
