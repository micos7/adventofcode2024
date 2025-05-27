package main

import (
	"container/heap"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
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
	simulatePart2Dijkstra(grid)
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

func simulatePart2Dijkstra(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	var start, end Point
	foundS, foundE := false, false
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				start = Point{r, c}
				foundS = true
			} else if grid[r][c] == 'E' {
				end = Point{r, c}
				foundE = true
			}
		}
	}
	if !foundS || !foundE {
		return -1
	}

	const INF = 1000000000
	dist := make([][][]int, rows)
	for r := range dist {
		dist[r] = make([][]int, cols)
		for c := range dist[r] {
			dist[r][c] = make([]int, 4)
			for d := range dist[r][c] {
				dist[r][c][d] = INF
			}
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	startState := State{Point: start, Dir: DirRight}
	dist[start.r][start.c][DirRight] = 0
	heap.Push(pq, &Node{State: startState, Cost: 0})
	minEndCost := INF

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		state := current.State
		cost := current.Cost
		if cost > dist[state.Point.r][state.Point.c][state.Dir] {
			continue
		}
		if state.Point.r == end.r && state.Point.c == end.c {
			minEndCost = min(minEndCost, cost)
			continue
		}
		dr, dc := dirVectors[state.Dir].r, dirVectors[state.Dir].c
		nr := state.Point.r + dr
		nc := state.Point.c + dc
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] != '#' {
			newCost := cost + 1
			if newCost < dist[nr][nc][state.Dir] {
				dist[nr][nc][state.Dir] = newCost
				newState := State{Point: Point{nr, nc}, Dir: state.Dir}
				heap.Push(pq, &Node{State: newState, Cost: newCost})
			}
		}

		var leftDir, rightDir Direction
		switch state.Dir {
		case DirUp:
			leftDir, rightDir = DirLeft, DirRight
		case DirDown:
			leftDir, rightDir = DirRight, DirLeft
		case DirLeft:
			leftDir, rightDir = DirDown, DirUp
		case DirRight:
			leftDir, rightDir = DirUp, DirDown
		}
		for _, newDir := range []Direction{leftDir, rightDir} {
			newCost := cost + 1000
			if newCost < dist[state.Point.r][state.Point.c][newDir] {
				dist[state.Point.r][state.Point.c][newDir] = newCost
				newState := State{Point: state.Point, Dir: newDir}
				heap.Push(pq, &Node{State: newState, Cost: newCost})
			}
		}
	}

	if minEndCost == INF {
		return -1
	}

	onPath := make(map[Point]bool)
	queue := []State{}
	visited := make(map[State]bool)

	for dir := Direction(0); dir < 4; dir++ {
		if dist[end.r][end.c][dir] == minEndCost {
			state := State{Point: end, Dir: dir}
			queue = append(queue, state)
			visited[state] = true
			onPath[end] = true
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentCost := dist[current.Point.r][current.Point.c][current.Dir]

		dr, dc := dirVectors[current.Dir].r, dirVectors[current.Dir].c
		prevR := current.Point.r - dr
		prevC := current.Point.c - dc
		if prevR >= 0 && prevR < rows && prevC >= 0 && prevC < cols &&
			grid[prevR][prevC] != '#' && dist[prevR][prevC][current.Dir] == currentCost-1 {
			prevState := State{Point: Point{prevR, prevC}, Dir: current.Dir}
			if !visited[prevState] {
				visited[prevState] = true
				queue = append(queue, prevState)
				onPath[Point{prevR, prevC}] = true
			}
		}

		for prevDir := Direction(0); prevDir < 4; prevDir++ {
			if prevDir == current.Dir {
				continue
			}
			canTurn := false
			switch prevDir {
			case DirUp, DirDown:
				canTurn = (current.Dir == DirLeft || current.Dir == DirRight)
			case DirLeft, DirRight:
				canTurn = (current.Dir == DirUp || current.Dir == DirDown)
			}
			if canTurn && dist[current.Point.r][current.Point.c][prevDir] == currentCost-1000 {
				prevState := State{Point: current.Point, Dir: prevDir}
				if !visited[prevState] {
					visited[prevState] = true
					queue = append(queue, prevState)
					onPath[current.Point] = true
				}
			}
		}
	}

	fmt.Printf("Number of tiles on best paths: %d\n", len(onPath))
	return len(onPath)
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
