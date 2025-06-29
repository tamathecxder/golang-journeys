package stdlib

import (
	"fmt"
	"sort"
)

type Product struct {
	Name     string
	Quantity int
}

type ProductList []Product

func (pl ProductList) Len() int {
	return len(pl)
}

func (pl ProductList) Less(i, j int) bool {
	return pl[i].Quantity < pl[j].Quantity
}

func (pl ProductList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func Stdlib_Sort() {
	products := ProductList{
		Product{Name: "Mineral Water", Quantity: 230},
		Product{Name: "Tapioca", Quantity: 12240},
		Product{Name: "Coffee Bean", Quantity: 25},
		Product{Name: "Straw", Quantity: 1521},
		Product{Name: "Milk", Quantity: 149},
		Product{Name: "Matcha Powder", Quantity: 9},
		Product{Name: "Packaging", Quantity: 4519},
	}

	fmt.Println("Before sorted", products)

	sort.Sort(ProductList(products))

	fmt.Println("After sorted", products)
}
