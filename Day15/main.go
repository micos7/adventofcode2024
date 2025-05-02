package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Use the actual filename for your input
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", fileName, err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	inputStr := strings.TrimSpace(string(content))
	parts := strings.Split(strings.ReplaceAll(inputStr, "\r\n", "\n"), "\n\n")
	if len(parts) != 2 {
		fmt.Println("Error: Input should have grid and directions separated by a blank line.")
		return
	}

	grid := make([][]rune, 0)
	rowsStr := strings.Split(parts[0], "\n")
	if len(rowsStr) == 0 || len(rowsStr[0]) == 0 {
		fmt.Println("Error: Grid part of input is empty.")
		return
	}
	for _, rowStr := range rowsStr {
		if len(rowStr) == 0 {
			continue
		}
		runes := []rune(rowStr)
		grid = append(grid, runes)
	}
	// Keep the hard-coded test grid commented out unless you explicitly need it
	// grid = [][]rune{
	// 	[]rune("##########"),
	// 	[]rune("#.O.O.OOO#"),
	// 	[]rune("#........#"),
	// 	[]rune("#OO......#"),
	// 	[]rune("#OO@.....#"),
	// 	[]rune("#O#.....O#"),
	// 	[]rune("#O.....OO#"),
	// 	[]rune("#O.....OO#"),
	// 	[]rune("#OO....OO#"),
	// 	[]rune("##########"),
	// }

	if len(grid) == 0 {
		fmt.Println("Error: Failed to parse grid.")
		return
	}
	if len(grid[0]) == 0 {
		fmt.Println("Error: Grid has zero columns.")
		return
	}

	directionsStr := strings.TrimSpace(parts[1])
	directions := make([]rune, 0, len(directionsStr))
	for _, r := range directionsStr {
		if r == '<' || r == '>' || r == '^' || r == 'v' {
			directions = append(directions, r)
		}
	}
	if len(directions) == 0 {
		fmt.Println("Warning: No valid directions found.")
	}

	simulate(grid, directions)

	gpsSum := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'O' {
				gpsSum += 100*r + c
			}
		}
	}

	fmt.Println("Sum of GPS coordinates:", gpsSum)
}

func simulate(grid [][]rune, directions []rune) {
	var r, c int
	found := false
	rows := len(grid)
	if rows == 0 {
		fmt.Println("Error in simulate: Grid has no rows.")
		return
	}
	cols := len(grid[0])

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				r, c = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		fmt.Println("Start '@' not found in grid")
		return
	}

	for _, dir := range directions {
		dr, dc := 0, 0
		switch dir {
		case '<':
			dc = -1
		case '>':
			dc = 1
		case '^':
			dr = -1
		case 'v':
			dr = 1
		default:
			continue
		}

		newR, newC := r+dr, c+dc

		// Bounds Check: Player's Target Cell
		if newR < 0 || newR >= rows || newC < 0 || newC >= cols {
			continue
		}

		targetCell := grid[newR][newC]

		switch targetCell {
		case '#':
			continue

		case '.':
			grid[r][c] = '.'
			r, c = newR, newC
			grid[r][c] = '@'

		case 'O':

			stackEndR, stackEndC := newR, newC
			canPush := true
			var emptySpotR, emptySpotC int

			for {
				checkR := stackEndR + dr
				checkC := stackEndC + dc

				if checkR < 0 || checkR >= rows || checkC < 0 || checkC >= cols {
					canPush = false // Cannot push off the grid
					break
				}

				cellAtCheck := grid[checkR][checkC]

				if cellAtCheck == '#' {
					canPush = false
					break
				} else if cellAtCheck == 'O' {
					stackEndR, stackEndC = checkR, checkC
				} else if cellAtCheck == '.' {
					emptySpotR, emptySpotC = checkR, checkC
					break
				} else {
					canPush = false
					break
				}
			}

			if canPush {
				grid[emptySpotR][emptySpotC] = 'O'

				currentBoxR := emptySpotR - dr
				currentBoxC := emptySpotC - dc

				for currentBoxR != newR || currentBoxC != newC {
					grid[currentBoxR][currentBoxC] = 'O'
					currentBoxR -= dr
					currentBoxC -= dc
				}

				grid[newR][newC] = '@'

				grid[r][c] = '.'
				r, c = newR, newC

			} else {
				continue
			}
		}
	}
}

func printGrid(grid [][]rune) {
	if grid == nil {
		fmt.Println("Cannot print nil grid")
		return
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
