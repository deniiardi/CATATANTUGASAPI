package database

import (
	"CatatanTugas/config"
	"CatatanTugas/model"
)

func GetPengingats() []model.Pengingat {
	var pengingats []model.Pengingat
	config.DB.Find(&pengingats)
	return pengingats
}

func GetPengingatByID(id string) model.Pengingat {
	var pengingat model.Pengingat
	config.DB.Where("id = ?", id).Find(&pengingat)
	return pengingat
}

func CreatePengingat(pengingat model.Pengingat) model.Pengingat {
	config.DB.Create(&pengingat)
	return pengingat
}

func DeletePengingatByID(id string) {
	var pengingat model.Pengingat
	config.DB.Where("id = ?", id).Delete(&pengingat)
}

func UpdatePengingatByID(id string, pengingat model.Pengingat) {
	config.DB.Where("id = ?", id).Updates(&pengingat)
}
