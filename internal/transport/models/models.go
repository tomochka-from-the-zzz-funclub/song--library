package transport

type SongRequest struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	ReleaseDate string `json:"release"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongRequestID struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	ReleaseDate string `json:"release"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
