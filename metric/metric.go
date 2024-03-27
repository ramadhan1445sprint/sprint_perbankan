package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ServerMetric struct {
	ReqDurationHist *prometheus.HistogramVec
}

func NewServerMetric() *ServerMetric {
	reqDurationHist := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "perbankan_http_request_duration",
		Help:    "Histogram of http request duration",
		Buckets: prometheus.LinearBuckets(1, 1, 10),
	}, []string{"method", "path", "status"})

	return &ServerMetric{
		ReqDurationHist: reqDurationHist,
	}
}
