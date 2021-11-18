package webMng

import (
	. "github.com/LJTian/Tools/log"
	"github.com/LJTian/Tools/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

var maxRead int = 20480

func InitWebMng() {
	LogPath := tools.GetCurrentDirectory()
	LogFile := LogPath + "/../log/WebMng.log"
	InitLog(LogFile, LOG_DEBUG)
	TlogPrintf(LOG_INFO, "日志注册成功,日志文件目录为:[%s]\n", LogFile)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.static", "")
}

func StartWebMng(addr string) {
	// 1.创建路由
	ginr := gin.Default()
	ginr.GET("/", index)

	// 加载静态文件
	ginr.StaticFS("/public", http.Dir("/../static/"))

	shellHead := ginr.Group("/shell").Use(ShellRecord)
	{
		shellHead.POST("/GetVersion", GetVersion )
		shellHead.POST("/GetRequest", GetRequest )
		shellHead.POST("/GetClientInfo", GetClientInfo )
		shellHead.POST("/GetHealthz", GetHealthz )
	}
	MsgHead := ginr.Group("/Msg").Use(MsgRecord)
	{
		MsgHead.POST("/SendMsg", sendMsg)
	}
	ginr.LoadHTMLFiles(tools.GetCurrentDirectory() + "/../static/index.static")
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	TlogPrintf(LOG_INFO, "WEB管理页面加载成功, 监听地址为:[%s]\n", addr)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	ginr.Run(addr)

	// 接受信号，优雅退出
}