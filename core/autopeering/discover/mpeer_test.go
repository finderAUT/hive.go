package discover

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iotaledger/hive.go/core/autopeering/peer"
	"github.com/iotaledger/hive.go/core/autopeering/peer/peertest"
	"github.com/iotaledger/hive.go/core/identity"
)

func TestWrapPeers(t *testing.T) {
	m := make([]*mpeer, 5)
	p := make([]*peer.Peer, 5)
	for i := range m {
		p[i] = peertest.NewPeer(testNetwork, testIP, i)
		m[i] = newMPeer(p[i])
	}

	wrapP := wrapPeers(p)
	assert.Equal(t, m, wrapP, "wrapPeers")
}

func TestUnwrapPeers(t *testing.T) {
	m := make([]*mpeer, 5)
	p := make([]*peer.Peer, 5)
	for i := range m {
		p[i] = peertest.NewPeer(testNetwork, testIP, i)
		m[i] = newMPeer(p[i])
	}

	unwrapP := unwrapPeers(m)
	assert.Equal(t, p, unwrapP, "unwrapPeers")
}

func TestContainsPeer(t *testing.T) {
	m := make([]*mpeer, 5)
	p := make([]*peer.Peer, 5)
	k := peertest.NewPeer(testNetwork, testIP, 999)
	for i := range m {
		p[i] = peertest.NewPeer(testNetwork, testIP, i)
		m[i] = newMPeer(p[i])
	}

	for i := range m {
		assert.Equal(t, true, containsPeer(m, p[i].ID()), "Contains")
	}
	assert.Equal(t, false, containsPeer(m, k.ID()), "Contains")
}

func TestUnshiftPeer(t *testing.T) {
	m := make([]*mpeer, 5)
	for i := range m {
		m[i] = newMPeer(peertest.NewPeer(testNetwork, testIP, i))
	}

	type testCase struct {
		input    []*mpeer
		toAdd    *mpeer
		expected []*mpeer
	}

	tests := []testCase{
		{
			input:    []*mpeer{},
			toAdd:    m[0],
			expected: []*mpeer{m[0]},
		},
		{
			input:    []*mpeer{m[0]},
			toAdd:    m[1],
			expected: []*mpeer{m[1], m[0]},
		},
		{
			input:    []*mpeer{m[0], m[1]},
			toAdd:    m[2],
			expected: []*mpeer{m[2], m[0], m[1]},
		},
		{
			input:    []*mpeer{m[0], m[1], m[2], m[3]},
			toAdd:    m[4],
			expected: []*mpeer{m[4], m[0], m[1], m[2]},
		},
	}

	for _, test := range tests {
		test.input = unshiftPeer(test.input, test.toAdd, len(m)-1)
		assert.Equal(t, test.expected, test.input, "unshiftPeer")
	}
}

func TestDeletePeer(t *testing.T) {
	m := make([]*mpeer, 5)
	for i := range m {
		m[i] = newMPeer(peertest.NewPeer(testNetwork, testIP, i))
	}

	type testCase struct {
		input    []*mpeer
		toRemove int
		expected []*mpeer
		deleted  *mpeer
	}

	tests := []testCase{
		{
			input:    []*mpeer{m[0]},
			toRemove: 0,
			expected: []*mpeer{},
			deleted:  m[0],
		},
		{
			input:    []*mpeer{m[0], m[1], m[2], m[3]},
			toRemove: 2,
			expected: []*mpeer{m[0], m[1], m[3]},
			deleted:  m[2],
		},
	}

	for _, test := range tests {
		var deleted *mpeer
		test.input, deleted = deletePeer(test.input, test.toRemove)
		assert.Equal(t, test.expected, test.input, "deletePeer_list")
		assert.Equal(t, test.deleted, deleted, "deletePeer_peer")
	}
}

func TestDeletePeerByID(t *testing.T) {
	m := make([]*mpeer, 5)
	p := make([]*peer.Peer, 5)
	for i := range m {
		p[i] = peertest.NewPeer(testNetwork, testIP, i)
		m[i] = newMPeer(p[i])
	}

	type testCase struct {
		input    []*mpeer
		toRemove identity.ID
		expected []*mpeer
		deleted  *mpeer
	}

	tests := []testCase{
		{
			input:    []*mpeer{m[0]},
			toRemove: p[0].ID(),
			expected: []*mpeer{},
			deleted:  m[0],
		},
		{
			input:    []*mpeer{m[0], m[1], m[2], m[3]},
			toRemove: p[2].ID(),
			expected: []*mpeer{m[0], m[1], m[3]},
			deleted:  m[2],
		},
	}

	for _, test := range tests {
		var deleted *mpeer
		test.input, deleted = deletePeerByID(test.input, test.toRemove)
		assert.Equal(t, test.expected, test.input, "deletePeerByID_list")
		assert.Equal(t, test.deleted, deleted, "deletePeerByID_peer")
	}
}

func TestPushPeer(t *testing.T) {
	m := make([]*mpeer, 5)
	max := len(m) - 1
	for i := range m {
		m[i] = newMPeer(peertest.NewPeer(testNetwork, testIP, i))
	}

	type testCase struct {
		input    []*mpeer
		toPush   *mpeer
		expected []*mpeer
	}

	tests := []testCase{
		{
			input:    []*mpeer{},
			toPush:   m[0],
			expected: []*mpeer{m[0]},
		},
		{
			input:    []*mpeer{m[0], m[1]},
			toPush:   m[2],
			expected: []*mpeer{m[0], m[1], m[2]},
		},
		{
			input:    []*mpeer{m[0], m[1], m[2], m[3]},
			toPush:   m[4],
			expected: []*mpeer{m[1], m[2], m[3], m[4]},
		},
	}

	for _, test := range tests {
		test.input = pushPeer(test.input, test.toPush, max)
		assert.Equal(t, test.expected, test.input, "pushPeer")
	}
}
