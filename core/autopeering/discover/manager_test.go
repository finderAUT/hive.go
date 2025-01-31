package discover

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/iotaledger/hive.go/core/autopeering/peer"
	"github.com/iotaledger/hive.go/core/autopeering/peer/peertest"
	"github.com/iotaledger/hive.go/core/autopeering/server"
	"github.com/iotaledger/hive.go/core/kvstore/mapdb"
)

func TestMgrClose(t *testing.T) {
	_, _, teardown := newManagerTest(t)
	defer teardown()

	time.Sleep(graceTime)
}

func TestMgrVerifyDiscoveredPeer(t *testing.T) {
	mgr, m, teardown := newManagerTest(t)
	defer teardown()

	p := peertest.NewPeer(testNetwork, testIP, 0)

	// expect Ping of peer p
	m.On("Ping", p).Return(nil).Once()
	// ignore DiscoveryRequest calls
	m.On("DiscoveryRequest", mock.Anything).Return([]*peer.Peer{}, nil).Maybe()

	// let the manager initialize
	time.Sleep(graceTime)

	mgr.addDiscoveredPeer(p)

	mgr.doReverify(make(chan struct{})) // manually trigger a verify
	m.AssertExpectations(t)
}

func TestMgrReverifyPeer(t *testing.T) {
	mgr, m, teardown := newManagerTest(t)
	defer teardown()

	p := peertest.NewPeer(testNetwork, testIP, 0)

	// expect Ping of peer p
	m.On("Ping", p).Return(nil).Once()
	// ignore DiscoveryRequest calls
	m.On("DiscoveryRequest", mock.Anything).Return([]*peer.Peer{}, nil).Maybe()

	// let the manager initialize
	time.Sleep(graceTime)

	mgr.addVerifiedPeer(p)

	mgr.doReverify(make(chan struct{})) // manually trigger a verify
	m.AssertExpectations(t)
}

func TestMgrRequestDiscoveredPeer(t *testing.T) {
	mgr, m, teardown := newManagerTest(t)
	defer teardown()

	p1 := peertest.NewPeer(testNetwork, testIP, 1)
	p2 := peertest.NewPeer(testNetwork, testIP, 2)

	// expect DiscoveryRequest on the discovered peer
	m.On("DiscoveryRequest", p1).Return([]*peer.Peer{p2}, nil).Once()
	// ignore any Ping
	m.On("Ping", mock.Anything).Return(nil).Maybe()

	mgr.addVerifiedPeer(p1)
	mgr.addDiscoveredPeer(p2)

	mgr.doQuery(make(chan time.Duration, 1)) // manually trigger a query
	m.AssertExpectations(t)
}

func TestMgrAddManyVerifiedPeers(t *testing.T) {
	mgr, m, teardown := newManagerTest(t)
	defer teardown()

	p := peertest.NewPeer(testNetwork, testIP, 999)

	// expect Ping of peer p
	m.On("Ping", p).Return(nil).Once()
	// ignore DiscoveryRequest calls
	m.On("DiscoveryRequest", mock.Anything).Return([]*peer.Peer{}, nil).Maybe()

	// let the manager initialize
	time.Sleep(graceTime)

	mgr.addVerifiedPeer(p)
	for i := 0; i < maxManaged+maxReplacements; i++ {
		mgr.addVerifiedPeer(peertest.NewPeer(testNetwork, testIP, i))
	}

	mgr.doReverify(make(chan struct{})) // manually trigger a verify
	ps := unwrapPeers(mgr.verifiedPeers())

	assert.Equal(t, maxManaged, len(ps))
	assert.Contains(t, ps, p)

	m.AssertExpectations(t)
}

func TestMgrDeleteUnreachablePeer(t *testing.T) {
	mgr, m, teardown := newManagerTest(t)
	defer teardown()

	p := peertest.NewPeer(testNetwork, testIP, 999)

	// expect Ping of peer p, but return error
	m.On("Ping", p).Return(server.ErrTimeout).Times(1)
	// ignore DiscoveryRequest calls
	m.On("DiscoveryRequest", mock.Anything).Return([]*peer.Peer{}, nil).Maybe()

	// let the manager initialize
	time.Sleep(graceTime)

	mgr.addVerifiedPeer(p)
	for i := 0; i < maxManaged; i++ {
		mgr.addVerifiedPeer(peertest.NewPeer(testNetwork, testIP, i))
	}

	mgr.doReverify(make(chan struct{})) // manually trigger a verify
	ps := unwrapPeers(mgr.verifiedPeers())

	assert.Equal(t, maxManaged, len(ps))
	assert.NotContains(t, ps, p)

	m.AssertExpectations(t)
}

type NetworkMock struct {
	mock.Mock

	loc *peer.Local
}

func newManagerTest(t require.TestingT) (*manager, *NetworkMock, func()) {
	db, err := peer.NewDB(mapdb.NewMapDB())
	require.NoError(t, err)
	local := peertest.NewLocal(testNetwork, net.IPv4zero, testPort, db)
	networkMock := &NetworkMock{
		loc: local,
	}
	mgr := newManager(networkMock, nil, log)
	return mgr, networkMock, mgr.close
}

func (m *NetworkMock) local() *peer.Local {
	return m.loc
}

func (m *NetworkMock) Ping(p *peer.Peer) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *NetworkMock) DiscoveryRequest(p *peer.Peer) ([]*peer.Peer, error) {
	args := m.Called(p)
	return args.Get(0).([]*peer.Peer), args.Error(1)
}
