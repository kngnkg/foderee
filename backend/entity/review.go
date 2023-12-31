package entity

import (
	"encoding/json"
	"time"
)

type PublishedStatus string

const (
	Published PublishedStatus = "published"
	Draft     PublishedStatus = "draft"
	UnListed  PublishedStatus = "unlisted"
)

type Review struct {
	ReviewId        string          `db:"review_id" validate:"required,uuid4"`
	PublishedStatus PublishedStatus `db:"published_status" validate:"required,published_status"`
	Author          *Author         `db:"author" validate:"required"`
	AlbumId         string          `db:"album_id" validate:"required,album_id"`
	Title           string          `db:"title" validate:"required,min=1,max=100"`
	Content         json.RawMessage `db:"content" validate:"required,json"`
	LikesCount      int
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
