package ipc

import (
	"testing"
	"fmt"
)

type EchoServer struct {

}

func (server *EchoServer) Handle(request string) string  {
	return "ECHO:" + request
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
        client1 := NewIpcClient(server)
        client2 := NewIpcClient(server)
        resp1, err1 := client1.Call("From Client1", "good")
        resp2, err2 := client2.Call("From Client2", "lucy")
	if err1 != nil{
		fmt.Println("error err1:", err1)
	}
	if err2 != nil{
		fmt.Println("error err2:", err2)
	}
	fmt.Println("resp1:", resp1)
	fmt.Println("resp2:", resp2)
	client1.Close()
	client2.Close()
}


