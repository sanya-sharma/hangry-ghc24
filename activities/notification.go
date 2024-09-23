package activities

import (
    "context"
    "time"
    "go.uber.org/cadence/activity"
)

func Notification(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Notification activity started")
	return "Notification for " + name + " sent!", nil
}
