package app1

import (
	"encoding/json"
	"net/http"
)

// Handler is the interface
type Handler interface {
	NotiPromotion(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	UCase UseCase
}

// NewHandler is new instance
func NewHandler(uCase UseCase) Handler {
	return &handler{
		UCase: uCase,
	}
}

func (h *handler) NotiPromotion(w http.ResponseWriter, r *http.Request) {

	model := Promotion{
		ID:   1,
		Name: "อิ่มฟินสุดคุ้ม รับ 600 บาท ฟรี",
		Desc: "เมื่อทานครบ 2,000 บาท ขึ้นไป",
	}

	err := h.UCase.NotiPromotion(model)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Notification Fail.",
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Notification Success",
		})
	}
}
