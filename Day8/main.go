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

	totalSumPart1 := int64(0)

	totalSumPart1 = countAntiNodes(input)

	fmt.Println("Total anti-node count:", totalSumPart1)
}
