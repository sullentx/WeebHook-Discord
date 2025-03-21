package infraestructure

import (
	"fmt"
	"os"
	"weebhook/application"
	"weebhook/infraestructure/controller"
	"weebhook/infraestructure/repositories"

	"github.com/joho/godotenv"
)

func Init() (*controller.WebhookHandler, *controller.ReviewHandler, *controller.StatusHandler) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	payloadRepo := repositories.NewPayloadRepository()

	discordWebhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	payloadUseCase := application.NewPayloadUseCase(payloadRepo, discordWebhookURL)
	reviewUseCase := application.NewReviewUseCase(payloadRepo, discordWebhookURL)

	webhookHandler := controller.NewWebhookHandler(*payloadUseCase)
	reviewHandler := controller.NewReviewHandler(*reviewUseCase)
	statusHandler := controller.NewStatusHandler()
	return webhookHandler, reviewHandler, statusHandler
}
