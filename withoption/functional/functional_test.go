// +build bench

package func_test

import (
	"testing"

	"github.com/zchee/go-sandbox/withoption"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/tap"
)

type ServerOption func(*withoption.Options)

// WriteBufferSize lets you set the size of write buffer, this determines how much data can be batched
// before doing a write on the wire.
func WriteBufferSize(s int) ServerOption {
	return func(o *withoption.Options) {
		o.WriteBufferSize = s
	}
}

// ReadBufferSize lets you set the size of read buffer, this determines how much data can be read at most
// for one read syscall.
func ReadBufferSize(s int) ServerOption {
	return func(o *withoption.Options) {
		o.ReadBufferSize = s
	}
}

// InitialWindowSize returns a ServerOption that sets window size for stream.
// The lower bound for window size is 64K and any value smaller than that will be ignored.
func InitialWindowSize(s int32) ServerOption {
	return func(o *withoption.Options) {
		o.InitialWindowSize = s
	}
}

// InitialConnWindowSize returns a ServerOption that sets window size for a connection.
// The lower bound for window size is 64K and any value smaller than that will be ignored.
func InitialConnWindowSize(s int32) ServerOption {
	return func(o *withoption.Options) {
		o.InitialConnWindowSize = s
	}
}

// KeepaliveParams returns a ServerOption that sets keepalive and max-age parameters for the server.
func KeepaliveParams(kp keepalive.ServerParameters) ServerOption {
	return func(o *withoption.Options) {
		o.KeepaliveParams = kp
	}
}

// KeepaliveEnforcementPolicy returns a ServerOption that sets keepalive enforcement policy for the server.
func KeepaliveEnforcementPolicy(kep keepalive.EnforcementPolicy) ServerOption {
	return func(o *withoption.Options) {
		o.KeepalivePolicy = kep
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
	return func(o *withoption.Options) {
		o.MaxReceiveMessageSize = m
	}
}

// MaxSendMsgSize returns a ServerOption to set the max message size in bytes the server can send.
// If this is not set, gRPC uses the default 4MB.
func MaxSendMsgSize(m int) ServerOption {
	return func(o *withoption.Options) {
		o.MaxSendMessageSize = m
	}
}

// MaxConcurrentStreams returns a ServerOption that will apply a limit on the number
// of concurrent streams to each ServerTransport.
func MaxConcurrentStreams(n uint32) ServerOption {
	return func(o *withoption.Options) {
		o.MaxConcurrentStreams = n
	}
}

// Creds returns a ServerOption that sets credentials for server connections.
func Creds(c credentials.TransportCredentials) ServerOption {
	return func(o *withoption.Options) {
		o.Creds = c
	}
}

// UnaryInterceptor returns a ServerOption that sets the UnaryServerInterceptor for the
// server. Only one unary interceptor can be installed. The construction of multiple
// interceptors (e.g., chaining) can be implemented at the caller.
func UnaryInterceptor(i grpc.UnaryServerInterceptor) ServerOption {
	return func(o *withoption.Options) {
		if o.UnaryInt != nil {
			panic("The unary server interceptor was already set and may not be reset.")
		}
		o.UnaryInt = i
	}
}

// StreamInterceptor returns a ServerOption that sets the StreamServerInterceptor for the
// server. Only one stream interceptor can be installed.
func StreamInterceptor(i grpc.StreamServerInterceptor) ServerOption {
	return func(o *withoption.Options) {
		if o.StreamInt != nil {
			panic("The stream server interceptor was already set and may not be reset.")
		}
		o.StreamInt = i
	}
}

// InTapHandle returns a ServerOption that sets the tap handle for all the server
// transport to be created. Only one can be installed.
func InTapHandle(h tap.ServerInHandle) ServerOption {
	return func(o *withoption.Options) {
		if o.InTapHandle != nil {
			panic("The tap handle was already set and may not be reset.")
		}
		o.InTapHandle = h
	}
}

// StatsHandler returns a ServerOption that sets the stats handler for the server.
func StatsHandler(h stats.Handler) ServerOption {
	return func(o *withoption.Options) {
		o.StatsHandler = h
	}
}

// UnknownServiceHandler returns a ServerOption that allows for adding a custom
// unknown service handler. The provided method is a bidi-streaming RPC service
// handler that will be invoked instead of returning the "unimplemented" gRPC
// error whenever a request is received for an unregistered service or method.
// The handling function has full access to the Context of the request and the
// stream, and the invocation bypasses interceptors.
func UnknownServiceHandler(streamHandler grpc.StreamHandler) ServerOption {
	return func(o *withoption.Options) {
		o.UnknownStreamDesc = &grpc.StreamDesc{
			StreamName: "unknown_service_handler",
			Handler:    streamHandler,
			// We need to assume that the users of the streamHandler will want to use both.
			ClientStreams: true,
			ServerStreams: true,
		}
	}
}

func New(opt ...ServerOption) *withoption.Server {
	opts := withoption.DefaultServerOptions
	for _, o := range opt {
		o(&opts)
	}

	return &withoption.Server{
		Opts: opts,
	}
}

func BenchmarkNewFunctional(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(MaxMsgSize(100), KeepaliveParams(withoption.DefaultKeepaliveParams), WriteBufferSize(1000))
	}
}
