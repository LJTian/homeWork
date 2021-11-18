package webMng

import (
	"github.com/LJTian/Tools/tools"
	. "github.com/LJTian/Tools/log"
	"github.com/gin-gonic/gin"
	"net/http"
)


// Msg 日志记录
func MsgRecord(c *gin.Context) {
	c.String(http.StatusOK, "活着呢~" );
}

func InitMsgLog() {
	LogPath := tools.GetCurrentDirectory()
	LogFile := LogPath + "/../log/WebMng.log"
	InitLog(LogFile, LOG_DEBUG)
	TlogPrintf(LOG_INFO, "日志注册成功,日志文件目录为:[%s]\n", LogFile)
}
