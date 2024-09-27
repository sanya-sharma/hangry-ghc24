package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func ValidateOrder(ctx context.Context, dish string, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ValidateOrder activity started")
	return dish + "Order validated for " + customer + "!", nil
}
