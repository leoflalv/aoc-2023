package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

// Set the path of the solution
const PATH = "day_3/problem_1/"

var xPos = []int{0, 1, 1, 1, 0, -1, -1, -1}
var yPos = []int{1, 1, 0, -1, -1, -1, 0, 1}

func checkX(x int, matrix [][]rune) bool {
	return x >= 0 && x < len(matrix[0])
}

func checkY(y int, matrix [][]rune) bool {
	return y >= 0 && y < len(matrix)
}

// Function to get the number
func getNumber(matrix [][]rune, xIndex int, yIndex int) (int, int) {
	startIndex := 0
	endIndex := len(matrix[yIndex]) - 1
	strNum := ""

	// Get number start position
	for x := xIndex; x >= 0; x-- {
		if !unicode.IsDigit(matrix[yIndex][x]) {
			startIndex = x + 1
			break
		}
	}

	// Get number end position
	for x := xIndex; x < len(matrix[yIndex]); x++ {
		if !unicode.IsDigit(matrix[yIndex][x]) {
			endIndex = x - 1
			break
		}
	}

	// Get number as string
	for x := startIndex; x <= endIndex; x++ {
		strNum += string(matrix[yIndex][x])
	}

	num, _ := strconv.Atoi(strNum)
	return num, endIndex
}

func getPiecesSum(matrix [][]rune) int {
	sum := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {

			for i := 0; i < 8; i++ {
				if unicode.IsDigit(matrix[y][x]) && checkX(x+xPos[i], matrix) && checkY(y+yPos[i], matrix) && matrix[y+yPos[i]][x+xPos[i]] != '.' && !unicode.IsDigit(matrix[y+yPos[i]][x+xPos[i]]) {
					num, lastIndex := getNumber(matrix, x, y)
					sum += num
					x = lastIndex
					break
				}
			}

		}
	}

	return sum
}

func parseEntry(entry []string) [][]rune {
	matrix := make([][]rune, len(entry))
	for i := 0; i < len(entry); i++ {
		matrix[i] = []rune(entry[i])
	}

	return matrix
}

func solve(entry []string) int {
	matrix := parseEntry(entry)
	return getPiecesSum(matrix)
}

func main() {
	var data []string

	fmt.Println("<-- Start -->")
	fmt.Println("Loading data....")

	inputPath := filepath.Join(PATH, "input.txt")
	file, err := os.Open(inputPath)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	fmt.Println("Loaded!!!")

	// Add here the function which solve the problem
	solution := solve(data)

	fmt.Println("Writing solution....")

	outputPath := filepath.Join(PATH, "output.txt")
	outFile, err := os.Create(outputPath)
	checkError(err)
	defer outFile.Close()

	_, err = outFile.WriteString(fmt.Sprint(solution))
	checkError(err)

	fmt.Println("Writed!!!!")
	fmt.Println("Done!!!!")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
