package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const MAX_SIZE int = 10

func main() {
	inp := `
3 1 2 6 9 7 9 9 1 7
1 1 1 7 3 4 0 6 3 9
7 1 3 6 9 9 1 9 3 6
4 0 8 0 1 0 3 8 3 0
7 8 1 0 7 6 3 3 3 3
7 7 6 0 5 2 4 0 3 3
2 0 0 0 0 0 0 0 3 9
5 3 2 0 4 0 5 5 3 0
3 9 9 0 4 8 8 2 7 5
7 8 4 0 8 4 7 3 0 7
	`
	inp = strings.TrimSpace(inp)
	matrix := toMatrix(inp)
	i, j := findThePlusSign(matrix)
	fmt.Println(inp)
	fmt.Println(i, j)
}

func toMatrix(inp string) [MAX_SIZE][MAX_SIZE]int {
	m := strings.Split(inp, "\n")
	matrix := [MAX_SIZE][MAX_SIZE]int{}

	for i := range m {
		row := m[i]
		numbers := strings.Split(row, " ")

		for j := range numbers {
			n, err := strconv.Atoi(numbers[j])
			if err != nil {
				log.Fatal(err.Error())
			}
			matrix[i][j] = int(n)
		}
	}
	return matrix
}

func findThePlusSign(arr [MAX_SIZE][MAX_SIZE]int) (n, m int) {
	maxLength := 0
	currentLenght := 0
	colHolder := -1
	rowHolder := -1

	for i := 0; i < len(arr); i++ {
		v := arr[i]
		colCounter := 0
		rowCounter := 0

		for j := 0; j < len(v); j++ {
			element := v[j]

			for {
				if colCounter > i ||
					rowCounter > j ||
					rowCounter+j >= MAX_SIZE ||
					colCounter+i >= MAX_SIZE ||
					arr[i-colCounter][j] != element ||
					arr[i+colCounter][j] != element ||
					v[j-rowCounter] != element ||
					v[j+rowCounter] != element {
					break
				}

				top := arr[i-colCounter][j]
				bottom := arr[i+colCounter][j]
				left := v[j-rowCounter]
				right := v[j+rowCounter]

				if top != element && bottom != element && left != element && right != element {
					break
				}

				if top == element || bottom == element {
					colCounter += 1
				}

				if left == element || right == element {
					rowCounter += 1
				}

				if colCounter > rowCounter {
					currentLenght = colCounter
				} else {
					currentLenght = rowCounter
				}

				if currentLenght > maxLength {
					maxLength = currentLenght
					if i != 0 || j != 0 {
						colHolder = i
						rowHolder = j
					}
				}
			}
		}
	}
	return colHolder, rowHolder
}
