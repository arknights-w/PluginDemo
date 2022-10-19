package test

import (
	"context"
	"fmt"
	"log"
	"plugin-demo-go/export"
	"plugin-demo-go/plugins/sms"
	"plugin-demo-go/plugins/sms/proto"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func TestX1(t *testing.T) {
	security := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(address, security)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	s := proto.NewSMSClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	res, err := s.Send(ctx, &proto.Msg{Phone: "5", Text: "wuhu~2"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %v\n", res)
}

func TestPlugin(t *testing.T) {

	cli := export.NewClient(address)
	sms := cli.GetSrv("sms").(sms.SM)

	res := sms.Send("1233213", "芜湖~起飞！")
	fmt.Printf("res: %v\n", res)

	cli.Kill()
}

func TestInstall(t *testing.T) {
	cli := export.NewClient(address)
	// cli.Install("","")
	cli.Install("D:\\BaiduNetdiskDownload\\Life\\sms_impl.zip")

	cli.Kill()

}

func TestUpdate(t *testing.T) {
	cli := export.NewClient(address)
	// cli.Install("","")
	cli.Update("conf")
	cli.Kill()
}

func TestUnInstall(t *testing.T) {
	cli := export.NewClient(address)
	cli.UnInstall("sms")
	cli.Kill()
}
func TestBanned(t *testing.T) {
	cli := export.NewClient(address)
	cli.Banned("sms")
	cli.Kill()
}
