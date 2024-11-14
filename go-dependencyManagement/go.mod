// go mod init gomod，初始化一个gomod文件

// go mod download github.com/gin-gonic/gin@1.9.0 下载第三方库
// pkg\mod\cache 缓存内容，写入硬盘，网络到磁盘
// 所下载的目录
// D:\software\IDE\goWorkingDictionary\pkg\mod\github.com\gin-gonic\gin@v1.10.0
// --download指令仅仅只下载指定的版本，不会安装它所需求的依赖

// go mod tidy 依赖对齐
// 确保go.mod文件中的依赖关系与项目实际使用的情况想匹配
// 添加使用的依赖，删除没使用的依赖
// 与download不同，会把所有依赖弄进来
// 这里啥都没用，会把我下面的一些关键字都干掉
// eg: main.go

// go help mod edit 对go.mod文件编辑
// go mod edit -require = ""
// ...

// go mod vendor, 将go mod文件中的依赖备份到go vendor里面
// 依赖读取是优先go vendor然后 gomod

// go mod verify 验证所依赖的mod有没有发生改变

// go mod why github.com/go-playground/validator/v10
// 查看项目依赖关系的最短路径

// go install 用来安装go包（包括可执行程序和库）
// --安装一些工具，可执行的插件
// 可执行文件安装在gopath/bin下，普通包安装在/gopath/pkg/mod下


// 模块名
module gomod

// golang sdk 版本
go 1.23.2

// 依赖配置、依赖关系
// eg：当两个模块依赖同一个库的不同版本时，选择最低的兼容版本

require github.com/gin-gonic/gin v1.10.0

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
