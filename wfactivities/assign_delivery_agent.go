package wfactivities

import (
	"context"
	"errors"
	"go.uber.org/cadence/activity"
	"time"
)

func AssignDeliveryAgent(ctx context.Context, customer string, shouldFail bool) (string, error) {
	if shouldFail {
		return "", errors.New("AssignDeliveryAgent flow failed")
	}
	logger := activity.GetLogger(ctx)
	logger.Info("AssignDeliveryAgent activity started")
	time.Sleep(10 * time.Second)
	return "Delivery agent for " + customer + " found!", nil
}
