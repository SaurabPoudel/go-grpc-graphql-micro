package order

import "time"

type service interface {
	PostOrder
	GetOrdersForAccount
}

type Order struct {
	ID         string
	CreatedAt  time.Time
	TotalPrice float64
	AccountID  string
	Products   []OrderedProduct
}