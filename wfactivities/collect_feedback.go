package wfactivities

import (
    "context"
    "go.uber.org/cadence/activity"
)

func CollectFeedback(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CollectFeedback activity started")
	return "CollectFeedback for " + name + " started!", nil
}
