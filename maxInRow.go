package main

/*
*************

Create a 2d array 5x6 filled with random numers.
Create a 2d array (slice, whatever) filled with the maximum number
in each row of the 2d array created earlier.

*************
*/

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().UnixNano())
	// actualy matrix is an array of 5 slices
	var matrix [5][]int

	for i:=0; i<5; i++ {
		for j:=0; j<6; j++ {
			matrix[i] = append(matrix[i], rand.Intn(100))
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Print("\n")
	}


	maxies := maxInRow(matrix[:])

	fmt.Printf("%v\n", maxies)

}


// [][]int recives a slice of 2d array?
func maxInRow(matrix [][]int) []int {
	array := []int{}

	for i:=0; i<5; i++ {
		max := matrix[i][0]
		for j:=0; j<6; j++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
		array = append(array, max)
	}

	return array
}
