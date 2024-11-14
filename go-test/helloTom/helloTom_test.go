package hellotom

// go test -v (当前目录下)
// go test -v (绝对路径，相对路径)
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* testing.T 是 Go 语言标准库 testing 包中的一个类型
 * 用于在测试函数中报告测试结果和辅助测试执行。*/
func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	// if output != expectOutput {
	// 报告一个错误，但不会立即停止测试。测试将继续执行，直到所有测试函数都完成。
	// t.Errorf("Expected %s do not match actual %s", expectOutput, output)
	// }
	// 借助第三方包来判断
	assert.Equal(t, expectOutput, output)
}