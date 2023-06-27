package main

import (
	"Net-Penetration/constant"
	"Net-Penetration/helper"
	"io"
	"log"
)

// 内网穿透客户端，用于将本地服务映射到公网，使得公网可以访问本地服务，实现内网穿透，
func main() {
	// 连接服务端
	conn, err := helper.CreateConnect(constant.ServerAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("连接成功，连接地址为：%s\n", conn.RemoteAddr().String())

	// 保持连接，读取数据
	for {
		// 从连接中读取数据
		data, err := helper.GetDataFromConnection(constant.BufSize, conn)
		if err != nil {
			log.Printf("读取数据失败，错误信息为：%s\n", err.Error())
			continue
		}
		log.Printf("接收到数据：%s\n", string(data))
		// 判断是否为新连接，如果是新连接，则连接隧道服务器，否则转发消息
		if string(data) == "New Connection" {
			// 连接隧道服务器
			go messgaeForward()
		}
	}
}

// 连接隧道服务器进行消息转发
func messgaeForward() {
	// 连接隧道服务器
	tunnelConn, err := helper.CreateConnect(constant.TunnelAddr)
	if err != nil {
		panic(err)
	}

	// 连接客户端服务
	clientConn, err := helper.CreateConnect(constant.AppPort)
	if err != nil {
		panic(err)
	}

	//	消息转发
	//	io.Copy()函数实现了数据的拷贝，可以将数据从一个接口拷贝到另一个接口，这里将客户端的数据拷贝到隧道服务器，将隧道服务器的数据拷贝到客户端
	//	io.Copy()函数会一直阻塞，直到两个接口中的数据全部拷贝完成
	go io.Copy(clientConn, tunnelConn)
	go io.Copy(tunnelConn, clientConn)
}
