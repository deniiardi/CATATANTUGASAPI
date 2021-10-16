package model

import "time"

type Pengingat struct {
	// gorm.Model
	ID             uint       `gorm:"primarykey" form:"id"`
	Pengingat_name string     `json:"pengingat_name"`
	Pengingat_time time.Time  `json:"pengingat_time"`
	CreatedAt      time.Time  `json:"createdAt"`
	Catatan        *Catatan   `json:"catatan,omitempty"`
	CatatanID      uint       `json:"catatanid"`
	DeletedAt      *time.Time `json:"deletedAt"`
}
