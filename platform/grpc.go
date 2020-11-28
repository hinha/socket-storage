package platform


import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

type ConnectionsGrpc interface {
	Open() *grpc.ClientConn
}

type connectionStringRpc struct {
	address string
	port    string
	domain  string
}

func InitializeGrpc(address, port, domain string) ConnectionsGrpc {
	return &connectionStringRpc{
		address: address,
		port:    port,
		domain:  domain,
	}
}

const MaxSizeFile = 1024 * 1024 * 25 // 25mb in bytes

func (cs *connectionStringRpc) Open() *grpc.ClientConn {
	var alive = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping back
		PermitWithoutStream: true,             // send pings even without active streams
	}

	logrus.WithFields(logrus.Fields{
		"platform": "gRPC",
		"domain":   cs.domain,
	}).Info("Connection rpc")

	portLine := fmt.Sprintf("%s:%s", cs.address, cs.port)
	logrus.Info(portLine)

	client, err := grpc.Dial(portLine,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxSizeFile)),
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(alive),
		grpc.WithBlock())

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"connection": cs.address,
			"password":   cs.port,
		}).Fatal(err)
		panic(err)
	}

	return client
}