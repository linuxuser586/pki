package rpc

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	ca "github.com/linuxuser586/pki/pkg/apis/ca/v1"
	client "github.com/linuxuser586/pki/pkg/apis/client/v1"
	"google.golang.org/grpc"
)

// Start the server
func Start() string {
	p := ":" + os.Getenv("PKI_GRPC_PORT")
	if p == ":" {
		p = ":10042"
	}
	log.Infof("starting PKI GRPC on %s", p)
	lis, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	s := grpc.NewServer()
	ca.Register(s)
	client.Register(s)
	go func() {
		for {
			<-c
			log.Info("stopping PKI GRPC server")
			s.Stop()
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return "stopped PKI GRPC server"
}
