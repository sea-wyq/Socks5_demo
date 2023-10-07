package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// 设置代理服务器地址
	proxyURL, err := url.Parse("socks5://root:1234@127.0.0.1:1080")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	// 创建一个自定义的Transport，使用代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// 创建一个自定义的Client，使用自定义的Transport
	client := &http.Client{
		Transport: transport,
	}

	// 发送GET请求
	resp, err := client.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应内容
	fmt.Println(string(body))
}
