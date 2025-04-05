package rapidapp

import (
	"context"
	"crypto/tls"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/rapidappio/rapidapp-go/pkg/proto/postgres"
)

const (
	apiUrl = "api.rapidapp.io:443"
	// apiUrl = "localhost:8080"
)

type Client struct {
	apiKey            string
	PostgresDatabases postgres.PostgresServiceClient
}

func NewClient(apiKey string) *Client {
	conn, err := grpc.Dial(apiUrl, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	// conn, err := grpc.Dial(apiUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Client{
		apiKey:            apiKey,
		PostgresDatabases: postgres.NewPostgresServiceClient(conn),
	}
}

func (c *Client) CreatePostgresDatabase(name string) (string, error) {
	id, err := c.PostgresDatabases.Create(c.newCtx(), &postgres.CreateRequest{Name: name})
	if err != nil {
		return "", err
	}
	return id.Id, err
}

func (c *Client) DeletePostgresDatabase(id string) error {
	_, err := c.PostgresDatabases.Delete(c.newCtx(), &postgres.DeleteRequest{Id: id})
	return err
}

func (c *Client) GetPostgresDatabase(id string) (*postgres.Postgres, error) {
	return c.PostgresDatabases.Get(c.newCtx(), &postgres.GetRequest{Id: id})
}

func (c *Client) ListPostgresDatabases() ([]*postgres.Postgres, error) {
	list, err := c.PostgresDatabases.List(c.newCtx(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}

func (c *Client) newCtx() context.Context {
	md := metadata.Pairs("x-api-key", c.apiKey)
	return metadata.NewOutgoingContext(context.Background(), md)
}
