代码中已包含大量注释，穿透的具体逻辑可前往进行查看：https://blog.csdn.net/qq_45808700/article/details/131417192

## 1 若没有公网ip，可进行本地测试，步骤如下：
1.首先运行服务端main.go
2. 运行客户端main.go
3. 运行应用端main.go
4. 此时可分别打开AppPort和AppTargetPort均可发现应用服务

![在这里插入图片描述](https://img-blog.csdnimg.cn/3c002c103dce4cd1904ae6d467b1cf31.png)

## 2 公网ip测试步骤：以Linux云服务器为例
1. 首先修改constant.go中的ServerIP为自己的公网ip
2. 将服务端main.go打包，上传到Linux云服务器，并添加权限。windows下将go程序打包为linux可执行程序需要配置go编译环境，具体参考[windows下将go程序打包为linux可执行程序教程](https://blog.csdn.net/qq_45808700/article/details/131419641)：
3. 在服务器运行上传的main
4. 在客户端运行客户端main.go
5. 运行应用端main.go
6. 此时可打开公网ip:AppTargetPort可发现内网的应用服务
