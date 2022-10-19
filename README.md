# 插件客户端

## build

``` sh
# 更新plugin文件
make protoc
# 更新数据库字段
make ent
```

## 需要修改的

```
1. db/connection.go 中的连接IP

2. export/manager.go 中的 target_path
```

## 使用方法

插件相关对外暴露的接口在 export 中

开发插件都在 plugins 中

详细使用方式可以看 test 文件夹

http服务器对外交流，收到插件相关需求后，使用 export 暴露的插件客户端接口，以 grpc 的形式请求插件服务器获得服务

这里面可以抽象出三类人:

1.与服务器交互的用户

2.编写http服务器的后端人员，他只允许调用 export暴露接口，不允许修改插件客户端的代码

3.编写插件客户端与插件服务器的人员，他负责编写插件客户端与插件服务器，提供服务

```
可以将目录结构进行划分：

http: main.go, httpServe

plugin: plugins, db, export, tools

```