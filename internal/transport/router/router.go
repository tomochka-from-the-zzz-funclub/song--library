package transport

import (
	"net/http"

	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/config"
	myLog "github.com/tomochka-from-the-zzz-funclub/song-library/internal/logger"
	builder "github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/handlers"
	middleware "github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/middlewares"
	"github.com/valyala/fasthttp"
)

type Router struct {
	builder builder.HandlersBuilder
	server  *fasthttp.Server
}

func HandlersCreate(cfg config.Config, middlewares ...middleware.Middleware) *Router {
	rout := router.New()

	// save path for metrics
	rout.SaveMatchedRoutePath = true

	r := Router{
		builder: builder.NewHandlersBuilder(cfg),
	}

	apiV1Group := rout.Group("/api/v1")

	// add setup probes
	rout.GET("/alive", r.builder.Alive(middlewares...))

	apiV1Group.POST("/add", r.builder.Add(middlewares...))
	apiV1Group.DELETE("/delete/{id}", r.builder.Delete(middlewares...))
	apiV1Group.GET("/get/filtration", r.builder.GetWithFiltration(middlewares...))
	apiV1Group.PUT("/update/", r.builder.UpdateSong(middlewares...))
	apiV1Group.GET("/get/text/pagina", r.builder.GetTextWithPagina(middlewares...))

	r.server = &fasthttp.Server{
		Handler: rout.Handler,
	}
	return &r
}

func (r *Router) WithMetrics() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9090", nil); err != nil && err != http.ErrServerClosed {
			myLog.Log.Fatalf("listen and serve of metrics stopped: %v", err)
		}
	}()
}

func (r *Router) Run() {
	go func() {
		myLog.Log.Infof("handlers start to listen")
		if err := r.server.ListenAndServe(":80"); err != nil && err != http.ErrServerClosed {
			myLog.Log.Fatalf("listen and serve of service: %v", err)
		}
	}()
}

func (r *Router) Shutdown() {
	myLog.Log.Infof("server stop")
	err := r.server.Shutdown()
	if err != nil {
		myLog.Log.Errorf("error during server stop: " + err.Error())
	}
}
