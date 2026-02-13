package orders

import (
	"net/http"

	"github.com/napat/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	json.Write(w, http.StatusCreated, nil)
}
