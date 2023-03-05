package backoffice

import "time"

type LineItem struct {
	SKU    string
	Amount int
	Price  int // ¢
}

type Invoice struct {
	ID       string
	Time     time.Time
	Customer string
	Items    []LineItem
}
