package impl

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/danilluk1/avito-tech/internal/db/models"
	"github.com/danilluk1/avito-tech/internal/dto"
	"github.com/danilluk1/avito-tech/internal/services/announcements"
	"github.com/jmoiron/sqlx"
)

type AnnouncementService struct {
	pgConn *sqlx.DB
}

func NewAnnouncementService(pgConn *sqlx.DB) announcements.AnnouncementService {
	return &AnnouncementService{
		pgConn: pgConn,
	}
}

func (as AnnouncementService) GetById(id int32, optional bool) (*models.Announcement, error) {
	query, args, err := sq.
		Select("announcements.id as id, name, description, array_agg(link) as photos").
		From("announcements").
		InnerJoin("photos ON photos.announcement_id = $1", id).
		GroupBy("announcements.id, link").
		ToSql()
	query = as.pgConn.Rebind(query)
	if err != nil {
		return nil, err
	}

	announcement := &models.Announcement{}
	row := as.pgConn.QueryRowx(query, args...)
	err = row.StructScan(announcement)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return announcement, nil
}

func (as AnnouncementService) GetMany(
	dto *dto.GetAnnouncementsQuery,
) (*[]models.Announcement, error) {
	return nil, nil
}

func (as AnnouncementService) Create(dto *dto.CreateAnnouncement) (*models.Announcement, error) {
	query, args, err := sq.
		Insert("announcements").
		Columns("name", "description").
		Values(dto.Name, dto.Description).Suffix("RETURNING id").ToSql()
	query = as.pgConn.Rebind(query)

	if err != nil {
		return nil, err
	}

	res := as.pgConn.QueryRowx(query, args...)
	var id int32
	err = res.Scan(&id)
	if err != nil {
		return nil, err
	}

	for _, image := range dto.Photos {
		query, args, err := sq.
			Insert("photos").
			Columns("link, announcement_id").
			Values(image, id).ToSql()
		query = as.pgConn.Rebind(query)
		if err != nil {
			return nil, err
		}
		_, err = as.pgConn.Queryx(query, args...)
		if err != nil {
			return nil, err
		}
	}

	announcement, err := as.GetById(id, false)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}
