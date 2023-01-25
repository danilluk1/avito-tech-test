package dto

import "time"

type Announcement struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Photos      []string  `json:"photos"`
}

type GetAnnouncementsQuery struct {
	Limit   int32  `json:"limit" validate:"required gte=0,lte=30"`
	Offset  int32  `json:"offset" validate:"required gte=0"`
	SortBy  string `json:"sortBy" validate:"required"`
	OrderBy string `json:"orderBy" validate:"required"`
}

type CreateAnnouncement struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Photos      []string  `json:"photos"`
}
