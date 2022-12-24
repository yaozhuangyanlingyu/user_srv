package main

import (
	"user_srv/api_handler"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/**
 * prometheus路由
 */
func PrometheusRouter(ginRouter *gin.Engine) {
	// Counter - 计数器
	ginRouter.Handle("GET", "/prom/counter", api_handler.Counter)

	// Gauge - 仪表盘
	ginRouter.Handle("GET", "/prom/gauge", api_handler.Gauge)

	// Histogram - 直方图
	ginRouter.Handle("GET", "/prom/histogram", api_handler.Histogram)

	// Summary 摘要
	ginRouter.Handle("GET", "/prom/summary", api_handler.Summary)

	// 监控接口
	ginRouter.Handle("GET", "/metrics", gin.WrapH(promhttp.Handler()))
}
