package solutions

import (
	"bufio"
	"fmt"
	"os"
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

		previousNumber, _ := strconv.Atoi(line[0])
		previousDifference := 0
		safe := true
		for i := 1; i < len(line); i++ {
			currentNumber, _ := strconv.Atoi(line[i])
			currentDifference := previousNumber - currentNumber

			if currentDifference == 0 || currentDifference > 3 || currentDifference < -3 {
				safe = false
				break
			}

			if (previousDifference < 0 && currentDifference > 0) || (previousDifference > 0 && currentDifference < 0) {
				safe = false
				break
			}

			previousNumber = currentNumber
			previousDifference = currentDifference
		}

		if safe {
			result++
		}
	}

	fmt.Println("Day 02 - Part 01:", result)
}
