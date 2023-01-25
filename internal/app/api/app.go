package api

import (
	"github.com/danilluk1/avito-tech/internal/services/announcements"
	"github.com/danilluk1/avito-tech/internal/services/logger"
)

type App struct {
	AnnouncementService announcements.AnnouncementService
	Logger              logger.Logger
}
