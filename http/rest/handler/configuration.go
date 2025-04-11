package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"mitra-kirim-be-mgmt/internal/configuration/model"
	"mitra-kirim-be-mgmt/internal/configuration/service"
	locationService "mitra-kirim-be-mgmt/internal/location/service"
	testimonialService "mitra-kirim-be-mgmt/internal/testimonial/service"
	"mitra-kirim-be-mgmt/pkg/response"
	"net/http"
)

type ConfigurationHandler struct {
	Svc            *service.Configuration
	LocSvc         *locationService.Location
	TestimonialSvc *testimonialService.Testimonial
	Log            *logrus.Logger
}

func NewConfigurationHandler(svc *service.Configuration, testimonialSvc *testimonialService.Testimonial, locSvc *locationService.Location, log *logrus.Logger) *ConfigurationHandler {
	return &ConfigurationHandler{
		Svc:            svc,
		LocSvc:         locSvc,
		TestimonialSvc: testimonialSvc,
		Log:            log,
	}
}

func (h *ConfigurationHandler) ListByTypes(c echo.Context) error {
	var req []string
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "failed to parse request body")
	}

	ctx := c.Request().Context()

	res, err := h.Svc.GetByTypes(ctx, req)
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
func (h *ConfigurationHandler) PublicConfig(c echo.Context) error {
	var req []string

	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return response.ErrorBadRequest(c, err, "failed to parse request body")
	}

	configs, err := h.Svc.GetByTypes(ctx, req)
	if err != nil {
		return response.ErrorInternal(c, err, "Failed to get public config")
	}

	locations, err := h.LocSvc.GetAll(ctx, 20, 0)
	if err != nil {
		return response.ErrorInternal(c, err, "Failed to get location")
	}

	testimonials, err := h.TestimonialSvc.GetSlide(ctx, 100, 0)
	if err != nil {
		return response.ErrorInternal(c, err, "Failed to get testimonials")
	}

	return response.SuccessOK(c, map[string]interface{}{
		"config":       configs,
		"location":     locations,
		"testimonials": testimonials,
	})
}
func (h *ConfigurationHandler) UpdateAppLogo(c echo.Context) error {
	ctx := c.Request().Context()
	appLogo, err := c.FormFile("appLogo")
	if err != nil || appLogo.Size == 0 {
		h.Log.Error(err)
		return response.ErrorBadRequest(c, err, "Validation error")
	}

	req := model.UpdateAppLogoRequest{
		AppLogo: appLogo,
	}

	res, err := h.Svc.UpdateLogoApp(ctx, &req)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}

	return response.SuccessOK(c, res)
}
func (h *ConfigurationHandler) UpdateHero(c echo.Context) error {
	ctx := c.Request().Context()

	heroDesc := c.FormValue("heroDesc")
	heroImg, err := c.FormFile("heroImg")
	if err != nil {
		h.Log.Error(err)
		heroImg = &multipart.FileHeader{
			Filename: "",
			Size:     0,
		}
	}
	req := &model.UpdateHeroLogoRequest{
		HeroDesc: heroDesc,
		HeroImg:  heroImg,
	}

	res, err := h.Svc.UpdateHero(ctx, req)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}

	return response.SuccessOK(c, res)
}
func (h *ConfigurationHandler) UpdateServices(c echo.Context) error {
	ctx := c.Request().Context()

	servicesJson := c.FormValue("services")

	var services []model.ServicesLayout
	if err := json.Unmarshal([]byte(servicesJson), &services); err != nil {
		h.Log.Error(err)
		return response.ErrorBadRequest(c, "Failed to parse services json")
	}

	for i := range services {
		fileKey := fmt.Sprintf("serviceFile%d", i)

		file, err := c.FormFile(fileKey)

		if err != nil || file == nil {
			continue
		}

		filename, err := h.Svc.SaveServiceImage(ctx, file)
		if err != nil {
			h.Log.Error(err)
			return response.ErrorInternal(c, "Failed to save service file")
		}

		services[i].Img = filename
	}

	updatedServicesJson, err := json.Marshal(services)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, "Failed to parse updated services json")
	}

	payload := &model.UpdateServicesRequest{Services: string(updatedServicesJson)}

	res, err := h.Svc.UpdateByKeys(ctx, payload)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}

	return response.SuccessOK(c, res)
}
func (h *ConfigurationHandler) UpdateConfiguration(c echo.Context) error {
	ctx := c.Request().Context()

	var req model.UpdateConfigurationRequest
	if err := c.Bind(&req); err != nil {
		c.Logger().Error("failed to parse request body")
		return response.ErrorBadRequest(c, err, "Validation error")
	}

	res, err := h.Svc.UpdateByKeys(ctx, req)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}

	return response.SuccessOK(c, res)
}
