package server

import (
	"context"
	"encoding/json"
	"github.com/patyukin/go-online-library/internal/converter"
	"github.com/patyukin/go-online-library/internal/repository/model"
	"github.com/patyukin/go-online-library/internal/server/dto"
	"github.com/patyukin/go-online-library/internal/service"
	"github.com/pkg/errors"
	"net/http"
)

type PromotionService interface {
	CreatePromotion(ctx context.Context, promotion model.Promotion, filters []model.Filter, directories []model.Directory) error
}

type Handler struct {
	promotionService PromotionService
}

func NewHandler(promotionService PromotionService) *Handler {
	return &Handler{
		promotionService: promotionService,
	}
}

func (h *Handler) CreatePromotionHandler(w http.ResponseWriter, r *http.Request) {
	var promotion dto.Promotion

	_ = json.NewDecoder(r.Body).Decode(&promotion)
	p := converter.ToPromotionFromDTO(&promotion)
	err := h.promotionService.CreatePromotion(r.Context(), p, converter.ToFiltersFromDTO(promotion.Filters), converter.ToDirectoriesFromDTO(promotion.Directories))
	if err != nil {
		h.SendResponseError(w, err)
		return
	}
}

func (h *Handler) SendResponseError(w http.ResponseWriter, err error) {
	var appErr *service.AppError
	if errors.As(err, &appErr) {
		if appErr.Code == service.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), 501)
		return
	}

	//
	http.Error(w, http.StatusText(http.StatusInternalServerError), 501)
	return
}

func (h *Handler) GetFilterHandler(w http.ResponseWriter, r *http.Request) {
	// unmarshal
	// validate
	// logic usecases
}

func (h *Handler) SetFilterHandler(w http.ResponseWriter, r *http.Request) {
	var promotion dto.Promotion
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// usecase
	r, err := sdfsdfs
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_ = w.Write([]byte(r))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
