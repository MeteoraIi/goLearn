package serverindex

import(
	"math/rand"
	// 记得查包在哪
	// 用fastrnd做rand的优化
	"github.com/bytedance/gopkg/lang/fastrand"
)

var ServerIndex [10]int

// 初始化服务器
func InitServerIndex() {
	for i := 0;i < 10;i ++ {
		ServerIndex[i] = i + 100
	}
}

// 选择执行服务器
func Select() int {
	return ServerIndex[rand.Intn(10)]
}

// 快速生成随机数，优化
func FastSelect() int {
	// 模拟随机情况
	return ServerIndex[fastrand.Intn(10)]
}