package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const PATH = "day_2/problem_1/"

const RED = "red"
const GREEN = "green"
const BLUE = "blue"

const RED_AMOUNT = 12
const GREEN_AMOUNT = 13
const BLUE_AMOUNT = 14

func isPossible(red int, green int, blue int) bool {
	return red <= RED_AMOUNT && green <= GREEN_AMOUNT && blue <= BLUE_AMOUNT
}

func getNumbers(str string) (blue int, red int, green int) {
	blue = 0
	red = 0
	green = 0

	for _, cubes := range strings.Split(str, ",") {
		withoutSpace := cubes[1:]
		if strings.Contains(withoutSpace, BLUE) {
			blue, _ = strconv.Atoi(strings.Split(withoutSpace, " ")[0])
		}
		if strings.Contains(withoutSpace, RED) {
			red, _ = strconv.Atoi(strings.Split(withoutSpace, " ")[0])
		}
		if strings.Contains(withoutSpace, GREEN) {
			green, _ = strconv.Atoi(strings.Split(withoutSpace, " ")[0])
		}
	}

	return
}

func solution(entries []string) int {
	sum := 0

	for _, entry := range entries {
		possible := true
		split := strings.Split(entry, ":")
		game, _ := strconv.Atoi(strings.Split(split[0], " ")[1])
		set := split[1]

		for _, subset := range strings.Split(set, ";") {
			blue, red, green := getNumbers(subset)
			possible = possible && isPossible(red, green, blue)
		}

		if possible {
			sum += game
		}
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

	// Add here the function which solve the problem
	solution := solution(data)

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
