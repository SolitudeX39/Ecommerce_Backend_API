package orders

import (
	"context"

	repo "github.com/napat/ecom/internal/adapters/postgresql/sqlc"
)

type orderItem struct {
	ProductID int64 `json:"product_id"`

	Quantity int64 `json:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}
