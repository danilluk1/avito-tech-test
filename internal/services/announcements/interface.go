package announcements

import (
	"github.com/danilluk1/avito-tech/internal/db/models"
	"github.com/danilluk1/avito-tech/internal/dto"
)

type AnnouncementService interface {
	GetById(id int32, optional bool) (*models.Announcement, error)
	GetMany(limit, offset uint64, sortBy dto.SortBy, orderBy dto.OrderBy) ([]models.Announcement, error)
	Create(dto *dto.CreateAnnouncement) (*models.Announcement, error)
}
