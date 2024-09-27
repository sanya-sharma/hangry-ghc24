package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func OrderReady(ctx context.Context, dish string, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderReady activity started")
	return dish + " Order ready for " + customer + "!", nil
}
