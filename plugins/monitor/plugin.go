package monitor

import (
	"context"
	"plugin-demo-go/plugins/monitor/proto"

	"google.golang.org/grpc"
)

type MonitorPlugin struct{
	cli proto.MonitorClient
}

func (m *MonitorPlugin) GetConn(conn *grpc.ClientConn){
	m.cli = proto.NewMonitorClient(conn)
	
}

func (m *MonitorPlugin) Command(cmd string,path string) CallBack{
	callback, err := m.cli.Command(context.Background(), &proto.Order{Cmd: cmd, Path: path})
	if err != nil {
		println(err)
		panic("remote callback failed")
	}
	return CallBack{Reslut: callback.Result,Msg: callback.Msg}
}
func (m *MonitorPlugin) Commands(cmd string,paths []string) CallBack{
	callback, err := m.cli.Commands(context.Background(), &proto.Orders{Cmd: cmd, Paths: paths})
	if err != nil {
		println(err)
		panic("remote callback failed")
	}
	return CallBack{Reslut: callback.Result,Msg: callback.Msg}
}
