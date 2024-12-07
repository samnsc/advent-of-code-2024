package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day02Part01() {
	file, error := os.Open("./solutions/day_02.txt")
	if error != nil {
		fmt.Println(error)
		panic(error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if CheckSafety(line) {
			result++
		}
	}

	fmt.Println("Day 02 - Part 01:", result)
}

func Day02Part02() {
	file, error := os.Open("./solutions/day_02.txt")
	if error != nil {
		fmt.Println(error)
		panic(error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if CheckSafety(line) {
			result++
			continue
		} else {
			for i := range line {
				// don't really like this solution but none of my other "smart" solutions seemed to work
				dampenedSlice := slices.Delete(slices.Clone(line), i, i+1)

				if CheckSafety(dampenedSlice) {
					result++
					break
				}
			}
		}
	}

	fmt.Println("Day 02 - Part 02:", result)
}

func CheckSafety(line []string) bool {
	previousNumber, _ := strconv.Atoi(line[0])
	previousDifference := 0
	safe := true
	for i := 1; i < len(line); i++ {
		currentNumber, _ := strconv.Atoi(line[i])
		currentDifference := previousNumber - currentNumber

		if currentDifference == 0 ||
			currentDifference > 3 ||
			currentDifference < -3 ||
			(previousDifference < 0 && currentDifference > 0) ||
			(previousDifference > 0 && currentDifference < 0) {
			safe = false
			break
		}

		previousNumber = currentNumber
		previousDifference = currentDifference
	}

	return safe
}
