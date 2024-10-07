package service

import "github.com/tomochka-from-the-zzz-funclub/song-library/internal/models"

type Service interface {
	AddSong(song models.Song) (int, error)
	DeleteSong(id int) error
	GetCoupletText(id int, number_couplet int, page int) ([]string, error)
	GetSongWithFiltre(name string, author string, release string, text string, link string, number_records int, page int) ([]models.Song, error)
	UpdateSong(song models.Song) error
}
