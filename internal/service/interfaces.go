package service

import "github.com/tomochka-from-the-zzz-funclub/song-library/internal/models"

type Service interface {
	AddSong(song models.Song) (string, error)
	DeleteSong(id string) error
	GetCoupletText(id string, number_couplet int, page int) ([]string, error)
	GetSongWithFiltre(name string, group string, release string, text string, link string, number_records int, page int) ([]models.Song, error)
	UpdateSong(song models.Song) error
}
