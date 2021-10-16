package model

import (
	"time"
)

type Catatan struct {
	// gorm.Model
	ID            uint           `gorm:"primarykey" form:"id"`
	PenggunaID    uint           `json:"penggunaid" form:"penggunaid"`
	Pengguna      *Pengguna      `json:"pengguna,omitempty" form:"pengguna"`
	Title         string         `json:"title" form:"title"`
	Content       string         `json:"content" form:"content"`
	Status        *string        `json:"status" form:"status"`
	Archive_date  *time.Time     `json:"archive_date" form:"archive_date"`
	LabelCatatans []LabelCatatan `gorm:"many2many:catatanlabelcatatans" json:"labelcatatans,omitempty"`
	Pengingats    []*Pengingat   `json:"pengingat" form:"pengingat"`
	Gambars       []*Gambar      `json:"gambar" form:"gambar"`
	CreatedAt     time.Time      `json:"createdAt" form:"createdAt"`
	DeletedAt     *time.Time     `json:"deletedAt" form:"deletedAt"`
}
