//IPC Server
//简单的IPC(进程间通信)框架,通信包的编码细节,让使用者可以专于业务
//这里用channel作为模块之间的通信方式, 传递JSON格式的字符串类型数据

package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct  {
	Method string
	Params string
}

type Response struct  {
	Code string
	Body string
}

//todo: 接口赋值
type Server interface  {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct  {
	Server
}

func NewIpcServer(server Server) *IpcServer  {
	return &IpcServer{server}
}

func (sever *IpcServer) Connect() chan string  {
	session := make(chan string, 0)

	go func(c chan string) {
		for{
			request := <- c
			if request == "CLOSE" {		// 关闭连接
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil{
				fmt.Println("Invaild request format:", request)
			}
			resp := sever.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)

			c <- string(b)	// 返回结果
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("A new session has been created successfully.")
	return session
}

