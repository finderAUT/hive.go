package node

import (
	"github.com/finderAUT/hive.go/v2/daemon"
)

type NodeOptions struct {
	plugins []*Plugin
	daemon  daemon.Daemon
}

func newNodeOptions(optionalOptions []NodeOption) *NodeOptions {
	result := &NodeOptions{}

	for _, optionalOption := range optionalOptions {
		optionalOption(result)
	}

	if result.daemon == nil {
		result.daemon = daemon.New()
	}

	return result
}

type NodeOption func(*NodeOptions)

func Plugins(plugins ...*Plugin) NodeOption {
	return func(args *NodeOptions) {
		args.plugins = append(args.plugins, plugins...)
	}
}

func Daemon(daemon daemon.Daemon) NodeOption {
	return func(args *NodeOptions) {
		args.daemon = daemon
	}
}
