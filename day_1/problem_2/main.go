package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

const PATH = "day_1/problem_2/"

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

var STR_NUMBERS = map[string]rune{
	"cero":  '0',
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func checkForStrNumber(str string, index int) (rune, bool) {
	end := int(math.Min(float64(len(str)), float64(index)+5))
	word := ""

	for i := index; i < end; i++ {
		word = word + string(str[i])

		if number, ok := STR_NUMBERS[word]; ok {
			return number, true
		}
	}

	return ' ', false
}

func getNumber(str string) int {
	var current rune
	first := ' '

	for i, char := range str {
		if _, ok := NUMBERS[char]; ok {
			current = char

			if first == ' ' {
				first = char
			}
		} else if strChar, ok := checkForStrNumber(str, i); ok {
			current = strChar

			if first == ' ' {
				first = strChar
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
