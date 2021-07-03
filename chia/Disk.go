package chia

import (
	"chianoob_linux/models"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"syscall"
)

var Nowhhd = models.DiskStatus{} //硬盘信息

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// disk usage of path/disk
//读取单块硬盘目录的信息
func DiskUsage(path string) (disk models.DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return disk
	}
	disk.Path = path
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return disk
}

//获取所有SSD及HHD信息
func GetAllDiskInfo() (ssd, hhd []models.DiskStatus) {
	//获取所有SSD信息
	for i := 0; i < 5; i++ {
		path := "/home/cykj/SSD" + strconv.Itoa(i)
		disk := DiskUsage(path)
		if disk.All != 0 {
			disk.Path = path + "/"
			ssd = append(ssd, disk)
		}
	}
	//获取所有HHD信息
	for i := 0; i < 12; i++ {
		path := "/home/cykj/HHD" + strconv.Itoa(i)
		disk := DiskUsage(path)
		if disk.All != 0 && float64(disk.All)/float64(GB) > 2048 {
			disk.Path = path + "/"
			fmt.Println("Path: ", disk.Path)
			fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
			fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
			fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
			hhd = append(hhd, disk)
		}
	}
	return ssd, hhd
}

//获取可用硬盘
func GetRunPath(idx string) (ssdpt string, ssdp2 string, hhdp string, err error) {
	ssd, hhd := GetAllDiskInfo()

	if len(ssd) == 0 || len(hhd) == 0 {
		return ssdpt, ssdp2, hhdp, errors.New("error : disk number error ")
	}

	if len(ssd) == 1 {
		ssdpt = ssd[0].Path
		ssdp2 = ssd[0].Path
	}

	if len(ssd) > 1 {
		if len(ssd) > 2 {
			if idx == "1" {
				ssdpt = ssd[0].Path
			} else {
				ssdpt = ssd[2].Path
			}
		}
		ssdp2 = ssd[1].Path
	}
	//float64(temp.Free)/float64(chia.GB)
	Nowhhd = models.DiskStatus{} //硬盘信息
	for _, temp := range hhd {
		if float64(temp.Free)/float64(GB) > 120 && float64(temp.All)/float64(GB) > 2248 {
			hhdp = temp.Path
			Nowhhd = temp
			break
		}
	}
	if hhdp == "" {
		return ssdpt, ssdp2, hhdp, errors.New("error ：hhd path nil")
	}
	return ssdpt, ssdp2, hhdp, nil
}

//检查已满的硬盘
func HHDFinish() string {
	full := 0
	hhd := []models.DiskStatus{}
	for i := 0; i < 12; i++ {
		path := "/home/cykj/HHD" + strconv.Itoa(i)
		disk := DiskUsage(path)
		if disk.All != 0 && float64(disk.All)/float64(GB) > 2048 {
			disk.Path = path + "/"
			hhd = append(hhd, disk)
		}
	}

	for _, temp := range hhd {
		if strings.Index(temp.Path, Nowhhd.Path) != -1 {
			Nowhhd.Free = temp.Free
			Nowhhd.All = temp.All
			Nowhhd.Used = temp.Used
		}
		if float64(temp.Free)/float64(GB) < 106 {
			full++
		}
	}
	return strconv.Itoa(full) + "/" + strconv.Itoa(len(hhd))
}
