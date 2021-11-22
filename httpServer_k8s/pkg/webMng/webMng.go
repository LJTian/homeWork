package webMng

import (
	"context"
	. "github.com/LJTian/Tools/log"
	"github.com/LJTian/Tools/tools"
	"github.com/gin-gonic/gin"
	"httpServer_k8s/pkg/conFig"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var maxRead int = 20480

func InitWebMng( logCfg conFig.LogStruct) {
	LogPath := tools.GetCurrentDirectory() + logCfg.Path
	LogFile := LogPath + logCfg.Name
	InitLog( LogFile, logCfg.Level)
	TlogPrintf( LOG_INFO, "日志注册成功,日志文件目录为:[%s]\n", LogFile )
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func StartWebMng( addr string )  {
	// 1.创建路由
	ginr := gin.Default()
	ginr.GET("/", index)

	// 加载静态文件
	ginr.StaticFS("/static", http.Dir(tools.GetCurrentDirectory() + "/../static"))

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
	ginr.LoadHTMLFiles(tools.GetCurrentDirectory() + "/../static/index.html")
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response

	TlogPrintf(LOG_INFO, "WEB管理页面加载成功, 监听地址为:[%s]\n", addr )

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	// ginr.Run(addr)
	// 接受信号，优雅退出
	// 优雅启动终止

	srv := &http.Server{
		Addr:         addr,
		Handler:      ginr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		// 监听请求
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			TlogPrintf(LOG_ERROR,"listen: %s\n", err)
		}
	}()

	// 优雅Shutdown（或重启）服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) // syscall.SIGINT
	<-quit
	TlogPrintln(LOG_DEBUG,"Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		TlogPrintf(LOG_ERROR,"Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
	}
	TlogPrintln(LOG_ERROR,"Shutdown Server Done ")
}