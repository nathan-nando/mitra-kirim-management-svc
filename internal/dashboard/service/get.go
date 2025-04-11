package service

import (
	"context"
	"mitra-kirim-be-mgmt/internal/dashboard/model"
)

func (s *Dashboard) Get(ctx context.Context) (model.Dashboard, error) {
	var dashboard model.Dashboard
	suggestionCount, err := s.SuggestionSvc.Count(ctx)
	if err != nil {
		return dashboard, err
	}
	testimonialCount, err := s.TestimonialSvc.Count(ctx)
	if err != nil {
		return dashboard, err
	}

	locationCount, err := s.LocSvc.Count(ctx)
	if err != nil {
		return dashboard, err
	}

	dashboard = model.Dashboard{
		ViewerCount:      0,
		SuggestionCount:  suggestionCount,
		TestimonialCount: testimonialCount,
		LocationCount:    locationCount,
	}
	return dashboard, nil
}
