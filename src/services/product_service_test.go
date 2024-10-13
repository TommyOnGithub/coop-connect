package services_test

import (
	"testing"

	"coop-connect-go/src/config"
	"coop-connect-go/src/services"
	"coop-connect-go/src/types"
)

func TestGetById(t *testing.T) {
	productService, err := services.NewProductService(config.APP_ROOT + "/data/products.json")
	if err != nil {
		t.Errorf("failed to create product service: %v", err)
	}

	tests := []struct {
		id       string
		expected types.Product
		exists   bool
	}{
		{"1", types.Product{ID: "1", Name: "Product 1", VendorID: "1"}, true},
		{"2", types.Product{ID: "2", Name: "Product 2", VendorID: "1"}, true},
		{"3", types.Product{ID: "3", Name: "Product 3", VendorID: "2"}, true},
	}

	for _, test := range tests {
		result, exists := productService.GetById(test.id)
		if exists != test.exists {
			t.Errorf("Expected exists to be %v, got %v", test.exists, exists)
		}
		if result != test.expected {
			t.Errorf("Expected result to be %v, got %v", test.expected, result)
		}
	}
}
