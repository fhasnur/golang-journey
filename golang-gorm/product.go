package golanggorm

import "time"

type Product struct {
	ID          string    `gorm:"primary_key;column:id"`
	Name        string    `gorm:"column:name"`
	Price       int64     `gorm:"column:price"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	LikedByUser []User    `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id"`
}

func (p *Product) TableName() string {
	return "products"
}
