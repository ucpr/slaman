package models

type Profile struct {
	RealName string `json:"real_name"`
	DisplayName string `json:"display_name"`
	AvatarImage string `json:"image_original"`
	StatusText string `json:"status_text"`
	StatusEmoji string `json:"status_emoji"`
}

type User struct {
	ID string `json:"id"`
	IsAdmin bool `json:"is_admin"`
	IsOwner bool `json:"is_owner"`
	IsPrimaryOwner bool `json:"is_primary_owner"`
	IsBot bool `json:"is_bot"`
	IsAppUser bool `json:"is_app_user"`
	Profile `json:"profile"`
}