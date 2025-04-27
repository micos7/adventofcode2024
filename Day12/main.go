package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Open the input file.
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Ensure the file is closed when the function exits.
	defer file.Close()

	// Read all content from the file.
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Trim leading/trailing whitespace and split the content into lines.
	inputStr := strings.TrimSpace(string(content))
	lines := strings.Split(inputStr, "\n")

	// Create the grid from the lines of the file.
	grid := make([][]rune, len(lines))
	for i := range lines {
		grid[i] = []rune(lines[i])
	}

	// Get grid dimensions.
	rows := len(grid)
	if rows == 0 {
		fmt.Println("Grid is empty.")
		return
	}
	cols := len(grid[0])
	if cols == 0 {
		fmt.Println("Grid rows are empty.")
		return
	}

	// Initialize a visited array for tracking islands during the main loop traversal.
	visitedMain := make([][]bool, rows)
	for i := range visitedMain {
		visitedMain[i] = make([]bool, cols)
	}

	// Variables to accumulate the total prices for both parts.
	totalPricePart1 := 0 // Area * Flawed Perimeter (standard perimeter length)
	totalPricePart2 := 0 // Area * Corrected Perimeter (Distinct Sides)

	// Iterate through each cell in the grid.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if visitedMain[i][j] {
				continue
			}

			currentChar := grid[i][j]

			// --- Identify all cells in the current region ---
			visitedRegion := make([][]bool, rows)
			for k := range visitedRegion {
				visitedRegion[k] = make([]bool, cols)
			}
			findRegionCellsDFS(grid, i, j, visitedRegion, currentChar, rows, cols)

			// --- Calculate Area ---
			currentArea := 0
			for r := 0; r < rows; r++ {
				for c := 0; c < cols; c++ {
					if visitedRegion[r][c] {
						currentArea++
					}
				}
			}

			// Skip calculation if region is empty (should not happen with valid input)
			if currentArea == 0 {
				continue
			}

			// --- Calculate Flawed Perimeter (for Part 1) ---
			visitedPerimeterFlawed := make([][]bool, rows)
			for k := range visitedPerimeterFlawed {
				visitedPerimeterFlawed[k] = make([]bool, cols)
			}
			// We need to start perimeter DFS from a valid cell of the region
			currentPerimeterFlawed := perimeter(grid, i, j, visitedPerimeterFlawed, currentChar)

			// --- Calculate Corrected Number of Sides (for Part 2) ---
			currentDistinctSides := countDistinctSidesCorrected_CellBased(visitedRegion, rows, cols) // Pass only needed info

			// --- Update VisitedMain ---
			for r := 0; r < rows; r++ {
				for c := 0; c < cols; c++ {
					if visitedRegion[r][c] {
						visitedMain[r][c] = true
					}
				}
			}

			// --- Accumulate Totals ---
			totalPricePart1 += currentArea * currentPerimeterFlawed
			totalPricePart2 += currentArea * currentDistinctSides

			// Optional: Print details for each region found.
			// fmt.Printf("Found island of char '%c' starting at (%d, %d): Area = %d, Flawed Perimeter = %d, Corrected Sides = %d, Price Part 1 = %d, Price Part 2 = %d\n",
			// 	currentChar, i, j, currentArea, currentPerimeterFlawed, currentDistinctSides, currentArea * currentPerimeterFlawed, currentArea * currentDistinctSides)

		}
	}

	// Print the final total prices for both parts.
	fmt.Println("Total price of fencing all regions (Part 1 - Area * Flawed Perimeter):", totalPricePart1)
	fmt.Println("Total price of fencing all regions (Part 2 - Area * Number of Sides):", totalPricePart2)

}

// area calculates the area of a connected component using DFS.
// It counts the number of cells with the same character that are connected.
// It uses and updates the provided visited array.
func area(grid [][]rune, row int, col int, visited [][]bool, char rune) int {
	rows := len(grid)
	cols := len(grid[0])
	// Base cases:
	// - Out of bounds
	// - Already visited
	// - Character does not match the target
	if row < 0 || row >= rows || col < 0 || col >= cols || visited[row][col] || grid[row][col] != char {
		return 0 // Stop recursion and contribute 0 to the area
	}

	// Mark the current cell as visited.
	visited[row][col] = true

	// Recursively call for neighbors and add 1 for the current cell.
	return area(grid, row+1, col, visited, char) +
		area(grid, row-1, col, visited, char) +
		area(grid, row, col+1, visited, char) +
		area(grid, row, col-1, visited, char) + 1
}

// perimeter calculates the standard perimeter length using DFS.
// Counts edges adjacent to different characters or boundaries.
// Included to replicate Part 1 calculation.
func perimeter(grid [][]rune, row int, col int, visited [][]bool, char rune) int {
	rows := len(grid)
	cols := len(grid[0])

	// If out of bounds or different character, it's considered an edge boundary.
	if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != char {
		return 1 // Count this as contributing 1 unit to the perimeter
	}
	// If already visited within this perimeter calculation, don't count again and stop recursion.
	if visited[row][col] {
		return 0
	}

	// Mark as visited for this perimeter calculation.
	visited[row][col] = true

	// Recursively call for neighbors. The sum of boundary encounters (return 1) is the perimeter.
	return perimeter(grid, row+1, col, visited, char) +
		perimeter(grid, row-1, col, visited, char) +
		perimeter(grid, row, col+1, visited, char) +
		perimeter(grid, row, col-1, visited, char)
}

// --- Functions for Corrected Side Counting (Part 2) ---

// Helper function to check if a cell coordinates (r, c) is inside the *current* region.
// Uses the 'regionVisited' map which must be pre-populated for the current region.
// Handles boundary checks.
func isInRegion(r, c, rows, cols int, regionVisited [][]bool) bool {
	if r < 0 || r >= rows || c < 0 || c >= cols {
		return false // Out of bounds is not in the region
	}
	return regionVisited[r][c] // Check the visited map for this specific region
}

// Helper DFS function to find all cells of the current region and mark them in 'visited'.
// This is used to identify the full extent of the region before counting sides.
func findRegionCellsDFS(grid [][]rune, r, c int, visited [][]bool, char rune, rows, cols int) {
	// Base cases for DFS recursion
	if r < 0 || r >= rows || c < 0 || c >= cols || visited[r][c] || grid[r][c] != char {
		return
	}
	// Mark current cell as visited (part of the region)
	visited[r][c] = true
	// Recurse to neighbors
	findRegionCellsDFS(grid, r+1, c, visited, char, rows, cols)
	findRegionCellsDFS(grid, r-1, c, visited, char, rows, cols)
	findRegionCellsDFS(grid, r, c+1, visited, char, rows, cols)
	findRegionCellsDFS(grid, r, c-1, visited, char, rows, cols)
}

// countDistinctSidesCorrected calculates the number of "sides" (straight fence sections)
// for the region identified by the pre-populated 'regionVisited' map.
// countDistinctSidesCorrected - Iterates through cells *in* the region
// and checks 4 boundary start conditions.
func countDistinctSidesCorrected_CellBased(regionVisited [][]bool, rows, cols int) int {
	sides := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !regionVisited[r][c] {
				continue
			} // Only process cells within the current region

			// --- Check Top Boundary Start ---
			if !isInRegion(r-1, c, rows, cols, regionVisited) && // Boundary Above
				(!isInRegion(r, c-1, rows, cols, regionVisited) || isInRegion(r-1, c-1, rows, cols, regionVisited)) { // Doesn't continue left
				sides++
			}
			// --- Check Left Boundary Start ---
			if !isInRegion(r, c-1, rows, cols, regionVisited) && // Boundary Left
				(!isInRegion(r-1, c, rows, cols, regionVisited) || isInRegion(r-1, c-1, rows, cols, regionVisited)) { // Doesn't continue up
				sides++
			}
			// --- Check Bottom Boundary Start ---
			if !isInRegion(r+1, c, rows, cols, regionVisited) && // Boundary Below
				(!isInRegion(r, c-1, rows, cols, regionVisited) || isInRegion(r+1, c-1, rows, cols, regionVisited)) { // Doesn't continue left
				sides++
			}
			// --- Check Right Boundary Start ---
			if !isInRegion(r, c+1, rows, cols, regionVisited) && // Boundary Right
				(!isInRegion(r-1, c, rows, cols, regionVisited) || isInRegion(r-1, c+1, rows, cols, regionVisited)) { // Doesn't continue up
				sides++
			}
		}
	}
	// This counted 4 for 1x1, 4 for 2x2, 8 for C. THIS is the logic that seemed correct.
	// If THIS produced "too high", then the dry runs must be wrong or something else is happening.
	// Let's assume this IS the correct logic and there's no overcounting inherent to it.
	return sides
}
