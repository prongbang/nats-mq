package app2

import "log"

// Handler is the interface
type Handler interface {
	SubNotiPromotion()
}

type handler struct {
	UCase UseCase
}

// NewHandler is new instance
func NewHandler(uc UseCase) Handler {
	return &handler{
		UCase: uc,
	}
}

func (h *handler) SubNotiPromotion() {
	err := h.UCase.SubNotiPromotion(func(p *Promotion) {
		log.Println("Subscribe <= ", p)
	})

	if err != nil {
		log.Println("err <= ", err)
	}
}
