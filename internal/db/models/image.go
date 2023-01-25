package models

type Photo struct {
	Link           string `db:"link" json:"link"`
	AnnouncementID int32  `db:"announcement_id" json:"announcement_id"`
}
