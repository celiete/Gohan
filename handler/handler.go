package handler

import (
	"fmt"
	"net/http"

	gh "github.com/bukalapak/gohan"
	"github.com/julienschmidt/httprouter"
)

// Handler controls request flow from client to service
type Handler struct {
	gohan *gh.Gohan
}

// Meta is used to consolidate all meta statuses
type Meta struct {
	HTTPStatus int `json:"http_status"`
}

// NewHandler returns a pointer of Handler instance
func NewHandler(gohan *gh.Gohan) *Handler {
	return &Handler{
		gohan: gohan,
	}
}

// Healthz is used to control the flow of GET /healthz endpoint
func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}
