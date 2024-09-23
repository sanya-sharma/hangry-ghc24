package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func DeliveryAgent(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("DeliveryAgent activity started")
	return "DeliveryAgent for " + name + " found!", nil
}
