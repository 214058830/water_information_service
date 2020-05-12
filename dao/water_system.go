package dao

import (
	"github.com/astaxie/beego/logs"
	"water_information_service/lib"
)

type WaterSystemInformation struct {
	Id               int64    `json:"id" orm:"pk;size(11);column(id)"`
	SystemName       string   `json:"system_name" orm:"size(64);column(system_name)"`
	Name             string   `json:"name" orm:"size(64);column(name)"`
	Level            float32   `json:"level" orm:"size(12);column(level)"`
	Flow             float32  `json:"flow" orm:"size(12);column(flow)"`
	Potential        string   `json:"potential" orm:"size(32);column(potential)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

const water_system_information_tableName = "water_system"

func (u *WaterSystemInformation) TableName() string {
	return water_system_information_tableName
}

// 查询所有水利信息
func SelectAllWaterSystem() (r []WaterSystemInformation, err error) {
	_, err = db.QueryTable(water_system_information_tableName).Filter("flag", false).All(&r)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 插入水利信息
func InsertSystemInformation(systemName, name string, level, flow float32, Potential string) (err error) {
	var data WaterSystemInformation
	data.SystemName = systemName
	data.Name = name
	data.Level = level
	data.Flow = flow
	data.Potential = Potential
	_, err = db.Insert(&data)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 查询所有水利信息
func SelectWaterSystemByName(system_name, name string) (r WaterSystemInformation, err error) {
	r.SystemName = system_name
	r.Name = name
	err = db.Read(&r, "system_name","Name")
	if err != nil {
		logs.Error(err)
	}
	return
}

// 修改水利信息
func AlterSystemInformation(id int64, systemName, name string, level, flow float32, Potential string) (err error) {
	var data WaterSystemInformation
	data.Id = id
	data.SystemName = systemName
	data.Name = name
	data.Level = level
	data.Flow = flow
	data.Potential = Potential
	_, err = db.Update(&data, "system_name", "name", "level", "flow", "potential")
	if err != nil {
		logs.Error(err)
	}
	return
}