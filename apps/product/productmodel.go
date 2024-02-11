package product

import (
	"time"
)

type Product struct {
	Id          int64     `gorm:"primaryKey" json:"id"`
	NamaProduct string    `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi   string    `gorm:"type:text" json:"deskripsi"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
