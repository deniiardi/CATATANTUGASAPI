package database

import (
	"CatatanTugas/config"
	"CatatanTugas/model"
	"fmt"
	"time"
)

func GetLabelCatatans() []model.LabelCatatan {
	var LabelCatatans []model.LabelCatatan

	config.DB.Where("deleted_at is null").Find(&LabelCatatans)
	return LabelCatatans
}

func GetLabelCatatansByID(labelcatatanId int) (interface{}, error) {
	var labelcatatan model.LabelCatatan

	if e := config.DB.Where("deleted_at is null").Find(&labelcatatan, labelcatatanId).Error; e != nil {
		return nil, e
	}
	return labelcatatan, nil
}

func CreateLabelCatatan(labelcatatan model.LabelCatatan) model.LabelCatatan {
	err := config.DB.Create(&labelcatatan).Error
	fmt.Println(err)
	for _, catatan := range labelcatatan.Catatans {
		config.DB.Raw("insert into catatanlabelcatatans values(?,?)", catatan.ID, labelcatatan.ID)
	}
	return labelcatatan
}

func DeleteLabelCatatanByID(id string) {
	var labelcatatan model.LabelCatatan
	deletLabelCatatan := time.Now()
	labelcatatan.DeletedAt = &deletLabelCatatan
	config.DB.Where("id = ?", id).Updates(&labelcatatan)
}

func UpdateLabelCatatanByID(id string, labelcatatan model.LabelCatatan) {
	config.DB.Where("id = ?", id).Updates(&labelcatatan)
}
