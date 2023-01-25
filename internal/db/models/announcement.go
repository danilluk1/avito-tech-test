package models

import (
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"time"
)

type Announcement struct {
	ID          int32           `db:"id" json:"id"`
	Name        string          `db:"name" json:"name"`
	Description string          `db:"description" json:"description"`
	Price       decimal.Decimal `db:"price" json:"price"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	Photos      pq.StringArray  `db:"photos" json:"photos"`
}
