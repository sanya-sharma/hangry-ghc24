package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func DeliveryConfirmation(ctx context.Context, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("DeliveryConfirmation activity started")
	return "Order delivered to " + customer + "!", nil
}
