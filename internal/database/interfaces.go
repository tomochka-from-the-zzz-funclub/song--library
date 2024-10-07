package database

import (
	"time"

	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/models"
)

type ISong interface {
	SetName(string)
	SetAuthor(string)
	SetReleaseDate(time.Time)
	SetText(string)
	SetLink(string)

	GetAuthor() string
	GetName() string
	GetReleaseDateS() string
	GetReleaseDateT() time.Time
	GetText() string
	GetLink() string
}

type Database interface {
	CreateSong(song ISong) (int, error)
	DeleteSong(id int) error
	FindIDByNameAndAuthor(name string, author string) (bool, error)
	GetFiltreSong(name, author, release, text, link string, number_records, page int) ([]models.Song, error)
	UpdateSong(song models.Song) error
	GetText(id int) (string, error)
}
