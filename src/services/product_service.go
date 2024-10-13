package services

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"coop-connect-go/src/types"
)

type ProductService struct {
	products map[string]types.Product
}

func NewProductService(filePath string) (*ProductService, error) {
	log.Printf("File Path: %s", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var products []types.Product
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		return nil, err
	}

	productMap := make(map[string]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	return &ProductService{products: productMap}, nil
}

func (ps *ProductService) GetById(id string) (types.Product, bool) {
	log.Printf("ID: %s, Products: %v", id, ps.products)
	product, exists := ps.products[id]
	return product, exists
}

func (ps *ProductService) GetProductsByVendorID(vendorID string) []types.Product {
	products := []types.Product{}
	for _, product := range ps.products {
		if product.VendorID == vendorID {
			products = append(products, product)
		}
	}
	return products
}
