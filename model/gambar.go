package model

import "time"

type Gambar struct {
	// gorm.Model
	ID          uint       `gorm:"primarykey" form:"id"`
	Gambar_name string     `json:"gambar" form:"gambar"`
	CreatedAt   time.Time  `json:"createdAt" form:"createdAt"`
	Catatan     *Catatan   `json:"catatan,omitempty" form:"catatan,omitempty"`
	CatatanID   uint       `json:"catatanid" form:"catatanid"`
	DeletedAt   *time.Time `json:"deletedAt"`
}
