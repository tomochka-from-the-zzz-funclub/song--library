package metrics

import (
	"strconv"
	"time"

	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/valyala/fasthttp"
)

var RequestCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "request_counter",
	Help: "Total number of requests",
}, []string{"method", "status"})

var TimeCounter = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "time_request",
	Help: "Total",
}, []string{"method", "status"})

func Middleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		Now := time.Now()
		handler(ctx)
		TimeWorkF := time.Now().Sub(Now)
		path := ctx.UserValue(router.MatchedRoutePathParam).(string)
		RequestCounter.WithLabelValues(path, strconv.Itoa(ctx.Response.StatusCode())).Inc()
		TimeCounter.WithLabelValues(path, strconv.Itoa(ctx.Response.StatusCode())).Add(float64(TimeWorkF))
	}

}
