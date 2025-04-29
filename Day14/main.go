package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

	re := regexp.MustCompile(`-?\d+`)

	c1 := 0
	c2 := 0
	c3 := 0
	c4 := 0

	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		if len(matches) < 4 {
			fmt.Println("Not enough numbers found")
			return
		}

		c, _ := strconv.Atoi(matches[0])
		r, _ := strconv.Atoi(matches[1])
		cv, _ := strconv.Atoi(matches[2])
		rv, _ := strconv.Atoi(matches[3])

		newc := negativeMod(c+(cv*100), 101)
		newr := negativeMod(r+(rv*100), 103)

		if newc < 50 && newr < 51 {

			c1++
		} else if newc > 50 && newr < 51 {

			c2++
		} else if newc < 50 && newr > 51 {

			c3++
		} else if newc > 50 && newr > 51 {
			c4++
		}
	}

	fmt.Println(c1 * c2 * c3 * c4)

	type Point struct {
		c, r   int
		cv, rv int
	}
	var points []Point

	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		if len(matches) < 4 {
			fmt.Println("Not enough numbers found")
			return
		}

		c, _ := strconv.Atoi(matches[0])
		r, _ := strconv.Atoi(matches[1])
		cv, _ := strconv.Atoi(matches[2])
		rv, _ := strconv.Atoi(matches[3])

		points = append(points, Point{c, r, cv, rv})
	}
	// outputFile, err := os.Create("all_frames.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer outputFile.Close()
	const width = 101
	const height = 103

	var uniquePositionTime int

	for t := 1; t <= width*height; t++ {
		for i := range points {
			points[i].c = (points[i].c + points[i].cv + width) % width
			points[i].r = (points[i].r + points[i].rv + height) % height
		}

		positionMap := make(map[string]int)
		for _, p := range points {
			key := fmt.Sprintf("%d,%d", p.c, p.r)
			positionMap[key]++
		}

		duplicates := 0
		for _, count := range positionMap {
			if count > 1 {
				duplicates += count - 1
			}
		}
		if duplicates == 0 {
			uniquePositionTime = t
			break
		}
	}

	if uniquePositionTime > 0 {
		fmt.Printf("\nPart 2: Secs = %d\n", uniquePositionTime)
	} else {
		fmt.Println("No solution.")
	}

}

func negativeMod(a int, b int) int {
	return (a%b + b) % b
}
