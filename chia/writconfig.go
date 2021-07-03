package chia

import (
	"chianoob_linux/models"
	"chianoob_linux/myfunc"
	"fmt"
	"github.com/astaxie/beego"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var ChiaConfig = `/home/cykj/chia-plotter/build/chia_plot -n -1 -r THR -u BUK -t SSDPT -2 SSDP2 -d HHDP  -p PPK -f FPK`

var HpoolYaml = `token: ""
path:
- /home/cykj/HHD0
- /home/cykj/HHD1
- /home/cykj/HHD2
- /home/cykj/HHD3
- /home/cykj/HHD4
- /home/cykj/HHD5
- /home/cykj/HHD6
- /home/cykj/HHD7
- /home/cykj/HHD8
- /home/cykj/HHD9
- /home/cykj/HHD10
- /home/cykj/HHD11
minerName: Pname
apiKey: Poolkey
cachePath: ""
deviceId: ""
extraParams: {}
log:
  lv: info
  path: ./log/
  name: miner.log
url:
  proxy: ""
scanPath: false
scanMinute: 60
`

func Initconfig(chiaConfig string) (commend string, err error) {
	pool()
	myfunc.GetTodesk()
	if len(models.Todesk) > 4 {
		models.Todesk = models.Todesk[:len(models.Todesk)-4]
	}
	fmt.Println("Todesk：", models.Todesk)
	models.Pnum, _ = strconv.Atoi(models.Conf.GetValue("chia", "today_pnum"))

	ssdpt, ssdp2, hhdp, err := GetRunPath("1")
	if err != nil {
		fmt.Println(err)
		cmd := exec.Command("sh", "-c", `/www/server/chiabee/mount_disk.sh`)
		cmd.Output()
		time.Sleep(30 * time.Second)
		return "", err
	}
	ssdpt = ssdpt + "chia1/"
	ssdp2 = ssdp2 + "chia1/"

	myfunc.CMD("rm -rf", "/var/log/syslog")
	myfunc.CMD("rm -rf", ssdpt)
	myfunc.CMD("rm -rf", ssdp2)
	myfunc.CMD("mkdir", ssdpt)
	myfunc.CMD("mkdir", ssdp2)
	THR := beego.AppConfig.String("THR")
	FPK := beego.AppConfig.String("FPK")
	PPK := beego.AppConfig.String("PPK")
	BUK := beego.AppConfig.String("BUK")
	fmt.Println("get path ssdt:", ssdpt, " ssdp2:", ssdp2, " hhdp:", hhdp)
	chiaConfig = strings.Replace(ChiaConfig, "FPK", FPK, -1)
	chiaConfig = strings.Replace(chiaConfig, "PPK", PPK, -1)
	chiaConfig = strings.Replace(chiaConfig, "THR", THR, -1)
	chiaConfig = strings.Replace(chiaConfig, "BUK", BUK, -1)
	chiaConfig = strings.Replace(chiaConfig, "SSDPT", ssdpt, -1)
	chiaConfig = strings.Replace(chiaConfig, "SSDP2", ssdp2, -1)
	commend = strings.Replace(chiaConfig, "HHDP", hhdp, -1)
	fmt.Println("commend : ", commend)
	return commend, nil
}

func pool() {
	fmt.Println("Phool start")
	Hpool := beego.AppConfig.String("Hpool")
	poolpath := beego.AppConfig.String("hpoolmincof")
	if poolpath != "" {
		pname := beego.AppConfig.String("porname")
		if pname == "" {
			pname = myfunc.LocalIp()
		}
		hpoolyaml := strings.Replace(HpoolYaml, "Pname", pname, -1)
		hpoolyaml = strings.Replace(hpoolyaml, "Poolkey", Hpool, -1)
		myfunc.WriteToFile(poolpath, hpoolyaml)
	}
}
