package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func OrderCollected(ctx context.Context, dish string, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderCollected activity started")
	return dish + " Order collected for " + customer + "!", nil
}
