//nolint:unused
package tests

import (
	"context"
	"testing"
	"time"

	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/stretchr/testify/require"

	"github.com/sunriselayer/sunrise-da/api/rpc/client"
	"github.com/sunriselayer/sunrise-da/libs/authtoken"
	"github.com/sunriselayer/sunrise-da/nodebuilder"
)

func getAdminClient(ctx context.Context, nd *nodebuilder.Node, t *testing.T) *client.Client {
	t.Helper()

	signer := nd.AdminSigner
	listenAddr := "ws://" + nd.RPCServer.ListenAddr()

	jwt, err := authtoken.NewSignedJWT(signer, []auth.Permission{"public", "read", "write", "admin"})
	require.NoError(t, err)

	client, err := client.NewClient(ctx, listenAddr, jwt)
	require.NoError(t, err)

	return client
}

func setTimeInterval(cfg *nodebuilder.Config, interval time.Duration) {
	cfg.P2P.RoutingTableRefreshPeriod = interval
	cfg.Share.Discovery.AdvertiseInterval = interval
}
