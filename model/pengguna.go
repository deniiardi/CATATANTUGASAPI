package model

import "time"

type Pengguna struct {
	// gorm.Model
	ID        uint       `gorm:"primarykey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	Catatans  []*Catatan `json:"catatan"`
	Token     string     `json:"token,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
