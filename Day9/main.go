package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type SpaceSegment struct {
	Start  int
	Length int
}

type FileSegment struct {
	ID     int
	Start  int
	Length int
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

	// fmt.Println(disk)

	diskCopy := make([]interface{}, len(disk))
	copy(diskCopy, disk)

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

	maxFileID := fileCount - 1

	for currentFileID := maxFileID; currentFileID >= 0; currentFileID-- {

		fileStart, fileLength := findFileSegmentInterface(diskCopy, currentFileID)

		if fileStart != -1 {

			spaceStart, _ := findLeftmostSuitableSpaceInterface(diskCopy, fileLength, fileStart-1)

			if spaceStart != -1 {

				movedFileBlocks := make([]interface{}, fileLength)
				for i := 0; i < fileLength; i++ {
					movedFileBlocks[i] = diskCopy[fileStart+i]
				}

				for i := 0; i < fileLength; i++ {
					diskCopy[fileStart+i] = "."
				}

				for i := 0; i < fileLength; i++ {
					diskCopy[spaceStart+i] = movedFileBlocks[i]
				}
			}
		}
	}

	checksumPart2 := 0
	for position, blockValue := range diskCopy {
		if fileID, ok := blockValue.(int); ok {
			checksumPart2 += position * fileID
		}
	}

	fmt.Println("Resulting checksum part 2:", checksumPart2)

	fmt.Println("Total check count part 1:", checksum)
}

func findFileSegmentInterface(disk []interface{}, fileID int) (start int, length int) {
	start = -1
	length = 0

	firstBlockIdx := -1
	for i := 0; i < len(disk); i++ {
		if val, ok := disk[i].(int); ok && val == fileID {
			firstBlockIdx = i
			break
		}
	}

	if firstBlockIdx == -1 {
		return -1, 0
	}

	start = firstBlockIdx
	length = 0
	for i := start; i < len(disk); i++ {
		if val, ok := disk[i].(int); ok && val == fileID {
			length++
		} else {
			break
		}
	}

	return start, length
}

func findLeftmostSuitableSpaceInterface(disk []interface{}, minLength int, maxIndex int) (start int, length int) {
	start = -1
	length = 0

	currentSpaceStart := -1
	currentSpaceLength := 0
	for i := 0; i <= maxIndex && i < len(disk); i++ {
		if disk[i] == "." {
			if currentSpaceStart == -1 {
				currentSpaceStart = i
				currentSpaceLength = 1
			} else {
				currentSpaceLength++
			}
		} else {
			if currentSpaceStart != -1 {
				if currentSpaceLength >= minLength {
					return currentSpaceStart, currentSpaceLength
				}
				currentSpaceStart = -1
				currentSpaceLength = 0
			}
		}
	}

	if currentSpaceStart != -1 && currentSpaceLength >= minLength {
		return currentSpaceStart, currentSpaceLength
	}

	return -1, 0
}
