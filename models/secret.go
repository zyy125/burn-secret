package models

import (
	"time"
)

const IDLenth = 5

type Secret struct {
	ID string `json:"id"`
	Content string `json:"content"`
	MaxViews int `json:"maxViews"`
	ViewsCount int `json:"viewsCount"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiryMinutes int `json:"expiryMinutes"`
}

type CreateRequest struct {
	Content string `json:"content"`
	MaxViews int `json:"maxViews"`
	ExpiryMinutes int `json:"expiryMinutes"`
}
