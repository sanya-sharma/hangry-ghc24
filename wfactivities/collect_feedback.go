package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func CollectFeedback(ctx context.Context, dish string, customer string, shouldFail bool) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CollectFeedback activity started")
	time.Sleep(10 * time.Second)
	return "Collecting feedback from " + customer + " for" + dish, nil
}
