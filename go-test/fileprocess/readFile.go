package fileprocess

import (
	"bufio"
	//"go/scanner"
	"os"
	"strings"
)

func ReadFirstLine() string {
	open, err := os.Open("log")
	if err != nil {
		return ""
	}
	defer open.Close()
	scanner := bufio.NewScanner(open)
	for scanner.Scan(){
		// 读完第一行之际返回
		return scanner.Text()
	}
	return ""
}

func ProcessFirstLine() string {
	// 此处测试依赖log文件
	line := ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}