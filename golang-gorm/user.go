package golanggorm

import "time"

type User struct {
	ID        string    `gorm:"primary_key;column:id"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

// Jika ingin mengubah nama tabel mapping-nya
func (u *User) TableName() string {
	return "users"
}
