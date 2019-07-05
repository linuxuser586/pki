package v1

import (
	"context"

	ca "github.com/linuxuser586/apis/grpc/pki/ca/v1"
	"github.com/linuxuser586/pki/pkg/cert"
	"google.golang.org/grpc"
)

// Register the GRPC server
func Register(s *grpc.Server) {
	ca.RegisterCertServiceServer(s, &service{})
}

type service struct {
}

func (service) Get(ctx context.Context, in *ca.CertRequest) (*ca.CertResponse, error) {
	return &ca.CertResponse{Cert: cert.CAPublic()}, nil
}
