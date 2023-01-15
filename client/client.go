package client

import (
	"context"
	"fmt"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"google.golang.org/grpc/metadata"
)

// Client for auth.
type Client struct {
	client v1.ServiceClient
	config *v1c.Config
}

// NewClient for auth.
func NewClient(client v1.ServiceClient, config *v1c.Config) *Client {
	return &Client{client: client, config: config}
}

// GenerateAccessToken for client.
func (c *Client) GenerateAccessToken(ctx context.Context, length uint32) (string, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("Basic %s", c.config.Admin))
	req := &v1.GenerateAccessTokenRequest{Length: length}

	resp, err := c.client.GenerateAccessToken(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Token.Bearer, nil
}

// GenerateServiceToken for client.
func (c *Client) GenerateServiceToken(ctx context.Context, kind, audience string) (string, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("Bearer %s", c.config.Access))
	req := &v1.GenerateServiceTokenRequest{Kind: kind, Audience: audience}

	resp, err := c.client.GenerateServiceToken(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Token.Bearer, nil
}

// VerifyServiceToken for client.
func (c *Client) VerifyServiceToken(ctx context.Context, kind, action, token string) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", fmt.Sprintf("Bearer %s", token))
	req := &v1.VerifyServiceTokenRequest{Kind: kind, Action: action}

	_, err := c.client.VerifyServiceToken(ctx, req)

	return err
}
