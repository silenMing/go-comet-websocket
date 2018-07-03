package conf

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host         string
	Port         string
	HeartbeatNum int
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func GetConfig() {
	JsonParse := NewJsonStruct()
	v := Config{}
	JsonParse.GetFile("./conf.json", &v)
}

func (jst *JsonStruct) GetFile(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
