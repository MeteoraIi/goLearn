// go mod tidy 测试
// 使用之后会产生一个go.sum
// go/sum记录了依赖，版本，哈希值，以及依赖所对应的依赖
package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.New()
	r.Run()
}