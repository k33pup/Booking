package notificationsvc

import (
	"booking/notification_svc/internal/pkg/app"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	nService, err := app.NewNotificationService()
	if err != nil {
		panic(err)
	}

	if err := nService.Start(ctx); err != nil {
		panic(err)
	}

	cancel()
}
