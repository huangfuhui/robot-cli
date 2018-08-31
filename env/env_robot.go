package main

import (
    "fmt"
    "os"
    "strings"
    "io/ioutil"
)

func main() {
    defer func() {
            if err := recover(); err != nil {
                // fmt.Println(err)
                fmt.Println("发生了一些未知错误...")
            }
           }()

    argsQuantity := len(os.Args)
    if argsQuantity < 3 {
        fmt.Println("请输入以下命令参数 env[dev sit uat deploy] env_path")
        return
    }

    env := os.Args[1]
    if !strings.Contains("dev sit uat deploy", env) {
        fmt.Println("请选择正确的环境值 -> [dev sit uat deploy]")
        return
    }
    envPath := os.Args[2]
    if envPath == "" {
        envPath = "./"
    } else if !isFileExist(envPath) {
        fmt.Println("文件夹 " + envPath + " 不存在")
        return
    }

    // 读取配置文件
    var config []byte
    var err error
    switch strings.TrimSpace(env) {
        case "dev":
            if !isFileExist("dev.env") {
                fmt.Println("文件 ./dev.env 不存在")
                return
            }
            config, err = ioutil.ReadFile("dev.env")
        case "sit":
            if !isFileExist("sit.env") {
                fmt.Println("文件 ./sit.env 不存在")
                return
            }
            config, err = ioutil.ReadFile("sit.env")
        case "uat":
            if !isFileExist("uat.env") {
                fmt.Println("文件 ./uat.env 不存在")
                return
            }
            config, err = ioutil.ReadFile("uat.env")
        case "deploy":
            if !isFileExist("deploy.env") {
                fmt.Println("文件 ./deploy.env 不存在")
                return
            }
            config, err = ioutil.ReadFile("deploy.env")
    }
    if err != nil {
        panic(err)
    }

    // 覆盖已有配置
    err = ioutil.WriteFile(envPath + "/.env", config, 0644)
    if err != nil {
        panic(err)
    }
}

// 判断文件是否存在
func isFileExist(filePath string) bool {
    _, err := os.Stat(filePath)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return false
}
