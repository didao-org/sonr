package host

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/pkg/errors"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/tools/emitter"
)

type SHost struct {
	host.Host
	*emitter.Emitter

	// Properties
	ctxHost      context.Context
	ctxTileAuth  context.Context
	ctxTileToken context.Context
	privKey      crypto.PrivKey

	// Libp2p
	id   peer.ID
	disc *dsc.RoutingDiscovery
	kdht *dht.IpfsDHT
	mdns discovery.Service

	// Rooms
	pubsub *psub.PubSub
}

// Start Begins Assigning Host Parameters ^
func NewHost(ctx context.Context, kc device.Keychain) (*SHost, error) {
	// Initialize DHT
	var kdhtRef *dht.IpfsDHT
	privKey, err := kc.GetPrivKey(device.Account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get private key")
	}

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(privKey),
		libp2p.DefaultTransports,
		libp2p.ConnectionManager(connmgr.NewConnManager(
			10,          // Lowwater
			15,          // HighWater,
			time.Minute, // GracePeriod
		)),
		libp2p.DefaultStaticRelays(),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			// Create DHT
			kdht, err := dht.New(ctx, h)
			if err != nil {
				return nil, err
			}

			// Set DHT
			kdhtRef = kdht
			return kdht, err
		}),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to initialize host")
	}

	// Create Host
	hn := &SHost{
		ctxHost: ctx,
		Emitter: emitter.New(2048),
		id:      h.ID(),
		Host:    h,
		kdht:    kdhtRef,
		privKey: privKey,
	}

	// Bootstrap Host
	err = hn.Bootstrap()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to bootstrap host")
	}
	return hn, nil
}

// ** ─── Host Info ────────────────────────────────────────────────────────
// Returns Host Node MultiAddr
func (hn *SHost) Pubsub() *psub.PubSub {
	return hn.pubsub
}
