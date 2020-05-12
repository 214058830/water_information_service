package models

import (
	"github.com/astaxie/beego/orm"
	"water_information_service/dao"

	"github.com/astaxie/beego/logs"
)

// 查询所有河道信息
func SelectAllWaterRiver() (r []dao.WaterRiverInformation, code string) {
	var err error
	r, err = dao.SelectAllWaterRiver()
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

// 查询所有水库信息
func SelectAllWaterReservoir() (r []dao.WaterReservoirInformation, code string) {
	var err error
	r, err = dao.SelectAllWaterReservoir()
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

// 查询所有水系信息
func SelectAllWaterSystem() (r []dao.WaterSystemInformation, code string) {
	var err error
	r, err = dao.SelectAllWaterSystem()
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

type RiverReservoirReq struct {
	Name            string  `json:"name"`
	Address         string  `json:"address"`
	RiverName       string  `json:"river_name"`
	RiverLevel      float32 `json:"river_level"`
	Flow            float32 `json:"flow"`
	AlertRiverLevel float32 `json:"alert_river_level"`
}

// 录入河道信息
func UploadRiverInformation(uploadInfo RiverReservoirReq) (code string) {
	logs.Debug(uploadInfo)
	if err := dao.InsertRiverInformation(uploadInfo.Name, uploadInfo.Address, uploadInfo.RiverName, uploadInfo.RiverLevel, uploadInfo.Flow, uploadInfo.AlertRiverLevel); err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

type AlterRiverReq struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	Address         string  `json:"address"`
	RiverName       string  `json:"river_name"`
	RiverLevel      float32 `json:"river_level"`
	Flow            float32 `json:"flow"`
	AlertRiverLevel float32 `json:"alert_river_level"`
}

// 录入河道信息
func AlterRiverInformation(uploadInfo AlterRiverReq) (code string) {
	logs.Debug(uploadInfo)
	if err := dao.AlterWaterInformation(uploadInfo.Id, uploadInfo.Name, uploadInfo.Address, uploadInfo.RiverName, uploadInfo.RiverLevel, uploadInfo.Flow, uploadInfo.AlertRiverLevel); err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

type AlterReservoirReq struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	Address         string  `json:"address"`
	RiverName       string  `json:"river_name"`
	RiverLevel      float32 `json:"river_level"`
	Storage         float32 `json:"storage"`
	AlertRiverLevel float32 `json:"alert_river_level"`
}

// 录入水库信息
func UploadReservoirInformation(uploadInfo RiverReservoirReq) (code string) {
	logs.Debug(uploadInfo)
	if err := dao.InsertReservoirInformation(uploadInfo.Name, uploadInfo.Address, uploadInfo.RiverName, uploadInfo.RiverLevel, uploadInfo.Flow, uploadInfo.AlertRiverLevel); err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

// 修改水库信息
func AlterReservoirInformation(uploadInfo AlterReservoirReq) (code string) {
	logs.Debug(uploadInfo)
	if err := dao.AlterReservoirInformation(uploadInfo.Id, uploadInfo.Name, uploadInfo.Address, uploadInfo.RiverName, uploadInfo.RiverLevel, uploadInfo.Storage, uploadInfo.AlertRiverLevel); err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

type SystemReq struct {
	SystemName string  `json:"system_name"`
	Name       string  `json:"name"`
	Level      float32 `json:"level"`
	Flow       float32 `json:"flow"`
	Potential  string  `json:"potential"`
}

// 录入水系信息
func UploadSystemInformation(uploadInfo SystemReq) (code string) {
	logs.Debug(uploadInfo)
	if err := dao.InsertSystemInformation(uploadInfo.SystemName, uploadInfo.Name, uploadInfo.Level, uploadInfo.Flow, uploadInfo.Potential); err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

type AlterSystemReq struct {
	Id         int64   `json:"id"`
	SystemName string  `json:"system_name" orm:"size(64);column(system_name)"`
	Name       string  `json:"name" orm:"size(64);column(name)"`
	Level      float32 `json:"level" orm:"size(12);column(level)"`
	Flow       float32 `json:"flow" orm:"size(12);column(flow)"`
	Potential  string  `json:"potential" orm:"size(32);column(potential)"`
}

// 修改全国水系信息
func AlterSystemInformation(uploadInfo AlterSystemReq) (code string) {
	logs.Debug(uploadInfo)
	if err := dao.AlterSystemInformation(uploadInfo.Id, uploadInfo.SystemName, uploadInfo.Name, uploadInfo.Level, uploadInfo.Flow, uploadInfo.Potential); err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

// 查询全国重点河道信息
func SelectWaterRiverInfo(name string) (r dao.WaterRiverInformation, code string) {
	logs.Debug(name)
	var err error
	code = "2000"
	r, err = dao.SelectWaterRiverByName(name)
	if err != nil {
		if err == orm.ErrNoRows {
			code = "4000"
		} else {
			logs.Error(err)
			code = "5000"
		}
	}
	return
}

// 查询全国重点水库信息
func SelectWaterReservoirInfo(name string) (r dao.WaterReservoirInformation, code string) {
	logs.Debug(name)
	var err error
	code = "2000"
	r, err = dao.SelectWaterReservoirByName(name)
	if err != nil {
		if err == orm.ErrNoRows {
			code = "4000"
		} else {
			logs.Error(err)
			code = "5000"
		}
	}
	return
}

// 查询全国七大水系信息
func SelectWaterSystemInfo(system_name, name string) (r dao.WaterSystemInformation, code string) {
	logs.Debug(name)
	var err error
	code = "2000"
	r, err = dao.SelectWaterSystemByName(system_name, name)
	if err != nil {
		if err == orm.ErrNoRows {
			code = "4000"
		} else {
			logs.Error(err)
			code = "5000"
		}
	}
	return
}
