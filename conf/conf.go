package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Host         string
	Port         int
	HeartbeatNum int
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func GetConfig() map[interface{}]interface{} {
	JsonParse := NewJsonStruct()
	v := Config{}
	JsonParse.GetFile(&v)
	// log.Print(v)
	m := make(map[interface{}]interface{})
	m["Host"] = v.Host
	m["Port"] = v.Port
	m["HeartbeatNum"] = v.HeartbeatNum
	return m
}

func (jst *JsonStruct) GetFile(v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	dir := "conf"
	for i := 0; i < 3; i++ {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			break
		}
		dir = filepath.Join("..", dir)
	}

	data, err := ioutil.ReadFile(filepath.Join(dir, "conf.json"))
	if err != nil {
		panic(err)
	}

	//读取的数据为json格式，需要进行解码
	jsonData := json.Unmarshal(data, v)

	if jsonData != nil {
		return
	}
}

//断言处理
func GetElemetValue(key string, mapValue map[interface{}]interface{}) string {
	value, ok := mapValue[key]
	if ok {
		return fmt.Sprint(value)
	}

	return ""
}
