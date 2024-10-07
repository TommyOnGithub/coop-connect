package types

type CustomerOrder struct {
	LineItems   []LineItem
	TotalAmount float64
}
