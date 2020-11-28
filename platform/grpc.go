package platform


import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

func (cs *connectionStringRpc) Open() *grpc.ClientConn {
	logrus.WithFields(logrus.Fields{
		"platform": "gRPC",
		"domain":   cs.domain,
	}).Info("Connection rpc")

	portLine := fmt.Sprintf("%s:%s", cs.address, cs.port)
	client, err := grpc.Dial(portLine, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"connection": cs.address,
			"password":   cs.port,
		}).Fatal(err)
		panic(err)
	}

	return client
}