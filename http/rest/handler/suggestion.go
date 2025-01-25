package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/suggestion/model"
	"mitra-kirim-be-mgmt/internal/suggestion/service"
	"mitra-kirim-be-mgmt/pkg/response"
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
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success:   false,
			Message:   "TEST",
			RequestID: "TEST",
			Internal:  err,
		})
	}
	return response.SuccessOK(c, res)
}

func (h *SuggestionHandler) Create(c echo.Context) error {
	var req model.SuggestionCreate
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "TEST",
			RequestID: "TEST",
			Internal:  err,
		})
	}

	res, err := h.Svc.Create(&req)
	if err != nil {
		h.Log.Error(err)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *SuggestionHandler) ReplyEmail(c echo.Context) error {
	var req model.SuggestionReply
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "TEST",
			RequestID: "TEST",
			Internal:  err,
		})
	}

	//res, err := h.Svc.ReplyEmail(&req)
	//if err != nil {
	//	h.Log.Error(err)
	//}
	return c.JSON(http.StatusOK, "OK")
}
