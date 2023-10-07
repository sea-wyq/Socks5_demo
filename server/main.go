package main

import (
	"context"
	"log"
	"net"

	socks5 "github.com/armon/go-socks5"
)

func main() {
	// 创建一个新的SOCKS5服务器实例
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{
			socks5.UserPassAuthenticator{
				Credentials: socks5.StaticCredentials{
					"root": "1234",
				},
			},
		},
		Rules: &PermitDestPort{
			Ports: []int{80},
		}, // 只容许转发到80端口
	}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	// 监听在本地的1080端口
	listener, err := net.Listen("tcp", ":1080")
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务器
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

type PermitDestPort struct {
	Ports []int
}

func (p *PermitDestPort) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	for _, port := range p.Ports {
		if req.DestAddr.Port == port {
			return ctx, true
		}
	}
	return ctx, false
}
