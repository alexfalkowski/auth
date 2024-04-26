package client

import (
	"context"

	"github.com/alexfalkowski/go-service/security/token"
)

// Token for client.
type Token struct {
	client *Client
}

// NewToken for client.
func NewToken(client *Client) *Token {
	return &Token{client: client}
}

// Generator for token.
func (t *Token) Generator(kind, audience string) token.Generator {
	return &generator{kind: kind, audience: audience, client: t.client}
}

// Verifier for token.
func (t *Token) Verifier(kind, audience, action string) token.Verifier {
	return &verifier{kind: kind, audience: audience, action: action, client: t.client}
}

type generator struct {
	kind, audience string
	client         *Client
}

func (g *generator) Generate(ctx context.Context) (context.Context, []byte, error) {
	t, err := g.client.GenerateServiceToken(ctx, g.kind, g.audience)

	return ctx, []byte(t), err
}

type verifier struct {
	kind, audience, action string
	client                 *Client
}

func (v *verifier) Verify(ctx context.Context, token []byte) (context.Context, error) {
	return ctx, v.client.VerifyServiceToken(ctx, string(token), v.kind, v.audience, v.action)
}
