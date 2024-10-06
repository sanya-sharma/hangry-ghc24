package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func DeliveryConfirmation(ctx context.Context, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("DeliveryConfirmation activity started")
	time.Sleep(10 * time.Second)
	return "Order delivered to " + customer + "!", nil
}
