package MQPassword

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	STORAGE_NAME 	= "mq.bin"
	APP_DEF_PASSWD 	= "e10adc3949ba59abbe56e057f20f883e"
	APP_NAME 		= "MQPassword管理器"
	APP_VERSION 	= "0.1"
	APP_AUTHOR 		= "soeasyd@gmail.com"
)

//数据属性
type InputContent struct {
	ID 		 string
	Name     string
	Account  string
	Password string
	Domain   string
	Comment  string
	CreateAt string
}

type Model struct {
	Version 	string 			`版本`
	Time 		int64			`更新时间`
	Count 		int				`数据量`
	Data 		[]InputContent	`数据`
}

//数据存储
var db_path string

//加密密钥(至少16位)
var MAIN_KEY string

//系统初始化
var APP_FIRST_INIT = false

//全局数据
var DATABASE Model

//系统初始化
func (this *Model)EnvInit() {
	path, _ := os.Getwd()

	//设置当前路径
	db_path = path + "/" + STORAGE_NAME

	_, err := os.Stat(db_path)

	if err != nil {
		//不存在则创建
		fd, err := os.Create(db_path)
		defer fd.Close()

		if err != nil {
			log.Fatal(err)
		}

		//初始化内容
		data := Model{Version:dataVersion(), Time:time.Now().Unix(), Count:0, Data:[]InputContent{}}

		//默认密码
		MAIN_KEY = APP_DEF_PASSWD

		cryptInstance := new(CryptService)

		fmt.Println(string(jsonEncode(data)), data)

		database := cryptInstance.Encrypt(jsonEncode(data))

		writeToFile(database, db_path)

		fmt.Println("加密内容:", database)

		APP_FIRST_INIT = true
	}
}

//密码校验
func (this *Model)PasswordCheck(password string) bool {
	if APP_FIRST_INIT == true {
		if md5Encrypt(password) == APP_DEF_PASSWD {
			return true
		}

		return false
	}

	//设置密钥
	MAIN_KEY = md5Encrypt(password)

	database := readFromFile(db_path)

	fmt.Println("读取加密:", database)

	cryptInstance := new(CryptService)
	databaseDecrypt := cryptInstance.Decrypt(database)

	//解密失败
	if string(databaseDecrypt) == "" {
		MAIN_KEY = ""

		return false
	}
	log.Println("解码:", string(databaseDecrypt))

	//解码
	jsonDecode(databaseDecrypt, &DATABASE)

	log.Println("数据:", DATABASE, DATABASE.Time)

	return true
}

//数据保存
func (this *Model)Save(data InputContent) bool {
	data.CreateAt = string(time.Now().Format(DATE_FORMAT))

	if len(DATABASE.Data) == 0 {
		//新增数据
		data.ID = dataGenID()
		DATABASE.Data = append(DATABASE.Data, data)
		DATABASE.Count++
	} else {
		isFind := false
		for index, item := range DATABASE.Data {
			if DATABASE.Data[index].Name == data.Name {
				DATABASE.Data[index] = data
				isFind = true
			}

			fmt.Println("创建：", item)
		}

		if isFind == false {
			data.ID = dataGenID()
			DATABASE.Data = append(DATABASE.Data, data)
			DATABASE.Count++
		}
	}


	log.Println("数据库:", DATABASE)

	return true
}

//数据列表
func (this *Model)List() []InputContent{
	if len(DATABASE.Data) == 0 {
		return []InputContent{}
	}
	return DATABASE.Data
}

//属性
func (this *Model) Item(name string) InputContent {
	for _, item := range DATABASE.Data {
		if item.Name == name {
			return item
		}
	}

	return InputContent{}
}

//存储数据
func (this *Model)Store() bool {
	if MAIN_KEY == "" {
		return false
	}

	DATABASE.Version = dataVersion()
	DATABASE.Time = time.Now().Unix()
	if DATABASE.Count == 0 {
		DATABASE.Count = len(DATABASE.Data)
	}

	//加密
	cryptInstance := new(CryptService)
	database := cryptInstance.Encrypt(jsonEncode(DATABASE))
	//写入文件
	writeToFile(database, db_path)

	return true
}

//密码更新
func (this *Model)PasswordChange(password string) (r bool, msg string) {
	if md5Encrypt(password) != MAIN_KEY {
		return false, "原密码错误"
	}

	MAIN_KEY = md5Encrypt(password)

	return true, "更新成功"
}
