package pojo

import (
	"encoding/json"
	"io"
	"time"
)

type Products []*Product

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	SKU         string  `json:"sku" validate:"sku"`
	CreatedDate string  `json:"-"`
	DeletedDate string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProductList() []*Product {
	return productList
}
func AddProduct(newProduct *Product) {
	newProduct.ID = generatePK()
	productList = append(productList, newProduct)
}

func UpdateProduct(id int, product *Product) error {
	_, position, err := findProductById(id)
	if err != nil {
		return err
	}
	product.ID = id
	productList[position] = product
	return nil
}

func findProductById(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, nil
}
func generatePK() int {
	last := productList[len(productList)-1]
	return last.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedDate: time.Now().UTC().String(),
		DeletedDate: time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedDate: time.Now().UTC().String(),
		DeletedDate: time.Now().UTC().String(),
	},
}
