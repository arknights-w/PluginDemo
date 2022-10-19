package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	std_sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"plugin-demo-go/db/gen"
	"time"
)

var client *gen.Client

// Connection
// 这是整合了 sql.DB 的写法
// 可以去官网搜索不进行整合的写法，以及其他整合手法
// 整合原因：可以通过对 sql.DB 进行参数配置，来达到自定义ORM
func Connection() *gen.Client {
	if client == nil {
		// ent sql driver
		dsn := "root:123456@tcp(172.26.48.1:3306)/gosql?charset=utf8mb4&parseTime=True"
		drv, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}

		// turn to std sql and set config
		db := drv.DB()
		Config(db)

		// new client
		client = gen.NewClient(gen.Driver(drv))

		// 运行自动迁移工具来创建所有Schema资源
		// 也就是说如果 当前数据库 没有 对应的表，执行以下两行代码 可创建
		ctx := context.Background()
		// client.Schema
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return client
}
func Config(db *std_sql.DB) {
	// 最大空闲连接数
	db.SetMaxIdleConns(10)
	// 最大连接数
	db.SetMaxOpenConns(100)
	// 最长连接时长
	db.SetConnMaxLifetime(time.Hour)
}
