package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func ValidateOrder(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ValidateOrder activity started")
	return "ValidateOrder for " + name + " found!", nil
}
