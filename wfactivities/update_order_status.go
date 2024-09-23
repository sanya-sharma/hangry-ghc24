package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func UpdateOrderStatus(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("UpdateOrderStatus activity started")
	return "UpdateOrderStatus " + name + " Received!", nil
}
