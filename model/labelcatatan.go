package model

import "time"

type LabelCatatan struct {
	// gorm.Model
	ID                uint       `gorm:"primarykey" json:"id" form:"id"`
	Catatans          []Catatan  `gorm:"many2many:catatanlabelcatatans" json:"labelcatatans,omitempty"`
	LabelCatatanTitle string     `json:"labelcatatantitle" form:"labelcatatantitle"`
	CreatedAt         time.Time  `json:"createdAt"`
	DeletedAt         *time.Time `json:"deletedAt"`
}
