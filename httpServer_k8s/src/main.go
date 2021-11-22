package main

import (
	"fmt"
	"github.com/LJTian/Tools/tools"
	"httpServer_k8s/pkg/conFig"
	"httpServer_k8s/pkg/webMng"
)

func main() {
	// 启动程序
	fmt.Println("启动")
	Cfg, _ := conFig.GetConfig(tools.GetCurrentDirectory() + "/../etc/config.yml")
	webMng.InitWebMng(Cfg.Spec.LogInfo)
	webMng.StartWebMng(Cfg.Spec.NetInfo.Port)
}