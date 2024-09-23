package activities

import (
    "context"
    "time"
    "go.uber.org/cadence/activity"
)

func OrderReceived(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderReceived activity started")
	return "Order " + name + " Received!", nil
}
