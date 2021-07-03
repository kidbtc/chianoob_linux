package myfunc

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//type Study struct{
//	CourseName string `yaml:"CourseName"`
//	Score 		int	   `yaml:"Score"`
//}
//type Student struct{
//	Name 		 string 	`yaml:"name"`
//	Address		 string 	`yaml:"addr"`
//	ScoreList	 []Study    `yaml:"ScoreList"`
//}

type Hpool struct {
	minerName string
	apiKey    string
}

func WriteToXml(src, minername, key string) {
	if minername == "" {
		minername = LocalIp()
	}
	hp := &Hpool{
		minerName: minername,
		apiKey:    key,
	}
	data, err := yaml.Marshal(hp) // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	checkError(err)
	err = ioutil.WriteFile(src, data, 0777)
	checkError(err)
}

//func readFromXml(src string){
//	content,err := ioutil.ReadFile(src)
//	checkError(err)
//	newStu := &Student{}
//	err = yaml.Unmarshal(content,&newStu)
//	checkError(err)
//	ScoreList := newStu.ScoreList
//	fmt.Println(newStu.Name+"的学习情况")
//	for _,v := range ScoreList {
//		fmt.Println("Course:" + v.CourseName + "\tScore:" + strconv.Itoa(v.Score))
//	}
//}
