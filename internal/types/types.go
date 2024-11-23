package types

import "github.com/srmbackisdeveloper/test-music-info/internal/models"

type AddSongRequest struct {
	Group string `json:"group" binding:"required"`
	Title string `json:"song" binding:"required"`
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

// ---

type PaginatedVersesResponse struct {
	Page        int      `json:"page"`
	Limit       int      `json:"limit"`
	TotalPages  int      `json:"totalPages"`
	TotalVerses int      `json:"totalVerses"`
	Data        []string `json:"data"`
}

type PaginatedSongsResponse struct {
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"totalPages"`
	TotalSongs int            `json:"totalSongs"`
	Data       []models.Music `json:"data"`
}
