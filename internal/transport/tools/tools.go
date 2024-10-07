package tools

import (
	middleware "github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/middlewares"
	"github.com/valyala/fasthttp"
)

func ApplyMiddleware(handler fasthttp.RequestHandler, middlewares ...middleware.Middleware) fasthttp.RequestHandler {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}
