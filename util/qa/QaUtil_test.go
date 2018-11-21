package QaUtil

import (
	"fmt"
	"testing"
	"time"
	"github.com/satori/go.uuid"
)

func TestRecQa(t *testing.T) {

	sid,_ := uuid.NewV4()

	fmt.Println(sid.String())

	fmt.Println(100 * time.Millisecond)

	qaReqInfo := QaReqInfo{Addr: "qa.linglongtech.local:8309", Uri: "/recognizeServer/?type=musicqa", ServiceType: "vbox_4.0"}
	var paras map[string]string = make(map[string]string)
	paras["yuninfo"] = `{"serviceType": "cqa", "uuid": "4233b0331c8e3d8dde33664ab35ad81cquery",
"service": "chat", "man_intv": "", "text": "\u5531\u4e00\u9996\u6b4c", "score": 1.001000046730042, "pkscore": 0.0,
"wstext": "\u5531/VI//  \u4e00/NM//  \u9996/AA//  \u6b4c/NN//", "rc": 0, "answer": {"topicID": "2002502",
"isCustomized": true, "text": "\u60f3\u542c\u6b4c\uff0c\u968f\u4fbf\u70b9\uff0c\u53ef\u4ee5\u5bf9\u6211\u8bf4\uff1a
\u6211\u60f3\u542c\u738b\u83f2\u7684\u6b4c\u3002", "topic": "chat_\u5e38\u89c4\u6027\u4fe1\u606f_\u5531\u6b4c",
"answerType": "chat", "type": "T"}, "operation": "ANSWER", "match_info": {"aid": "2450713", "type": "sim_sent",
"value": "\u5531\u4e00\u9996\u6b4c", "condition": "u}\";vbox\"", "qid": "2248613"}}\r\n`
	paras["resource"] = "#ABNF 1.0 GB2312;mode IND;meta;#ABNF HEAD-END;$main#userSonglistName=喜欢;$smartHome#userDeviceName_general=窗帘|电视|电视机|空调|插座|风扇|电风扇|加湿器|净化器|空气净化器|灯|热水器|博联插座|多位开关;$instruction#userCollectiontAlbumName=null;$instruction#userCollectiontSoundName=null;$haierHome#userDeviceName_Ac=空调;\r\n"
	paras["cfg"] = `{"qaCfg":{"useCopyRight":"false","newSchedule":"true","newNews":"true","useCountDown":"true","useNewJingping":"true","useKnowledge":"true","musicSrc":"migu"},"tvCfg":{},"msgCfg":{},"stateCfg":{}}\r\n`
	paras["province"] = "福建省\r\n"
	paras["appid"] = "5a0930ed\r\n"

	resp := QaRec(sid.String(), "热歌榜", &qaReqInfo)

	fmt.Println(resp)

	//conn, err := net.Dial("tcp", "qa.linglongtech.local:8309")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//var input bytes.Buffer
	//
	//input.WriteString("POST /recognizeServer/?type=musicqa HTTP/1.1\r\n")
	//input.WriteString("Accept: text/plain\r\n")
	//input.WriteString("ServiceType:vbox_4.0\r\n")
	//input.WriteString(`yuninfo:{"serviceType": "cqa", "uuid": "4233b0331c8e3d8dde33664ab35ad81cquery", "service": "chat", "man_intv": "", "text": "\u5531\u4e00\u9996\u6b4c", "score": 1.001000046730042, "pkscore": 0.0, "wstext": "\u5531/VI//  \u4e00/NM//  \u9996/AA//  \u6b4c/NN//", "rc": 0, "answer": {"topicID": "2002502", "isCustomized": true, "text": "\u60f3\u542c\u6b4c\uff0c\u968f\u4fbf\u70b9\uff0c\u53ef\u4ee5\u5bf9\u6211\u8bf4\uff1a\u6211\u60f3\u542c\u738b\u83f2\u7684\u6b4c\u3002", "topic": "chat_\u5e38\u89c4\u6027\u4fe1\u606f_\u5531\u6b4c", "answerType": "chat", "type": "T"}, "operation": "ANSWER", "match_info": {"aid": "2450713", "type": "sim_sent", "value": "\u5531\u4e00\u9996\u6b4c", "condition": "u}\";vbox\"", "qid": "2248613"}}`)
	//input.WriteString("resource:#ABNF 1.0 GB2312;mode IND;meta;#ABNF HEAD-END;$main#userSonglistName=喜欢;$smartHome#userDeviceName_general=窗帘|电视|电视机|空调|插座|风扇|电风扇|加湿器|净化器|空气净化器|灯|热水器|博联插座|多位开关;$instruction#userCollectiontAlbumName=null;$instruction#userCollectiontSoundName=null;$haierHome#userDeviceName_Ac=空调;\r\n")
	//input.WriteString(`cfg:{"qaCfg":{"useCopyRight":"false","newSchedule":"true","newNews":"true","useCountDown":"true","useNewJingping":"true","useKnowledge":"true","musicSrc":"migu"},"tvCfg":{},"msgCfg":{},"stateCfg":{}}\r\n`)
	//input.WriteString("province:福建省\r\n")
	//input.WriteString("appid:5a0930ed\r\n")
	//input.WriteString("Content-Length: 0\r\n")
	//input.WriteString("usercontent:热歌榜\r\n")
	//input.WriteString("\r\n")
	//
	//encoder := mahonia.NewEncoder("GBK")
	//
	//conn.Write([]byte(encoder.ConvertString(input.String())))
	//
	//var respBytes []byte
	//for {
	//	tempBytes := make([]byte, 32)
	//	size, error := conn.Read(tempBytes[:])
	//	fmt.Println(size)
	//	if error != nil {
	//		break
	//	}
	//	respBytes = append(respBytes, tempBytes...)
	//}
	//
	//decoder := mahonia.NewDecoder("GBK")
	//resp := decoder.ConvertString(string(respBytes))
	//
	//fmt.Println(resp)
}
