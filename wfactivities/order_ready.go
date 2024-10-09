package wfactivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"time"
)

func OrderReady(ctx context.Context, dish string, customer string, shouldFail bool) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderReady activity started")
	time.Sleep(10 * time.Second)
	return dish + " Order ready for " + customer + "!", nil
}
