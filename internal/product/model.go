package product

import "time"

type Product struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	TenantID    uint `gorm:"index"`
	Name        string
	Description string
	Price       float64
	CategoryID  uint      `gorm:"index"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
