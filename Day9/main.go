package main

import (
	"fmt"
	"io"
	"os"
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

	var lengths []int
	for _, char := range inputStr {
		if char >= '0' && char <= '9' {
			digit := int(char - '0')
			lengths = append(lengths, digit)
		}
	}

	var disk []interface{} = make([]interface{}, 0)
	fileCount := 0

	isFileLength := true

	for _, length := range lengths {
		if isFileLength {
			for i := 0; i < length; i++ {
				disk = append(disk, fileCount)
			}
			fileCount++
		} else {
			for i := 0; i < length; i++ {
				disk = append(disk, ".")
			}
		}
		isFileLength = !isFileLength
	}

	l := 0
	r := len(disk) - 1

	for l < r {
		for l < len(disk) && disk[l] != "." {
			l++
		}

		for r >= 0 && disk[r] == "." {
			r--
		}

		if l < r {
			temp := disk[r]
			disk[l] = temp
			disk[r] = "."

			l++
			r--
		}
	}

	checksum := 0
	for position, blockValue := range disk {
		if fileID, ok := blockValue.(int); ok {
			checksum += position * fileID
		}
	}

	fmt.Println("Total check count part 1:", checksum)
}
