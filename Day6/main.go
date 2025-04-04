package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Point represents a coordinate on the grid
type point struct {
	r, c int
}

// State represents the guard's position and direction
type state struct {
	r, c int
	dir  string
}

func main() {
	inputFilename := "input.txt" // Make sure you have an input.txt file
	originalGrid, err := readGrid(inputFilename)
	if err != nil {
		log.Fatalf("Failed to read grid: %v", err)
	}
	if len(originalGrid) == 0 {
		fmt.Println("Input grid is empty or could not be read.")
		return
	}
	// Basic check for rows, column check happens where grid is used
	if len(originalGrid[0]) == 0 {
		fmt.Println("Input grid has zero columns.")
		return
	}

	fmt.Println("Original Grid:")
	printGrid(originalGrid)

	loopCausingCount := 0
	rows := len(originalGrid)
	// Assume rectangular grid based on original code's intent and first row
	// If grid can be jagged, logic inside simulation needs careful bounds checking per row
	cols := len(originalGrid[0])

	fmt.Println("\nTesting potential obstacle placements...")

	// Create a mutable copy for modifications
	grid := make([][]rune, rows)
	for r := range originalGrid {
		grid[r] = make([]rune, cols) // Assuming rectangular based on previous lines
		copy(grid[r], originalGrid[r])
	}

	startR, startC, _, startFound := findStart(grid)
	if !startFound {
		fmt.Println("No start symbol found in the grid.")
		return
	}
	visitedCells := 0
	for r := 0; r < rows; r++ {
		// Ensure column iteration respects actual row length if potentially jagged
		// but for this loop, we use the 'cols' derived earlier based on assumption
		for c := 0; c < cols; c++ {
			// Check bounds just in case grid WAS jagged despite earlier assumption
			if c >= len(grid[r]) {
				continue // Skip if this column doesn't exist for this row
			}

			originalChar := grid[r][c]

			// Check if this is a valid spot for a new obstacle:
			// Not the start position and not already an obstacle.
			isStartPos := (r == startR && c == startC)
			// Allow placing obstacles on empty paths (e.g., '.')
			// Add other non-obstacle characters here if needed (e.g., ' ')
			isValidPlacement := !isStartPos && originalChar != '#' && originalChar == '.'

			if isValidPlacement {
				// --- Modify ---
				grid[r][c] = '#'
				// fmt.Printf("Testing obstacle at (%d, %d)...\n", r, c) // Uncomment for verbose output

				// --- Simulate ---
				visitedPositions, loopDetected := simulateGuardPathWithLoopDetection(grid) // Pass the modified grid

				// --- Check Result ---
				if loopDetected {
					loopCausingCount++
					fmt.Printf("  -> Loop DETECTED by placing '#' at (%d, %d). Path length before/during loop: %d\n", r, c, len(visitedPositions))
				} else {
					// fmt.Printf("  -> No loop for obstacle at (%d, %d). Path length: %d\n", r, c, len(visitedPositions)) // Uncomment for verbose output
				}
				visitedCells = len(visitedPositions)

				// --- Revert ---
				grid[r][c] = originalChar // Put back the original character
			}
		}
	}

	fmt.Printf("\nTotal number of single obstacle placements causing a loop: %d\n", loopCausingCount)
	fmt.Printf("\nTotal number of cells: %d\n", visitedCells)

}

func readGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", filename, err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		grid = append(grid, runes)
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file '%s': %w", filename, err)
	}
	return grid, nil
}

func printGrid(grid [][]rune) {
	if len(grid) == 0 {
		fmt.Println("[]")
		return
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func findStart(grid [][]rune) (int, int, string, bool) {
	for r, row := range grid {

		if r >= len(grid) {
			continue
		}
		for c, cell := range row {

			if c >= len(grid[r]) {
				continue
			}
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
	return -1, -1, "", false
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

		return dir
	}
}

func isOutOfBounds(r, c, rows, cols int) bool {
	return r < 0 || r >= rows || c < 0 || c >= cols
}

func simulateGuardPathWithLoopDetection(grid [][]rune) (map[point]bool, bool) {
	if len(grid) == 0 {
		return make(map[point]bool), false
	}
	rows := len(grid)
	if rows > 0 && len(grid[0]) == 0 {
		fmt.Println("Warning: Grid has rows but zero columns.")
		return make(map[point]bool), false // Treat as unable to simulate
	}

	cols := 0
	if rows > 0 {
		cols = len(grid[0])
	}
	if cols == 0 && rows == 0 {
		return make(map[point]bool), false
	}

	startR, startC, currentDir, found := findStart(grid)
	if !found {
		fmt.Println("Warning: No start symbol (^, v, <, >) found during simulation.")
		return make(map[point]bool), false
	}

	currentRow, currentCol := startR, startC
	visitedPositions := make(map[point]bool)
	visitedStates := make(map[state]bool)

	maxSteps := rows * cols * 4

	for step := 0; step < maxSteps; step++ {

		if currentRow < 0 || currentRow >= rows || currentCol < 0 || currentCol >= len(grid[currentRow]) {
			fmt.Printf("Warning: Guard out of bounds at step %d (%d, %d). Terminating.\n", step, currentRow, currentCol)
			return visitedPositions, false
		}

		currentState := state{r: currentRow, c: currentCol, dir: currentDir}
		if visitedStates[currentState] {
			visitedPositions[point{r: currentRow, c: currentCol}] = true
		}
		visitedStates[currentState] = true
		visitedPositions[point{r: currentRow, c: currentCol}] = true

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

		if isOutOfBounds(nextR, nextC, rows, cols) {
			return visitedPositions, false // EXIT NORMALLY
		}

		if grid[nextR][nextC] == '#' {
			currentDir = turnRight(currentDir)
		} else {
			currentRow = nextR
			currentCol = nextC
		}
	}

	fmt.Printf("Warning: Max steps (%d) reached. Assuming loop.\n", maxSteps)
	return visitedPositions, true
}
