package main

import (
	"golang_task4/router"
	"golang_task4/utils"
)

func main() {

	// 读取配置文件
	utils.InitConfig()

	// 初始化数据库
	utils.InitDB()

	r := router.Router()

	err := r.Run()
	if err != nil {
		return
	}
}
