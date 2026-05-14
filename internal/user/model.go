package user

import "time"

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	TenantID  uint `gorm:"index"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      string    `gorm:"type:enum('CUSTOMER', 'SHOP_ADMIN', 'SUPER_ADMIN');default:'CUSTOMER'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
