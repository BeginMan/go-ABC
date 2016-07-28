package httpcom

import (
	"fmt"
	"net"
	"os"

	"netcom"
	"utilities"
)

/*
通过net包向网络主机发送HTTP Head请求
读取网络主机返回的信息

	Usage:
	command qiniu.com:80

	HTTP/1.1 301 Moved Permanently
	Server: nginx
	Date: Thu, 21 Jul 2016 03:30:49 GMT
	Content-Type: text/html
	Content-Length: 178
	Connection: close
	Location: http://www.qiniu.com/
*/

func Head() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	//Dial()函数是对DialTCP(),DialUDP(), DialIP(),DialUnix()的封装
	//我们也可以直接调用这些函数，它们的功能是一致的
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", service) // 用于解析地址和端口号
	//conn, err := net.DialTCP("tcp", nil, tcpAddr)		// 建立连接
	utilities.CheckError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	utilities.CheckError(err)

	result, err := netcom.ReadFully(conn)
	utilities.CheckError(err)

	fmt.Println(string(result))
	os.Exit(0)
}
