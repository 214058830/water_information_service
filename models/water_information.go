package models

import (
	"github.com/astaxie/beego"
	"os"
	"water_information_service/dao"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 查询所有水利信息
func SelectAllWaterInformation() (r []dao.WaterInformation, code string) {
	var err error
	r, err = dao.SelectAllWaterInformation()
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

type UploadReq struct {
	CompanyName     string `json:"company_name"`
	Principal       string `json:"principal"`
	TelephoneNumber string `json:"telephone_number"`
	FaxNumber       string `json:"fax_number"`
	PostCode        string `json:"post_code"`
	Address         string `json:"address"`
	Level           string `json:"level"`
	Introduction    string `json:"introduction"`
	Path            string
}

// 录入水利信息
func Upload(uploadInfo UploadReq) (code string) {
	var err error
	if err = dao.SelectWaterInformationByCompanyName(uploadInfo.CompanyName); err != nil {
		if (err == orm.ErrNoRows) {
			var id int64
			id, err = dao.InsertWaterInformationDetail(uploadInfo.Introduction, uploadInfo.Path)
			if err != nil {
				logs.Error(err)
				code = "5000"
				return
			} else {
				if err = dao.InsertWaterInformation(uploadInfo.CompanyName, uploadInfo.Principal, uploadInfo.TelephoneNumber, uploadInfo.FaxNumber, uploadInfo.PostCode, uploadInfo.Address, uploadInfo.Level, id); err != nil {
					logs.Error(err)
					code = "5000"
					if err = dao.DeleteWaterInformationDetail(id); err != nil {
						logs.Error(err)
					}
					return
				}
			}
		} else {
			logs.Error(err)
			code = "5000"
			return
		}
	} else {
		// 已存在
		code = "2006"
		return
	}
	code = "2000"
	return
}

// 获取水利信息详情
func GetWaterInformationDetail(id int64) (r dao.WaterInformationDetail, code string) {
	var err error
	r, err = dao.SelectWaterInformationDetail(id)
	if err != nil {
		if err == orm.ErrNoRows {
			code = "2007"
		} else {
			logs.Error(err)
			code = "5000"
		}
	} else {
		code = "2000"
	}
	return
}

type AlterReq struct {
	Id              int64  `json:"id"`
	CompanyName     string `json:"company_name"`
	Principal       string `json:"principal"`
	TelephoneNumber string `json:"telephone_number"`
	FaxNumber       string `json:"fax_number"`
	PostCode        string `json:"post_code"`
	Address         string `json:"address"`
	Level           string `json:"level"`
	Introduction    string `json:"introduction"`
	Path            string
}

// 修改水利信息
func Alter(alterReq AlterReq) (code string) {
	// 检查id值得信息是否存在
	if tempInfo, err := dao.SelectWaterInformationById(alterReq.Id); err != nil {
		if (err != orm.ErrNoRows) {
			code = "5000"
		} else {
			// 不存在情况
			code = "5001"
		}
	} else {
		err = dao.AlterWaterInformation(alterReq.Id, alterReq.CompanyName, alterReq.Principal, alterReq.TelephoneNumber, alterReq.FaxNumber, alterReq.PostCode, alterReq.Address, alterReq.Level)
		if err != nil {
			logs.Error(err)
			code = "5000"
			return
		}
		var tempDetail dao.WaterInformationDetail
		tempDetail, err = dao.SelectWaterInformationDetail(tempInfo.DetailId)
		if err != nil {
			logs.Error(err)
			code = "5000"
			return
		} else {
			if err := os.Remove(beego.AppConfig.String("waterInformationImagePath") + tempDetail.ImagePath); err != nil {
				logs.Error(err)
			}
			if err = dao.AlterWaterInformationDetail(tempDetail.Id, alterReq.Introduction, alterReq.Path); err != nil {
				logs.Error(err)
				code = "5000"
				return
			}
		}
		code = "2000"
	}
	return
}
