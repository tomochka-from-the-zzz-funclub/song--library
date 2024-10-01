package main

import (
	config "song-library/config"
	db "song-library/database"
)

func main() {
	cfg := config.LoadConfig()
	//InitLogger() сделаем в транспорте
	db.ConnectDB(cfg)
	//Routes() в транспорте
}
