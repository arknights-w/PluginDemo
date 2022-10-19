package plugins

import (
	"google.golang.org/grpc"
)

type Plugin interface {
	GetConn(conn *grpc.ClientConn)
}
