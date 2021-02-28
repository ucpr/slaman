package models

type Channel struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Members   []string `json:"members"`
	IsArchive bool     `json:"is_archived"`
}
