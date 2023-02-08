package main


/*
**************

Calculate the sum of all positive elements of the array

**************
*/


import "fmt"
import "math/rand"
import "math"
import "time"

func main() {
	rand.Seed(time.Now().UnixNano())

	var array [15]int

	// filling the array with random integers in range -20:20
	// negative numbers are gotten by multiplying randInt1 with minus one powered randInt2
	// if randInt2 is odd -- number is negative
	for i:=0; i<15; i++ {
		array[i] = rand.Intn(20) * int(math.Pow(-1, float64(rand.Intn(10))))
	}

	fmt.Println(array)

	summ := summPositive(array)

	fmt.Println(summ)
}


func summPositive(array [15]int) int {
	sum := 0
	for i:=0; i<15; i++ {
		if array[i] > 0 {
			fmt.Printf(" *%d ", array[i])
			sum += array[i]
		}
	}
	fmt.Print("\n")

	return sum
}
