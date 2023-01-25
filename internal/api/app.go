package api

import "github.com/danilluk1/avito-tech/internal/services/announcements"

type App struct {
	AnnouncementService announcements.AnnouncementService
}
