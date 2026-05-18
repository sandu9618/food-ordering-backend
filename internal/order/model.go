package order

import "time"

type Order struct {
	ID          uint        `gorm:"primaryKey;autoIncrement"`
	TenantID    uint        `gorm:"index"`
	UserID      uint        `gorm:"index"`
	Status      string      `gorm:"type:enum('PENDING','CONFIRMED','PREPARING','READY','DELIVERED','CANCELLED');default:'PENDING'"`
	TotalAmount float64
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
	Items       []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	OrderID     uint `gorm:"index"`
	ProductID   uint `gorm:"index"`
	ProductName string
	Price       float64
	Quantity    int `gorm:"default:1"`
	Subtotal    float64
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
