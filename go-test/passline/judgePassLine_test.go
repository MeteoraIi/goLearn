// 计算覆盖率go test judgePassLine_test.go judgePassLine.go --cover
package passline

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 这里可以多个测试案例
// eg: 一个false，一个true
func TestJudgePassLingTrue(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}

func TestJudgePassLingFalse(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, false, isPass)
}