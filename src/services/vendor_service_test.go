package services_test

import (
	"coop-connect-go/src/config"
	"coop-connect-go/src/services"
	"coop-connect-go/src/types"
	"reflect"
	"testing"
)

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
				ID:   "1",
				Name: "Vendor 1",
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
			if !reflect.DeepEqual(vendor.Products, tt.expected.Products) {
				t.Errorf("expected vendor products %v, got %v", tt.expected.Products, vendor.Products)
			}
		})
	}
}
