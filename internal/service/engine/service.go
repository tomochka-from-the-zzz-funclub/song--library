package service

import (
	"strings"

	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/config"
	db "github.com/tomochka-from-the-zzz-funclub/song-library/internal/database"
	pg "github.com/tomochka-from-the-zzz-funclub/song-library/internal/database/engine"
	myErrors "github.com/tomochka-from-the-zzz-funclub/song-library/internal/errors"
	"github.com/tomochka-from-the-zzz-funclub/song-library/internal/models"
)

type ServiceMusic struct {
	base db.Database
}

func NewServiceMusic(cfg config.Config) *ServiceMusic {
	m := ServiceMusic{
		base: pg.NewPostgres(cfg),
	}
	return &m
}

func (s *ServiceMusic) AddSong(song models.Song) (string, error) {
	check, err := s.base.FindIDByNameAndGroup(song.Name, song.Group)
	if err != nil {
		return "", err
	}
	if !check {
		id, err := s.base.CreateSong(&song)
		return id, err
	}
	return "", myErrors.ErrAddSong
}

func (s *ServiceMusic) GetSongWithFiltre(name, author, release, text, link string, number_records, page int) ([]models.Song, error) {
	array_song, err := s.base.GetFiltreSong(name, author, release, text, link, number_records, page)
	return array_song, err
}

func (s *ServiceMusic) DeleteSong(id string) error {
	err := s.base.DeleteSong(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceMusic) UpdateSong(song models.Song) error {
	err := s.base.UpdateSong(song)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceMusic) GetCoupletText(id string, number_couplet, page int) ([]string, error) {
	pag_couplets := make([]string, 0)
	text, err := s.base.GetText(id)
	if err != nil {
		return pag_couplets, err
	}
	couplets := strings.Split(text, "\n\n")

	if len(couplets) < number_couplet*page {
		return pag_couplets, myErrors.ErrValidationParams
	}
	for i := 0; i < number_couplet; i++ {
		pag_couplets = append(pag_couplets, couplets[(page-1)*number_couplet+i])
	}

	return pag_couplets, nil
}
