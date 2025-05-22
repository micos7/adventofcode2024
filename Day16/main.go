package main

import (
	"container/heap"
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

	grid := make([][]rune, 0)
	rowsStr := strings.Split(inputStr, "\n")
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

	simulatePart1Dijkstra(grid)

	// printGrid(grid)
}

type Point struct {
	r, c int
}

type Direction int

const (
	DirUp    Direction = 0
	DirDown  Direction = 1
	DirLeft  Direction = 2
	DirRight Direction = 3
)

// Mapping from Direction to dr, dc
var dirVectors = map[Direction]Point{
	DirUp:    {-1, 0},
	DirDown:  {1, 0},
	DirLeft:  {0, -1},
	DirRight: {0, 1},
}

type State struct {
	Point Point
	Dir   Direction
}

type Node struct {
	State State
	Cost  int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func simulatePart1Dijkstra(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	// Find the starting point 'S'
	var start Point
	foundS := false
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				start = Point{r, c}
				foundS = true
				break
			}
		}
		if foundS {
			break
		}
	}

	if !foundS {
		return -1
	}

	const infinity = 1_000_000_000
	minCost := make([][][]int, rows)
	for r := range minCost {
		minCost[r] = make([][]int, cols)
		for c := range minCost[r] {
			minCost[r][c] = make([]int, 4)
			for d := range minCost[r][c] {
				minCost[r][c][d] = infinity
			}
		}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	initialState := State{Point: start, Dir: DirRight}
	initialNode := &Node{State: initialState, Cost: 0}

	heap.Push(&pq, initialNode)
	minCost[start.r][start.c][DirRight] = 0

	minOverallCost := infinity

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*Node)
		currentState := currentNode.State
		currentPoint := currentState.Point
		currentDir := currentState.Dir
		currentCost := currentNode.Cost

		if currentCost > minCost[currentPoint.r][currentPoint.c][currentDir] {
			continue
		}

		if grid[currentPoint.r][currentPoint.c] == 'E' {
			if currentCost < minOverallCost {
				minOverallCost = currentCost
			}
		}

		for _, nextDir := range []Direction{DirUp, DirDown, DirLeft, DirRight} {
			dr, dc := dirVectors[nextDir].r, dirVectors[nextDir].c

			newR := currentPoint.r + dr
			newC := currentPoint.c + dc

			if newR >= 0 && newR < rows && newC >= 0 && newC < cols && grid[newR][newC] != '#' {
				newCost := currentCost + 1

				if nextDir != currentDir {
					newCost += 1000
				}

				newState := State{Point: Point{newR, newC}, Dir: nextDir}

				if newCost < minCost[newState.Point.r][newState.Point.c][newState.Dir] {
					minCost[newState.Point.r][newState.Point.c][newState.Dir] = newCost
					heap.Push(&pq, &Node{State: newState, Cost: newCost})
				}
			}
		}
	}

	if minOverallCost == infinity {
		fmt.Println("No path found from 'S' to 'E'.")
		return -1
	} else {
		fmt.Printf("Lowest score to reach 'E': %d\n", minOverallCost)
		return minOverallCost
	}
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}
