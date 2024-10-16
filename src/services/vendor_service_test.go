package services_test

import (
	"coop-connect-go/src/config"
	"coop-connect-go/src/services"
	"coop-connect-go/src/types"
	"testing"
)

func TestCreateVendor(t *testing.T) {
	tests := []struct {
		name     string
		vendor   types.Vendor
		expected types.Vendor
	}{
		{
			name: "Basic test",
			vendor: types.Vendor{
				ID:    "",
				Name:  "Vendor 3",
				Email: "vendor3@gmail.com",
			},
			expected: types.Vendor{
				Name:  "Vendor 3",
				Email: "vendor3@gmail.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vendorService, err := services.NewVendorService(config.APP_ROOT + "/data/vendors.json")
			if err != nil {
				t.Errorf("failed to create vendor service: %v", err)
			}
			vendor := vendorService.Insert(tt.vendor)
			if vendor.Name != tt.expected.Name {
				t.Errorf("expected vendor name %s, got %s", tt.expected.Name, vendor.Name)
			}
			if vendor.Email != tt.expected.Email {
				t.Errorf("expected vendor email %s, got %s", tt.expected.Email, vendor.Email)
			}
		})
	}
}

func TestGetVendorById(t *testing.T) {
	tests := []struct {
		name     string
		vendorId string
		expected types.Vendor
	}{
		{
			name:     "Vendor exists",
			vendorId: "1",
			expected: types.Vendor{
				ID:    "1",
				Name:  "Vendor 1",
				Email: "vendor1@gmail.com",
			},
		},
		{
			name:     "Vendor does not exist",
			vendorId: "999",
			expected: types.Vendor{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vendorService, err := services.NewVendorService(config.APP_ROOT + "/data/vendors.json")
			if err != nil {
				t.Errorf("failed to create vendor service: %v", err)
			}
			vendor := vendorService.GetById(tt.vendorId)
			if vendor.ID != tt.expected.ID {
				t.Errorf("expected vendor ID %s, got %s", tt.expected.ID, vendor.ID)
			}
			if vendor.Name != tt.expected.Name {
				t.Errorf("expected vendor name %s, got %s", tt.expected.Name, vendor.Name)
			}
			if vendor.Email != tt.expected.Email {
				t.Errorf("expected vendor email %s, got %s", tt.expected.Email, vendor.Email)
			}
		})
	}
}

func TestGetVendorByProductId(t *testing.T) {
	tests := []struct {
		name      string
		productId string
		expected  types.Vendor
	}{
		{
			name:      "Basic test",
			productId: "1",
			expected: types.Vendor{
				ID:    "1",
				Name:  "Vendor 1",
				Email: "vendor1@gmail.com",
			},
		},
		{
			name:      "Basic test 2",
			productId: "3",
			expected: types.Vendor{
				ID:    "2",
				Name:  "Vendor 2",
				Email: "vendor2@gmail.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vendorService, err := services.NewVendorService(config.APP_ROOT + "/data/vendors.json")
			if err != nil {
				t.Errorf("failed to create vendor service: %v", err)
			}
			vendor, err := vendorService.GetVendorByProductID(tt.productId)
			if err != nil {
				t.Errorf("failed to get vendor by product id: %v", err)
			}
			if vendor.ID != tt.expected.ID {
				t.Errorf("expected vendor ID %s, got %s", tt.expected.ID, vendor.ID)
			}
			if vendor.Name != tt.expected.Name {
				t.Errorf("expected vendor name %s, got %s", tt.expected.Name, vendor.Name)
			}
		})
	}
}
