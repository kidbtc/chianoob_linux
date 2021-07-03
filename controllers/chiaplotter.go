package controllers

import (
	"github.com/astaxie/beego"
)

type PlotterController struct {
	beego.Controller
}

func (c *PlotterController) Pstatic() {
	//pthread := models.Chia_Thread{}
	//fmt.Println(pthread)
	//chiaDStatus := myfunc.GetDiskPercent()
	//chia.GetPstatic1()
	//c.Data["CpuStatus"] = myfunc.GetCpuPercent()
	//c.Data["MemStatus"] = myfunc.GetMemPercent()
	//c.Data["DiskStatus"] = chiaDStatus
	//c.Data["chia"] = models.Chia
	c.TplName = "index.html"
}
