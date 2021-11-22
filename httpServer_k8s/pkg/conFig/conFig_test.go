package conFig

import (
	"fmt"
	"testing"
)

func TestInitLog(t *testing.T) {

	fmt.Println("开始测试！")
	config, err := GetConfig("E:\\github.com\\LJTian\\ubuntu\\homeWork\\httpServer_k8s\\etc\\config.yml");
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)

	fmt.Println("结束测试！")
}
