/* 1.package进行包的声明，建议：这个包和所在文件夹同名
 * 2.main包是程序的入口包，一般main函数放在该包下
 */
package main

/* 这他妈的是以前老go的方式，现在gomod不用这货了
 * 第一步：Go会先去GOROOT的scr目录中查找，很显然它不是标准库的包，没找到。
 * 第二步：继续在GOPATH的src目录去找，准确说是GOPATH/src/YOUR DIR这个目录。如果该目录不存在，会报错找不到package。
 * 在使用GOPATH管理项目时，需要按照GO寻找package的规范来合理地保存和组织Go代码。
 */
import (
	"fmt"
	// 这里如果不用gomod，必须报改源码放在gopath的src下，不然包红温
	// "dbutilss"

	// gomod方式，这里[go-package]其实是 module [go-package]
	"go-package/src/dbutilss"
	"go-package/mytool"
)

func main() {
	fmt.Println("fuck you")
	mytool.Add(2, 3)
	dbutilss.GetConn()
}
