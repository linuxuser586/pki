package v1

import (
	"context"

	client "github.com/linuxuser586/apis/grpc/pki/client/v1"
	"github.com/linuxuser586/pki/pkg/cert"
	"google.golang.org/grpc"
)

// Register the GRPC server
func Register(s *grpc.Server) {
	client.RegisterClientServiceServer(s, &service{})
}

type service struct {
}

func (service) NewCert(ctx context.Context, in *client.CertRequest) (*client.CertResponse, error) {
	c, err := cert.GenerateRequest(in.Subjects)
	if err != nil {
		return nil, err
	}
	return &client.CertResponse{Key: c.Key, Cert: c.Cert}, nil
}
