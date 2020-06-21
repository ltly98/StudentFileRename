package operation

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Setting struct {
	StudentPath string
	OperationPath string
}

type Student struct {
	Name string
	Num  string
}

func CheckOperationPath() (Setting,bool) {

	var ReadSetting Setting

	SettingInfo, err := ioutil.ReadFile("setting.json")

	if err != nil {
		fmt.Printf("打开setting文件失败：%v\n,请重试!", err)
		return ReadSetting,false
	}

	err2 := json.Unmarshal(SettingInfo, &ReadSetting)
	if err2 != nil {
		fmt.Println("反序列化失败：%v\n",err2)
		return ReadSetting,false
	}

	_,err3 := os.Stat(ReadSetting.OperationPath)

	if err3 != nil{
		fmt.Println("操作路径出错！")
		return ReadSetting,false
	}else{
		return ReadSetting,true
	}
}

func RenameOperation(student Student) {
	Setting,IsExist := CheckOperationPath()

	if IsExist {
		File,err := ioutil.ReadDir(Setting.OperationPath)

		var IsModify bool = false

		if err != nil {
			fmt.Printf("读取路径失败：%v\n", err)
			return
		}

		for _,file := range File{
			match,_ := regexp.MatchString("("+student.Num+"|"+student.Name+")", file.Name())

			fileType := strings.Split(file.Name(), ".")

			if match {
				os.Rename(Setting.OperationPath+"\\"+file.Name(), Setting.OperationPath+"\\"+student.Num+student.Name+"."+fileType[len(fileType)-1])
				fmt.Println(student.Name+"的文件修改成功！")
				IsModify = true
			}
		}
		
		if !IsModify{
				fmt.Println(student.Name+"的文件不存在！")
			}	
	}

}