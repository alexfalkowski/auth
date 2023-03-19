package client_test

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"
	"time"

	v1 "github.com/alexfalkowski/auth/api/auth/v1"
	"github.com/alexfalkowski/auth/client"
	v1c "github.com/alexfalkowski/auth/client/v1/config"
	"github.com/alexfalkowski/auth/cmd"
	"github.com/alexfalkowski/auth/config"
	"github.com/alexfalkowski/go-service/logger"
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/metrics"
	"github.com/alexfalkowski/go-service/transport"
	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"google.golang.org/grpc"
)

var options = []fx.Option{
	fx.NopLogger, marshaller.Module,
	cmd.Module, config.Module,
	logger.ZapModule, metrics.PrometheusModule, transport.Module,
	client.Module, fx.Invoke(register),
}

func TestValidSetup(t *testing.T) {
	Convey("Given I have a app", t, func() {
		So(os.Setenv("CONFIG_FILE", "../test/.config/client.yaml"), ShouldBeNil)

		l, err := net.Listen("tcp", "localhost:8080")
		So(err, ShouldBeNil)

		server := grpc.NewServer()

		go server.Serve(l) //nolint:errcheck

		time.Sleep(time.Second)

		app := fxtest.New(t, options...)

		Convey("When I start the app", func() {
			app.RequireStart()

			Convey("Then I should have a started app", func() {
				app.RequireStop()
			})
		})

		server.Stop()

		So(os.Unsetenv("CONFIG_FILE"), ShouldBeNil)
	})
}

func register(_ *client.Client) {}

func TestValidGenerateAccessToken(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		cfg := &v1c.Config{}
		v := &validClient{}
		c := client.NewClient(v, cfg)

		Convey("When I generate an access token", func() {
			t, err := c.GenerateAccessToken(context.Background(), 0)
			So(err, ShouldBeNil)

			Convey("Then I should have a valid token", func() {
				So(t, ShouldEqual, "test")
			})
		})
	})
}

func TestInvalidGenerateAccessToken(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		cfg := &v1c.Config{}
		v := &invalidClient{}
		c := client.NewClient(v, cfg)

		Convey("When I generate an access token", func() {
			_, err := c.GenerateAccessToken(context.Background(), 0)

			Convey("Then I should have an error", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestValidGenerateServiceToken(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		cfg := &v1c.Config{}
		v := &validClient{}
		c := client.NewClient(v, cfg)

		Convey("When I generate an service token", func() {
			t, err := c.GenerateServiceToken(context.Background(), "test", "test")
			So(err, ShouldBeNil)

			Convey("Then I should have a valid token", func() {
				So(t, ShouldEqual, "test")
			})
		})
	})
}

func TestInvalidServiceAccessToken(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		cfg := &v1c.Config{}
		v := &invalidClient{}
		c := client.NewClient(v, cfg)

		Convey("When I generate an service token", func() {
			_, err := c.GenerateServiceToken(context.Background(), "test", "test")

			Convey("Then I should have an error", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

func TestValidVerifyServiceToken(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		cfg := &v1c.Config{}
		v := &validClient{}
		c := client.NewClient(v, cfg)

		Convey("When I verify a service token", func() {
			err := c.VerifyServiceToken(context.Background(), "test", "test", "test")

			Convey("Then I should have a valid token", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestInvalidVerifyServiceToken(t *testing.T) {
	Convey("Given I have a valid client", t, func() {
		cfg := &v1c.Config{}
		v := &invalidClient{}
		c := client.NewClient(v, cfg)

		Convey("When I generate an access token", func() {
			err := c.VerifyServiceToken(context.Background(), "test", "test", "test")

			Convey("Then I should have an error", func() {
				So(err, ShouldBeError)
			})
		})
	})
}

type validClient struct{}

func (*validClient) GenerateAccessToken(_ context.Context, _ *v1.GenerateAccessTokenRequest, _ ...grpc.CallOption) (*v1.GenerateAccessTokenResponse, error) {
	return &v1.GenerateAccessTokenResponse{Token: &v1.AccessToken{Bearer: "test"}}, nil
}

func (*validClient) GenerateKey(_ context.Context, _ *v1.GenerateKeyRequest, _ ...grpc.CallOption) (*v1.GenerateKeyResponse, error) {
	return &v1.GenerateKeyResponse{}, nil
}

func (*validClient) GeneratePassword(_ context.Context, _ *v1.GeneratePasswordRequest, _ ...grpc.CallOption) (*v1.GeneratePasswordResponse, error) {
	return &v1.GeneratePasswordResponse{}, nil
}

func (*validClient) GenerateServiceToken(_ context.Context, _ *v1.GenerateServiceTokenRequest, _ ...grpc.CallOption) (*v1.GenerateServiceTokenResponse, error) {
	return &v1.GenerateServiceTokenResponse{Token: &v1.ServiceToken{Bearer: "test"}}, nil
}

func (*validClient) GetPublicKey(_ context.Context, _ *v1.GetPublicKeyRequest, _ ...grpc.CallOption) (*v1.GetPublicKeyResponse, error) {
	return &v1.GetPublicKeyResponse{}, nil
}

func (*validClient) VerifyServiceToken(_ context.Context, _ *v1.VerifyServiceTokenRequest, _ ...grpc.CallOption) (*v1.VerifyServiceTokenResponse, error) {
	return &v1.VerifyServiceTokenResponse{}, nil
}

type invalidClient struct{}

func (*invalidClient) GenerateAccessToken(_ context.Context, _ *v1.GenerateAccessTokenRequest, _ ...grpc.CallOption) (*v1.GenerateAccessTokenResponse, error) {
	return &v1.GenerateAccessTokenResponse{}, errors.New("an issue")
}

func (*invalidClient) GenerateKey(_ context.Context, _ *v1.GenerateKeyRequest, _ ...grpc.CallOption) (*v1.GenerateKeyResponse, error) {
	return &v1.GenerateKeyResponse{}, errors.New("an issue")
}

func (*invalidClient) GeneratePassword(_ context.Context, _ *v1.GeneratePasswordRequest, _ ...grpc.CallOption) (*v1.GeneratePasswordResponse, error) {
	return &v1.GeneratePasswordResponse{}, errors.New("an issue")
}

func (*invalidClient) GenerateServiceToken(_ context.Context, _ *v1.GenerateServiceTokenRequest, _ ...grpc.CallOption) (*v1.GenerateServiceTokenResponse, error) {
	return &v1.GenerateServiceTokenResponse{}, errors.New("an issue")
}

func (*invalidClient) GetPublicKey(_ context.Context, _ *v1.GetPublicKeyRequest, _ ...grpc.CallOption) (*v1.GetPublicKeyResponse, error) {
	return &v1.GetPublicKeyResponse{}, errors.New("an issue")
}

func (*invalidClient) VerifyServiceToken(_ context.Context, _ *v1.VerifyServiceTokenRequest, _ ...grpc.CallOption) (*v1.VerifyServiceTokenResponse, error) {
	return &v1.VerifyServiceTokenResponse{}, errors.New("an issue")
}
