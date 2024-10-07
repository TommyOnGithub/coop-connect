package services

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"coop-connect-go/src/types"
)

type VendorService struct {
	vendors map[string]types.Vendor
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

	return &VendorService{vendors: vendorMap}, nil
}

func (vs *VendorService) GetVendorByProductID(id string) (types.Vendor, error) {
	vendor := types.Vendor{}
	for _, v := range vs.vendors {
		for _, product := range v.Products {
			if product.ID == id {
				vendor = v
				break
			}
		}
	}
	if vendor.ID == "" {
		return types.Vendor{}, errors.New("vendor not found")
	}
	return vendor, nil
}
