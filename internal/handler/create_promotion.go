package handler

import (
	"encoding/json"
	"github.com/patyukin/go-online-library/internal/handler/reqdto"
	"github.com/patyukin/go-online-library/internal/usecase/converter"
	"net/http"
)

func (h *Handler) CreatePromotionHandler(w http.ResponseWriter, r *http.Request) {
	var promotion reqdto.Promotion

	_ = json.NewDecoder(r.Body).Decode(&promotion)

	p := converter.ToPromotionModelFromReqDTO(promotion)
	f := converter.ToFiltersModelFromReqDTO(promotion.Filters)
	d := converter.ToDirectoriesModelFromReqDTO(promotion.Directories)
	err := h.uc.CreatePromotion(r.Context(), p, f, d)
	if err != nil {
		h.SendResponseError(w, err)
		return
	}
}
