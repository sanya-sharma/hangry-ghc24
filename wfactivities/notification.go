package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func Notification(ctx context.Context, dish string, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Notification activity started")
	return "Notification for " + customer + "'s "+ dish + " sent!", nil
}
