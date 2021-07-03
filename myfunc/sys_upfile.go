package myfunc

import (
	"chianoob_linux/models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type files struct {
	filename     string
	filemd5      string
	downloadpath string
}

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Read(p)
	r.Current += int64(n)
	fmt.Println("\r进度 %.2f%%", float64(r.Current*10000/r.Total*100)/100)
	return
}

//upfile|/dsads/dasdsa/home_md5_www.baidu.com/down|/dsads/dasdsa/home_md5_www.baidu.com/down

//-----------------------------------文件更新--------------------------------------//
func Upfile(url string) bool {
	fmt.Println("start upfile")
	reb := 0
	if url == "" {
		url = models.UpFileip
	}
	var upfile []models.UpFiles
	url = "http://" + url + "/filemd5"
	fmt.Println("upfile url:", url)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	//转换为json格式
	if err := json.Unmarshal(body, &upfile); err == nil {
		//循环解析
		for _, temp := range upfile {
			if GetFilesMd5(temp.PathName) != temp.Md5 { //GetFilesMd5（）取本地文件的MD5与服务器返回的MD5值相比
				fmt.Println(temp.PathName, " Md5 Key is Different,Start Download")
				CMD("rm -rf ", temp.PathName)                //因为占用问题先删除本地文件，直接覆盖会提示占用
				err := DwonLoadFile(temp.Url, temp.PathName) //下载文件
				if err != nil {
					fmt.Println("Download err:", err, "  ", temp.PathName)
					break
				}
				fmt.Println("DownLoad finish !  ", temp.PathName)
				CMD("chmod 777 ", temp.PathName) //修改文件权限
				reb++
			}
		}
		time.Sleep(10 * time.Second)
		if reb > 0 {
			fmt.Println(" Restart chianoob_linux")
			KillProcess("chianoob_linux") //杀死自进程 让宝塔守护程序自启
		} else {
			fmt.Println(" No UpFile")
		}
		//fmt.Println(dat["status"])
	} else {
		fmt.Println(err)
	}
	return true
}

//比对MD5下载文件
func DwonLoadFile(url, filename string) error {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	if strings.Index(url, "http://") == -1 {
		url = "http://" + url
	} //当地址仅为IP时会报错
	fmt.Println("Download Url:", url, " FilePatch:", filename)
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = r.Body.Close() }()

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	n, err := io.Copy(f, r.Body)
	panic("下载异常")
	fmt.Println(n, err, " reboot!")
	return nil
	//sudo chmod  777
}

//-----------------------------------文件更新--------------------------------------//
