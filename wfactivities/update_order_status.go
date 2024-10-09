package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func UpdateOrderStatus(ctx context.Context, dish string, customer string, shouldFail bool) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("UpdateOrderStatus activity started")
	time.Sleep(10 * time.Second)
	return "order status updated for " + customer + "'s " + dish, nil
}
