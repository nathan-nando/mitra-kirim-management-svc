package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/suggestion/service"
	"net/http"
)

type SuggestionHandler struct {
	Svc *service.Suggestion
	Log *logrus.Logger
}

func NewSuggestionHandler(svc *service.Suggestion, log *logrus.Logger) *SuggestionHandler {
	return &SuggestionHandler{
		Svc: svc,
		Log: log,
	}
}

func (h *SuggestionHandler) List(c echo.Context) error {
	res, err := h.Svc.Get(c.Request().Context())
	if err != nil {
		h.Log.Error(err)
	}
	return c.JSON(http.StatusOK, res)
}
