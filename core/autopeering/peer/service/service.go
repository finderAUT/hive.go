package service

// Endpoint defines where a service can be reached.
type Endpoint interface {
	Network() string
	Port() int
}

// Service is a read-only interface to access services.
type Service interface {
	// Get returns the network end point address of the given service or nil if not supported.
	Get(Key) Endpoint

	// CreateRecord creates a modifyable Record from the services.
	CreateRecord() *Record
}

// Key is the type of keys used to look-up a service.
type Key string

const (
	// PeeringKey is the key for the auto peering service.
	PeeringKey Key = "peering"
	// FPCKey is the key for the FPC service.
	FPCKey Key = "fpc"
	// P2PKey is the key for the p2p service.
	P2PKey Key = "p2p"
)
