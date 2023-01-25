package dto

import "time"

type SortBy string

const (
	SortByPrice SortBy = "price"
	SortByDate  SortBy = "date"
)

type OrderBy string

const (
	OrderByAsc  OrderBy = "asc"
	OrderByDesc OrderBy = "desc"
)

type Announcement struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Photos      []string  `json:"photos"`
}

type GetAnnouncementsQuery struct {
	Limit   int32   `json:"limit" validate:"gte=0,lte=30"`
	Offset  int32   `json:"offset" validate:"gte=0"`
	SortBy  SortBy  `json:"sortBy" validate:""`
	OrderBy OrderBy `json:"orderBy" validate:""`
}

type CreateAnnouncement struct {
	Name        string   `json:"name" validate:"required,min=5,max=200"`
	Description string   `json:"description" validate:"required,min=5,max=1000"`
	Photos      []string `json:"photos" validate:"required"`
}
