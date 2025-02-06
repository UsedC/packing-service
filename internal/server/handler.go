package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/UsedC/packing-service/internal/app"
)

type Handler struct {
	app *app.App
}

func NewHandler(app *app.App) *Handler {
	return &Handler{app: app}
}

func (s *Handler) HandleGetPackSizes(w http.ResponseWriter, r *http.Request) {
	slog.Info("HandleGetPackSizes request", slog.String("method", r.Method))

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	packSizes := s.app.GetPackSizes(r.Context())

	response, err := json.Marshal(packSizes)
	if err != nil {
		slog.Error("error marshalling pack sizes", slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		slog.Error("error writing response", slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *Handler) HandleSetPackSizes(w http.ResponseWriter, r *http.Request) {
	slog.Info("HandleSetPackSizes request", slog.String("method", r.Method))

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SetPackSizesRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error decoding request", slog.String("error", err.Error()))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = req.Validate()
	if err != nil {
		slog.Info("error validating request", slog.String("error", err.Error()), slog.Any("request", req))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	s.app.SetPackSizes(r.Context(), req.PackSizes)

	w.WriteHeader(http.StatusOK)
}

func (s *Handler) HandleCalculatePacks(w http.ResponseWriter, r *http.Request) {
	slog.Info("HandleCalculatePacks request", slog.String("method", r.Method))

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculatePacksRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error decoding request", slog.String("error", err.Error()))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = req.Validate()
	if err != nil {
		slog.Info("error validating request", slog.String("error", err.Error()), slog.Any("request", req))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	packs, err := s.app.CalculatePacks(r.Context(), req.Total)
	if err != nil {
		slog.Error("error calculating packs", slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(packs)
	if err != nil {
		slog.Error("error marshalling packs", slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		slog.Error("error writing response", slog.String("error", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
