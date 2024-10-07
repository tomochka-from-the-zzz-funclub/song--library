package panics

import (
	"github.com/fasthttp/router"
	myLog "github.com/tomochka-from-the-zzz-funclub/song-library/internal/logger"
	"github.com/valyala/fasthttp"
)

func Middleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		handler(ctx)
		if r := recover(); r != nil {
			path := ctx.UserValue(router.MatchedRoutePathParam).(string)
			myLog.Log.Errorf("catch panic in method %s, panic message: %v", path, r)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBody([]byte{})
		}
	}

}
