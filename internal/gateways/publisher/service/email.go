package service

import (
	"context"
	"encoding/json"
	"mitra-kirim-be-mgmt/internal/gateways/publisher/model"
	"mitra-kirim-be-mgmt/pkg/contants"
	"time"
)

func (s *Publisher) Publish(context context.Context, request *model.PublisherEmail) error {
	jsonMessage, err := json.Marshal(request)
	if err != nil {
		return err
	}

	result, err := s.client.Publish(context, contants.EmailChannelNotification, jsonMessage).Result()
	if err != nil {
		return err
	}

	if result == 0 {
		s.Logger.Println("Warning: No subscribers received the message.")
	} else {
		s.Logger.Printf("Message published successfully to %d subscribers\n", result)
	}
	return nil
}

func (s *Publisher) PublishWithRetry(context context.Context, request *model.PublisherEmail) {
	jsonMessage, err := json.Marshal(request)
	if err != nil {
		s.Logger.Println("Error parsing email message:", err)
	}

	for i := 0; i < s.MaxRetry; i++ {
		result, err := s.client.Publish(context, contants.EmailChannelNotification, jsonMessage).Result()
		if err != nil {
			s.Logger.Printf("Attempt %d: Failed to publish, error: %v\n", i+1, err)
			time.Sleep(2 * time.Second)
		} else if result == 0 {
			s.Logger.Println("Warning: No subscribers received the message.")
		} else {
			s.Logger.Printf("Message published successfully to %d subscribers\n", result)
			return
		}
	}
	s.Logger.Println("Failed to publish after maximum retries.")
	return
}
