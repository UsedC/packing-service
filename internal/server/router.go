package server

import (
	"net/http"
)

func NewRouter(handler *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/pack-sizes", handler.HandleGetPackSizes)
	mux.HandleFunc("/set-pack-sizes", handler.HandleSetPackSizes)
	mux.HandleFunc("/calculate-packs", handler.HandleCalculatePacks)

	return mux
}
