package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func GetLogger(ctx context.Context) *log.Entry {
	loggerValue := GetContextValue(ctx, "logger")

	logger := log.WithFields(log.Fields{})
	if loggerValue != nil {
		logger = loggerValue.(*log.Entry)
	}

	return logger
}
