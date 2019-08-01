package models

// Comments SQL Model
type Comments struct {
	ID       int64  `json:"id"`
	Data     string `json:"name"`
	AuthorID string `json:"author_id"`
}
