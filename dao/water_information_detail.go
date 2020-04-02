package dao

import (
	"github.com/astaxie/beego/orm"
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type WaterInformationDetail struct {
	Id               int64    `json:"id" orm:"pk;size(11);column(id)"`
	Introduction     string   `json:"introduction" orm:"size(512);column(introduction)"`
	ImagePath        string   `json:"image_path" orm:"size(64);column(image_path)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const water_information_detail_tableName = "water_information_detail"

func (u *WaterInformationDetail) TableName() string {
	return water_information_detail_tableName
}

// 查询水利信息详情
func SelectWaterInformationDetail(id int64) (r WaterInformationDetail, err error) {
	r.Id = id
	err = db.Read(&r, "id")
	if err != nil && err != orm.ErrNoRows {
		logs.Error(err)
	}
	return
}

// 插入水利介绍信息
func InsertWaterInformationDetail(introduction string, imagePath string) (id int64, err error) {
	var data WaterInformationDetail
	data.Introduction = introduction
	data.ImagePath = imagePath
	id, err = db.Insert(&data)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 删除水利介绍信息
func DeleteWaterInformationDetail(id int64) (err error) {
	if _, err = db.Delete(&WaterInformationDetail{Id: id}); err != nil {
		logs.Error(err)
	}
	return
}

// 修改水利介绍信息
func AlterWaterInformationDetail(id int64, introduction string, imagePath string) (err error) {
	var data WaterInformationDetail
	data.Id = id
	data.Introduction = introduction
	data.ImagePath = imagePath
	_, err = db.Update(&data, "introduction", "image_path")
	if err != nil {
		logs.Error(err)
	}
	return
}