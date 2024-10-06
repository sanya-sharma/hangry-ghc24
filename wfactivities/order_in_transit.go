package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func OrderInTransit(ctx context.Context, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderInTransit activity started")
	time.Sleep(10 * time.Second)
	return "Order in transit for " + customer + "!", nil
}
