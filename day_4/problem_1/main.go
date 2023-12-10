package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Set the path of the solution
const PATH = "day_4/problem_1/"

func getLists(entry string) (winList map[string]struct{}, yourList []string) {
	winList = make(map[string]struct{})

	listString := strings.Split(entry, ":")[1]
	list := strings.Split(listString, "|")

	for _, num := range strings.Split(list[0], " ") {
		trimmed := strings.TrimSpace(num)
		if trimmed != "" {
			winList[trimmed] = struct{}{}
		}
	}

	for _, num := range strings.Split(list[1], " ") {
		trimmed := strings.TrimSpace(num)
		if trimmed != "" {
			yourList = append(yourList, trimmed)
		}
	}

	return
}

func solution(entries []string) int {
	sum := 0

	for _, entry := range entries {
		winList, yourList := getLists(entry)
		lineSum := 0

		for _, card := range yourList {
			_, ok := winList[card]

			if ok {
				if lineSum == 0 {
					lineSum = 1
				} else {
					lineSum *= 2
				}
			}
		}

		sum += lineSum
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
