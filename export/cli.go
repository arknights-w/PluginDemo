package export

import (
	"log"
	"plugin-demo-go/plugins"
	"plugin-demo-go/plugins/monitor"
	"plugin-demo-go/plugins/sms"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 这里配置需要插入的插件
var plugins_ = map[string]plugins.Plugin{
	"sms": &sms.SmsPlugin{},
	"mon": &monitor.MonitorPlugin{},
}

type Client struct {
	plugins map[string]plugins.Plugin
	conn    *grpc.ClientConn
	done	chan bool
}

// create function
// input address of grpc server
// return Client connected to server
func NewClient(address string) *Client {
	security := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(address, security)
	if err != nil {
		log.Fatal(err)
	}

	cli := new(Client)
	cli.done = make(chan bool, 1)
	cli.conn = conn
	cli.plugins = plugins_

	go cli.AutoManage()

	return cli
}

// Get service
// input service name
// return the implement of service interface
func (c *Client) GetSrv(srv string) interface{} {
	c.plugins[srv].GetConn(c.conn)
	return c.plugins[srv]
}

// kill the Client
// break the connection between the cli and server
func (c *Client) Kill() {
	// 关闭远程连接
	c.conn.Close()
	// 关闭巡回访问服务器的线程
	c.done<-true
}
