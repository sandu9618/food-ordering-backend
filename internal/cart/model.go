package cart

import "time"

type Cart struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	TenantID  uint       `gorm:"index"`
	UserID    uint       `gorm:"index"`
	Items     []CartItem `gorm:"foreignKey:CartID"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

type CartItem struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CartID    uint      `gorm:"index"`
	ProductID uint      `gorm:"index"`
	Quantity  int       `gorm:"default:1"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
