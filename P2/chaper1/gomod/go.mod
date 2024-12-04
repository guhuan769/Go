// 模块名
module gomod

// go SDK 版本
go 1.22.9

// 当前module(项目) 依赖的包
//dependency latest

retract v1.1.0

replace golang.org/x/crypto v0.0.0 => github.com/golang/crypto v0.29.0

exclude github.com/gin-gonic/gin v1.9.0
