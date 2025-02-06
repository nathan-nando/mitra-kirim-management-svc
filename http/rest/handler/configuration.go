package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/internal/configuration/service"
	fileUploaderSvc "mitra-kirim-be-mgmt/internal/file-uploader/service"
	"mitra-kirim-be-mgmt/pkg/response"
	"net/http"
)

type ConfigurationHandler struct {
	Svc *service.Configuration
	Log *logrus.Logger
}

func NewConfigurationHandler(svc *service.Configuration, fileSvc *fileUploaderSvc.FileUploader, log *logrus.Logger) *ConfigurationHandler {
	return &ConfigurationHandler{
		Svc: svc,
		Log: log,
	}
}

func (h *ConfigurationHandler) ListByTypes(c echo.Context) error {
	var req []string
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success:   false,
			Message:   "TEST",
			RequestID: "TEST",
			Internal:  err,
		})
	}

	res, err := h.Svc.GetByTypes(req)
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

func (h *ConfigurationHandler) UpdateApp(c echo.Context) error {
	var req model.UpdateAppRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "Validation error")
	}

	res, err := h.Svc.UpdateApp(req)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}

	return response.SuccessOK(c, res)
}
func (h *ConfigurationHandler) UpdateAppLogo(c echo.Context) error {

	appLogo, err := c.FormFile("appLogo")
	if err != nil || appLogo.Size == 0 {
		h.Log.Error(err)
		return response.ErrorBadRequest(c, err, "Validation error")
	}

	req := model.UpdateAppLogoRequest{
		AppLogo: appLogo,
	}

	res, err := h.Svc.UpdateLogoApp(&req)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}

	return response.SuccessOK(c, res)
}
func (h *ConfigurationHandler) UpdateSocial(c echo.Context) error {
	var req model.UpdateSocialRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "Validation error")
	}

	res, err := h.Svc.UpdateSocial(req)

	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}

	return response.SuccessOK(c, res)
}
func (h *ConfigurationHandler) UpdateToko(c echo.Context) error {
	var req model.UpdateTokoRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "Validation error")
	}

	res, err := h.Svc.UpdateToko(req)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}

	return response.SuccessOK(c, res)
}
