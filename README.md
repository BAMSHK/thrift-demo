#### 1. 项目描述
是一个用go写的thrift的小demo，使用到了multiplyexecdProcess注册多个服务，多个服务使用一个main函数启动，在thrift client端连接thift服务时需要指定serviceName，值为thrift服务端在main函数中定义的
#### 2. 目录说明
--cmd ：可执行文件
  --server ：main中启动thriftService ，支持同时添加多个handle
  --client ：thrift 客户端开启入口，这里值测试了timeService
-- rpc：该目录中存放的是thriftService代码，使用MutilpyexecdProcess支持注册服务
--thrift：存放thrift 定义文件
--internal/service：存放thrift的service实现
--gen-go:存放thrift生成的go代码，也可以进入到上述thrift目录之后使用for f in `find *.thrift` ; do thrift --gen go -o ../ $f ; done 

#### 3. 使用方法
* 进入到rpc目录 修改thrift_service中的ip
* 进入到 cmd/server目录中使用 go run main.go 启动服务端
* 进入到 cmd/client 目录中使用 go run main.go 启动客户端
#### 4. 启动之后的效果是
* 服务端效果
![](https://gitee.com/BiAn-MoShangHuaKai/img/raw/master/data/20210117001055.png)
* 客户端效果
![](https://gitee.com/BiAn-MoShangHuaKai/img/raw/master/data/20210117001150.png)
