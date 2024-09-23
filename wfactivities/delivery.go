package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func Delivery(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Delivery activity started")
	return "Delivery for " + name + " started!", nil
}
