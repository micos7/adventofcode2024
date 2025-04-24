package main

import (
	"fmt"
	"io"
	"os"
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

	numbers := strings.Fields(inputStr)

	var nums []int
	for _, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			continue
		}
		nums = append(nums, n)
	}

	final := transformRecursive(nums, 25)

	fmt.Println(len(final))

}

func transformRecursive(nums []int, steps int) []int {
	if steps <= 0 {
		return nums
	}

	var next []int
	for _, n := range nums {
		next = append(next, transformFn(n)...)
	}

	return transformRecursive(next, steps-1)
}

func transformFn(n int) []int {
	if n == 0 {
		return []int{1}
	}
	left, right, ok := splitIfEvenDigits(n)
	if ok {
		return []int{left, right}
	} else {
		return []int{n * 2024}
	}
}

func splitIfEvenDigits(n int) (int, int, bool) {
	str := strconv.Itoa(n)
	length := len(str)

	if length%2 != 0 {
		return 0, 0, false
	}

	half := length / 2
	leftStr := str[:half]
	rightStr := str[half:]

	left, err1 := strconv.Atoi(leftStr)
	right, err2 := strconv.Atoi(rightStr)

	if err1 != nil || err2 != nil {
		return 0, 0, false
	}

	return left, right, true
}
