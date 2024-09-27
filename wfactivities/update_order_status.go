package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func UpdateOrderStatus(ctx context.Context, dish string, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("UpdateOrderStatus activity started")
	return "order status updated for " + customer + "'s " + dish, nil
}
