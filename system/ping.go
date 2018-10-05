package system

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
)

type ICMP struct {
	Type uint8
	Code uint8
	Checksum uint16
	Identifier uint16
	SequenceNum uint16
}

func CheckSum(data []byte) uint16  {
	var (
		sum uint32
		length int = len(data)
		index int
	)

	for length >1 {
		sum += uint32(data[index])<<8+uint32(data[index+1])
		index +=2
		length -= 2
	}
	if length >0 {
		sum += uint32(data[index])
	}
	sum +=(sum>>16)
	return uint16(^sum)
}

func Ping(ip string)  {
	var (
		icmp ICMP
		buffer bytes.Buffer
	)
	conn,err := net.Dial("ip4:icmp",ip)
	if err != nil {
		logs.Error(err)
		return
	}
	defer conn.Close()
	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0
	binary.Write(&buffer,binary.BigEndian,icmp)
	icmp.Checksum = CheckSum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer,binary.BigEndian,icmp)
	if _,err := conn.Write(buffer.Bytes());err != nil {
		logs.Error(err)
		return
	}
	fmt.Println("ping成功")


}
