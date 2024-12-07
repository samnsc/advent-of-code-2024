package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day01Part01() {
	file, error := os.Open("./solutions/day_01.txt")
	if error != nil {
		fmt.Println(error)
		panic(error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var firstGroup []int
	var secondGroup []int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		firstGroupNumber, _ := strconv.Atoi(line[0])
		firstGroup = append(firstGroup, firstGroupNumber)

		secondGroupNumber, _ := strconv.Atoi(line[1])
		secondGroup = append(secondGroup, secondGroupNumber)
	}

	slices.SortFunc(firstGroup, func(a, b int) int {
		return a - b
	})

	slices.SortFunc(secondGroup, func(a, b int) int {
		return a - b
	})

	result := 0
	for i := 0; i < len(firstGroup); i++ {
		distance := firstGroup[i] - secondGroup[i]
		if distance < 0 {
			distance -= 2 * distance
		}
		result += distance
	}

	fmt.Println("Day 01 - Part 01:", result)
}
