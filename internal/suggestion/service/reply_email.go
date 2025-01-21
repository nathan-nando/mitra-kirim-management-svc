package service

import (
	"mitra-kirim-be-mgmt/internal/suggestion/model"
)

func (s *Suggestion) ReplyEmail(request *model.SuggestionReply) (string, error) {
	//suggestion, err := s.Repo.(request)
	//if err != nil {
	//	return "", err
	//}
	//
	//go func() {
	//	bgCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//	defer cancel()
	//	s.Publisher.PublishWithRetry(
	//		bgCtx,
	//		&modelPub.PublisherEmail{
	//			Name:    request.Name,
	//			Email:   request.Email,
	//			Message: request.Message,
	//		})
	//}()
	//
	//return suggestion.ID, nil
}
