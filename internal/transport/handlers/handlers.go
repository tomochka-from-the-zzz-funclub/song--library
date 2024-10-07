package builder

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"strconv"

	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/config"
	myErrors "github.com/tomochka-from-the-zzz-funclub/song-library/internal/errors"
	myLog "github.com/tomochka-from-the-zzz-funclub/song-library/internal/logger"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/models"
	srvInterface "github.com/tomochka-from-the-zzz-funclub/song-library/internal/service"
	service "github.com/tomochka-from-the-zzz-funclub/song-library/internal/service/engine"
	middlewares "github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/middlewares"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/tools"
	"github.com/valyala/fasthttp"
)

type HandlersBuilder struct {
	srv srvInterface.Service
}

func NewHandlersBuilder(cfg config.Config) HandlersBuilder {
	hb := HandlersBuilder{
		srv: service.NewServiceMusic(cfg),
	}
	return hb
}

func (hb *HandlersBuilder) Alive(middlewares ...middlewares.Middleware) func(ctx *fasthttp.RequestCtx) {
	myLog.Log.Infof("Start func Alive")
	return tools.ApplyMiddleware((func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(http.StatusOK)
	}), middlewares...)
}

func (hb *HandlersBuilder) Add(middlewares ...middlewares.Middleware) func(ctx *fasthttp.RequestCtx) {
	myLog.Log.Infof("Start func Add")
	return tools.ApplyMiddleware((func(ctx *fasthttp.RequestCtx) {
		song, err := ParseJsonSong(ctx)
		if err != nil {
			WriteJsonErr(ctx, err)
			myLog.Log.Errorf("Error parsing song from request body: %v", song)
		} else {
			myLog.Log.Debugf("Trying to add a song to the library: %v", song)
			id, err := hb.srv.AddSong(song)
			if err != nil {
				WriteJsonErr(ctx, err)
				myLog.Log.Errorf("Song not added")
			} else {
				myLog.Log.Debugf("song added successfully")
				err = WriteJsonID(ctx, id)
				if err != nil {
					myLog.Log.Errorf("error write in json in func Add(): %v", err.Error())
				}
			}
		}
	}), middlewares...)
}

func (hb *HandlersBuilder) Delete(middlewares ...middlewares.Middleware) func(ctx *fasthttp.RequestCtx) {
	myLog.Log.Infof("Start func Delete")
	return tools.ApplyMiddleware((func(ctx *fasthttp.RequestCtx) {
		id, ok := ctx.UserValue("id").(string)
		if !ok {
			WriteJsonErr(ctx, myErrors.ErrUnknownTypeParams)
			myLog.Log.Errorf("Error unknown type id")
		}
		//id, err := strconv.Atoi(idStr)
		// if err != nil {
		// 	WriteJsonErr(ctx, err)
		// 	myLog.Log.Errorf("Error parsing song name")
		// } else {
		myLog.Log.Debugf("Trying to delete a song")
		err := hb.srv.DeleteSong(id)
		if err != nil {
			WriteJsonErr(ctx, err)
			myLog.Log.Errorf("The song specified to delete is not in the library with id: %v", id)
			return
		}
		myLog.Log.Debugf("Song successfully deleted")
		//}

	}), middlewares...)
}

func (hb *HandlersBuilder) GetWithFiltration(middlewares ...middlewares.Middleware) func(ctx *fasthttp.RequestCtx) {
	myLog.Log.Infof("Start func GetWithFiltration")
	return tools.ApplyMiddleware((func(ctx *fasthttp.RequestCtx) {
		filtre_name := string(ctx.QueryArgs().Peek("name"))
		filtre_group := string(ctx.QueryArgs().Peek("group"))
		filtre_release := string(ctx.QueryArgs().Peek("release"))
		filtre_text := string(ctx.QueryArgs().Peek("text"))
		filtre_link := string(ctx.QueryArgs().Peek("link"))
		number_recordsS := string(ctx.QueryArgs().Peek("records"))
		pageS := string(ctx.QueryArgs().Peek("page"))
		number_recordsI, err := strconv.Atoi(number_recordsS)
		if err != nil {
			myLog.Log.Errorf("message from func GetWithFiltration with get number_recordsI %v", err.Error())
			return
		}
		pageI, err := strconv.Atoi(pageS)
		if err != nil {
			myLog.Log.Errorf("message from func GetWithFiltration with get pageI %v", err.Error())
			return
		}
		if (pageI < 1) || (number_recordsI < 1) {
			WriteJsonErr(ctx, myErrors.ErrValidationParams)
			myLog.Log.Errorf("Parameters did not pass initial validation")
			return
		}
		myLog.Log.Debugf("Trying to get a list of songs that have passed the filter")
		songs, err := hb.srv.GetSongWithFiltre(filtre_name, filtre_group, filtre_release, filtre_text, filtre_link, number_recordsI, pageI)
		if err != nil {
			WriteJsonErr(ctx, err)
			myLog.Log.Errorf("message from func GetWithFiltration %v", err.Error())
			return
		}
		if len(songs) == 0 {
			myLog.Log.Errorf("There are no songs in the library")
		} else {
			ctx.SetContentType("application/json")
			ctx.Response.BodyWriter()
			err := json.NewEncoder((*ctx).Response.BodyWriter()).Encode(songs)
			if err != nil {
				WriteJsonErr(ctx, err)
				myLog.Log.Errorf("error write in json in func GetWithFiltration() %v", err.Error())
			}
		}

	}), middlewares...)
}

func (hb *HandlersBuilder) UpdateSong(middlewares ...middlewares.Middleware) func(ctx *fasthttp.RequestCtx) {
	myLog.Log.Infof("Start func Update")
	return tools.ApplyMiddleware((func(ctx *fasthttp.RequestCtx) {
		song, err := ParseJsonSongWithID(ctx)
		if err != nil {
			WriteJsonErr(ctx, err)
			myLog.Log.Errorf("Error parsing song from request body: %v", err.Error())
		} else {
			myLog.Log.Debugf("Trying to change song information")
			err = hb.srv.UpdateSong(models.Song{
				ID:          song.ID,
				Name:        song.Name,
				Group:       song.Group,
				Text:        song.Text,
				ReleaseDate: song.ReleaseDate,
				Link:        song.Link,
			})
			if err != nil {
				WriteJsonErr(ctx, err)
				myLog.Log.Errorf("Song update error: ", err.Error())
				return
			}
			myLog.Log.Debugf("successfully updated song information with id: %v", song.ID)
		}

	}), middlewares...)
}

func (hb *HandlersBuilder) GetTextWithPagina(middlewares ...middlewares.Middleware) func(ctx *fasthttp.RequestCtx) {
	myLog.Log.Infof("Start func GetTextWithPagina")
	return tools.ApplyMiddleware((func(ctx *fasthttp.RequestCtx) {
		number_couplet, err := strconv.Atoi(string(ctx.QueryArgs().Peek("couplet")))
		if err != nil {
			WriteJsonErr(ctx, myErrors.ErrParseURL)
			myLog.Log.Errorf("Error parsing parameter number_couplet: %v", err.Error())
			return
		}
		page, err := strconv.Atoi(string(ctx.QueryArgs().Peek("page")))
		if err != nil {
			WriteJsonErr(ctx, myErrors.ErrParseURL)
			myLog.Log.Errorf("Error parsing parameter page: %v", err.Error())
			return
		}
		//id, err := strconv.Atoi(string(ctx.QueryArgs().Peek("id")))
		id := string(ctx.QueryArgs().Peek("id"))
		// if err != nil {
		// 	WriteJsonErr(ctx, myErrors.ErrParseURL)
		// 	myLog.Log.Errorf("Error parsing parameter id: %v", err.Error())
		// 	return
		// }
		if (page < 1) || (number_couplet < 1) {
			WriteJsonErr(ctx, myErrors.ErrValidationParams)
			myLog.Log.Errorf("Parameters not validated")
			return
		}
		myLog.Log.Debugf("Trying to get song lyrics with verse pagination")
		text_couplet, err := hb.srv.GetCoupletText(id, number_couplet, page)
		if err != nil {
			WriteJsonErr(ctx, err)
			myLog.Log.Errorf("Error getting verses of song with pagination")
			return
		}
		myLog.Log.Debugf("Received song lyrics with pagination by verses")
		for i := 0; i < number_couplet; i++ {
			err = WriteJsonText(ctx, text_couplet[i])
		}

	}), middlewares...)
}
