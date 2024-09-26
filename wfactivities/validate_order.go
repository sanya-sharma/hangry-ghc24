package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
)

func ValidateOrder(ctx context.Context, name string) (string, error) {

	info := activity.GetInfo(ctx)

	// Get a logger to log retry attempts
	logger := activity.GetLogger(ctx)
	logger.Info("ValidateOrder activity started")

	// Simulate a failure for first 2 attempts and succeed on the 3rd attempt
	if info.Attempt < 3 {
		logger.Info("ValidateOrder activity failed")
		return "", nil
	}

	// Succeed on the 3rd attempt
	logger.Info("ValidateOrder activity succeeded")
	return "ValidateOrder for " + name + " found!", nil
}
