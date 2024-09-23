package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func OrderCollected(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderCollected activity started")
	return "Order " + name + " Collected!", nil
}
