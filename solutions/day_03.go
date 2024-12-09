package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Day03Part01() {
	file, error := os.Open("./solutions/day_03.txt")
	if error != nil {
		fmt.Println(error)
		panic(error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		text := scanner.Text()

		for i := 0; i < len(text); i++ {
			if text[i] == 'm' && text[i+1] == 'u' && text[i+2] == 'l' && text[i+3] == '(' {
				i += 3 // first value after the initial '('
				var tokens []int

				var aux []rune
				for j := i + 1; j < i+9; j++ { // 8 is the maximum size that it can be
					if unicode.IsDigit(rune(text[j])) {
						aux = append(aux, rune(text[j]))
					} else if rune(text[j]) == ',' {
						if tokens == nil && aux != nil && len(aux) <= 3 {
							value, _ := strconv.Atoi(string(aux))

							tokens = append(tokens, value)
							aux = nil
						} else {
							break
						}
					} else if rune(text[j]) == ')' {
						if len(tokens) == 1 && aux != nil && len(aux) <= 3 {
							value, _ := strconv.Atoi(string(aux))

							tokens = append(tokens, value)
						}
						break
					} else {
						break
					}
				}

				if len(tokens) == 2 {
					result += tokens[0] * tokens[1]
				}
			}
		}
	}

	fmt.Println("Day 03 - Part 01:", result)
}
