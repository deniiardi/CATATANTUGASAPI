package database

import (
	"CatatanTugas/config"
	"CatatanTugas/middleware"
	"CatatanTugas/model"
	"errors"
	"time"
)

//ada tambahan where deleted_at is null
func GetPenggunas() []model.Pengguna {
	var penggunas []model.Pengguna
	config.DB.Preload("Catatans").Where("deleted_at is null").Find(&penggunas)
	return penggunas
}

// func GetPenggunaByID(id string) model.Pengguna {
// 	var pengguna model.Pengguna
// 	config.DB.Where("id = ?", id).Preload("Catatans").Find(&pengguna)
// 	return pengguna
// }

func CreatePengguna(pengguna model.Pengguna) model.Pengguna {
	config.DB.Create(&pengguna)
	return pengguna
}

//Di bawah ada editan dari mentor utk delete
func DeletePenggunaByID(id string) {
	var pengguna model.Pengguna
	deletPengguna := time.Now()
	pengguna.DeletedAt = &deletPengguna
	config.DB.Where("id = ?", id).Updates(&pengguna)
}

func UpdatePenggunaByID(id string, pengguna model.Pengguna) {
	config.DB.Where("id = ?", id).Updates(&pengguna)
}

// func IsValid(nm string, em string) bool {
// 	var pengguna model.Pengguna
// 	if err := config.DB.Where("email = ?", em).Find(&pengguna).Error; err != nil {
// 		return false
// 	}

// 	return nm == pengguna.Name
// }

func GetDetailPengguna(penggunaId int) (interface{}, error) {
	var pengguna model.Pengguna

	if e := config.DB.Where("deleted_at is null").Find(&pengguna, penggunaId).Error; e != nil {
		return nil, e
	}
	return pengguna, nil
}

func LoginPengguna(pengguna *model.Pengguna) (interface{}, error) {
	var penggunaFound model.Pengguna
	var err error
	if err = config.DB.Where("email = ?", pengguna.Email).First(&penggunaFound).Error; err != nil {
		return nil, err
	}

	if pengguna.Password != penggunaFound.Password {
		return nil, errors.New("pengguna not found")
	}

	pengguna.Token, err = middleware.CreateToken(int(pengguna.ID))
	if err != nil {
		return nil, err
	}

	return pengguna.Token, nil
}
