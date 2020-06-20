package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

type (
	loggerKey struct {}
)

func GetLogger(ctx context.Context) *logrus.Entry {
	var logger *logrus.Entry

	existingLogger := ctx.Value(loggerKey{})
	if existingLogger == nil {
		logger = logrus.NewEntry(logrus.StandardLogger())
	} else {
		logger = existingLogger.(*logrus.Entry)
	}
	return logger.WithContext(ctx)
}
