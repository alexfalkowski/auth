package client

import (
	"context"

	"github.com/alexfalkowski/go-service/security/token"
)

type contextKey string

var (
	// Kind for client.
	Kind = contextKey("kind")

	// Audience for client.
	Audience = contextKey("audience")

	// Audience for client.
	Action = contextKey("action")
)

// Param is used to pass parameters through context.
type Param string

// NewGenerator for client.
func NewGenerator(client *Client) token.Generator {
	return &generator{client: client}
}

// NewVerifier for client.
func NewVerifier(client *Client) token.Verifier {
	return &verifier{client: client}
}

type generator struct {
	client *Client
}

func (g *generator) Generate(ctx context.Context) (context.Context, []byte, error) {
	kind := ctx.Value(Kind).(string)
	audience := ctx.Value(Audience).(string)

	t, err := g.client.GenerateServiceToken(ctx, kind, audience)

	return ctx, []byte(t), err
}

type verifier struct {
	client *Client
}

func (v *verifier) Verify(ctx context.Context, token []byte) (context.Context, error) {
	kind := ctx.Value(Kind).(string)
	audience := ctx.Value(Audience).(string)
	action := ctx.Value(Action).(string)

	return ctx, v.client.VerifyServiceToken(ctx, string(token), kind, audience, action)
}
