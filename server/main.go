package main

import (
	"Net-Penetration/constant"
	"Net-Penetration/helper"
	"io"
	"log"
	"net"
	"sync"
)

// serverConn 服务端连接
var serverConn *net.TCPConn

// appConn 目的服务端连接
var appConn *net.TCPConn

// wg 用于等待所有协程结束
var wg sync.WaitGroup

// 内网穿透服务端
func main() {
	//监听服务端
	go serverListen()
	//监听目的服务端
	go appListen()
	//启动隧道服务
	go tunnelListen()

	//等待所有协程结束
	wg.Add(1)
	wg.Wait()
}

func serverListen() {
	//监听服务端，用于接收客户端连接
	tcpListener, err := helper.CreateListen(constant.ServerAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("服务端监听地址为：%s\n", tcpListener.Addr().String())

	//接收客户端连接
	for {
		serverConn, err = tcpListener.AcceptTCP()
		if err != nil {
			log.Printf("接收连接失败，错误信息为：%s\n", err.Error())
			return
		}
		//保持连接
		go helper.KeepAlive(serverConn)
	}
}

// 监听隧道服务，用于接收隧道客户端连接，隧道客户端连接用于转发目的服务端和客户端之间的消息，实现内网穿透
func tunnelListen() {
	tcpListener, err := helper.CreateListen(constant.TunnelAddr)
	if err != nil {
		panic(err)
	}
	log.Printf("隧道监听地址为：%s\n", tcpListener.Addr().String())
	for {
		tunnelConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Printf("接收连接失败，错误信息为：%s\n", err.Error())
			return
		}
		// 数据转发
		go io.Copy(appConn, tunnelConn)
		go io.Copy(tunnelConn, appConn)
	}
}

// 监听目的服务端，用于接收目的服务端连接，目的服务端可以是本地转发的端口，也可以是远程服务器的端口
func appListen() {
	//监听目的服务端
	tcpListener, err := helper.CreateListen(constant.AppTargetPort)
	if err != nil {
		panic(err)
	}
	log.Printf("应用目的服务端监端口地址为：%s\n", tcpListener.Addr().String())

	for {
		appConn, err = tcpListener.AcceptTCP()
		if err != nil {
			log.Printf("接收连接失败，错误信息为：%s\n", err.Error())
			return
		}
		_, err := serverConn.Write([]byte("New Connection"))
		if err != nil {
			log.Printf("发送消息失败，错误信息为：%s\n", err.Error())
		}
	}
}
