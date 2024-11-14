package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type DictResponseBD struct {
	Errno int `json:"errno"`
	Data  []struct {
		K string `json:"k"`
		V string `json:"v"`
	} `json:"data"`

	Logid int `json:"logid"`
}

type DictRequestHS struct {
	Source         string   `json:"source"`
	Words          []string `json:"words"`
	SourceLanguage string   `json:"source_language"`
	TargetLanguage string   `json:"target_language"`
}

type DictResponseHS struct {
	Details []struct {
		Detail string `json:"detail"`
		Extra  string `json:"extra"`
	} `json:"details"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

type DictResponseHSData struct {
	Result []struct {
		Ec struct {
			Basic struct {
				Explains []struct {
					Pos   string `json:"pos"`
					Trans string `json:"trans"`
				} `json:"explains"`
			} `json:"basic"`
		} `json:"ec"`
	} `json:"result"`
}

func Query() string {
	reader := bufio.NewReader(os.Stdin)
	word, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	word = strings.TrimSuffix(word, "\r\n")

	return word
}

func QueryBD(word string) {
	client := &http.Client{}

	var data = strings.NewReader("kw=" + word)

	// 创建一个Http请求
	req, err := http.NewRequest("POST", "https://fanyi.baidu.com/sug", data)
	if err != nil {
		log.Fatal(err)
	}

	// 设置http头部
	// "Content - Type"，服务器根据这个字段（表单）来解析请求体中的数据。
	// "application/x - www - form - urlencoded"，这是Content - Type头部的值。
	// 是一种常见的内容类型，表示请求体中的数据是经过 URL 编码的表单数据形式。
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 执行http请求，用resp接受返回
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// 延迟调用resp.Body.Close()
	/* resp的Body字段是一个io.ReadCloser类型，它实现了io.Reader接口用于读取响应体的内容，
	   同时也实现了io.Closer接口用于关闭响应体的资源。*/
	/* defer resp.Body.Close()的作用就是确保在函数执行结束之前，
	   resp.Body会被正确关闭，从而避免资源泄漏 */
	defer resp.Body.Close()

	/* 这行代码的主要目的是从http.Response对象（resp）的响应体（resp.Body）中读取全部内容，
	   并将其存储为字节切片（[]byte）类型的变量bodyText，同时返回可能出现的错误信息。*/
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 正确返回码应该是200
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode: ", resp.StatusCode, "body", string(bodyText))
	}

	var dictResponse DictResponseBD
	// 反序列化, http请求返回的是一个JSON
	// 这行代码用于将字节切片bodyText中的 JSON 数据解析并反序列化为dictResponse变量所对应的 Go 结构体类型。
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dictResponse.Data[0].V)

}
func QueryHS(word string) {
	client := &http.Client{}

	dictRequest := DictRequestHS{Source: "youdao", Words: []string{word},
		SourceLanguage: "en",
		TargetLanguage: "zh"}
	/* 这行代码主要是将dictRequest结构体对象转换为 JSON 格式的字节切片（[]byte）。
	   这个过程在 Go 语言中称为序列化.*/
	buf, err := json.Marshal(dictRequest)
	if err != nil {
		log.Fatal(err)
	}
	/* 这行代码创建了一个bytes.Reader类型的对象data，
	   它可以用于从字节切片buf中读取数据，就好像从一个可读流（io.Reader）中读取一样。
	   转换成这种东西，作为HTTP请求的请求体*/
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/detail/v1/?msToken=&X-Bogus=DFSzKwVLQDaZV6ChtsiBTUkX95zT&_signature=_02B4Z6wo00001Zmxi-wAAIDBrpVPs8XtxJmZoY9AAAFB8Hg5w3erQ-qacwD5t8PYJspBusaj5GvwceNIWw8Frq4aQspKSQgaHkUdhXImB8kC.WZ0DrWQKSki1uzKj864ybWRE1hvOhXeKsc9bf", data)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode: ", resp.StatusCode, "body: ", string(bodyText))
	}

	var response DictResponseHS
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Fatal(err)
	}

	item := response.Details[0]
	jsonStr := item.Detail

	var HSData DictResponseHSData
	err = json.Unmarshal([]byte(jsonStr), &HSData)
	if err != nil {
		panic(err)
	}

	for _, item := range HSData.Result[0].Ec.Basic.Explains {
		fmt.Println(item.Pos, item.Trans)
	}

}

func main() {
	for {
		fmt.Println("\n1.百度翻译")
		fmt.Println("2.火山翻译")
		fmt.Println("3.退出")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSuffix(input, "\r\n")
		your_select, err := strconv.Atoi(input)

		if your_select == 3 {
			break
		}

		fmt.Println("请输入要查询的单词: ")
		word := Query()

		switch your_select {
		case 1:
			QueryBD(word)
		case 2:
			QueryHS(word)
		default:
			fmt.Println("Fuck you, you son of bitch!")
		}

	}
}
