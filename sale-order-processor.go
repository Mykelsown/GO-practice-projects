package gopracticeprojects

import "fmt"

var productPrice = map[string]float64{
	"T-Shirt": 20.00,
	"Jug":     10.5,
	"Fan":     30.50,
	"Mouse":   5.00,
}

func CheckForSale(item string) (bool, string) {
	check := ""
	get := ""
	for i, l := range item {
		if i < len(item)-5 {
			get += string(l)
		}

		if i > len(item)-6 {
			check += string(l)
		}

	}

	if check == "_SALE" {
		return true, get
	}

	return false, ""
}

func SaleOrderProcessor(orderItems []string) float64 {
	var totalDiscountPrice float64

	for _, item := range orderItems {
		isSale, product := CheckForSale(string(item))
		if isSale == true {
			productPrice[product] = productPrice[product] * 0.9
			totalDiscountPrice += productPrice[product]
		}
		totalDiscountPrice += productPrice[item]
	}

	return totalDiscountPrice
}

func DataCollation() [][]string{
	data := [][]string{
		{"Products", "Initial Price", "Discounted Price"},
		{""}, 
	}
}

func TableConstructor(data []string) string{

}
