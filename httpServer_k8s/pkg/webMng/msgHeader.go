package webMng

import (
	. "github.com/LJTian/Tools/log"
	"github.com/LJTian/Tools/tools"
	"github.com/gin-gonic/gin"
)

// Msg 日志记录
func MsgRecord() gin.HandlerFunc {
	return func(c *gin.Context){
		TlogPrintf(LOG_DEBUG,"【MSG】服务收到 请求IP is [%s] 的访问 URL为 [%s]", c.ClientIP(), c.Request.RequestURI )
		c.Next()
		TlogPrintf(LOG_DEBUG,"【MSG】服务收到 请求IP is [%s] 的访问 URL为 [%s] status is [%d]", c.ClientIP(),
			c.Request.RequestURI, c.Writer.Status() )
	}
}

func InitMsgLog() {
	LogPath := tools.GetCurrentDirectory()
	LogFile := LogPath + "/../log/WebMng.log"
	InitLog(LogFile, LOG_DEBUG)
	TlogPrintf(LOG_INFO, "日志注册成功,日志文件目录为:[%s]\n", LogFile)
}
