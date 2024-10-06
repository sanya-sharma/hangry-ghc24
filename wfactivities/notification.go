package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func Notification(ctx context.Context, dish string, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Notification activity started")
	time.Sleep(10 * time.Second)
	return "Notification for " + customer + "'s " + dish + " sent!", nil
}
