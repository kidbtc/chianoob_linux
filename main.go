package main

import (
	"chianoob_linux/network"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), ":Task init")
	//定时任务初始化
	network.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()
	go network.Chiasocket()
	beego.Run()
}
