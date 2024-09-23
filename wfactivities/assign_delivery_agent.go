package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func AssignDeliveryAgent(ctx context.Context, name string) (string, error) {
	////////////////////////////////////////
	// TODO: WF failure here
	////////////////////////////////////////
	logger := activity.GetLogger(ctx)
	logger.Info("AssignDeliveryAgent activity started")
	return "AssignDeliveryAgent for " + name + " found!", nil
}
