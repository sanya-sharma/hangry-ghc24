package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func OrderInTransit(ctx context.Context, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderInTransit activity started")
	return "Order in transit for " + customer + "!", nil
}
