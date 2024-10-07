package myErrors

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

type Error struct {
	httpCode int
	cause    string
}

func NewError(code int, cause string) Error {
	return Error{
		httpCode: code,
		cause:    cause,
	}
}

func (e Error) GetHttpCode() int {
	return e.httpCode
}

func (e Error) GetCause() string {
	return e.cause
}
func (e Error) Error() string {
	return "Status code: " + strconv.Itoa(e.httpCode) + " cause: " + e.cause
}

//json
var ErrParseJSON = NewError(fasthttp.StatusBadRequest, "error decoding json")
var ErrEqualJSON = NewError(fasthttp.StatusBadRequest, "error read information in JSON format: empty")
var ErrParseJSONTime = NewError(fasthttp.StatusBadRequest, "errors with parse time relize")
var ErrParseJSONNameAndGroup = NewError(fasthttp.StatusBadRequest, "error read name in JSON")
var ErrWriteJSON = NewError(fasthttp.StatusBadRequest, "errors with write responce")
var ErrWriteJSONID = NewError(fasthttp.StatusBadRequest, "errors with write id in json ctreted song")
var ErrWriteJSONerr = NewError(fasthttp.StatusInternalServerError, "error writing response to JSON")

//serv
var ErrAddSong = NewError(fasthttp.StatusInternalServerError, "the song already exists in the library")
var ErrValidationParams = NewError(fasthttp.StatusInternalServerError, "error with valid parameters")

// db
var ErrFindSongDB = NewError(fasthttp.StatusInternalServerError, "error checking the existence of the song")
var ErrCreateSongDB = NewError(fasthttp.StatusInternalServerError, "error creating song in data base")
var ErrNotDeleteDB = NewError(fasthttp.StatusInternalServerError, "could not delete the song in database")
var ErrUpdateSongDB = NewError(fasthttp.StatusInternalServerError, "could not update the song from the database")
var ErrGetTextDB = NewError(fasthttp.StatusInternalServerError, "could not get text from database")
var ErrGetSongWithFiltreDB = NewError(fasthttp.StatusInternalServerError, "could not get song from database with pagin and filtre")
var NotFoundDB = NewError(fasthttp.StatusNotFound, "not found song with such id")

var ErrParseURL = NewError(fasthttp.StatusBadRequest, "error with parse parameters URL")
var ErrUnknownTypeParams = NewError(fasthttp.StatusBadRequest, "unknown type id in func Delete()")
