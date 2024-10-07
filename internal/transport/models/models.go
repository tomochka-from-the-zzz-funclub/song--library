package transport

type SongRequest struct {
	Name        string `json:"name"`
	Group       string `json:"group"`
	ReleaseDate string `json:"release"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongRequestID struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Group       string `json:"group"`
	ReleaseDate string `json:"release"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
