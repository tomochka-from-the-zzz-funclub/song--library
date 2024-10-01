package transport

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"

	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

// func Routes() {
// 	fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
// 		switch string(ctx.Path()) {
// 		case "/songs":
// 			switch ctx.Method() {
// 			case "GET":
// 				GetSongs(ctx)
// 			case "POST":
// 				AddSong(ctx)
// 			case "DELETE":
// 				DeleteSong(ctx)
// 			case "PUT":
// 				UpdateSong(ctx)
// 			}
// 		case "/info":
// 			GetSongInfo(ctx)
// 		default:
// 			ctx.NotFound()
// 		}
// 	})
// }

type HandlersBuilder struct {
	lg   zerolog.Logger
	rout *router.Router
}

func HandleCreate() {
	hb := HandlersBuilder{
		lg:   zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.UnixDate}),
		rout: router.New(),
	}
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8090", nil)
	}()

	hb.rout.GET("/shortlink/get", hb.GetShortLink())
	fmt.Println(fasthttp.ListenAndServe(":80", hb.rout.Handler))
}

func (hb *HandlersBuilder) GetShortLink() func(ctx *fasthttp.RequestCtx) {
	hb.lg.Info().
		Msgf("Start func GetShortLink")
	return metrics(func(ctx *fasthttp.RequestCtx) {
		if ctx.IsPost() {
			// 	longlink, timelife, err := ParseJsonL(ctx)
			// 	if err != nil {
			// 		err_ := WriteJsonErr(ctx, err.Error())
			// 		if err_ != nil {
			// 			hb.lg.Warn().
			// 				Msgf("message from func GetShortLink %v", err_.Error())
			// 		}
			// 		hb.lg.Warn().
			// 			Msgf("message from func GetShortLink %v", err.Error())
			// 	} else {
			// 		slink, err := hb.s.CreateShortLink(longlink, timelife)
			// 		if err != nil {
			// 			WriteJsonErr(ctx, err.Error())
			// 			hb.lg.Warn().
			// 				Msgf("message from func GetShortLink with Set %v", err.Error())
			// 		}
			// 		err = WriteJson(ctx, slink)
			// 		if err != nil {
			// 			WriteJsonErr(ctx, err.Error())
			// 			hb.lg.Warn().
			// 				Msgf("message from func GetShortLink %v", err.Error())
			// 		}
			// 	}
			// } else {
			// 	WriteJsonErr(ctx, my_errors.ErrMethodNotAllowed.Error())
			// 	ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			// 	hb.lg.Warn().
			// 		Msgf("message from func GetShortLink %v", my_errors.ErrMethodNotAllowed.Error())

		}
	}, "GetShortLink")
}
