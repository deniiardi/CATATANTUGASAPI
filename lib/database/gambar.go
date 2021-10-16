package database

import (
	"CatatanTugas/config"
	"CatatanTugas/model"
)

func GetGambars() []model.Gambar {
	var gambars []model.Gambar
	config.DB.Find(&gambars)
	return gambars
}

func GetGambarByID(id string) model.Gambar {
	var gambar model.Gambar
	config.DB.Where("id = ?", id).Find(&gambar)
	return gambar
}

func CreateGambar(gambar model.Gambar) model.Gambar {
	config.DB.Create(&gambar)
	return gambar
}

func DeleteGambarByID(id string) {
	var gambar model.Gambar
	config.DB.Where("id = ?", id).Delete(&gambar)
}

func UpdateGambarByID(id string, gambar model.Gambar) {
	config.DB.Where("id = ?", id).Updates(&gambar)
}
