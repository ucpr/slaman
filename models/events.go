package models

type Event struct {
	Type   string `json:"type"`
	UserID string `json:"user"`
}
