package dao

import (
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type WaterInformation struct {
	Id               int64    `json:"id" orm:"pk;size(11);column(id)"`
	CompanyName      string   `json:"company_name" orm:"size(64);column(company_name)"`
	Principal        string   `json:"principal" orm:"size(12);column(principal)"`
	TelephoneNumber  string   `json:"telephone_number" orm:"size(12);column(telephone_number)"`
	FaxNumber        string   `json:"fax_number" orm:"size(12);column(fax_number)"`
	PostCode         string   `json:"post_code" orm:"size(6);column(post_code)"`
	Address          string   `json:"address" orm:"size(32);column(address)"`
	Level            string   `json:"level" orm:"size(12);column(level)"`
	DetailId         int64    `json:"detail_id" orm:"size(11); column(detail_id)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

const water_information_tableName = "water_information"

func (u *WaterInformation) TableName() string {
	return water_information_tableName
}

// 查询单位信息
func SelectWaterInformationByCompanyName(company_name string) (err error) {
	var data WaterInformation
	data.CompanyName = company_name
	err = db.Read(&data, "company_name")
	if err != nil && err != orm.ErrNoRows {
		logs.Error(err)
	}
	return
}

// 查询单位信息
func SelectWaterInformationById(id int64) (data WaterInformation, err error) {
	data.Id = id
	err = db.Read(&data, "id")
	if err != nil && err != orm.ErrNoRows {
		logs.Error(err)
	}
	return
}

// 查询所有水利信息
func SelectAllWaterInformation() (r []WaterInformation, err error) {
	_, err = db.QueryTable(water_information_tableName).Filter("flag", false).All(&r)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 插入水利信息
func InsertWaterInformation(company_name string, principal string, telephone_number string, fax_number string, post_code string, address string, level string, id int64) (err error) {
	var data WaterInformation
	data.CompanyName = company_name
	data.Principal = principal
	data.TelephoneNumber = telephone_number
	data.FaxNumber = fax_number
	data.PostCode = post_code
	data.Address = address
	data.Level = level
	data.DetailId = id
	_, err = db.Insert(&data)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 修改水利信息
func AlterWaterInformation(id int64, company_name string, principal string, telephone_number string, fax_number string, post_code string, address string, level string) (err error) {
	var data WaterInformation
	data.Id = id
	data.CompanyName = company_name
	data.Principal = principal
	data.TelephoneNumber = telephone_number
	data.FaxNumber = fax_number
	data.PostCode = post_code
	data.Address = address
	data.Level = level
	_, err = db.Update(&data, "company_name", "principal", "telephone_number", "fax_number", "post_code", "address", "level")
	if err != nil {
		logs.Error(err)
	}
	return
}