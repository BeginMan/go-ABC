package ipc

import "encoding/json"

//简单的IPC(进程间通信)框架,通信包的编码细节,让使用者可以专于业务
//这里用channel作为模块之间的通信方式, 传递JSON格式的字符串类型数据

type IpcClient struct  {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error)  {
	req := &Request{method, params}
	var b []byte
	b, err = json.Marshal(req)  // load json
	if err != nil {
		return
	}

	client.conn <- string(b)	// json data send to channel
	str := <- client.conn		// wait response

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1
	return
}

func (client *IpcClient) Close()  {
	client.conn <- "CLOSE"
}