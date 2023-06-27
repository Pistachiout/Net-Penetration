package constant

const (

	// 需要内网穿透的服务端口，运行内网穿透客户端后，该端口将映射到公网，公网可以访问该端口
	AppPort = ":4000"

	// ServerPort 服务端监听端口，用于连接客户端,一般为公网端口,服务端监听该端口，接收客户端的数据，然后转发给服务端本地服务
	ServerPort = ":8081"
	// TunnelPort 隧道端口，用于客户端和服务端建立隧道，客户端和服务端通过该端口通信
	TunnelPort = ":8082"
	// AppTargetPort 穿透后的目的服务端口，在服务端端口运行客户端应用，用户可通过访问公网ip的该端口访问内网服务
	AppTargetPort = ":8083"

	// ServerIP 服务端公网ip地址，用于连接服务端，一般为公网IP
	ServerIP = ""
	// ServerAddr 服务端监听地址，用于接收客户端的数据，然后转发给服务端本地服务
	ServerAddr = ServerIP + ServerPort
	// TunnelAddr 隧道地址，用于客户端和服务端建立隧道，客户端和服务端通过该地址通信
	TunnelAddr = ServerIP + TunnelPort

	// BufSize 缓冲区大小
	BufSize = 1024
)
