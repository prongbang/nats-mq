package app1

import (
	"fmt"
	"net/http"
)

// Route is the interface
type Route interface {
	Initial(port string)
}

type route struct {
	Handle Handler
}

// NewRoute is new instance
func NewRoute(handle Handler) Route {
	return &route{
		Handle: handle,
	}
}

func (r *route) Initial(port string) {
	http.HandleFunc("/api/v1/noti-promotion", r.Handle.NotiPromotion)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
