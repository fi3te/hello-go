package main

import (
	"strconv"
	"time"

	"go.uber.org/zap"
)

func demo(logger *zap.Logger) {
	logger.Info("hello world")
	logger.Error("error message " + strconv.Itoa(78))
	logger.Info("message with additional fields",
		zap.String("url", "https://..."),
		zap.Int64("size", 32768),
		zap.Duration("duration", time.Second),
		zap.Ints("array", []int{2, 4, 8}),
	)
}

func demoSugared(logger *zap.Logger) {
	sugar := logger.Sugar()
	sugar.Infow("test message", "key1", "value1", "key2", 3)
	sugar.Infof("formatted message: %s", "abc")
}
