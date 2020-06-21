package main

import (
	"manage/setting"
	"manage/student"
	"manage/operation"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func main() {

	var InputCmd string

	fmt.Println("----------------文件批量修改-----------------")
	fmt.Println("相关说明：")
	fmt.Println("学生名单请手动修改")
	fmt.Println("第一次使用请修改操作路径，学生路径不建议修改")
	fmt.Println("---------------------------------------------")

	setting.CheckSetting()
	student.CheckStudent()

	fmt.Println("---------------------------------------------")

	for {
		
		fmt.Println("需要修改输入1 不需要输入非空任意值")
		fmt.Print("是否要修改路径？：")
		fmt.Scan(&InputCmd)

		if InputCmd != "1" {
			break
		}else{
			setting.ModifySetting()
			break
		}

		
	}

	fmt.Println("---------------------------------------------")

	for {
		fmt.Println("需要修改文件名输入1 不需要输入非空任意值")
		fmt.Print("是否要进行修改文件名？：")
		fmt.Scan(&InputCmd)

		if InputCmd != "1" {
			break
		}else{

			var student []operation.Student

			ReadStudentInfo,err := ioutil.ReadFile("student.json")

			if err != nil {
				fmt.Printf("读取学生名单出错：%v\n", err)
				return 
			}

			err2 := json.Unmarshal(ReadStudentInfo, &student)

			if err2 != nil {
				fmt.Printf("学生名单反序列化失败：%v\n", err2)
				return 
			}

			for i:=0;i<len(student);i++{
				operation.RenameOperation(student[i])
			}

			break

		}
		
	}

	fmt.Println("---------------------------------------------")

}
