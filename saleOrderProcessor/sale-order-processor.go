package gopracticeprojects

import (
	"fmt"
	"strconv"
	"strings"
)

var productPrice = map[string]float64{
	"T-Shirt": 20.00,
	"Jug":     10.5,
	"Fan":     30.50,
	"Mouse":   5.00,
}

func CollateMapKeys(m map[string]float64) []string{
	keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
	return keys
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

// ------My-Function-----
// func TextSlicer(t []string) []string {
// 	var refined []string

// 	for i, el := range t {
// 		get := ""
// 		check := ""
// 		for l, char := range el {
// 			if l < len(el)-5 {
// 				get += string(char)
// 			}

// 			if i > len(el)-6 {
// 				check += string(l)
// 			}
// 		}

// 		if check == "_SALE" {
// 			refined[i] = get
// 		}
// 	}

// 	return refined
// }

//--------AI MODIFICATION------- 
func TextSlicer(t []string) []string {
	refined := make([]string, len(t))
	for i, el := range t {
		refined[i] = strings.TrimSuffix(el, "_SALE")
	}
	return refined
}


// ------My-Function-----

// func DataCollation(items []string) [][]string {
// 	row := len(items) + 1
// 	col := 3

// 	data := make([][]string, row)
// 	for i := range data{
// 		data[i] = make([]string, col)
// 	}

// 	keysInSlice := CollateMapKeys(productPrice)

// 	data[0][0] = "Product-Name"
// 	data[0][1] = "Initial-Price"
// 	data[0][2] = "Discounted-Price"
	
// 	for i := 1; i <= len(items); i++ {
// 		for j:=0; j < col; j++ {
// 			if j == 0{
// 				data[i][j] = items[i-1]
// 			}

// 			if j == 1 {
// 				for _, els := range keysInSlice {
// 					if items[i-1] == els{
// 						product := productPrice[els]
// 						data[i][j] = strconv.FormatFloat(product, 'f', 2, 64 )
// 					}
// 				}
// 			}

// 			if j == 2 {
// 				for _, els := range keysInSlice {
// 					if items[i-1] == els{
// 						product := productPrice[els] * 0.9
// 						data[i][j] = strconv.FormatFloat(product, 'f', 2, 64 )
// 					}
// 				}
// 			}

// 		} 
// 	}

// 	return data	
// }

//--------AI MODIFICATION------- 
func DataCollation(items []string) [][]string {
	rows := len(items) + 1
	data := make([][]string, rows)

	for i := range data {
		data[i] = make([]string, 3)
	}

	data[0] = []string{"Product-Name", "Initial-Price", "Discounted-Price"}

	for i, item := range items {
		if price, ok := productPrice[item]; ok {
			data[i+1][0] = item
			data[i+1][1] = strconv.FormatFloat(price, 'f', 2, 64)
			data[i+1][2] = strconv.FormatFloat(price*0.9, 'f', 2, 64)
		}
	}

	return data
}

// TableConstructor converts 2D slice of strings into a formatted table string
func TableConstructor(data [][]string) string {
	if len(data) == 0 {
		return ""
	}

	cols := len(data[0])
	colWidths := make([]int, cols)

	// Step 1: calculate max width for each column
	for _, row := range data {
		for j, cell := range row {
			if len(cell) > colWidths[j] {
				colWidths[j] = len(cell)
			}
		}
	}

	var sb strings.Builder

	// Step 2: build each row
	for i, row := range data {
		sb.WriteString("|")
		for j, cell := range row {
			// left-align the content
			sb.WriteString(" " + fmt.Sprintf("%-*s", colWidths[j], cell) + " |")
		}
		sb.WriteString("\n")

		// After header row, add separator
		if i == 0 {
			sb.WriteString("|")
			for _, w := range colWidths {
				sb.WriteString(strings.Repeat("-", w+2) + "|")
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

