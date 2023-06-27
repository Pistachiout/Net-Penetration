package helper

import (
	"log"
	"net"
	"time"
)

// CreateListen 监听，参数为监听地址listenAddr，返回 TCPListener，通过 net.ResolveTCPAddr 解析地址，通过 net.ListenTCP 监听端口
//
//	监听是指服务端监听某个端口，等待客户端的连接，一旦客户端连接上来，服务端就会创建一个新的goroutine处理客户端的请求。
//
// ResolveTCPAddr是一个解析TCP地址的函数，addr为域名或者IP地址加端口号，返回一个TCPAddr，该结构体包含了ip和port
// ListenTCP函数监听TCP地址，addr则是一个TCP地址，如果addr的端口字段为0，函数将选择一个当前可用的端口，返回值l是一个net.Listener接口，可以用来接收连接。
func CreateListen(listenAddr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	return tcpListener, err
}

// CreateConnect 连接，参数为服务端地址connectAddr，返回 TCPConn，通过 net.ResolveTCPAddr 解析地址，通过 net.DialTCP 连接服务端
// 连接是指客户端连接服务端，连接成功后，客户端就可以向服务端发送数据了，与监听不同的是，连接是客户端发起的，而监听是服务端发起的。
// DialTCP函数在网络协议tcp上连接本地地址laddr和远端地址raddr，如果laddr为nil，则自动选择本地地址，如果raddr为nil，则函数在建立连接之前不会尝试解析地址，一般用于客户端。
func CreateConnect(connectAddr string) (*net.TCPConn, error) {
	// 解析地址,返回TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", connectAddr)
	if err != nil {
		return nil, err
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	return tcpConn, err
}

// KeepAlive 保持连接,参数为连接conn，通过循环向连接中写入数据，保持连接,每隔3秒写入一次,如果写入失败，说明连接已经断开，退出循环
func KeepAlive(conn *net.TCPConn) {
	for {
		_, err := conn.Write([]byte("KeepAlive"))
		if err != nil {
			log.Printf("[KeepAlive] Error %s", err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}

// GetDataFromConnection for循环获取Connection中的数据
func GetDataFromConnection(bufSize int, conn *net.TCPConn) ([]byte, error) {
	b := make([]byte, 0)
	for {
		// 读取数据
		data := make([]byte, bufSize)
		n, err := conn.Read(data)
		if err != nil {
			return nil, err
		}
		b = append(b, data[:n]...)
		if n < bufSize {
			break
		}
	}
	return b, nil
}
