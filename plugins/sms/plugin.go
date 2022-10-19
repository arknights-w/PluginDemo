package sms

import (
	"context"
	"plugin-demo-go/plugins/sms/proto"

	"google.golang.org/grpc"
)

type SmsPlugin struct {
	cli proto.SMSClient
}
func (s *SmsPlugin) GetConn(conn *grpc.ClientConn) {
	s.cli = proto.NewSMSClient(conn)
}

func (s *SmsPlugin) Send(phone string,text string) Res {
	res, err := s.cli.Send(context.Background(), &proto.Msg{Phone: phone, Text: text})
	if err != nil {
		panic("remote callback failed")
	}
	return Res{Reslut: res.Result,Msg: res.Msg}
}

