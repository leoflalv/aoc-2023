package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Set the path of the solution
const PATH = "day_4/problem_2/"

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
	copyNumber := make([]int, len(entries))

	for i, entry := range entries {
		if copyNumber[i] == 0 {
			copyNumber[i] = 1
		}

		winList, yourList := getLists(entry)
		lineSum := 0

		for _, card := range yourList {
			_, ok := winList[card]

			if ok {
				lineSum += 1
			}
		}

		for rep := 1; rep <= lineSum; rep++ {
			if i+rep < len(copyNumber) {
				if copyNumber[i+rep] == 0 {
					copyNumber[i+rep] = 1
				}

				copyNumber[i+rep] += copyNumber[i]
			}

		}

	}

	for _, num := range copyNumber {
		sum += num
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
