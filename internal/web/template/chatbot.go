package template

import "github.com/tyrm/megabot/internal/models"

// ChatbotTemplate contains the variables for the "chatbot" template.
type ChatbotTemplate struct {
	Common

	Sidebar Sidebar
}

// ChatbotServiceTemplate contains the variables for the "chatbot_service" template.
type ChatbotServiceTemplate struct {
	Common

	ChatbotServices           *[]models.ChatbotService
	ChatbotServicesPagination Pagination
	Sidebar                   Sidebar
}
