package database

import (
	"CatatanTugas/config"
	"CatatanTugas/model"
	"fmt"
	"time"
)

func GetCatatans() []model.Catatan {
	var Catatans []model.Catatan
	config.DB.Debug().Preload("LabelCatatans").Preload("Gambars").Where("deleted_at is null").Find(&Catatans)
	return Catatans
}

func GetCatatansByID(catatanId int) (interface{}, error) {
	var catatan model.Catatan
	if e := config.DB.Where("deleted_at is null").Find(&catatan, catatanId).Error; e != nil {
		return nil, e
	}
	return catatan, nil
}

func CreateCatatan(catatan model.Catatan) model.Catatan {
	config.DB.Create(&catatan)
	for _, labelcatatan := range catatan.LabelCatatans {
		err := config.DB.Debug().Raw("insert into catatanlabelcatatans values(?)", labelcatatan.ID).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	return catatan
}

func DeleteCatatanByID(id string) {
	var catatan model.Catatan
	deletCatatan := time.Now()
	catatan.DeletedAt = &deletCatatan
	config.DB.Where("id = ?", id).Updates(&catatan)
}

func UpdateCatatanByID(id string, catatan model.Catatan) {
	config.DB.Where("id = ?", id).Updates(&catatan)
}
