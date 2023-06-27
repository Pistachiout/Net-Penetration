代码中已包含大量注释，穿透的具体逻辑可前往进行查看：https://blog.csdn.net/qq_45808700/article/details/131417192

## 1 若没有公网ip，可进行本地测试，步骤如下：
1. 首先运行服务端main.go
2. 运行客户端main.go
3. 运行应用端main.go
4. 此时可前往浏览器进行测试，分别打开AppPort和AppTargetPort进行查看
5. 在第一次打开AppTargetPort时会建立连接，故需要刷新后才会进行消息转发，看到内网的服务

![image](https://github.com/Pistachiout/Net-Penetration/assets/63298680/7ec916ca-8240-4b2b-9b8c-e35f57aec432)


## 2 公网ip测试步骤：以Linux云服务器为例
1. 首先修改constant.go中的ServerIP为自己的公网ip
2. 将服务端main.go打包，上传到Linux云服务器，并添加权限。windows下将go程序打包为linux可执行程序需要配置go编译环境，具体参考[windows下将go程序打包为linux可执行程序教程](https://blog.csdn.net/qq_45808700/article/details/131419641)：
3. 在服务器运行上传的main
4. 在客户端运行客户端main.go
5. 运行应用端main.go
6. 此时可打开公网ip:AppTargetPort，并进行刷新，即可发现内网的应用服务
