package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func countAntiNodes(input string) int64 {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	antennas := make(map[rune][][2]int)
	antiNodes := make(map[[2]int]bool)

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != '.' {
				antennas[grid[row][col]] = append(antennas[grid[row][col]], [2]int{row, col})
			}
		}
	}

	for _, positions := range antennas {
		for i := range positions {
			for j := i + 1; j < len(positions); j++ {
				pos1 := positions[i]
				pos2 := positions[j]

				deltaRow := pos2[0] - pos1[0]
				deltaCol := pos2[1] - pos1[1]

				dist1 := math.Sqrt(float64(deltaRow*deltaRow + deltaCol*deltaCol))

				if dist1 == 0 {
					continue
				}

				antiNode1 := [2]int{pos1[0] - deltaRow, pos1[1] - deltaCol}
				antiNode2 := [2]int{pos2[0] + deltaRow, pos2[1] + deltaCol}

				if antiNode1[0] >= 0 && antiNode1[0] < len(grid) && antiNode1[1] >= 0 && antiNode1[1] < len(grid[0]) {
					antiDist1 := math.Sqrt(float64((antiNode1[0]-pos1[0])*(antiNode1[0]-pos1[0]) + (antiNode1[1]-pos1[1])*(antiNode1[1]-pos1[1])))
					antiDist2 := math.Sqrt(float64((antiNode1[0]-pos2[0])*(antiNode1[0]-pos2[0]) + (antiNode1[1]-pos2[1])*(antiNode1[1]-pos2[1])))

					if antiDist1*2 == antiDist2 || antiDist2*2 == antiDist1 {
						antiNodes[antiNode1] = true
					}

				}

				if antiNode2[0] >= 0 && antiNode2[0] < len(grid) && antiNode2[1] >= 0 && antiNode2[1] < len(grid[0]) {
					antiDist1 := math.Sqrt(float64((antiNode2[0]-pos1[0])*(antiNode2[0]-pos1[0]) + (antiNode2[1]-pos1[1])*(antiNode2[1]-pos1[1])))
					antiDist2 := math.Sqrt(float64((antiNode2[0]-pos2[0])*(antiNode2[0]-pos2[0]) + (antiNode2[1]-pos2[1])*(antiNode2[1]-pos2[1])))

					if antiDist1*2 == antiDist2 || antiDist2*2 == antiDist1 {
						antiNodes[antiNode2] = true
					}
				}
			}
		}
	}

	output := len(antiNodes)
	return int64(output)
}

func countAntiNodesUpdated(input string) int64 {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	rows, cols := len(grid), len(grid[0])
	antennas := make(map[rune][][2]int)
	antiNodes := make(map[[2]int]bool)

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != '.' {
				antennas[grid[row][col]] = append(antennas[grid[row][col]], [2]int{row, col})
			}
		}
	}

	for _, positions := range antennas {
		if len(positions) < 2 {
			continue
		}

		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				currentPos := [2]int{r, c}

				if isCollinearWithAtLeastTwo(currentPos, positions) {
					antiNodes[currentPos] = true
				}
			}
		}
	}

	return int64(len(antiNodes))
}

func isCollinearWithAtLeastTwo(pos [2]int, antennas [][2]int) bool {
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {

			if areCollinear(antennas[i], antennas[j], pos) {
				return true
			}
		}
	}
	return false
}

func areCollinear(A, B, C [2]int) bool {
	area := (B[0]-A[0])*(C[1]-A[1]) - (C[0]-A[0])*(B[1]-A[1])
	return area == 0
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

	input := string(content)

	totalAntiNodesPart1 := int64(0)
	totalAntiNodesPart2 := int64(0)

	totalAntiNodesPart1 = countAntiNodes(input)

	totalAntiNodesPart2 = countAntiNodesUpdated(input)

	fmt.Println("Total anti-node count part 1:", totalAntiNodesPart1)
	fmt.Println("Total anti-node count part 2:", totalAntiNodesPart2)
}
