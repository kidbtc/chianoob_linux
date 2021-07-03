package chia

import (
	"chianoob_linux/models"
	"chianoob_linux/myfunc"
	"fmt"
	_ "fmt"
	"os/exec"
	"time"
	_ "time"
)

func init() {
	Chiainit()
}

func Chiainit() {
	fmt.Println("The system Version:" + models.Version)   //打印版本号
	fmt.Println("The system LocalIp:" + myfunc.LocalIp()) //本地ip
	myfunc.KillProcess("chia_plot")                       //p盘残余
	myfunc.KillProcess("hpool-miner-chia")                //关闭挖矿进程 依赖宝塔重启 读取配置
	if myfunc.CheckFile("/conf.ini") {                    //ini 初始化
		myfunc.CMD(`cat>>conf.ini`, "")
	} //检查配置文件 不存在就创建
	time.Sleep(30 * time.Second)
	Initconfig(ChiaConfig) //初始化
	for i := 0; i < 3; i++ {
		cmd := exec.Command("sh", "-c", "/www/server/chiabee/mount_disk.sh")
		cmd.Output()
		fmt.Println("硬盘加载中", i)
		time.Sleep(5 * time.Second)
	}

	go func() {
		for {
			PlooterRun() //P盘开始
			time.Sleep(10 * time.Second)
		}
	}()
}
