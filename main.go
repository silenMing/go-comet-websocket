package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
	"strings"
)

//全局信息
var datas Datas
var users map[*websocket.Conn] string

type UserMsg struct {  
    UserName string  
    Msg      string  
    DataType string  
} 

type UserData struct {  
    UserName string  
}

type Datas struct {  
    UserMsgs  []UserMsg  
    UserDatas []UserData  
}

func main() {
	fmt.Println("启动时间")
	fmt.Println(time.Now())

	//初始化  
    datas = Datas{}  
    users = make(map[*websocket.Conn]string)

    http.HandleFunc("/", h_index)

    http.Handle("/webSocket", websocket.Handler(h_websocket))
    
    http.ListenAndServe(":3333", nil)

}

func h_index(w http.ResponseWriter,r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func h_websocket(ws *websocket.Conn) {
	var userMsg UserMsg
	var data string

	for {
		if _, ok := users[ws]; !ok {  
            users[ws] = "匿名"  
        }

        userMsgsLen := len(datas.UserMsgs)
		fmt.Println("UserMsgs", userMsgsLen, "users长度：", len(users))

		//有消息时候
		if userMsgsLen > 0 {
       	b, errMarshl := json.Marshal(datas)
       	if errMarshl != nil{
       	    	fmt.Println("全局消息内容异常...")  
            	break 
       	}

       	for key, _ := range users {
       	  	errMarshl := websocket.Message.Send(key, string(b))
       	  	if errMarshl != nil {
       	  	        	//移除出错的链接  
                	delete(users, key)  
                	fmt.Println("发送出错...")  
                	break 
       	  	}
       	}
       	datas.UserMsgs = make([]UserMsg, 0)
		}

		fmt.Println("开始解析数据...")
		err := websocket.Message.Receive(ws,&data)
		fmt.Println("data：", data)

    	if err != nil {  
        	//移除出错的链接  
        	delete(users, ws)  
        	fmt.Println("接收出错...")  
        	break  
    	}

    	data = strings.Replace(data, "\n", "", 0)
    	err = json.Unmarshal([]byte(data), &userMsg)
    
    	if err != nil {
    		fmt.Println("解析数据异常...")
    		break
    	}

    	fmt.Println("请求数据类型：", userMsg.DataType)

    	switch userMsg.DataType{
    	case "send":
    		if _,ok := users[ws]; ok{
           	users[ws] = userMsg.UserName

           	datas.UserDatas = make([]UserData,0)

            	for _, item := range users {  
  
                	userData := UserData{UserName: item}  
                	datas.UserDatas = append(datas.UserDatas, userData)  
            	}

    		}
    		datas.UserMsgs = append(datas.UserMsgs, userMsg)
    	} 
	}
}
