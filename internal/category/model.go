package category

import "time"

type Category struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	TenantID  uint `gorm:"index"`
	Name      string
	Code      string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
