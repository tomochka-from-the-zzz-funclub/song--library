package builder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	myErrors "github.com/tomochka-from-the-zzz-funclub/song-library/internal/errors"
	myLog "github.com/tomochka-from-the-zzz-funclub/song-library/internal/logger"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/models"
	httpModels "github.com/tomochka-from-the-zzz-funclub/song-library/internal/transport/models"
	"github.com/valyala/fasthttp"
)

func ParseJsonSong(ctx *fasthttp.RequestCtx) (models.Song, error) {
	var song httpModels.SongRequest
	err := json.NewDecoder(bytes.NewReader(ctx.Request.Body())).Decode(&song)
	if err != nil {
		myLog.Log.Errorf("error in parse json string", err.Error())
		return models.Song{}, myErrors.ErrParseJSON
	}
	if song.Name == "" || song.Group == "" || song.Link == "" || song.ReleaseDate == "" || song.Text == "" {
		return models.Song{}, myErrors.ErrEqualJSON
	}

	time, err := time.Parse("2006/01/02", song.ReleaseDate)
	if err != nil {
		return models.Song{}, myErrors.ErrParseJSONTime
	}

	return models.Song{
		Name:        song.Name,
		Group:       song.Group,
		ReleaseDate: time,
		Text:        song.Text,
		Link:        song.Link,
	}, nil
}

func ParseJsonSongWithID(ctx *fasthttp.RequestCtx) (models.Song, error) {
	var song httpModels.SongRequestID
	err := json.NewDecoder(bytes.NewReader(ctx.Request.Body())).Decode(&song)
	fmt.Println(song.ID)
	if err != nil {
		return models.Song{}, myErrors.ErrParseJSON
	}
	if song.Name == "" || song.Group == "" || song.Link == "" || song.ReleaseDate == "" || song.Text == "" {
		return models.Song{}, myErrors.ErrEqualJSON
	}

	time, err := time.Parse("2006/01/02", song.ReleaseDate)
	if err != nil {
		return models.Song{}, myErrors.ErrParseJSONTime
	}
	return models.Song{
		ID:          song.ID,
		Name:        song.Name,
		Group:       song.Group,
		ReleaseDate: time,
		Text:        song.Text,
		Link:        song.Link,
	}, nil
}

func ParseJsonNameAndAuthorSong(ctx *fasthttp.RequestCtx) (string, string, error) {
	var namesong struct {
		Name   string `json:"name"`
		Author string `json:"author"`
	}
	err := json.NewDecoder(bytes.NewReader(ctx.Request.Body())).Decode(&namesong)
	if err != nil {
		return "", "", myErrors.ErrParseJSONNameAndGroup
	}
	return namesong.Name, namesong.Author, nil
}

func WriteJson(ctx *fasthttp.RequestCtx, s string) error {
	ctx.SetContentType("application/json")
	ctx.Response.BodyWriter()
	err := json.NewEncoder((*ctx).Response.BodyWriter()).Encode(s)
	if err != nil {
		return myErrors.ErrWriteJSON
	}
	return nil
}

func WriteJsonID(ctx *fasthttp.RequestCtx, id string) error {
	var idsong struct {
		ID string `json:"id"`
	}
	ctx.SetContentType("application/json")
	ctx.Response.BodyWriter()
	idsong.ID = id
	err := json.NewEncoder((*ctx).Response.BodyWriter()).Encode(idsong)
	if err != nil {
		return myErrors.ErrWriteJSON
	}
	return nil
}

func WriteJsonText(ctx *fasthttp.RequestCtx, text string) error {
	var idsong struct {
		Text string `json:"text"`
	}
	ctx.SetContentType("application/json")
	ctx.Response.BodyWriter()
	idsong.Text = text
	err := json.NewEncoder((*ctx).Response.BodyWriter()).Encode(idsong)
	if err != nil {
		return myErrors.ErrWriteJSON
	}
	return nil
}

func WriteJsonErr(ctx *fasthttp.RequestCtx, err error) error {
	var err_mes struct {
		ErrorMessage string
	}
	ctx.SetContentType("application/json")
	if myerror, ok := err.(myErrors.Error); ok {
		err_mes.ErrorMessage = myerror.GetCause()
		if myerror.GetHttpCode() != 0 {
			ctx.SetStatusCode(myerror.GetHttpCode())
		} else {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		}
	} else {
		err_mes.ErrorMessage = err.Error()
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}

	err_ := json.NewEncoder((*ctx).Response.BodyWriter()).Encode(err_mes)
	if err_ != nil {
		return myErrors.ErrWriteJSONerr
	}
	return nil
}
