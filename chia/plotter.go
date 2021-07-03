package chia

import (
	"bufio"
	"chianoob_linux/models"
	"chianoob_linux/myfunc"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

func PlooterRun() error {
	str, err := Initconfig(ChiaConfig)
	if err != nil {
		fmt.Println(err)
		return err
	}
	Command(str)
	return nil
}

func Command(cmd string) error {
	//c := exec.Command("cmd", "/C", cmd) 	// windows
	c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				fmt.Println("err :cmd read err")
				break
			}
			fmt.Print(readString)
			models.Uptime = time.Now().Format("2006-01-02 15:04:05")

			if strings.Index(readString, "fwrite() failed") != -1 {
				myfunc.KillProcess("chia_plot")
				fmt.Println("err : disk full return")
				return
			}
			if strings.Index(readString, "Total plot creation time was") != -1 {
				tempstr := myfunc.GetBetweenStr1(readString, "Total plot creation time was ", " sec")
				tempint, _ := strconv.ParseFloat(tempstr, 64)
				models.Ptime = strconv.FormatFloat(tempint/float64(60), 'f', 0, 64)
				fmt.Println("第", models.Pnum, "轮运行结束：", models.Ptime)
				pnumAdd()
			}
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}

func pnumAdd() {

	if models.PDate != time.Now().Format("2006-01-02") {
		models.Pnum = 1
		models.PDate = time.Now().Format("2006-01-02")
		models.Conf.SetValue("chia", "pdate", models.PDate)
		models.Conf.SetValue("chia", "today_pnum", strconv.Itoa(models.Pnum))
		return
	} else {
		models.Pnum++
	}
	models.Conf.SetValue("chia", "today_pnum", strconv.Itoa(models.Pnum))
}
