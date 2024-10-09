package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func DeliveryConfirmation(ctx context.Context, dish string, customer string, shouldFail bool) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("DeliveryConfirmation activity started")
	time.Sleep(10 * time.Second)
	return "Order delivered to " + customer + "!", nil
}
