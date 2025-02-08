package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/testimonial/model"
	"mitra-kirim-be-mgmt/internal/testimonial/service"
	"mitra-kirim-be-mgmt/pkg/converter"
	"mitra-kirim-be-mgmt/pkg/response"
)

type TestimonialHandler struct {
	Svc *service.Testimonial
	Log *logrus.Logger
}

func NewTestimonialHandler(svc *service.Testimonial, log *logrus.Logger) *TestimonialHandler {
	return &TestimonialHandler{
		Svc: svc,
		Log: log,
	}
}

func (h *TestimonialHandler) List(c echo.Context) error {
	limit := converter.GetQueryInt(c, "limit", 50)
	offset := converter.GetQueryInt(c, "offset", 50)

	res, err := h.Svc.GetAll(limit, offset)
	if err != nil {
		h.Log.Error(err)
		return response.ErrorInternal(c, err)
	}
	return response.SuccessOK(c, res)
}

func (h *TestimonialHandler) Create(c echo.Context) error {

	file, err := c.FormFile("img")
	name := c.FormValue("nama")
	description := c.FormValue("deskripsi")

	if err != nil || file.Size == 0 {
		h.Log.Error(err)
		return response.ErrorBadRequest(c, err, "Validation error")
	}

	if name == "" || description == "" {
		return response.ErrorBadRequest(c, err, "Validation error")
	}

	req := model.TestimonialCreate{
		Name:        name,
		Description: description,
		Img:         "",
	}

	userID := c.Get("userID").(string)

	res, err := h.Svc.Create(file, &req, userID)

	if err != nil {
		return response.ErrorUnauthorized(c, err, "Server error")
	}
	return response.SuccessOK(c, res)
}
