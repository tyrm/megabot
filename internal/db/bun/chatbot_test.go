package bun

import (
	"context"
	"fmt"
	"github.com/tyrm/megabot/internal/models"
	"github.com/tyrm/megabot/internal/testdata"
	"testing"
)

func TestChatbotDB_ReadChatbotServiceByID(t *testing.T) {
	client, err := testNewTestClient()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}

	tables := []struct {
		chatbotService *models.ChatbotService
		expectedConfig string
		exists         bool
	}{
		{
			chatbotService: testdata.TestChatbotServices[0],
			expectedConfig: testdata.TestChatbotServicesConfigs[0],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[1],
			expectedConfig: testdata.TestChatbotServicesConfigs[1],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[2],
			expectedConfig: testdata.TestChatbotServicesConfigs[2],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[3],
			expectedConfig: testdata.TestChatbotServicesConfigs[3],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[4],
			expectedConfig: testdata.TestChatbotServicesConfigs[4],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[5],
			expectedConfig: testdata.TestChatbotServicesConfigs[5],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[6],
			expectedConfig: testdata.TestChatbotServicesConfigs[6],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[7],
			expectedConfig: testdata.TestChatbotServicesConfigs[7],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[8],
			expectedConfig: testdata.TestChatbotServicesConfigs[8],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[9],
			expectedConfig: testdata.TestChatbotServicesConfigs[9],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[10],
			expectedConfig: testdata.TestChatbotServicesConfigs[10],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[11],
			expectedConfig: testdata.TestChatbotServicesConfigs[11],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[12],
			expectedConfig: testdata.TestChatbotServicesConfigs[12],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[13],
			expectedConfig: testdata.TestChatbotServicesConfigs[13],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[14],
			expectedConfig: testdata.TestChatbotServicesConfigs[14],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[15],
			expectedConfig: testdata.TestChatbotServicesConfigs[15],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[16],
			expectedConfig: testdata.TestChatbotServicesConfigs[16],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[17],
			expectedConfig: testdata.TestChatbotServicesConfigs[17],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[18],
			expectedConfig: testdata.TestChatbotServicesConfigs[18],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[19],
			expectedConfig: testdata.TestChatbotServicesConfigs[19],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[20],
			expectedConfig: testdata.TestChatbotServicesConfigs[20],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[21],
			expectedConfig: testdata.TestChatbotServicesConfigs[21],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[22],
			expectedConfig: testdata.TestChatbotServicesConfigs[22],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[23],
			expectedConfig: testdata.TestChatbotServicesConfigs[23],
			exists:         true,
		},
		{
			chatbotService: testdata.TestChatbotServices[24],
			expectedConfig: testdata.TestChatbotServicesConfigs[24],
			exists:         true,
		},
		{
			chatbotService: &models.ChatbotService{
				ID: 1000,
			},
			exists: false,
		},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running ReadChatbotServiceByID %v", i, table.chatbotService.ID)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			chatbotService, err := client.ReadChatbotServiceByID(context.Background(), table.chatbotService.ID)
			if err != nil {
				t.Errorf("[%d] got error reading chatbot service %d: %s", i, table.chatbotService.ID, err.Error())
				return
			}

			if table.exists {
				if chatbotService == nil {
					t.Errorf("[%d] expected chatbot service: got 'nil'", i)
					return
				}

				if chatbotService.ID != table.chatbotService.ID {
					t.Errorf("[%d] wrong id for chatbot service: got '%d', want '%d'", i, chatbotService.ID, table.chatbotService.ID)
				}
				if chatbotService.Description != table.chatbotService.Description {
					t.Errorf("[%d] wrong description for chatbot service: got '%s', want '%s'", i, chatbotService.Description, table.chatbotService.Description)
				}
				if chatbotService.ServiceType != table.chatbotService.ServiceType {
					t.Errorf("[%d] wrong type for chatbot service: got '%s', want '%s'", i, chatbotService.ServiceType, table.chatbotService.ServiceType)
				}

				config, err := chatbotService.GetConfig()
				if err != nil {
					t.Errorf("[%d] got error reading chatbot service config %d: %s", i, table.chatbotService.ID, err.Error())
					return
				}
				if config != table.expectedConfig {
					t.Errorf("[%d] wrong config for chatbot service: got '%s', want '%s'", i, chatbotService.Config, table.chatbotService.Config)
				}
			} else {
				if chatbotService != nil {
					t.Errorf("[%d] unexpected chatbot service: got '%T'", i, chatbotService)
					return
				}
			}

		})
	}
}
