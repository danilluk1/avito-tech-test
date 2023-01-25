package announcements

import (
	"github.com/danilluk1/avito-tech/internal/db/models"
	"github.com/danilluk1/avito-tech/internal/dto"
)

type AnnouncementService interface {
	GetById(id int32) (*models.Announcement, error)
	GetMany(dto dto.GetAnnouncementsQuery) (*[]models.Announcement, error)
	Create(dto dto.CreateAnnouncement) (*models.Announcement, error)
}
