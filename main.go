package main

import (
    "os"

    "github.com/gin-gonic/gin"

    "github.com/hd2yao/giligili/conf"
    "github.com/hd2yao/giligili/server"
)

func main() {
    // 从配置文件读取配置
    conf.Init()

    // 装载路由
    gin.SetMode(os.Getenv("GIN_MODE"))
    r := server.NewRouter()
    r.Run(":3000")
}
