package fileprocess

import (
	"testing"
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestProcessFirstLineWithMock(t *testing.T) {
	// 对ReadProcessLine进行打桩, 将函数ReadFirstLine替换
	// 实际上在这个测试中，ReadFirstLine不会执行
	monkey.Patch(ReadFirstLine, func()string {
		return "line110"
	} )

	// 测试结束后恢复原实现
	defer monkey.Unpatch(ReadFirstLine())
	line := ProcessFirstLine()
	assert.Equal(t, "line000", line)
}