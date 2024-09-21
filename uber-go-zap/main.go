package main

import "go.uber.org/zap"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	demo(logger)
	demoSugared(logger)
}
