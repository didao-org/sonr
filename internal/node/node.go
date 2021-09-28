package node

import (
	"container/list"
	"context"

	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/exchange"
	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/transfer"
	"github.com/sonr-io/core/tools/logger"
	"github.com/sonr-io/core/tools/state"
	"go.uber.org/zap"
)

// Node Emission Events
const (
	Event_STATUS = "status"
)

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	// Emitter is the event emitter for this node
	*state.Emitter

	// Host and context
	host *host.SNRHost

	// Properties
	ctx context.Context

	// Queue - the transfer queue
	queue *list.List

	// Profile - the node's profile
	profile *common.Profile

	// TransferProtocol - the transfer protocol
	*transfer.TransferProtocol

	// ExchangeProtocol - the exchange protocol
	*exchange.ExchangeProtocol

	// LobbyProtocol - The lobby protocol
	*lobby.LobbyProtocol
}

// NewNode Creates a node with its implemented protocols
func NewNode(ctx context.Context, host *host.SNRHost, loc *common.Location) *Node {
	// Initialize Node
	node := &Node{
		Emitter: state.NewEmitter(2048),
		host:    host,
		ctx:     ctx,
		queue:   list.New(),
	}

	// Set Transfer Protocol
	node.TransferProtocol = transfer.NewProtocol(ctx, host, node.Emitter)

	// Set Exchange Protocol
	exch, err := exchange.NewProtocol(ctx, host, node.Emitter)
	if err != nil {
		logger.Error("Failed to start ExchangeProtocol", zap.Error(err))
		return node
	}
	node.ExchangeProtocol = exch

	// Set Lobby Protocol
	lobby, err := lobby.NewProtocol(host, loc, node.Emitter)
	if err != nil {
		logger.Error("Failed to start LobbyProtocol", zap.Error(err))
		return node
	}
	node.LobbyProtocol = lobby
	return node
}

// Edit method updates Node's profile
func (n *Node) Edit(p *common.Profile) error {
	// Set Profile and Get Peer
	n.profile = p
	peer, err := n.Peer()
	if err != nil {
		return err
	}

	// Push Update to Exchange
	err = n.ExchangeProtocol.Update(peer)
	if err != nil {
		logger.Error("Failed to update Exchange", zap.Error(err))
		return err
	}

	// Push Update to Lobby
	err = n.LobbyProtocol.Update(peer)
	if err != nil {
		logger.Error("Failed to update Lobby", zap.Error(err))
		return err
	}
	return nil
}

// Supply a transfer item to the queue
func (n *Node) Supply(paths []string) error {
	// Create Transfer
	payload, err := common.NewPayload(n.profile, paths)
	if err != nil {
		logger.Error("Failed to Supply Paths", zap.Error(err))
		return err
	}

	// Add items to transfer
	n.queue.PushBack(payload)
	return nil
}

// Share a peer to have a transfer
func (n *Node) Share(to *common.Peer) error {
	// Create New ID for Invite
	id, err := n.host.NewID()
	if err != nil {
		logger.Error("Failed to create new id for Shared Invite", zap.Error(err))
		return err
	}

	// Create new Metadata
	meta, err := n.host.NewMetadata()
	if err != nil {
		logger.Error("Failed to create new metadata for Shared Invite", zap.Error(err))
		return err
	}

	// Fetch Peer Data
	from, err := n.Peer()
	if err != nil {
		return err
	}

	// Create Invite Request
	req := &transfer.InviteRequest{
		Payload:  n.queue.Front().Value.(*common.Payload),
		Metadata: meta,
		To:       to,
		From:     from,
		Uuid:     id,
	}

	// Fetch Peer ID from Exchange
	entry, err := n.ExchangeProtocol.Query(exchange.NewQueryRequestFromSName(from.GetSName()))
	if err != nil {
		logger.Error("Failed to search peer", zap.Error(err))
		return err
	}

	// Invite peer
	err = n.TransferProtocol.Request(entry.PeerID, req)
	if err != nil {
		logger.Error("Failed to invite peer", zap.Error(err))
		n.Emit(Event_STATUS, err)
		return err
	}
	return nil
}

// Respond to an invite request
func (n *Node) Respond(decs bool, to *common.Peer) error {
	// Create new Metadata
	meta, err := n.host.NewMetadata()
	if err != nil {
		logger.Error("Failed to create new metadata for Shared Invite", zap.Error(err))
		return err
	}

	// Fetch Peer Data
	from, err := n.Peer()
	if err != nil {
		return err
	}

	// Create Invite Response
	resp := &transfer.InviteResponse{
		Decision: decs,
		Metadata: meta,
		From:     from,
		To:       to,
	}

	// Respond on TransferProtocol
	err = n.TransferProtocol.Respond(resp)
	if err != nil {
		logger.Error("Failed to respond to invite", zap.Error(err))
		n.Emit(Event_STATUS, err)
		return err
	}
	return nil
}
