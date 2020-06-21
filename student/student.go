package student

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Num  string
}

func ShowStudent() {
	StudentInfo, err := ioutil.ReadFile("student.json")

	if err != nil {
		fmt.Printf("打开student文件失败：%v\n", err)
		return 
	} 

	var ReadStudent []Student

	err2 := json.Unmarshal(StudentInfo, &ReadStudent)

	if err2 != nil {
		fmt.Printf("反序列化student.json文件失败：%v\n", err2)
		return 
	} 

	fmt.Println("---------------学生信息-----------------")
	for i:=0;i<len(ReadStudent);i++{
		fmt.Printf("%v: %v %v\n", i+1,ReadStudent[i].Num,ReadStudent[i].Name)
	}
	fmt.Println("----------------------------------------")
}

func CheckStudent(){

	//标记文件是否存在格式错误等问题
	var IsErr bool = false

	StudentInfo, err := ioutil.ReadFile("student.json")

	if err != nil {
		fmt.Printf("打开student文件失败：%v\n", err)
		IsErr = true
	} 

	//如果文件存在，检查文件
	if !IsErr {
		var ReadStudent []Student

		err2 := json.Unmarshal(StudentInfo, &ReadStudent)

		if err2 != nil {
			fmt.Printf("文件内容出错：%v\n", err2)
			IsErr = true 
		}
	}

	if IsErr {
		fmt.Println("检查student.json文件出错，开始创建或重置文件！")

		InitInfo := [...]Student{{Name:"zhangsan",Num:"001"},{Name:"lisi",Num:"002"}}

		WriteInfo,err3 := json.Marshal(InitInfo)
		if err3 != nil {
			fmt.Println("序列化失败：%v\n",err3)
			return 
		}

		err4 := ioutil.WriteFile("student.json", []byte(WriteInfo), 0666)

		if err3 != nil {
			fmt.Printf("重置student.json失败：%v\n", err4)
			return 
		}
	}
	fmt.Println("student.json检查完毕!")
	ShowStudent()
}