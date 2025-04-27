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
	lines := strings.Split(inputStr, "\n\n")

	re := regexp.MustCompile(`(?m)(Button A|Button B|Prize): X[\+=](\d+), Y[\+=](\d+)`)

	part1Tokens := 0

	for _, machine := range lines {
		matches := re.FindAllStringSubmatch(machine, -1)

		var ax, ay, bx, by, prizeX, prizeY int

		for _, match := range matches {
			label := match[1]
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])

			switch label {
			case "Button A":
				ax = x
				ay = y
			case "Button B":
				bx = x
				by = y
			case "Prize":
				prizeX = x
				prizeY = y
			}
		}

		minTokens := -1

		for a := 0; a <= 1000; a++ {
			for b := 0; b <= 1000; b++ {
				x := (a * ax) + (b * bx)
				y := (a * ay) + (b * by)

				if x == prizeX && y == prizeY {
					tokens := (a * 3) + (b * 1)
					if minTokens == -1 || tokens < minTokens {
						minTokens = tokens
					}
				}
			}
		}

		if minTokens != -1 {
			part1Tokens += minTokens
		} else {
			fmt.Println("No solution found.")
		}
	}

	fmt.Println(part1Tokens)
}
