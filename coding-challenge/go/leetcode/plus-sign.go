package main

import "log"

func findThePlusSignOrig(arr [MAX_SIZE][MAX_SIZE]int) (n, m int) {
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
					log.Println(i, j)
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
