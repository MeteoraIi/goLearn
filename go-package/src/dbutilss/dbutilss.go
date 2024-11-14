package dbutilss

import "fmt"

// 能否被外界访问
// 首字母大写可以被其他包访问
func GetConn() {
	fmt.Println("dbutils.getConn()")
}
