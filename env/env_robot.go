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
                fmt.Println(err)
                fmt.Println("发生了一些未知错误...")
            }
           }()

    argsQuantity := len(os.Args)
    if argsQuantity < 4 {
        fmt.Println("请输入以下命令参数 env[dev sit uat deploy] new_env_path old_env_path")
        return
    }

    env := os.Args[1]
    if !strings.Contains("dev sit uat deploy", env) {
        fmt.Println("请选择正确的环境值 -> [dev sit uat deploy]")
        return
    }
    newEnvPath := os.Args[2]
    if newEnvPath == "" {
        newEnvPath = "./"
    } else if !isFileExist(newEnvPath) {
        fmt.Println("文件夹 " + newEnvPath + " 不存在")
        return
    }
    oldEnvPath := os.Args[3]
    if oldEnvPath == "" {
        oldEnvPath = "./"
    } else if !isFileExist(oldEnvPath) {
        fmt.Println("文件夹 " + oldEnvPath + " 不存在")
        return
    }

    // 读取配置文件
    var config []byte
    var err error
    switch strings.TrimSpace(env) {
        case "dev":
            if !isFileExist(newEnvPath + "/dev.env") {
                fmt.Println("文件 " + newEnvPath + "/dev.env 不存在")
                return
            }
            config, err = ioutil.ReadFile(newEnvPath + "/dev.env")
        case "sit":
            if !isFileExist(newEnvPath + "/sit.env") {
                fmt.Println("文件 " + newEnvPath + "/sit.env 不存在")
                return
            }
            config, err = ioutil.ReadFile(newEnvPath + "/sit.env")
        case "uat":
            if !isFileExist(newEnvPath + "/uat.env") {
                fmt.Println("文件 " + newEnvPath + "/uat.env 不存在")
                return
            }
            config, err = ioutil.ReadFile(newEnvPath + "/uat.env")
        case "deploy":
            if !isFileExist(newEnvPath + "/deploy.env") {
                fmt.Println("文件 " + newEnvPath + "/deploy.env 不存在")
                return
            }
            config, err = ioutil.ReadFile(newEnvPath + "/deploy.env")
    }
    if err != nil {
        panic(err)
    }

    // 覆盖已有配置
    err = ioutil.WriteFile(oldEnvPath + "/.env", config, 0644)
    if err != nil {
        panic(err)
    }

    fmt.Println("切换环境完成...")
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
