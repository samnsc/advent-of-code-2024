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
		safe, _ := CheckSafety(line)

		if safe {
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
		safe, problemIndex := CheckSafety(line)

		if safe {
			result++
			continue
		} else {
			// I only have to teste the string without the problem index and the two indices that preceed it
			for i := problemIndex; i > problemIndex-3 && i >= 0; i-- {
				dampenedSlice := slices.Delete(slices.Clone(line), i, i+1)
				safe, _ := CheckSafety(dampenedSlice)

				if safe {
					result++
					break
				}
			}
		}
	}

	fmt.Println("Day 02 - Part 02:", result)
}

func CheckSafety(line []string) (bool, int) {
	previousNumber, _ := strconv.Atoi(line[0])
	previousDifference := 0
	safe := true
	problemIndex := -1
	for i := 1; i < len(line); i++ {
		currentNumber, _ := strconv.Atoi(line[i])
		currentDifference := previousNumber - currentNumber

		if currentDifference == 0 ||
			currentDifference > 3 ||
			currentDifference < -3 ||
			(previousDifference < 0 && currentDifference > 0) ||
			(previousDifference > 0 && currentDifference < 0) {
			safe = false
			problemIndex = i
			break
		}

		previousNumber = currentNumber
		previousDifference = currentDifference
	}

	return safe, problemIndex
}
