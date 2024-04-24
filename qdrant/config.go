package qdrant

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Host string
	Port int

	ApiToken string

	UseTls    bool
	tlsConfig *tls.Config

	GrpcOptions []grpc.DialOption
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) GetTransportCreds() grpc.DialOption {
	if c.UseTls {
		tlsConfig := c.tlsConfig
		if tlsConfig == nil {
			tlsConfig = &tls.Config{}
		}
		return grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))
	}
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}
