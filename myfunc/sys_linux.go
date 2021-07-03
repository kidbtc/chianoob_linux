package myfunc

import (
	"bufio"
	"chianoob_linux/models"
	"io"
	"os/exec"
	"strings"
	"sync"
)

func CMD(str, path string) {
	cmd := exec.Command("sh", "-c", str+" "+path)
	cmd.Output()
}

func KillProcess(name string) {
	//kill -9 $(ps -ef|grep 进程名关键字|grep -v grep|awk '{print $2}')
	cmd := exec.Command("sh", "-c", `kill -9 $(ps -ef|grep `+name+`|grep -v grep|awk '{print $2}')`)
	cmd.Output()
}
func Reboot() {
	cmd := exec.Command("sh", "-c", "reboot")
	cmd.Output()
}
func RestHpool() {
	KillProcess("hpool-miner-chia") //关闭挖矿进程 依赖宝塔重启 读取配置
}

func GetTodesk() {
	//c := exec.Command("cmd", "/C", cmd) 	// windows
	c := exec.Command("sh", "-c", "todesk") //初始化Cmd
	stdout, err := c.StdoutPipe()
	if err != nil {
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			readString = strings.Replace(readString, " ", "", -1)
			readString = strings.Replace(readString, "\r", "", -1)
			readString = strings.Replace(readString, "\n", "", -1)
			if strings.Index(readString, "ID:") != -1 {
				models.Todesk = readString
			}
			if strings.Index(readString, "Password:") != -1 {
				readString = strings.Replace(readString, "Password:", "PW:", -1)
				models.Todesk = models.Todesk + " / " + readString
				return
			}
		}
	}()

	err = c.Start()
	wg.Wait()
}
