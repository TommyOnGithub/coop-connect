package services

import (
	"coop-connect-go/src/types"
)

type OrderPlexerService struct {
	vendorService VendorService
}

func NewOrderPlexerService(vendorService VendorService) *OrderPlexerService {
	return &OrderPlexerService{vendorService: vendorService}
}

func (op *OrderPlexerService) CreatePurchaseOrdersFromCustomerOrder(order types.CustomerOrder) []types.PurchaseOrder {
	purchaseOrders := make(map[string]types.PurchaseOrder)
	for _, lineItem := range order.LineItems {
		vendor, err := op.vendorService.GetVendorByProductID(lineItem.ProductID)
		if err != nil {
			panic("Failed to get Vendor by ID: " + err.Error())
		}
		if _, exists := purchaseOrders[vendor.ID]; exists {
			var purchaseOrder = purchaseOrders[vendor.ID]
			purchaseOrder.LineItems = append(purchaseOrder.LineItems, lineItem)
			purchaseOrders[vendor.ID] = purchaseOrder
		} else {
			purchaseOrder := types.PurchaseOrder{
				LineItems: []types.LineItem{lineItem},
				VendorID:  vendor.ID,
			}
			purchaseOrders[vendor.ID] = purchaseOrder
		}
	}
	var result []types.PurchaseOrder
	for _, po := range purchaseOrders {
		result = append(result, po)
	}
	return result
}
