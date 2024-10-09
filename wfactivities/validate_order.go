package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func ValidateOrder(ctx context.Context, dish string, customer string, shouldFail bool) (string, error) {

	// Get a logger to log retry attempts
	logger := activity.GetLogger(ctx)
	logger.Info("ValidateOrder activity started")
	time.Sleep(5 * time.Second)

	// Simulate a failure for first 2 attempts and succeed on the 3rd attempt

	logger.Info("ValidateOrder activity succeeded")
	return "Order validated for " + customer + "!", nil
}
