package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func OrderInTransit(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderInTransit activity started")
	return "OrderInTransit for " + name + " found!", nil
}
