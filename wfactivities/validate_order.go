package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
)

func ValidateOrder(ctx context.Context, name string) (string, error) {
	////////////////////////////////////////
	// TODO: show multiple retries and the last one should succeed
	////////////////////////////////////////
	logger := activity.GetLogger(ctx)
	logger.Info("ValidateOrder activity started")
	return dish + "Order validated for " + customer + "!", nil
}
