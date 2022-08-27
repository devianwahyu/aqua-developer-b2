package model

import (
	"sort"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

type Products struct {
	Product []Product
}

type ProductInterface interface {
	GetMostProduct(money int)
	GetMinMaxProductPrice()
	GetProductBySpecificPrice(price int)
}

// Question 1
func (p Products) GetMostProduct(money int) ([]Product, int) {
	var result []Product
	total := 0

	sort.Slice(p.Product, func(i, j int) bool {
		return p.Product[i].Price < p.Product[j].Price
	})

	for _, value := range p.Product {
		if total < money {
			total += value.Price
			result = append(result, value)
		} else {
			break
		}
	}

	return result, len(result)
}

// Question 2
func (p Products) GetMinMaxProductPrice() ([]Product, int) {
	var max int = p.Product[0].Price
	var min int = p.Product[0].Price
	for _, value := range p.Product {
		if max < value.Price {
			max = value.Price
		}
		if min > value.Price {
			min = value.Price
		}
	}

	var result []Product
	for _, value := range p.Product {
		if value.Price == max || value.Price == min {
			result = append(result, value)
		}
	}

	return result, len(result)
}

// Question 3
func (p Products) GetProductBySpecificPrice(price int) ([]Product, int) {
	var result []Product

	for _, value := range p.Product {
		if value.Price == price {
			result = append(result, value)
		}
	}

	return result, len(result)
}
