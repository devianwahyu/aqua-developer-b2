package utils

import (
	"fmt"
	"week1/model"
)

func PrintProduct(message string, p []model.Product, count int) {
	fmt.Printf("%s\n------------------\n", message)
	for _, value := range p {
		fmt.Printf("Product %d\nName: %s\nPrice: %d\n------------------\n", value.ID, value.Name, value.Price)
	}
	fmt.Printf("Banyak barang: %d\n\n", count)
}
