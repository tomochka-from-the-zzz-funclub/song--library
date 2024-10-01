package service

import (
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"
)

type Song struct {
	ID    int    `json:"id"`
	Group string `json:"group"`
	Song  string `json:"song"`
}

func GetSongs(ctx *fasthttp.RequestCtx) {
	// Реализация получения данных с пагинацией
	ctx.SetContentType("application/json")
	songs := []Song{} // Здесь выборка из БД
	json.NewEncoder(ctx).Encode(songs)
}

func AddSong(ctx *fasthttp.RequestCtx) {
	var song Song
	if err := json.Unmarshal(ctx.PostBody(), &song); err != nil {
		//logger.Error().Err(err).Msg("Failed to unmarshal song")
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	// Запрос к Swagger API
	// Пример: resp, err := http.Get(SWAGGER_URL + "?group=" + song.Group + "&song=" + song.Song)
	// Обработка ответа и добавление в БД

	ctx.SetStatusCode(http.StatusCreated)
	// Отправка успешного ответа
}

func DeleteSong(ctx *fasthttp.RequestCtx) {
	// Логика удаления песни
}

func UpdateSong(ctx *fasthttp.RequestCtx) {
	// Логика изменения данных песни
}

func GetSongInfo(ctx *fasthttp.RequestCtx) {
	// Логика получения информации о песне по запросу к Swagger API
}
