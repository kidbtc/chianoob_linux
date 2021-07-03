package myfunc

import (
	"chianoob_linux/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func EnumBetweenStr(str, start, end string) []string {
	ret := []string{}
	i := 0
	for {
		n := strings.Index(str, start)
		if n == -1 {
			n = 0
			return ret
		}
		s1 := str
		s1 = string([]byte(s1)[n:])

		m := strings.Index(s1, end)
		if m == -1 {
			m = len(s1)
		}
		ret = append(ret, string([]byte(s1)[:m]))
		str = s1[m:]
		i++
		if i > 100 {
			return ret
		}
	}
	return ret
}
func GetBetweenStr1(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start) // 增加了else，不加的会把start带上
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
func GetBetweenStr(str, start, end string, contain int) string {
	n := strings.Index(str, start)
	if n == -1 {
		return ""
	}

	if contain == 1 {
		str = string([]byte(str)[n:])
	} else {
		str = string([]byte(str)[n+len(start):])
	}

	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
func GetLastBetweenStr(str, start, end string, contain int) string {
	n := strings.LastIndex(str, start)
	if n == -1 {
		return ""
	}

	if contain == 1 {
		str = string([]byte(str)[n:])
	} else {
		str = string([]byte(str)[n+len(start):])
	}

	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
func substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return dir + "\\"
}

//文件是否存在
func CheckFile(filepath string) (exist bool) {
	// 是一个文件的指针
	fileInfo, e := os.Stat(filepath)
	if fileInfo != nil && e == nil {
		exist = true
		// 判断文件是否不存在
	} else if os.IsNotExist(e) {
		exist = false
	}
	// 这里返回可以带返回值的名，也可以不带
	return
	//return exist
}
func Changetext(filepath, str1, str2, str3 string) (bool bool) {
	s := ""
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error", err)
		return false
	}

	str := string(data)
	temp_array := strings.Split(str, "\n")
	for i := 0; i < len(temp_array); i++ {
		temp_array[i] = strings.Replace(temp_array[i], "\r", "", -1)
		if strings.Index(temp_array[i], str1) != -1 {
			temp_array[i] = str1 + str2 + str3
			fmt.Println("changetetx: ", temp_array[i])
		}
		s = s + temp_array[i] + "\n"
	}

	dstFile, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer dstFile.Close()

	dstFile.WriteString(s)
	return true
}

//覆盖写入
func WriteToFile(fileName string, content string) bool {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		fmt.Println("write succeed!")
		defer f.Close()
	}
	return true
}

//获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) string {
	var t string
	//var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Format("2006-01-02 15:04:05")
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Format("2006-01-02 15:04:05")
	}
	t = fi.ModTime().Format("2006-01-02 15:04:05")
	return t
}

//获取文件信息
func GetFileInfo(disk string) map[string]string {
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(disk + ":\\")
	if err != nil {
		log.Fatal(err)
	}
	//存储文件名 和 时间
	mp := make(map[string]string)
	for _, fi := range fileInfoList {
		if !fi.IsDir() { // 遍历 文件
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".fgcc")
			if ok {
				//当前时间减一天
				m, _ := time.ParseDuration("-24h")
				beforDay := time.Now().Add(m).Format("2006-01-02")
				//转换成时间戳
				beginTimeNum := handleTime(beforDay)
				//本地文件创建时间
				fileDate := fi.ModTime().Format("2006-01-02")
				fileNum := handleTime(fileDate)
				//当前时间
				endTimeNum := handleTime(time.Now().Format("2006-01-02"))
				//将今日和昨日生成的文件打印出来
				if fileNum >= beginTimeNum && fileNum <= endTimeNum {
					fullName := fi.Name()
					//fmt.Println(fi.Size())
					fileTime := fi.ModTime().Format("2006-01-02")
					mp[fullName] = fileTime
				}
			}
		}
	}
	return mp
}

//转时间戳方法
func handleTime(tm string) int64 {
	t1, _ := time.ParseInLocation("2006-01-02", tm, time.Local)
	fileNum := t1.Unix()
	return fileNum
}

//根据路径找到该路径下的所有文件，并返回
func GetFilesFromDir(path string) []models.UpFiles {
	files := []models.UpFiles{}
	mfile := models.UpFiles{}
	//从目录path下找到所有的目录文件
	allDir, err := ioutil.ReadDir(path)
	//如果有错误发生就返回nil，不再进行查找
	if err != nil {
		return files
	}
	//遍历获取到的每一个目录信息
	for _, dir := range allDir {
		if dir.IsDir() {
			//如果还是目录，就继续向下进行递归查找,并追加到返回切片中
			retfiles := GetFilesFromDir(path + "/" + dir.Name())
			for _, temp := range retfiles {
				files = append(files, temp)
			}
			continue
		}
		//如果不是目录,就读取该文件并加入到files中
		fileName := path + "/" + dir.Name()
		file, fileErr := ioutil.ReadFile(fileName)
		if fileErr != nil {
			return files
		}
		mfile.PathName = fileName
		mfile.Md5 = ByteToMd5(file)
		//打印文件名和MD5
		fmt.Println(mfile.PathName + " 的MD5为 " + mfile.Md5)
		files = append(files, mfile)
	}
	return files
}

//根据路径找到该文件，并返回
func GetFilesMd5(path string) string {
	mfile := models.UpFiles{}
	file, fileErr := ioutil.ReadFile(path)
	if fileErr != nil {
		return ""
	}
	mfile.PathName = path
	mfile.Md5 = ByteToMd5(file)
	//打印文件名和MD5
	fmt.Println(mfile.PathName + " 的MD5为 " + mfile.Md5)
	return mfile.Md5
}

//[]byte转md5
func ByteToMd5(fileByte []byte) string {
	has := md5.Sum(fileByte)
	//用新的切片存放
	has2 := has[:]
	md5Str := hex.EncodeToString(has2)
	return md5Str
}
