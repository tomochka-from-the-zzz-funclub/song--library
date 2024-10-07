package models

import "time"

type Song struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Group       string    `json:"group"`
	ReleaseDate time.Time `json:"release"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

func (song *Song) SetName(n string) {
	song.Name = n
}

func (song *Song) SetGroup(g string) {
	song.Group = g
}

func (song *Song) SetText(t string) {
	song.Text = t
}
func (song *Song) SetLink(l string) {
	song.Link = l
}
func (song *Song) SetReleaseDate(t time.Time) {
	song.ReleaseDate = t
}

func (song *Song) GetID() string {
	return song.ID
}

func (song *Song) GetName() string {
	return song.Name
}

func (song *Song) GetGroup() string {
	return song.Group
}

func (song *Song) GetText() string {
	return song.Text
}

func (song *Song) GetLink() string {
	return song.Link
}

func (song *Song) GetReleaseDateT() time.Time {
	return song.ReleaseDate
}

func (song *Song) GetReleaseDateS() string {
	return song.ReleaseDate.Format("2006/01/02")
}
