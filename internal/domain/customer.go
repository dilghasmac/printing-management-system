package domain

import "time"

type Customer struct {
	CustomerID    uint64 `gorm:"primaryKey;autoIncrement"`
	CompanyName   string `gorm:"size:150;not null"`
	ContactPerson string `gorm:"size:100"`
	Phone         string `gorm:"size:20"`
	Email         string `gorm:"size:150"`
	Address       string `gorm:"type:text"`
	Remarks       string `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
