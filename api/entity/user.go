package entity

import "time"

type UserId string

type User struct {
	UserId          UserId
	DisplayId       string
	Name            string
	AvatarUrl       string
	Bio             string
	FollowersCount  int
	FollowingCount  int
	FollowingGenres []string
	// Followed        bool
	// Following       bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
