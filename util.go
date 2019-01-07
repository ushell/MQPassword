package MQPassword

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const DATE_FORMAT = "2006-01-02 15:04:05"

func jsonEncode(data interface{}) []byte {
	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}

func jsonDecode(data []byte, jsonData *Model) {
	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		log.Fatal(err)
	}
}

func md5Encrypt(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func dataVersion() string {
	date := time.Now()
	return date.Format("20060102150405")
}

func dataGenID() string {
	data := time.Now().Unix()
	return md5Encrypt(string(data))
}

//存数数据
func writeToFile(data []byte, filename string) bool {
	if filename == "" {
		log.Fatal("密码文件不存在")
	}

	fd, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer fd.Close()

	_, wErr := fd.Write(data)
	if wErr != nil {
		log.Fatal(wErr)
	}

	return true
}

//读取数据
func readFromFile(filename string) []byte {
	var data []byte

	if db_path == "" {
		log.Fatal("数据未正常初始化!")
	}

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func store(data interface{}, filename string){
	buffer := new(bytes.Buffer)

	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(data)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)

	if err !=nil {
		panic(err)
	}
}