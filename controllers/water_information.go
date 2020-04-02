package controllers

import (
	"strconv"
	"strings"
	"time"

	"water_information_service/log"
	"water_information_service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type WaterInformationController struct {
	beego.Controller
}

// 通用响应格式
type contributeWaterAmountRes struct {
	Data interface{} `json:"data"`
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
}

// @Title Get
// @查询所有水利信息
// @Success 200 success
// @Failure 403 is empty
// @router / [get]
func (this *WaterInformationController) GetAllWaterInformation() {
	var r contributeWaterAmountRes
	r.Data, r.Code = models.SelectAllWaterInformation()
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Post
// @上传水利信息
// @Success 200 success
// @Failure 403 is empty
// @router /upload [post]
func (this *WaterInformationController) Upload() {
	var uploadInfo models.UploadReq
	var r contributeWaterAmountRes
	this.parsingParam(&uploadInfo)
	r.Code = models.Upload(uploadInfo)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// 处理参数 文件存入磁盘 以path作为返回值
func (this *WaterInformationController) parsingParam(uploadInfo *models.UploadReq) {
	uploadInfo.CompanyName = this.GetString("company_name")
	uploadInfo.Principal = this.GetString("principal")
	uploadInfo.TelephoneNumber = this.GetString("telephone_number")
	uploadInfo.FaxNumber = this.GetString("fax_number")
	uploadInfo.PostCode = this.GetString("post_code")
	uploadInfo.Address = this.GetString("address")
	uploadInfo.Level = this.GetString("level")
	uploadInfo.Introduction = this.GetString("introduction")
	//  处理文件参数
	tmpfile, fheader, err := this.Ctx.Request.FormFile("file")
	if err != nil {
		return
	}
	defer tmpfile.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	index := strings.LastIndexByte(fheader.Filename, '.')
	path := beego.AppConfig.String("waterInformationImagePath") + timeStr + fheader.Filename[index:]
	err = this.SaveToFile("file", path)
	if err != nil {
		logs.Error(err)
	}
	uploadInfo.Path = timeStr + fheader.Filename[index:]
	return
}

// @Title Get
// @获取水利信息详情
// @Success 200 success
// @Failure 403 is empty
// @router /detail [get]
func (this *WaterInformationController) GetWaterInformationDetail() {
	var r contributeWaterAmountRes
	id, err := this.GetInt64("id")
	if err != nil {
		logs.Error(err)
	}
	r.Data, r.Code = models.GetWaterInformationDetail(id)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Post
// @修改水利信息
// @Success 200 success
// @Failure 403 is empty
// @router /alter [post]
func (this *WaterInformationController) Alter() {
	var alterInfo models.AlterReq
	var r contributeWaterAmountRes
	this.parsingAlterParam(&alterInfo)
	r.Code = models.Alter(alterInfo)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// 处理参数 文件存入磁盘 以path作为返回值
func (this *WaterInformationController) parsingAlterParam(uploadInfo *models.AlterReq) {
	uploadInfo.Id, _ = this.GetInt64("id")
	uploadInfo.CompanyName = this.GetString("company_name")
	uploadInfo.Principal = this.GetString("principal")
	uploadInfo.TelephoneNumber = this.GetString("telephone_number")
	uploadInfo.FaxNumber = this.GetString("fax_number")
	uploadInfo.PostCode = this.GetString("post_code")
	uploadInfo.Address = this.GetString("address")
	uploadInfo.Level = this.GetString("level")
	uploadInfo.Introduction = this.GetString("introduction")
	//  处理文件参数
	tmpfile, fheader, err := this.Ctx.Request.FormFile("file")
	if err != nil {
		return
	}
	defer tmpfile.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	index := strings.LastIndexByte(fheader.Filename, '.')
	path := beego.AppConfig.String("waterInformationImagePath") + timeStr + fheader.Filename[index:]
	err = this.SaveToFile("file", path)
	if err != nil {
		logs.Error(err)
	}
	uploadInfo.Path = timeStr + fheader.Filename[index:]
	return
}
