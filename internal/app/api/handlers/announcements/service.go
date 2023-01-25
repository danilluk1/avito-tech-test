package announcements

import (
	"github.com/danilluk1/avito-tech/internal/app/api"
	"github.com/danilluk1/avito-tech/internal/db/models"
	"github.com/danilluk1/avito-tech/internal/dto"
)

func AddAnnouncement(app *api.App, dto *dto.CreateAnnouncement) (*models.Announcement, error) {
	announcement, err := app.AnnouncementService.Create(dto)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

func GetAnnouncementById(app *api.App, id int32) (*models.Announcement, error) {
	announcement, err := app.AnnouncementService.GetById(id)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}
