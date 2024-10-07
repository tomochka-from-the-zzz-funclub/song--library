package cors

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

func Middleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(c *fasthttp.RequestCtx) {
		c.Response.Header.Add("Access-Control-Allow-Origin", "*")
		c.Response.Header.Add("Access-Control-Allow-Credentials", "true")
		c.Response.Header.Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, X-User-Id, origin")
		c.Response.Header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		if string(c.Request.Header.Method()) == "OPTIONS" {
			c.Response.SetStatusCode(http.StatusNoContent)
			return
		}
		handler(c)
	}
}
