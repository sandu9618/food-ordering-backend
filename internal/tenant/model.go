package tenant

import "time"

type Tenant struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Subdomain string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
