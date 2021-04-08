package handler

import (
	"github.com/r4lly99/echo-snack/service"
)

type Handler struct {
	userService service.Service
}

func NewHandler(us service.Service) *Handler {
	return &Handler{
		userService: us,
	}
}
