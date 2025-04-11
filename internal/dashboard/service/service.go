package service

import (
	"github.com/sirupsen/logrus"
	locationService "mitra-kirim-be-mgmt/internal/location/service"
	suggestionService "mitra-kirim-be-mgmt/internal/suggestion/service"
	testimonialService "mitra-kirim-be-mgmt/internal/testimonial/service"
)

type Dashboard struct {
	LocSvc         *locationService.Location
	TestimonialSvc *testimonialService.Testimonial
	SuggestionSvc  *suggestionService.Suggestion
	Logger         *logrus.Logger
}

func NewDashboard(locSvc *locationService.Location, testimonialSvc *testimonialService.Testimonial, suggestionSvc *suggestionService.Suggestion, logger *logrus.Logger) *Dashboard {
	return &Dashboard{LocSvc: locSvc,
		TestimonialSvc: testimonialSvc,
		SuggestionSvc:  suggestionSvc,
		Logger:         logger,
	}
}
