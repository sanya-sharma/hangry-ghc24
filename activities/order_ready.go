package activities

import (
    "context"
    "time"
    "go.uber.org/cadence/activity"
)

func OrderReady(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderReady activity started")
	return "Order " + name + " Ready!", nil
}
