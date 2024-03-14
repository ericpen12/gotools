package log

import "go.uber.org/zap"

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	logger, _ = zap.NewProduction(zap.AddCallerSkip(1))
	defer logger.Sync() // flushes buffer, if any
	sugar = logger.Sugar()
}
