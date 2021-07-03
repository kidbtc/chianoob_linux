package network

import (
	"chianoob_linux/chia"
	"chianoob_linux/models"
	"chianoob_linux/myfunc"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"net"
	"strconv"
	"strings"
	"time"
)

func init() {
	porname = beego.AppConfig.String("porname")
	//go chia.Chiainit()
}

var Conn net.Conn
var SocketCount = 0
var HasData chan string
var Chiastr string
var porname string

//初始化定时任务 设置定时时长
func InitTask() {
	tk := toolbox.NewTask("generateWarning", "0/60 * * * * *", generateWarning)
	toolbox.AddTask("generateWarning", tk)
}

//更新状态 Chiastr全局字符串
func generateWarning() error {
	if models.Task == 0 {
		//fmt.Println("定时检测线程开启:", time.Now().Format("2006-01-02 15:04:05"))
		Chiastr = ChiaPost()
	}
	return nil
}

func ChiaPost() string {
	var st string
	clinet := models.Client{}                       //定义客户端变量
	clinet.Name = beego.AppConfig.String("porname") //读取配置文件中的自定名称
	clinet.IP = myfunc.LocalIp()                    //获取本地ip
	if clinet.Name == "" {                          //当未设置自定义名称时，用ip代替自定义名称
		clinet.Name = clinet.IP
	}
	clinet.Finish = chia.HHDFinish()              //获取当前完P盘成度
	clinet.NowHHD = chia.Nowhhd                   //当前在P的硬盘 数据   -盘符 -总 -已使用 -剩余
	clinet.PNum = strconv.Itoa(models.Pnum * 105) //单文件 105g * 轮数
	clinet.UpTime = models.Uptime                 //最后打印的时间
	clinet.Ptime = models.Ptime                   //一轮完成的时间
	//当前硬盘名称
	//float64(clinet.NowHHD.Used)/float64(GB)
	path := clinet.NowHHD.Path                                                                          //当前P盘路径
	path = path[strings.LastIndex(path, "/")+1:]                                                        //路径清洗  /home/cykj/HHD0 清洗成 HHD0
	used := strconv.FormatFloat(float64(clinet.NowHHD.Used)/float64(chia.GB), 'f', 0, 64)               //格式化使用磁盘空间
	all := strconv.FormatFloat(float64(clinet.NowHHD.All)/float64(chia.GB), 'f', 0, 64)                 //格式化总磁盘空间
	rate := strconv.FormatFloat(float64(clinet.NowHHD.Used)/float64(clinet.NowHHD.All)*100, 'f', 0, 64) //格式化使用百分比
	hhddisk := path + "/" + used + "/" + all + "/" + rate
	//名称 ip 状态 完成度 当前目标 今日量 矿池监控 设置 版本号 最后更新时间 单轮使用时间  /home
	st = "chia" + "|" + clinet.Name + "|" + clinet.IP + "|" + clinet.Error + "|" + clinet.Finish + "|" + hhddisk + "|" + clinet.PNum + "|" + clinet.PoolPro + "|" + models.Version + "|" + clinet.UpTime + "|" + clinet.Ptime + "|" + models.Todesk
	//st = "chia|192.168.1.2|192.168.1.2|正常|5/12|HHD1/1560/144000/16|32||1.16"
	return st
}

//建立套接字连接
func Chiasocket() {
	hostInfo := models.Serverip
	fmt.Println("server ip:", hostInfo)
	for {
		var err error
		Conn, err = net.Dial("tcp", hostInfo)
		fmt.Print("connect (", hostInfo)
		if err != nil {
			fmt.Println(") fail")
		} else {
			fmt.Println(") ok")
			defer Conn.Close()
			go DoTaskSend(Conn)
			DoTaskRec(Conn)
			time.Sleep(3 * time.Second)
		}
	}
}

//不为空时直接发送至服务器
func DoTaskSend(conn net.Conn) {
	for {
		if Chiastr != "" { //不为空时直接发送至服务器
			_, err := conn.Write([]byte(Chiastr)) //发送函数
			if err != nil {
				fmt.Println("recv data error")
				break
			} else {
				Chiastr = ""
			}
		}
		time.Sleep(1 * time.Second)
	}
}

//接收函数
func DoTaskRec(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("对方断开或者异常", err)
			return
		}
		//启动升级函数
		if string(buf[:n]) == "upfile" {
			fmt.Println("recv msg : upfile")
			myfunc.Upfile("")
		}
		//启动重启函数
		if string(buf[:n]) == "reboot" {
			fmt.Println("recv msg : reboot")
			myfunc.Reboot()
		}
		//启动挖矿程序
		if string(buf[:n]) == "resthpool" {
			fmt.Println("recv msg : resthpool")
			myfunc.RestHpool()
		}
		time.Sleep(1 * time.Second)
	}
}
