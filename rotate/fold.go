package rotate

import (
	"context"
)

type k string

type fn func(context.Context) (context.Context, error)

func fold(ctx context.Context, fns ...fn) (context.Context, error) {
	var err error

	for _, f := range fns {
		ctx, err = f(ctx)
		if err != nil {
			return ctx, err
		}
	}

	return ctx, nil
}
