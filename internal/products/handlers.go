package products

import (
	"log"
	"net/http"

	"github.com/napat/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewTextHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)

}
