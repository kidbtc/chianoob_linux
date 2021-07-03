package models

import (
	"github.com/astaxie/beego"
	"github.com/widuu/goini"
)

type DiskStatus struct {
	Path string `json:"path"`
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

var Version = "1.2.9"
var Todesk = ""
var Conf = goini.SetConfig("/conf.ini")
var Task = 0
var Color = []string{"amber", "blue", "blush", "coral", "cyan", "green", "khaki", "parpl", "pink", "salmon", "seagreen", "slategray", "turquoise"}
var ChiaDiskIO = make(map[string]interface{})
var CleanSize = 1024000 //字节
var Canused = uint64(180)
var Path = beego.AppConfig.String("chiapath") //取所在路径
//var Version = Conf.GetValue("chia", "Version")
var PDate = Conf.GetValue("chia", "pdate")
var Pnum = 0
var Ptime = ""
var Uptime = ""
var Serverip = beego.AppConfig.String("serverip")
var UpFileip = beego.AppConfig.String("upfileip")
var Reb = 0

//P盘线程

//硬盘

type UpFiles struct {
	PathName string
	Md5      string
	Url      string
}

//机械剩余                          昨日P图    今日P图            进程总数      最慢时间 值守进程 挖矿进程 异常  设置
//C-2.2/2.3/2.2/2.3/2.2/2.32.2/2.3   3T          5T                11            11       1       1        0    ---
type Client struct {
	IP        string
	Name      string
	Finish    string     //完成状态
	Todatwork string     //今日计数
	NowHHD    DiskStatus //硬盘信息
	UpTime    string
	PoolPro   string //挖矿进程
	Error     string //异常
	PNum      string //当前轮数
	Ptime     string //单轮时间
}
