package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/patyukin/go-online-library/internal/usecase"
	"github.com/patyukin/go-online-library/internal/usecase/model"
	"net/http"
)

type UseCase interface {
	CreatePromotion(ctx context.Context, promotion model.Promotion, filters []model.Filter, directories []model.Directory) error
}

type Handler struct {
	uc UseCase
}

func New(uc UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetPromotionHandler(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) DeletePromotionHandler(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) UpdatePromotionHandler(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) SendResponseError(w http.ResponseWriter, err error) {
	var appErr *usecase.AppError
	if errors.As(err, &appErr) {
		if appErr.Code == usecase.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusNotImplemented)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusNotImplemented)
	return
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func sendResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
