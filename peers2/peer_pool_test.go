package peers2

import (
	"fmt"
	"testing"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"

	"github.com/status-im/status-go/logutils"
	"github.com/status-im/status-go/peers"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/discv5"
)

func init() {
	if err := logutils.OverrideRootLog(true, "DEBUG", "", false); err != nil {
		panic(err)
	}
}

func newInMemoryCache() (*peers.Cache, error) {
	memdb, err := leveldb.Open(storage.NewMemStorage(), nil)
	if err != nil {
		return nil, err
	}
	return peers.NewCache(memdb), nil
}

func createServer() (*p2p.Server, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return &p2p.Server{
		Config: p2p.Config{
			PrivateKey:  key,
			MaxPeers:    10,
			NoDiscovery: true,
			ListenAddr:  "0.0.0.0:0",
		},
	}, nil
}

func TestPeerPoolRequestToAddPeer(t *testing.T) {
	var err error

	// start a peer
	peer1, err := createServer()
	require.NoError(t, err)
	require.NoError(t, peer1.Start())

	// start a server
	server, err := createServer()
	require.NoError(t, err)
	require.NoError(t, server.Start())

	// subscribe to server events
	events := make(chan *p2p.PeerEvent, 20)
	sub := server.SubscribeEvents(events)
	defer sub.Unsubscribe()

	// creat TopicPool and PeerPool
	topic := discv5.Topic("test-topic")
	topicPool := NewTopicPoolBase(nil, topic)
	peerPool := NewPeerPool([]TopicPool{topicPool}, nil)
	peerPool.Start(server)

	// add peer to server
	peerPool.RequestToAddPeer(
		topic,
		discv5.NewNode(discv5.NodeID(peer1.Self().ID), peer1.Self().IP, peer1.Self().TCP, peer1.Self().UDP),
	)
	// verify the peer was added and the PeerPool state
	nodeID, err := getPeerFromEvent(events, p2p.PeerEventTypeAdd)
	require.NoError(t, err)
	require.Equal(t, 1, server.PeerCount())
	require.Equal(t, peer1.Self().ID, discover.NodeID(nodeID))
	require.NotNil(t, peerPool.nodeIDToPeerInfo[nodeID])
	require.Equal(t, []TopicPool{topicPool}, peerPool.nodeIDToPeerInfo[nodeID].topics)
}

func TestPeerPoolCache(t *testing.T) {
	var err error

	// start a peer
	peer1, err := createServer()
	require.NoError(t, err)
	require.NoError(t, peer1.Start())

	// start a server
	server, err := createServer()
	require.NoError(t, err)
	require.NoError(t, server.Start())

	// subscribe to events
	events := make(chan *p2p.PeerEvent, 20)
	sub := server.SubscribeEvents(events)
	defer sub.Unsubscribe()

	// create TopicPool and PeerPool
	topic := discv5.Topic("test-topic")
	topicPool := NewTopicPoolBase(nil, topic)
	cache, err := newInMemoryCache()
	require.NoError(t, err)
	peerPool := NewPeerPool([]TopicPool{topicPool}, cache)
	peerPool.Start(server)

	// add peer to server
	peerPool.RequestToAddPeer(
		topic,
		discv5.NewNode(discv5.NodeID(peer1.Self().ID), peer1.Self().IP, peer1.Self().TCP, peer1.Self().UDP),
	)
	// verify the peer was added and the PeerPool state
	_, err = getPeerFromEvent(events, p2p.PeerEventTypeAdd)
	require.NoError(t, err)
	// TODO(adam): PeerPool should emit events or RequestToAddPeer() should return a channel
	time.Sleep(time.Millisecond * 100)
	require.Len(t, cache.GetPeersRange(topic, 10), 1)

	// TODO(adam): rest removal from cache
}

func TestPeerPoolWithTopicPoolWithLimits(t *testing.T) {
	var err error

	// start a peer
	var peers []*p2p.Server
	for i := 0; i < 2; i++ {
		peer, err := createServer()
		require.NoError(t, err)
		require.NoError(t, peer.Start())
		peers = append(peers, peer)
	}

	// start a server
	server, err := createServer()
	require.NoError(t, err)
	require.NoError(t, server.Start())

	// subscribe to events
	events := make(chan *p2p.PeerEvent, 20)
	sub := server.SubscribeEvents(events)
	defer sub.Unsubscribe()

	// create TopicPool and PeerPool
	topic := discv5.Topic("test-topic")
	topicPool := NewTopicPoolWithLimits(NewTopicPoolBase(nil, topic), 1, 1)
	cache, err := newInMemoryCache()
	require.NoError(t, err)
	peerPool := NewPeerPool([]TopicPool{topicPool}, cache)
	peerPool.Start(server)

	// add peers to server
	for _, peer := range peers {
		peerPool.RequestToAddPeer(
			topic,
			discv5.NewNode(discv5.NodeID(peer.Self().ID), peer.Self().IP, peer.Self().TCP, peer.Self().UDP),
		)
	}

	// verify the peer was added and the PeerPool state
	_, err = getPeerFromEvent(events, p2p.PeerEventTypeAdd)
	require.NoError(t, err)
	_, err = getPeerFromEvent(events, p2p.PeerEventTypeAdd)
	require.NoError(t, err)
	// because the max limit is reached
	_, err = getPeerFromEvent(events, p2p.PeerEventTypeDrop)
	require.NoError(t, err)
	// TODO(adam): PeerPool should emit events or RequestToAddPeer() should return a channel
	time.Sleep(time.Millisecond * 100)
	require.Len(t, cache.GetPeersRange(topic, 10), 2)

	// TODO(adam): rest removal from cache
}

func TestPeerPoolWithTopicPoolEphemeral(t *testing.T) {
	var err error

	// start a peer
	var peers []*p2p.Server
	for i := 0; i < 2; i++ {
		peer, err := createServer()
		require.NoError(t, err)
		require.NoError(t, peer.Start())
		peers = append(peers, peer)
	}

	// start a server
	server, err := createServer()
	require.NoError(t, err)
	require.NoError(t, server.Start())

	// subscribe to events
	events := make(chan *p2p.PeerEvent, 20)
	sub := server.SubscribeEvents(events)
	defer sub.Unsubscribe()

	// create TopicPool and PeerPool
	topic := discv5.Topic("test-topic")
	topicPool := NewTopicPoolEphemeral(
		NewTopicPoolWithLimits(NewTopicPoolBase(nil, topic), 1, 2),
	)
	cache, err := newInMemoryCache()
	require.NoError(t, err)
	peerPool := NewPeerPool([]TopicPool{topicPool}, cache)
	peerPool.Start(server)

	// add peers to server
	for _, peer := range peers {
		peerPool.RequestToAddPeer(
			topic,
			discv5.NewNode(discv5.NodeID(peer.Self().ID), peer.Self().IP, peer.Self().TCP, peer.Self().UDP),
		)

		_, err = getPeerFromEvent(events, p2p.PeerEventTypeAdd)
		require.NoError(t, err)
		_, err = getPeerFromEvent(events, p2p.PeerEventTypeDrop)
		require.NoError(t, err)
	}

	// TODO(adam): PeerPool should emit events or RequestToAddPeer() should return a channel
	time.Sleep(time.Millisecond * 100)
	require.Len(t, cache.GetPeersRange(topic, 10), 2)

	// TODO(adam): rest removal from cache
}

func getPeerFromEvent(events <-chan *p2p.PeerEvent, etype p2p.PeerEventType) (nodeID discover.NodeID, err error) {
	select {
	case ev := <-events:
		if ev.Type == etype {
			return ev.Peer, nil
		}
		return nodeID, fmt.Errorf("invalid event '%s' when expected '%s'", ev.Type, etype)
	case <-time.After(5 * time.Second):
		return nodeID, fmt.Errorf("timed out")
	}
}