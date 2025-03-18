package service

import (
	"context"
	"mitra-kirim-be-mgmt/pkg/contants"
	"time"
)

func (s *Testimonial) GetSlide(context context.Context, limit int, offset int) ([]string, error) {
	//list, err := s.CacheSvc.LRange(context, contants.CacheTestimonials, 0, -1)
	//if err == nil && len(list) > 0 {
	//	s.Logger.Info("Testimonials CACHE")
	//	return list, nil
	//}
	//if err != nil {
	//	return []string{}, err
	//}

	testimonials, err := s.Repository.GetSlide(limit, offset)
	s.Logger.Info("Testimonials DB")

	if err != nil {
		return []string{}, err
	}
	if len(testimonials) == 0 {
		return []string{}, nil
	}

	results := make([]string, 0, len(testimonials))

	pipe := s.CacheSvc.Pipeline()

	for _, testimonial := range testimonials {
		results = append(results, testimonial.Img)
	}

	pipe.RPush(context, contants.CacheTestimonials, results)
	pipe.Expire(context, contants.CacheTestimonials, time.Duration(s.CacheTime)*time.Minute)
	_, err = pipe.Exec(context)
	if err != nil {
		return []string{}, err
	}

	return results, nil
}
