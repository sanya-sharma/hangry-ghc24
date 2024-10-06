package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func OrderCollected(ctx context.Context, dish string, customer string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderCollected activity started")
	time.Sleep(10 * time.Second)
	return dish + " Order collected for " + customer + "!", nil
}
