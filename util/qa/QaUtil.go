package QaUtil

import (
	"net"
	"fmt"
	"bytes"
	"github.com/axgle/mahonia"
	"time"
)

const (
	QaEncode        string = "GBK"
	QaDecode        string = "GBK"
	QaPerReciveSize int    = 32
	QaTimeOut       int    = 200
)

type QaReqInfo struct {
	Addr        string
	Uri         string
	ServiceType string
	Paras       map[string]string
}

func QaRec(sid string, userContent string, reqInfo *QaReqInfo) string {
	conn, err := net.Dial("tcp", reqInfo.Addr)
	net.DialTimeout("tcp", reqInfo.Addr, QaTimeOut)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer conn.Close()

	var reqBuf bytes.Buffer

	reqBuf.WriteString(fmt.Sprintf("POST %s HTTP/1.1\r\n", reqInfo.Uri))
	reqBuf.WriteString("Accept: text/plain\r\n")
	reqBuf.WriteString(fmt.Sprintf("ServiceType:%s\r\n", reqInfo.ServiceType))
	if reqInfo.Paras != nil {
		for key, value := range reqInfo.Paras {
			reqBuf.WriteString(fmt.Sprintf("%s:%s\r\n", key, value))
		}
	}
	reqBuf.WriteString("Content-Length: 0\r\n")
	reqBuf.WriteString(fmt.Sprintf("usercontent:%s\r\n", userContent))
	reqBuf.WriteString("\r\n")

	encoder := mahonia.NewEncoder(QaEncode)

	reqStr := reqBuf.String()

	conn.Write([]byte(encoder.ConvertString(reqStr)))

	var respBytes []byte
	for {
		tempBytes := make([]byte, QaPerReciveSize)
		_, error := conn.Read(tempBytes[:])
		if error != nil {
			break
		}
		respBytes = append(respBytes, tempBytes...)
	}

	decoder := mahonia.NewDecoder(QaDecode)
	respStr := decoder.ConvertString(string(respBytes))

	return respStr

}
