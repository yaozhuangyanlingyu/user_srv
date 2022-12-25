package api_handler

import (
	"math/rand"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
)

// 计数
var AccessCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_requests_total",
	},
	[]string{"method", "path"},
)

// 仪表盘
var QueueGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "queue_num_total",
	},
	[]string{"name"},
)

// 直方图
var HttpDurationsHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_durations_histogram_seconds",
		Buckets: []float64{0.2, 0.5, 1, 2, 5, 10, 30},
	},
	[]string{"path"},
)

// 摘要
var HttpDurations = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "http_durations_seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"path"},
)

func init() {
	prometheus.MustRegister(AccessCounter, QueueGauge, HttpDurationsHistogram, HttpDurations)
}

/**
 * Counter（计数器）
 */
func Counter(c *gin.Context) {
	// URL请求计数
	for _, url := range []string{
		"/order/list",
		"/order/detail",
		"/order/delete",
		"/product/list",
		"/product/detail",
	} {
		AccessCounter.With(prometheus.Labels{
			"method": c.Request.RequestURI,
			"path":   url,
		}).Add(1)
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

/**
 * Gauge（仪表盘）
 */
func Gauge(c *gin.Context) {
	// URL请求计数
	num := c.Query("num")
	fNum := cast.ToFloat64(num)
	QueueGauge.With(
		prometheus.Labels(
			prometheus.Labels{
				"name": "queue_eddycjy",
			},
		),
	).Set(fNum)

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

/**
 * histogram（直方图）
 */
func Histogram(c *gin.Context) {
	// URL请求计数
	//purl, _ := url.Parse(c.Request.RequestURI)
	for _, url := range []string{
		"/order/list",
		"/order/detail",
		"/order/delete",
		"/product/list",
		"/product/detail",
	} {
		HttpDurationsHistogram.With(
			prometheus.Labels{
				"path": url,
			},
		).Observe(
			float64(rand.Intn(30)),
		)
	}

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

/**
 * Summary（摘要）
 */
func Summary(c *gin.Context) {
	// URL请求计数
	purl, _ := url.Parse(c.Request.RequestURI)
	HttpDurations.With(
		prometheus.Labels{
			"path": purl.Path,
		},
	).Observe(float64(rand.Intn(30)))

	// 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
