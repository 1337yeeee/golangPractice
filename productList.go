package main

/*
***************

Create a shoping list that consists of Product and Price
Print each product whose price is more then 10

***************
*/

import "fmt"

type Buying struct {
	product string;
	price int
}

func main() {
	var shoppingList []Buying

	fmt.Print("Enter the number of elements: ")
	var number int
	fmt.Scan(&number)

	for i:=0; i<number; i++ {
		fmt.Printf("Enter the name of the %dth product: ", i+1)
		var product string
		fmt.Scan(&product)

		fmt.Printf("Enter the price of the %dth product: ", i+1)
		var price int
		fmt.Scan(&price)

		shoppingList = append(shoppingList, Buying{product, price})
	}

	fmt.Println(shoppingList)

	printPriceMoreThan(10, shoppingList)
}


func printPriceMoreThan(price int, list []Buying) {
	for i, v := range list {
		if v.price > price {
			fmt.Printf("%d: %v {%d}\n", i, v.product, v.price)
		}
	}
}
