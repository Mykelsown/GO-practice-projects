package main

import (
	"fmt"
	"gopracticeprojects"
)


func main() {
	fmt.Println("---------------SALE ORDER PROCESSOR----------------\n ")

	items := []string{
		"Jug", "T-Shirt_SALE", "Mouse", "Fan_SALE",
	}

	a1 := gopracticeprojects.TextSlicer(items)
	a2 := gopracticeprojects.DataCollation(a1)
	fmt.Println(gopracticeprojects.TableConstructor(a2))
	fmt.Printf("The total discounted price is %v\nThanks for shopping with us\n", gopracticeprojects.SaleOrderProcessor(items))
	// fmt.Printf("Your total discounted price is ₦%v \n", gopracticeprojects.SaleOrderProcessor(items))
}
