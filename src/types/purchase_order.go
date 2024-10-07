package types

type PurchaseOrder struct {
	LineItems []LineItem
	VendorID  string
}
