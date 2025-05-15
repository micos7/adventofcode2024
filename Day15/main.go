package main

import (
	"fmt"
	"io"
	"os"
	"sort"
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

	// --- Grid Parsing ---
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

	// --- Directions Parsing ---
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

	grid_cp := deepCopyGrid(grid)

	// --- Simulate the movement ---
	Part1(grid, directions)

	// test := [][]rune{
	// 	[]rune("#######"),
	// 	[]rune("#...#.#"),
	// 	[]rune("#.....#"),
	// 	[]rune("#..OO@#"),
	// 	[]rune("#..O..#"),
	// 	[]rune("#.....#"),
	// 	[]rune("#######"),
	// }
	//
	// g := [][]rune{
	// 	[]rune("##########"),
	// 	[]rune("#..O..O.O#"),
	// 	[]rune("#......O.#"),
	// 	[]rune("#.OO..O.O#"),
	// 	[]rune("#..O@..O.#"),
	// 	[]rune("#O#..O...#"),
	// 	[]rune("#O..O..O.#"),
	// 	[]rune("#.OO.O.OO#"),
	// 	[]rune("#....O...#"),
	// 	[]rune("##########"),
	// }

	expandedGrid := expandGrid(grid_cp)

	// printGrid(expandedGrid)
	// return

	// dirs := []rune("<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^ vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v ><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv< <<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^ ^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^>< ^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^ >^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^ <><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<> ^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v> v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^")
	// g2 := [][]rune{
	// 	[]rune("####################"),
	// 	[]rune("##....[]....[]..[]##"),
	// 	[]rune("##............[]..##"),
	// 	[]rune("##..[][]....[]..[]##"),
	// 	[]rune("##....[]@.....[]..##"),
	// 	[]rune("##[]##....[]......##"),
	// 	[]rune("##[]....[]....[]..##"),
	// 	[]rune("##..[][]..[]..[][]##"),
	// 	[]rune("##........[]......##"),
	// 	[]rune("####################"),
	// }

	gridPart2 := simulatePart2(expandedGrid, directions)

	// printGrid(expandedGrid)
	// return

	// --- Calculate the sum of GPS coordinates of boxes ---
	// Using 0-based indexing for rows and columns
	gpsSum := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'O' {
				gpsSum += 100*r + c
			}
		}
	}
	// gpsSumPart2Left := 0
	// gpsSumPart2Right := 0
	// for r := 0; r < len(gridPart2); r++ {
	// 	for c := 0; c < len(gridPart2[r]); c++ {
	// 		if gridPart2[r][c] == '[' {
	// 			gpsSumPart2Left++
	// 		} else if gridPart2[r][c] == ']' {
	// 			gpsSumPart2Right++
	// 		}
	// 	}
	// }
	// fmt.Println("Sum of GPS coordinates part 2, Left brackets:", gpsSumPart2Left)
	// fmt.Println("Sum of GPS coordinates part 2, Right brackets:", gpsSumPart2Right)

	gpsLeft := 0
	gpsRight := 0
	for r := 0; r < len(expandedGrid); r++ {
		for c := 0; c < len(expandedGrid[r]); c++ {
			if expandedGrid[r][c] == '[' {
				gpsLeft++
			} else if expandedGrid[r][c] == ']' {
				gpsRight++
			}
		}
	}
	fmt.Println("Sum of GPS coordinates EXPANDED part 2, Left brackets:", gpsLeft)
	fmt.Println("Sum of GPS coordinates EXPANDED part 2, Right brackets:", gpsRight)

	gpsSumPart2 := 0
	for r := 0; r < len(gridPart2); r++ {
		for c := 0; c < len(gridPart2[r]); c++ {
			if gridPart2[r][c] == '[' {
				score := 100*r + c
				fmt.Printf("Found [ at (%d, %d)]: +%d\n", r, c, score)
				gpsSumPart2 += score
			}
		}
	}

	// --- Print the final result ---
	fmt.Println("Sum of GPS coordinates:", gpsSum)
	fmt.Println("Sum of GPS coordinates part 2:", gpsSumPart2)
}

func Part1(grid [][]rune, directions []rune) {
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
					canPush = false
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

func simulatePart2(gr [][]rune, directions []rune) [][]rune {
	copiedGrid := make([][]rune, len(gr))
	for i := range gr {
		copiedGrid[i] = make([]rune, len(gr[i]))
		copy(copiedGrid[i], gr[i])
	}
	currentGrid := copiedGrid

	var playerR, playerC int
	foundPlayer := false
	for rIdx := range currentGrid {
		for cIdx := range currentGrid[rIdx] {
			if currentGrid[rIdx][cIdx] == '@' {
				playerR, playerC = rIdx, cIdx
				foundPlayer = true
				break
			}
		}
		if foundPlayer {
			break
		}
	}

	if !foundPlayer {
		return currentGrid
	}

	// Process each movement direction
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

		potentialPlayerR, potentialPlayerC := playerR+dr, playerC+dc

		if potentialPlayerR < 0 || potentialPlayerR >= len(currentGrid) ||
			potentialPlayerC < 0 || potentialPlayerC >= len(currentGrid[0]) {
			continue
		}

		targetCellContent := currentGrid[potentialPlayerR][potentialPlayerC]

		if targetCellContent == '.' {
			currentGrid[playerR][playerC] = '.'
			playerR, playerC = potentialPlayerR, potentialPlayerC
			currentGrid[playerR][playerC] = '@'
		} else if targetCellContent == '[' || targetCellContent == ']' {
			boxContactR, boxContactC := potentialPlayerR, potentialPlayerC
			if targetCellContent == ']' {
				if potentialPlayerC-1 >= 0 && currentGrid[potentialPlayerR][potentialPlayerC-1] == '[' {
					boxContactC = potentialPlayerC - 1
				} else {
					continue
				}
			}

			boxes := collectBoxChainGroupPush(currentGrid, boxContactR, boxContactC, dr, dc)

			if len(boxes) > 0 && canPushChain(currentGrid, boxes, dr, dc) {
				boxesToClear := make([]BoxPosition, len(boxes))
				copy(boxesToClear, boxes)
				if dr > 0 {
					sort.Slice(boxesToClear, func(i, j int) bool { return boxesToClear[i].R > boxesToClear[j].R })
				} else if dr < 0 {
					sort.Slice(boxesToClear, func(i, j int) bool { return boxesToClear[i].R < boxesToClear[j].R })
				} else if dc > 0 {
					sort.Slice(boxesToClear, func(i, j int) bool { return boxesToClear[i].LeftC > boxesToClear[j].LeftC })
				} else if dc < 0 {
					sort.Slice(boxesToClear, func(i, j int) bool { return boxesToClear[i].LeftC < boxesToClear[j].LeftC })
				}
				for _, b := range boxesToClear {
					currentGrid[b.R][b.LeftC] = '.'
					currentGrid[b.R][b.LeftC+1] = '.'
				}

				currentGrid[playerR][playerC] = '.'
				playerR, playerC = potentialPlayerR, potentialPlayerC

				boxesToDraw := make([]BoxPosition, len(boxes))
				copy(boxesToDraw, boxes)
				if dr > 0 {
					sort.Slice(boxesToDraw, func(i, j int) bool { return boxesToDraw[i].R < boxesToDraw[j].R })
				} else if dr < 0 {
					sort.Slice(boxesToDraw, func(i, j int) bool { return boxesToDraw[i].R > boxesToDraw[j].R })
				} else if dc > 0 {
					sort.Slice(boxesToDraw, func(i, j int) bool { return boxesToDraw[i].LeftC < boxesToDraw[j].LeftC })
				} else if dc < 0 {
					sort.Slice(boxesToDraw, func(i, j int) bool { return boxesToDraw[i].LeftC > boxesToDraw[j].LeftC })
				}
				for _, b := range boxesToDraw {
					movedBoxR, movedBoxC := b.R+dr, b.LeftC+dc
					currentGrid[movedBoxR][movedBoxC] = '['
					currentGrid[movedBoxR][movedBoxC+1] = ']'
				}

				currentGrid[playerR][playerC] = '@'
			}
		}
		fmt.Println("Movement is ", string(dir))
		printGrid(currentGrid)
	}
	return currentGrid
}

type BoxPosition struct {
	R     int
	LeftC int
}

func collectBoxChainGroupPush(gr [][]rune, initialContactR, initialContactC, dr, dc int) []BoxPosition {
	var finalChainAsSlice []BoxPosition

	queue := []BoxPosition{} // Queue for BFS
	visitedAndInChain := make(map[string]BoxPosition)

	firstHitBox, isBox := getBoxFromGrid(gr, initialContactR, initialContactC)
	if !isBox {
		return finalChainAsSlice
	}

	firstHitBoxKey := fmt.Sprintf("%d,%d", firstHitBox.R, firstHitBox.LeftC)
	queue = append(queue, firstHitBox)
	visitedAndInChain[firstHitBoxKey] = firstHitBox

	head := 0
	for head < len(queue) {
		currentBox := queue[head]
		head++

		targetRowForCurrentBox := currentBox.R + dr

		targetColForLeftPartOfCurrentBox := currentBox.LeftC + dc

		targetColForRightPartOfCurrentBox := currentBox.LeftC + 1 + dc

		if collidingBox1, isBox1 := getBoxFromGrid(gr, targetRowForCurrentBox, targetColForLeftPartOfCurrentBox); isBox1 {
			collidingBox1Key := fmt.Sprintf("%d,%d", collidingBox1.R, collidingBox1.LeftC)
			if _, exists := visitedAndInChain[collidingBox1Key]; !exists {
				visitedAndInChain[collidingBox1Key] = collidingBox1
				queue = append(queue, collidingBox1)
			}
		}

		if collidingBox2, isBox2 := getBoxFromGrid(gr, targetRowForCurrentBox, targetColForRightPartOfCurrentBox); isBox2 {
			collidingBox2Key := fmt.Sprintf("%d,%d", collidingBox2.R, collidingBox2.LeftC)
			if _, exists := visitedAndInChain[collidingBox2Key]; !exists {
				visitedAndInChain[collidingBox2Key] = collidingBox2
				queue = append(queue, collidingBox2)
			}
		}
	}

	for _, box := range visitedAndInChain {
		finalChainAsSlice = append(finalChainAsSlice, box)
	}

	return finalChainAsSlice
}

func getBoxFromGrid(gr [][]rune, r, c int) (BoxPosition, bool) {
	if r < 0 || r >= len(gr) || c < 0 || c >= len(gr[0]) {
		return BoxPosition{}, false
	}
	if gr[r][c] == '[' && c+1 < len(gr[0]) && gr[r][c+1] == ']' {
		return BoxPosition{R: r, LeftC: c}, true
	}
	if gr[r][c] == ']' && c-1 >= 0 && gr[r][c-1] == '[' {
		return BoxPosition{R: r, LeftC: c - 1}, true // Return the position of '['
	}
	return BoxPosition{}, false
}

func canPushChain(gr [][]rune, boxes []BoxPosition, dr, dc int) bool {
	if len(boxes) == 0 {
		return false
	}

	movingBoxesSet := make(map[string]bool)
	for _, b := range boxes {
		movingBoxesSet[fmt.Sprintf("%d,%d", b.R, b.LeftC)] = true
	}

	for _, boxToMove := range boxes {
		targetR_Left := boxToMove.R + dr
		targetC_Left := boxToMove.LeftC + dc
		targetR_Right := boxToMove.R + dr
		targetC_Right := boxToMove.LeftC + 1 + dc

		if targetR_Left < 0 || targetR_Left >= len(gr) || targetC_Left < 0 || targetC_Left >= len(gr[0]) {
			return false // Out of bounds
		}
		if gr[targetR_Left][targetC_Left] == '#' {
			return false
		}
		if gr[targetR_Left][targetC_Left] != '.' {
			collidingBox, isBox := getBoxFromGrid(gr, targetR_Left, targetC_Left)
			if isBox {
				if !movingBoxesSet[fmt.Sprintf("%d,%d", collidingBox.R, collidingBox.LeftC)] {
					return false
				}
			} else if gr[targetR_Left][targetC_Left] != '@' { // '@' might be where player moves into
			}
		}

		if targetR_Right < 0 || targetR_Right >= len(gr) || targetC_Right < 0 || targetC_Right >= len(gr[0]) {
			return false
		}
		if gr[targetR_Right][targetC_Right] == '#' {
			return false
		}
		if gr[targetR_Right][targetC_Right] != '.' { // If not empty
			collidingBox, isBox := getBoxFromGrid(gr, targetR_Right, targetC_Right)
			if isBox {
				if !movingBoxesSet[fmt.Sprintf("%d,%d", collidingBox.R, collidingBox.LeftC)] {
					return false
				}
			} else if gr[targetR_Right][targetC_Right] != '@' {

			}
		}
	}
	return true
}

func expandGrid(grid [][]rune) [][]rune {
	expandedGrid := make([][]rune, 0, len(grid))
	for _, row := range grid {
		newRow := make([]rune, 0, len(row)*2)
		for _, cell := range row {
			switch cell {
			case '#':
				newRow = append(newRow, '#', '#')
			case 'O':
				newRow = append(newRow, '[', ']')
			case '.':
				newRow = append(newRow, '.', '.')
			case '@':
				newRow = append(newRow, '@', '.')
			default:
				// Optionally handle unknown runes
				newRow = append(newRow, cell, cell)
			}
		}

		expandedGrid = append(expandedGrid, newRow)
	}

	return expandedGrid
}

// Helper function to print the grid (optional)
func printGrid(grid [][]rune) {
	if grid == nil {
		fmt.Println("Cannot print nil grid")
		return
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func deepCopyGrid(original [][]rune) [][]rune {
	copyGrid := make([][]rune, len(original))
	for i := range original {
		copyGrid[i] = make([]rune, len(original[i]))
		copy(copyGrid[i], original[i])
	}
	return copyGrid
}
