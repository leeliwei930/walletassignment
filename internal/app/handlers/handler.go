package handlers

import (
	"github.com/leeliwei930/walletassignment/internal/interfaces"
)

type Handler struct {
	app interfaces.Application
}

func NewHandler(app interfaces.Application) *Handler {
	return &Handler{app: app}
}
