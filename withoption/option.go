package withoption

import (
	"math"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/tap"
)

type Server struct {
	Opts Options
}

const (
	DefaultServerMaxReceiveMessageSize = 1024 * 1024 * 4
	DefaultServerMaxSendMessageSize    = math.MaxInt32
)

var DefaultKeepaliveParams = keepalive.ServerParameters{
	Timeout: 10 * time.Second,
}

var DefaultServerOptions = Options{
	MaxReceiveMessageSize: DefaultServerMaxReceiveMessageSize,
	MaxSendMessageSize:    DefaultServerMaxSendMessageSize,
}

type Options struct {
	Creds                 credentials.TransportCredentials
	Codec                 grpc.Codec
	Cp                    grpc.Compressor
	Dc                    grpc.Decompressor
	UnaryInt              grpc.UnaryServerInterceptor
	StreamInt             grpc.StreamServerInterceptor
	InTapHandle           tap.ServerInHandle
	StatsHandler          stats.Handler
	MaxConcurrentStreams  uint32
	MaxReceiveMessageSize int
	MaxSendMessageSize    int
	UseHandlerImpl        bool // use http.Handler-based server
	UnknownStreamDesc     *grpc.StreamDesc
	KeepaliveParams       keepalive.ServerParameters
	KeepalivePolicy       keepalive.EnforcementPolicy
	InitialWindowSize     int32
	InitialConnWindowSize int32
	WriteBufferSize       int
	ReadBufferSize        int
}
