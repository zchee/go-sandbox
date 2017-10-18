// +build bench

package kv_test

import (
	"fmt"
	"testing"

	"github.com/zchee/go-sandbox/withoption"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

type key int

const (
	keyWriteBufferSize key = 1 << iota
	keyReadBufferSize
	keyInitialWindowSize
	keyInitialConnWindowSize
	keyKeepaliveParams
	keyKeepaliveEnforcementPolicy
	keyMaxMsgSize
	keyMaxRecvMsgSize
	keyMaxSendMsgSize
	keyMaxConcurrentStreams
	keyCreds
)

type ServerOption interface {
	Key() key
	Value() interface{}
}

type serverOption struct {
	key   key
	value interface{}
}

func (s serverOption) Key() key {
	return s.key
}

func (s serverOption) Value() interface{} {
	return s.value
}

// WriteBufferSize lets you set the size of write buffer, this determines how much data can be batched
// before doing a write on the wire.
func WriteBufferSize(s int) ServerOption {
	return &serverOption{
		key:   keyWriteBufferSize,
		value: s,
	}
}

// ReadBufferSize lets you set the size of read buffer, this determines how much data can be read at most
// for one read syscall.
func ReadBufferSize(s int) ServerOption {
	return &serverOption{
		key:   keyReadBufferSize,
		value: s,
	}
}

// InitialWindowSize returns a ServerOption that sets window size for stream.
// The lower bound for window size is 64K and any value smaller than that will be ignored.
func InitialWindowSize(s int32) ServerOption {
	return &serverOption{
		key:   keyInitialWindowSize,
		value: s,
	}
}

// InitialConnWindowSize returns a ServerOption that sets window size for a connection.
// The lower bound for window size is 64K and any value smaller than that will be ignored.
func InitialConnWindowSize(s int32) ServerOption {
	return &serverOption{
		key:   keyInitialConnWindowSize,
		value: s,
	}
}

// KeepaliveParams returns a ServerOption that sets keepalive and max-age parameters for the server.
func KeepaliveParams(kp keepalive.ServerParameters) ServerOption {
	return &serverOption{
		key:   keyKeepaliveParams,
		value: kp,
	}
}

// KeepaliveEnforcementPolicy returns a ServerOption that sets keepalive enforcement policy for the server.
func KeepaliveEnforcementPolicy(kep keepalive.EnforcementPolicy) ServerOption {
	return &serverOption{
		key:   keyKeepaliveEnforcementPolicy,
		value: kep,
	}
}

// MaxMsgSize returns a ServerOption to set the max message size in bytes the server can receive.
// If this is not set, gRPC uses the default limit. Deprecated: use MaxRecvMsgSize instead.
func MaxMsgSize(m int) ServerOption {
	return MaxRecvMsgSize(m)
}

// MaxRecvMsgSize returns a ServerOption to set the max message size in bytes the server can receive.
// If this is not set, gRPC uses the default 4MB.
func MaxRecvMsgSize(m int) ServerOption {
	return &serverOption{
		key:   keyMaxRecvMsgSize,
		value: m,
	}
}

// MaxSendMsgSize returns a ServerOption to set the max message size in bytes the server can send.
// If this is not set, gRPC uses the default 4MB.
func MaxSendMsgSize(m int) ServerOption {
	return &serverOption{
		key:   keyMaxSendMsgSize,
		value: m,
	}
}

// MaxConcurrentStreams returns a ServerOption that will apply a limit on the number
// of concurrent streams to each ServerTransport.
func MaxConcurrentStreams(n uint32) ServerOption {
	return &serverOption{
		key:   keyMaxConcurrentStreams,
		value: n,
	}
}

// Creds returns a ServerOption that sets credentials for server connections.
func Creds(c credentials.TransportCredentials) ServerOption {
	return &serverOption{
		key:   keyCreds,
		value: c,
	}
}

func New(opt ...ServerOption) *withoption.Server {
	opts := withoption.DefaultServerOptions
	for _, o := range opt {
		switch o.Key() {
		case keyWriteBufferSize:
			opts.WriteBufferSize = o.Value().(int)
		case keyReadBufferSize:
			opts.ReadBufferSize = o.Value().(int)
		case keyInitialWindowSize:
			opts.InitialWindowSize = o.Value().(int32)
		case keyInitialConnWindowSize:
			opts.InitialConnWindowSize = o.Value().(int32)
		case keyKeepaliveParams:
			opts.KeepaliveParams = o.Value().(keepalive.ServerParameters)
		case keyKeepaliveEnforcementPolicy:
			opts.KeepalivePolicy = o.Value().(keepalive.EnforcementPolicy)
		case keyMaxMsgSize, keyMaxRecvMsgSize:
			opts.MaxReceiveMessageSize = o.Value().(int)
		case keyMaxConcurrentStreams:
			opts.MaxConcurrentStreams = o.Value().(uint32)
		case keyCreds:
			opts.Creds = o.Value().(credentials.TransportCredentials)
		default:
			panic(fmt.Sprintf("unknown key: %s", o.Key()))
		}
	}

	return &withoption.Server{
		Opts: opts,
	}
}

func BenchmarkNewKeyValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(MaxMsgSize(100), KeepaliveParams(withoption.DefaultKeepaliveParams), WriteBufferSize(1000))
	}
}
