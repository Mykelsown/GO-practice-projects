package main

import (
	"fmt"
	"gopracticeprojects"
)

func main() {
	fmt.Println("---------Let it begin----------")

	items := []string{
		"Jug", "T-Shirt_SALE", "Mouse", "Fan_SALE",
	}
	fmt.Printf("Your total discounted price is ₦%v \n", gopracticeprojects.SaleOrderProcessor(items))
}
