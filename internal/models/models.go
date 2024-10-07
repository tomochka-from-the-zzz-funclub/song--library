package models

import "time"

type Song struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	ReleaseDate time.Time `json:"release"`
	Text        string    `json:"text"`
	Link        string    `json:"link"`
}

func (song *Song) SetName(n string) {
	song.Name = n
}

func (song *Song) SetAuthor(g string) {
	song.Author = g
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

func (song *Song) GetID() int {
	return song.ID
}

func (song *Song) GetName() string {
	return song.Name
}

func (song *Song) GetAuthor() string {
	return song.Author
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

func (song *Song) GetReleaseDateS() string { // мне не нужно проверять
	return song.ReleaseDate.Format("2006/01/02")
}
