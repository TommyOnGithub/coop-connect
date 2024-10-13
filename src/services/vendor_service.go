package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"coop-connect-go/src/config"
	"coop-connect-go/src/types"
)

type VendorService struct {
	vendors        map[string]types.Vendor
	productService ProductService
	filePath       string
}

func NewVendorService(filePath string) (*VendorService, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var vendors []types.Vendor
	err = json.Unmarshal(byteValue, &vendors)
	if err != nil {
		return nil, err
	}

	vendorMap := make(map[string]types.Vendor)
	for _, vendor := range vendors {
		vendorMap[vendor.ID] = vendor
	}

	productService, err := NewProductService(config.APP_ROOT + "/data/products.json")
	if err != nil {
		return nil, err
	}

	return &VendorService{vendors: vendorMap, productService: *productService, filePath: filePath}, nil
}

func (vs *VendorService) GetById(id string) types.Vendor {
	log.Printf("ID: %s, Vendors: %v", id, vs.vendors)
	vendor, exists := vs.vendors[id]
	if exists {
		return vendor
	}
	return types.Vendor{}
}

func (vs *VendorService) GetVendorByProductID(id string) (types.Vendor, error) {
	log.Printf("Product ID: %s", id)
	product, exists := vs.productService.GetById(id)
	if !exists {
		return types.Vendor{}, errors.New("product not found")
	}

	log.Printf("Product ID: %s", id)
	vendor := vs.GetById(product.VendorID)
	if vendor.ID == "" {
		return types.Vendor{}, errors.New("vendor not found")
	}
	return vendor, nil
}

func (vs *VendorService) getVendorValues() []types.Vendor {
	values := make([]types.Vendor, 0, len(vs.vendors))
	for _, value := range vs.vendors {
		values = append(values, value)
	}
	return values
}

func (vs *VendorService) commitVendorsToFile() error {
	file, err := os.Create(vs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := json.Marshal(vs.getVendorValues())
	if err != nil {
		return err
	}

	_, err = file.Write(byteValue)
	if err != nil {
		return err
	}

	return nil
}

func (vs *VendorService) Insert(vendor types.Vendor) types.Vendor {
	vendor.ID = generateRandomID()
	vs.vendors[vendor.ID] = vendor
	err := vs.commitVendorsToFile()
	if err != nil {
		log.Printf("Failed to commit vendors to file: %v", err)
	}
	return vendor
}

func generateRandomID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
