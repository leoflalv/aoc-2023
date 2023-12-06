package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const PATH = "day_1/problem_1/"

var NUMBERS = map[rune]struct{}{
	'0': {},
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
}

func getNumber(str string) int {
	var current rune
	first := ' '

	for _, char := range str {
		if _, ok := NUMBERS[char]; ok {
			current = char
			if first == ' ' {
				first = char
			}
		}
	}

	number := 10*(int(first)-48) + (int(current) - 48)

	return number
}

func getSum(entries []string) int {
	sum := 0
	for _, str := range entries {
		sum = sum + getNumber(str)
	}

	return sum
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

	solution := getSum(data)

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
