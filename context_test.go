package river

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestClientFromContext(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	client := &Client[pgx.Tx]{}
	ctx = withClient(ctx, client)

	require.Equal(t, client, ClientFromContext[pgx.Tx](ctx))

	result, err := ClientFromContextSafely[pgx.Tx](ctx)
	require.NoError(t, err)
	require.Equal(t, client, result)

	require.Panics(t, func() {
		ClientFromContext[pgx.Tx](context.Background())
	})

	result, err = ClientFromContextSafely[pgx.Tx](context.Background())
	require.Error(t, err)
	require.Nil(t, result)
}
