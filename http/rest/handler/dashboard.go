package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/dashboard/service"
	"mitra-kirim-be-mgmt/pkg/response"
	"net/http"
)

type DashboardHandler struct {
	Svc *service.Dashboard
	Log *logrus.Logger
}

func NewDashboardHandler(svc *service.Dashboard, log *logrus.Logger) *DashboardHandler {
	return &DashboardHandler{
		Svc: svc,
		Log: log,
	}
}

func (h *DashboardHandler) Get(c echo.Context) error {
	var req []string
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "failed to parse request body")
	}

	ctx := c.Request().Context()

	res, err := h.Svc.Get(ctx)
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
