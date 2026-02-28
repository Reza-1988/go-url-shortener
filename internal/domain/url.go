package domain

import "time"

type URL struct {
	ID          int64     `gorm:"primaryKey"`
	OwnerID     int64     `gorm:"index;not null"`
	OriginalURL string    `gorm:"not null"`
	ShortCode   string    `gorm:"type:varchar(8);uniqueIndex;not null"`
	ClickCount  int64     `gorm:"not null;default:0"`
	IsDisabled  bool      `gorm:"not null;default:false"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

func (URL) TableName() string {
	return "urls"
}
