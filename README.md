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

对外暴露的接口在 export 中

开发插件都在 plugins 中

详细使用方式可以看 test 文件夹