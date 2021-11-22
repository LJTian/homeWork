package webMng

import (
	"encoding/json"
	"fmt"
	. "github.com/LJTian/Tools/log"
	"github.com/LJTian/Tools/tools"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"syscall"
)

// shell 日志记录
func ShellRecord(c *gin.Context) {
	TlogPrintf(LOG_DEBUG,"【SHELL】服务收到 请求IP is [%s] 的访问 URL为 [%s]", c.ClientIP(), c.Request.RequestURI )
	//c.String(http.StatusOK, "活着呢~" );
}

func InitShellLog() {
	LogPath := tools.GetCurrentDirectory()
	LogFile := LogPath + "/../log/WebMng.log"
	InitLog(LogFile, LOG_DEBUG)
	TlogPrintf(LOG_INFO, "日志注册成功,日志文件目录为:[%s]\n", LogFile)
}

// 获取版本信息
func GetVersion(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("SERSION: %s", os.Getenv("VERSION")))
}

// 获取报文头信息
func GetRequest(c *gin.Context) {
	HeaderStr,_ := json.Marshal(c.Request.Header)
	c.String(http.StatusOK, fmt.Sprintf("Header: %s", string(HeaderStr)))
}

// 获取客户端信息
func GetClientInfo(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("【SHELL】服务收到 请求IP is [%s] 的访问 URL为 [%s]", c.ClientIP(), c.Request.RequestURI) )
}

// 获取心跳
func GetHealthz(c *gin.Context) {
	c.String(http.StatusOK, "活着呢~" );
}

//handleMsg 读取器信息
func handleMsg(length int, err error, msg []byte) []byte {
	var strBuff string
	if length > 0 {
		strBuff = fmt.Sprintf("%s", string(msg[2:length]))
	}
	return []byte(strBuff)
}

func sendMsg(c *gin.Context) {

	msg, ok := c.GetPostForm("msg")
	if !ok {
		TlogPrintln(LOG_ERROR, "err : GetPostForm Err")
		c.String(200, fmt.Sprintf("err : GetPostForm Err"))
	}

	var resBuf []byte
	ip, _ := c.GetPostForm("ip")
	port, _ := c.GetPostForm("port")
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	// 添加报文长度
	_, err = conn.Write(tools.StatisticalLen(msg,2))
	if err != nil {
		fmt.Println("conn.Write err=", err)
	}

	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			resBuf = handleMsg(length, err, ibuf)
			TlogPrintln(LOG_DEBUG, "收到的报文为:\n", string(resBuf))
			c.String(http.StatusOK, string(resBuf))
			conn.Close()
			break
		case syscall.EAGAIN: // try again
			continue
		default:
			TlogPrintln(LOG_ERROR, "未收到报文")
			return
		}
	}
}