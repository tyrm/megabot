package db

import (
	"context"
	"github.com/tyrm/megabot/internal/models"
)

// Chatbot contains functions related to the chatbot module.
type Chatbot interface {
	// ReadChatbotServiceByID returns one chatbot service.
	ReadChatbotServiceByID(ctx context.Context, id int64) (*models.ChatbotService, Error)
	// ReadChatbotServicesPage returns a page of chatbot services.
	ReadChatbotServicesPage(ctx context.Context, index, count int) (*[]models.ChatbotService, Error)
}
