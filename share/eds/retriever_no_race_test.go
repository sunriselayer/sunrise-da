// go:build !race

package eds

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/nmt"
	"github.com/celestiaorg/rsmt2d"
	"github.com/sunriselayer/sunrise/pkg/da"
	"github.com/sunriselayer/sunrise/pkg/wrapper"

	"github.com/sunriselayer/sunrise-da/share"
	"github.com/sunriselayer/sunrise-da/share/eds/byzantine"
	"github.com/sunriselayer/sunrise-da/share/eds/edstest"
	"github.com/sunriselayer/sunrise-da/share/ipld"
)

func TestRetriever_ByzantineError(t *testing.T) {
	const width = 8
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bserv := ipld.NewMemBlockservice()
	shares := edstest.RandEDS(t, width).Flattened()
	_, err := ipld.ImportShares(ctx, shares, bserv)
	require.NoError(t, err)

	// corrupt shares so that eds erasure coding does not match
	copy(shares[14][share.NamespaceSize:], shares[15][share.NamespaceSize:])

	// import corrupted eds
	batchAdder := ipld.NewNmtNodeAdder(ctx, bserv, ipld.MaxSizeBatchOption(width*2))
	attackerEDS, err := rsmt2d.ImportExtendedDataSquare(
		shares,
		share.DefaultRSMT2DCodec(),
		wrapper.NewConstructor(uint64(width),
			nmt.NodeVisitor(batchAdder.Visit)),
	)
	require.NoError(t, err)
	err = batchAdder.Commit()
	require.NoError(t, err)

	// ensure we rcv an error
	dah, err := da.NewDataAvailabilityHeader(attackerEDS)
	require.NoError(t, err)
	r := NewRetriever(bserv)
	_, err = r.Retrieve(ctx, &dah)
	var errByz *byzantine.ErrByzantine
	require.ErrorAs(t, err, &errByz)
}
