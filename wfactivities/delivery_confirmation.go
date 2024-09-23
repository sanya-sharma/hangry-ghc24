package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func DeliveryConfirmation(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("DeliveryConfirmation activity started")
	return "DeliveryConfirmation for " + name + " started!", nil
}
