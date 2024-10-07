package services_test

import (
	"coop-connect-go/src/config"
	"coop-connect-go/src/services"
	"coop-connect-go/src/types"
	"reflect"
	"testing"
)

func slicesEqual(a, b []types.PurchaseOrder) bool {
	return reflect.DeepEqual(a, b)
}

func TestConvertOrderToPurchaseOrders(t *testing.T) {
	tests := []struct {
		name     string
		order    types.CustomerOrder
		expected []types.PurchaseOrder
	}{
		{
			name: "Basic conversion",
			order: types.CustomerOrder{
				LineItems: []types.LineItem{
					{
						ProductID: "1",
						Quantity:  5,
					},
					{
						ProductID: "2",
						Quantity:  5,
					},
				},
				TotalAmount: 100,
			},
			expected: []types.PurchaseOrder{
				{
					LineItems: []types.LineItem{
						{
							ProductID: "1",
							Quantity:  5,
						},
						{
							ProductID: "2",
							Quantity:  5,
						},
					},
					VendorID: "1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vendorService, err := services.NewVendorService(config.APP_ROOT + "/data/vendors.json")
			if err != nil {
				t.Errorf("failed to create vendor service: %v", err)
			}
			var orderPlexer = services.NewOrderPlexerService(*vendorService)
			result := orderPlexer.CreatePurchaseOrdersFromCustomerOrder(tt.order)
			if !slicesEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
