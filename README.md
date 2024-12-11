# go-zero-demo

#### 介绍
go-zero(https://go-zero.dev)，根据官网教程的demo，仅方便学习go-zero使用

#### 软件架构
框架：go-zero(https://github.com/zeromicro/go-zero)


#### 安装教程

1.  安装etcd,mysql,redis
2.  安装go1.21.3,若为其他版本（需更改shorturl/go.mod中的版本号）
3.  若缺失依赖，请go mod tidy
4.  修改配置文件（路径：1.shorturl/api/etc/shorturl-api.yaml,2.shorturl/rpc/transform/etc/transform.yaml），替换本地的地址

#### 使用说明

1.  参考官网教程,使用cmd窗口发送命令即可验证，如下
2.  第一步：curl -i "http://localhost:8888/shorten?url=https://go-zero.dev"
3.  第二步：curl -i "http://localhost:8888/expand?shorten=b0434f"

#### 说明

1.仅方便学习go-zero使用

