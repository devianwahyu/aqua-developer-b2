package main

import (
	"strconv"
	"week1/model"
	"week1/utils"
)

func main() {
	products := model.Products{
		Product: []model.Product{
			{
				ID:    1,
				Name:  "Benih Lele",
				Price: 50000,
			}, {
				ID:    2,
				Name:  "Pakan Lele Cap Menara",
				Price: 25000,
			}, {
				ID:    3,
				Name:  "Probiotik A",
				Price: 75000,
			}, {
				ID:    4,
				Name:  "Probiotik Nila B",
				Price: 10000,
			}, {
				ID:    5,
				Name:  "Pakan Nila",
				Price: 20000,
			}, {
				ID:    6,
				Name:  "Benih Nila A",
				Price: 20000,
			}, {
				ID:    7,
				Name:  "Cupang",
				Price: 5000,
			}, {
				ID:    8,
				Name:  "Benih Nila B",
				Price: 30000,
			}, {
				ID:    9,
				Name:  "Benih Cupang",
				Price: 10000,
			}, {
				ID:    10,
				Name:  "Probiotik B",
				Price: 10000,
			},
		},
	}

	// Print Product Petani Can Buy
	money := 100000
	listMostProduct, countMostProduct := products.GetMostProduct(money)
	utils.PrintProduct("List Produk yang terbeli dengan "+strconv.Itoa(money), listMostProduct, countMostProduct)

	//Print Min and Max Price
	listMinMaxProduct, countMinMaxProduct := products.GetMinMaxProductPrice()
	utils.PrintProduct("List Produk dengan Harga Minimal dan Maksimal ", listMinMaxProduct, countMinMaxProduct)

	// Print Product With Specific Price
	price := 10000
	listProductSpecificPrice, countSpecificPrice := products.GetProductBySpecificPrice(price)
	utils.PrintProduct("List Produk dengan harga "+strconv.Itoa(price), listProductSpecificPrice, countSpecificPrice)
}
