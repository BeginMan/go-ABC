package netcom

import (
	"bytes"
	"io"
	"net"
)

//ReadFully其实就是标准库 io/ioutil的`ReadAll()`的写法
//result, err := ReadFully(conn)
//result, err := ioutil.ReadAll(conn)
//两者等效
func ReadFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

//校验和算法
func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}

	sum += (sum >> 16)
	return uint16(^sum)
}
