package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func CollectFeedback(ctx context.Context, dish string, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CollectFeedback activity started")
	return "Collecting feedback from " + customer + " for" + dish, nil
}
