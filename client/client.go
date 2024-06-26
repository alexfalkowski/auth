package client

import (
	"context"

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

// GenerateServiceToken for client.
func (c *Client) GenerateServiceToken(ctx context.Context, kind, audience string) (string, error) {
	k, err := c.config.GetAccess()
	if err != nil {
		return "", err
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+k)
	req := &v1.GenerateServiceTokenRequest{Kind: kind, Audience: audience}

	resp, err := c.client.GenerateServiceToken(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.GetToken().GetBearer(), nil
}

// VerifyServiceToken for client.
func (c *Client) VerifyServiceToken(ctx context.Context, token, kind, audience, action string) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	req := &v1.VerifyServiceTokenRequest{Kind: kind, Audience: audience, Action: action}

	_, err := c.client.VerifyServiceToken(ctx, req)

	return err
}
