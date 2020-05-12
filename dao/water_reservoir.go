package dao

import (
	"github.com/astaxie/beego/logs"
	"water_information_service/lib"
)

type WaterReservoirInformation struct {
	Id               int64    `json:"id" orm:"pk;size(11);column(id)"`
	Name             string   `json:"name" orm:"size(64);column(name)"`
	Address          string   `json:"address" orm:"size(32);column(address)"`
	RiverName        string   `json:"river_name" orm:"size(64);column(river_name)"`
	RiverLevel       float32  `json:"river_level" orm:"size(12);column(river_level)"`
	Storage          float32  `json:"storage" orm:"size(12);column(storage)"`
	AlertRiverLevel  float32  `json:"alert_river_level" orm:"size(12);column(alert_river_level)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

const water_reservoir_information_tableName = "water_reservoir"

func (u *WaterReservoirInformation) TableName() string {
	return water_reservoir_information_tableName
}

// 查询所有水利信息
func SelectAllWaterReservoir() (r []WaterReservoirInformation, err error) {
	_, err = db.QueryTable(water_reservoir_information_tableName).Filter("flag", false).All(&r)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 查询所有水利信息
func SelectWaterReservoirByName(name string) (r WaterReservoirInformation, err error) {
	r.Name = name
	err = db.Read(&r, "Name")
	if err != nil {
		logs.Error(err)
	}
	return
}

// 插入水利信息
func InsertReservoirInformation(name string, address string, riverName string, riverLevel float32, storage float32, alertRiverLevel float32) (err error) {
	var data WaterReservoirInformation
	data.Name = name
	data.Address = address
	data.RiverName = riverName
	data.RiverLevel = riverLevel
	data.Storage = storage
	data.AlertRiverLevel = alertRiverLevel
	_, err = db.Insert(&data)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 修改水利信息
func AlterReservoirInformation(id int64, name string, address string, riverName string, riverLevel float32, storage float32, alertRiverLevel float32) (err error) {
	var data WaterReservoirInformation
	data.Id = id
	data.Name = name
	data.Address = address
	data.RiverName = riverName
	data.RiverLevel = riverLevel
	data.Storage = storage
	data.AlertRiverLevel = alertRiverLevel
	_, err = db.Update(&data, "name", "address", "river_name", "river_level", "storage", "alert_river_level")
	if err != nil {
		logs.Error(err)
	}
	return
}