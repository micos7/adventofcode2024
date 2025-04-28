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

}

func negativeMod(a int, b int) int {
	return (a%b + b) % b
}
