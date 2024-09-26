package main

import (
	"context"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const patternLabel = "pattern"

var durationSummaryVec *prometheus.SummaryVec

type ctxKey int

const (
	ctxKeyPattern ctxKey = iota
)

func initMetrics() {
	if durationSummaryVec == nil {
		durationSummaryVec = promauto.NewSummaryVec(prometheus.SummaryOpts{
			Name:       "request_duration",
			Help:       "The duration of processed requests",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{patternLabel})
	}

}

func instrumentHandlerDuration(next http.HandlerFunc) http.HandlerFunc {
	if durationSummaryVec == nil {
		log.Fatal("duration metric is not initialized")
	}

	return addPatternToCtx(promhttp.InstrumentHandlerDuration(
		durationSummaryVec,
		next,
		withPatternFromCtxAsLabel(),
	))
}

func addPatternToCtx(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ctxKeyPattern, r.Pattern)
		next(w, r.WithContext(ctx))
	}
}

func withPatternFromCtxAsLabel() promhttp.Option {
	return promhttp.WithLabelFromCtx(patternLabel,
		func(ctx context.Context) string {
			return ctx.Value(ctxKeyPattern).(string)
		},
	)
}
