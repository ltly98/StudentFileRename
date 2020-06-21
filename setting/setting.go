package setting

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Setting struct {
	StudentPath string
	OperationPath string
}

//显示信息
func ShowSetting() {
	SettingInfo, err := ioutil.ReadFile("setting.json")

	if err != nil {
		fmt.Printf("打开setting文件失败：%v\n,请重试!", err)
		return 
	}

	var ReadSetting Setting 

	err2 := json.Unmarshal(SettingInfo, &ReadSetting)
	if err2 != nil {
		fmt.Println("反序列化失败：%v\n",err2)
		return 
	}
	fmt.Println("---------------路径信息-----------------")
	fmt.Printf("操作路径：%v\n", ReadSetting.OperationPath)
	fmt.Printf("学生路径：%v\n", ReadSetting.StudentPath)
	fmt.Println("----------------------------------------")
}


//检查setting文件
func CheckSetting() {

	//标记文件是否存在格式错误等问题
	var IsErr bool = false

	SettingInfo, err := ioutil.ReadFile("setting.json")

	if err != nil {
		fmt.Printf("打开setting文件失败：%v\n", err)
		IsErr = true
	} 

	//如果文件存在，用json解析来检查文件
	if !IsErr{

		var ReadSetting Setting

		err2 := json.Unmarshal(SettingInfo,&ReadSetting)

		if err2 != nil {
			fmt.Printf("解析setting错误：%v\n", err2)
			IsErr = true
		}
	}

	if IsErr {

		fmt.Println("检查setting.json文件出错，开始创建或重置文件！")

		InitInfo := Setting{StudentPath: "student.json",OperationPath:".\\"}

		WriteInfo,err3 := json.Marshal(InitInfo)
		if err3 != nil {
			fmt.Println("序列化失败：%v\n",err3)
			return 
		}

		err4 := ioutil.WriteFile("setting.json", []byte(WriteInfo), 0666)

		if err4 != nil {
			fmt.Printf("重置setting.json失败：%v\n", err4)
			return 
		}				
	}

	fmt.Println("setting.json文件检查完毕！")

	ShowSetting()
}

//修改setting文件
func ModifySetting() {

	var ModifyOperationPath string = ""
	var ModifyStudentPath string = ""

	for{
		fmt.Print("请输入要修改的操作路径（不想修改0代替，不能为空）：")
		fmt.Scan(&ModifyOperationPath)
		if ModifyOperationPath != ""{
			break
		}
	}

	//连续使用scanf后面会吸收前面的内容（c语言老问题了），所以我选择文字填充一下，序号意思是只有第一次才填充
	num := 0

	for{
		num++
		fmt.Print("请输入要修改的学生路径（不想修改0代替，不能为空）：")
		fmt.Scan(&ModifyStudentPath)
		if ModifyStudentPath != ""{
			break
		}
		if num == 1{
			fmt.Println("此处c语言老问题，用文字填充一下")
		}
	}


	SettingInfo, err := ioutil.ReadFile("setting.json")

	if err != nil {
		fmt.Printf("打开setting文件失败：%v\n,请重试!", err)
		return 
	}

	var WriteSetting Setting 

	err2 := json.Unmarshal(SettingInfo, &WriteSetting)
	if err2 != nil {
		fmt.Println("反序列化失败：%v\n",err2)
		return 
	}

	//此处else方便显示用

	if ModifyOperationPath != "0"{
		WriteSetting.OperationPath = ModifyOperationPath
	}else{
		ModifyOperationPath = WriteSetting.OperationPath
	}

	if ModifyStudentPath != "0"{
		WriteSetting.StudentPath = ModifyStudentPath
	}else{
		ModifyStudentPath = WriteSetting.StudentPath
	}
	
	WriteInfo,err3 := json.Marshal(WriteSetting)
	if err3 != nil {
		fmt.Println("序列化失败：%v\n",err3)
		return 
	}

	err4 := ioutil.WriteFile("setting.json", []byte(WriteInfo), 0666)
	if err4 != nil {
		fmt.Printf("写入setting错误：%v\n,请重试!", err4)
		return 
	}

	fmt.Println("---------------路径修改完毕-------------")
	fmt.Printf("当前操作路径：%v\n", WriteSetting.OperationPath)
	fmt.Printf("当前学生路径：%v\n", WriteSetting.StudentPath)
	fmt.Println("----------------------------------------")
}
