package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"os"
	"water_information_service/log"
	"water_information_service/models"

	"github.com/astaxie/beego"
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
// @查询重点河道信息
// @Success 200 success
// @Failure 403 is empty
// @router /river [get]
func (this *WaterInformationController) GetAllWaterRiver() {
	var r contributeWaterAmountRes
	r.Data, r.Code = models.SelectAllWaterRiver()
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @通过站名查询重点河道信息
// @Success 200 success
// @Failure 403 is empty
// @router /riverinfo [get]
func (this *WaterInformationController) GetWaterRiverByName() {
	var r contributeWaterAmountRes
	name := this.GetString("name")
	r.Data, r.Code = models.SelectWaterRiverInfo(name)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

//@Title Post
//@修改河道信息
//@Success 200 success
//@Failure 403 is empty
//@router /alter_river [post]
func (this *WaterInformationController) AlterRiver() {
	var uploadInfo models.AlterRiverReq
	var r contributeWaterAmountRes
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &uploadInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.AlterRiverInformation(uploadInfo)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @查询重点水库信息
// @Success 200 success
// @Failure 403 is empty
// @router /reservoir [get]
func (this *WaterInformationController) GetAllWaterReservoir() {
	var r contributeWaterAmountRes
	r.Data, r.Code = models.SelectAllWaterReservoir()
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @通过站名查询重点水库信息
// @Success 200 success
// @Failure 403 is empty
// @router /reservoirinfo [get]
func (this *WaterInformationController) GetWaterReservoirByName() {
	var r contributeWaterAmountRes
	name := this.GetString("name")
	r.Data, r.Code = models.SelectWaterReservoirInfo(name)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

//@Title Post
//@修改水库信息
//@Success 200 success
//@Failure 403 is empty
//@router /alter_reservoir [post]
func (this *WaterInformationController) AlterReservoir() {
	var uploadInfo models.AlterReservoirReq
	var r contributeWaterAmountRes
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &uploadInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.AlterReservoirInformation(uploadInfo)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @查询重点水系信息
// @Success 200 success
// @Failure 403 is empty
// @router /system [get]
func (this *WaterInformationController) GetAllWaterSystem() {
	var r contributeWaterAmountRes
	r.Data, r.Code = models.SelectAllWaterSystem()
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @通过站名查询重点水库信息
// @Success 200 success
// @Failure 403 is empty
// @router /systeminfo [get]
func (this *WaterInformationController) GetWaterSystemByName() {
	var r contributeWaterAmountRes
	systemName := this.GetString("system_name")
	name := this.GetString("name")
	r.Data, r.Code = models.SelectWaterSystemInfo(systemName, name)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

//@Title Post
//@修改水系信息
//@Success 200 success
//@Failure 403 is empty
//@router /alter_system [post]
func (this *WaterInformationController) AlterSystem() {
	var uploadInfo models.AlterSystemReq
	var r contributeWaterAmountRes
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &uploadInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.AlterSystemInformation(uploadInfo)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

//@Title Post
//@上传河道信息
//@Success 200 success
//@Failure 403 is empty
//@router /river [post]
func (this *WaterInformationController) RiverUpload() {
	var uploadInfo models.RiverReservoirReq
	var r contributeWaterAmountRes
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &uploadInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.UploadRiverInformation(uploadInfo)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

//@Title Post
//@上传水库信息
//@Success 200 success
//@Failure 403 is empty
//@router /reservoir [post]
func (this *WaterInformationController) ReservoirUpload() {
	var uploadInfo models.RiverReservoirReq
	var r contributeWaterAmountRes
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &uploadInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.UploadReservoirInformation(uploadInfo)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

//@Title Post
//@上传水系信息
//@Success 200 success
//@Failure 403 is empty
//@router /system [post]
func (this *WaterInformationController) SystemUpload() {
	var uploadInfo models.SystemReq
	var r contributeWaterAmountRes
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &uploadInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.UploadSystemInformation(uploadInfo)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

//@Title Post
//@上传热点信息
//@Success 200 success
//@Failure 403 is empty
//@router /file [post]
func (this *WaterInformationController) FileUpload() {
	var r contributeWaterAmountRes
	r.Code = "2000"
	path, err := this.parsingParam()
	if err != nil {
		logs.Error(err)
		r.Code = "5000"
	}
	logs.Debug(path)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// 处理参数 文件存入磁盘 以path作为返回值
func (this *WaterInformationController) parsingParam() (string, error) {
	//  处理文件参数
	tmpfile, _, err := this.Ctx.Request.FormFile("file")
	if err != nil {
		return "", err
	}
	defer tmpfile.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	//timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	//index := strings.LastIndexByte(fheader.Filename, '.')
	rootpath := beego.AppConfig.String("waterInformationImagePath")
	path := rootpath + "test.png"
	err = this.SaveToFile("file", path)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	if beego.AppConfig.String("runmode") == "dev" {
		// 删除6.jpg，1-5.jpg重命名，test.jpg重命名为1.jpg
		err = os.Remove(rootpath + "6.png")
		if err != nil {
			logs.Error(err)
		}
		err = os.Rename(rootpath+"5.png", rootpath+"6.png")
		if err != nil {
			logs.Error(err)
		}
		err = os.Rename(rootpath+"4.png", rootpath+"5.png")
		if err != nil {
			logs.Error(err)
		}
		err = os.Rename(rootpath+"3.png", rootpath+"4.png")
		if err != nil {
			logs.Error(err)
		}
		err = os.Rename(rootpath+"2.png", rootpath+"3.png")
		if err != nil {
			logs.Error(err)
		}
		err = os.Rename(rootpath+"1.png", rootpath+"2.png")
		if err != nil {
			logs.Error(err)
		}
		err = os.Rename(rootpath+"test.png", rootpath+"1.png")
		if err != nil {
			logs.Error(err)
		}
		return "1.png", nil
	}
	// 删除6.jpg，1-5.jpg重命名，test.jpg重命名为1.jpg
	err = os.Remove(rootpath + "6.35659cb8.png")
	if err != nil {
		logs.Error(err)
	}
	err = os.Rename(rootpath+"5.e98fb380.png", rootpath+"6.35659cb8.png")
	if err != nil {
		logs.Error(err)
	}
	err = os.Rename(rootpath+"4.7cf2c2f9.png", rootpath+"5.e98fb380.png")
	if err != nil {
		logs.Error(err)
	}
	err = os.Rename(rootpath+"3.55ccb3bd.png", rootpath+"4.7cf2c2f9.png")
	if err != nil {
		logs.Error(err)
	}
	err = os.Rename(rootpath+"2.5a85c2ac.png", rootpath+"3.55ccb3bd.png")
	if err != nil {
		logs.Error(err)
	}
	err = os.Rename(rootpath+"1.5d5bcc7d.png", rootpath+"2.5a85c2ac.png")
	if err != nil {
		logs.Error(err)
	}
	err = os.Rename(rootpath+"test.png", rootpath+"1.5d5bcc7d.png")
	if err != nil {
		logs.Error(err)
	}
	return "1.5d5bcc7d.png", nil
}

// @Title Get
// @获取水利信息详情
// @Success 200 success
// @Failure 403 is empty
// @router /detail [get]
//func (this *WaterInformationController) GetWaterInformationDetail() {
//	var r contributeWaterAmountRes
//	id, err := this.GetInt64("id")
//	if err != nil {
//		logs.Error(err)
//	}
//	r.Data, r.Code = models.GetWaterInformationDetail(id)
//	r.Msg = log.CodeMap[r.Code]
//	this.Data["json"] = r
//	this.ServeJSON()
//}

// @Title Post
// @修改水利信息
// @Success 200 success
// @Failure 403 is empty
// @router /alter [post]
//func (this *WaterInformationController) Alter() {
//	var alterInfo models.AlterReq
//	var r contributeWaterAmountRes
//	this.parsingAlterParam(&alterInfo)
//	r.Code = models.Alter(alterInfo)
//	r.Msg = log.CodeMap[r.Code]
//	this.Data["json"] = r
//	this.ServeJSON()
//}

// 处理参数 文件存入磁盘 以path作为返回值
//func (this *WaterInformationController) parsingAlterParam(uploadInfo *models.AlterReq) {
//	uploadInfo.Id, _ = this.GetInt64("id")
//	uploadInfo.CompanyName = this.GetString("company_name")
//	uploadInfo.Principal = this.GetString("principal")
//	uploadInfo.TelephoneNumber = this.GetString("telephone_number")
//	uploadInfo.FaxNumber = this.GetString("fax_number")
//	uploadInfo.PostCode = this.GetString("post_code")
//	uploadInfo.Address = this.GetString("address")
//	uploadInfo.Level = this.GetString("level")
//	uploadInfo.Introduction = this.GetString("introduction")
//	//  处理文件参数
//	tmpfile, fheader, err := this.Ctx.Request.FormFile("file")
//	if err != nil {
//		return
//	}
//	defer tmpfile.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
//	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
//	index := strings.LastIndexByte(fheader.Filename, '.')
//	path := beego.AppConfig.String("waterInformationImagePath") + timeStr + fheader.Filename[index:]
//	err = this.SaveToFile("file", path)
//	if err != nil {
//		logs.Error(err)
//	}
//	uploadInfo.Path = timeStr + fheader.Filename[index:]
//	return
//}
