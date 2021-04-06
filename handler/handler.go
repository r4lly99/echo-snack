package handler

import "github.com/r4lly99/echo-snack/user"

type Handler struct {
	userService user.Service
}

func NewHandler(us user.Service) *Handler {
	return &Handler{
		userService: us,
	}
}
