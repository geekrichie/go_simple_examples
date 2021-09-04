package tcp

import (
	"fmt"
	"io"
	"time"
)

var (
	//默认接收缓冲区大小
	DefaultRecvBufSize =  4 << 10 //4k
	//默认发送缓冲区大小
	DefaultSendBufSize = 1 << 10 //1k
	//是否允许异步写
	DefaultAsyncWrite = true
)

type StopMode uint8

const  (
	// StopImmediately mean stop directly, the cached data maybe will not send.
	StopImmediately StopMode = iota
	// StopGracefullyButNotWait stop and flush cached data.
	StopGracefullyBufNotWait
	// StopGracefullyAndWait stop and block until cached data sended.
	StopGracefullyAndWait
)

type LogLevel uint8

const (
	Panic LogLevel = iota
	Fatal
	Error
	Warn
	Info
	Debug
)

type Logger interface{
	Log(l LogLevel, v ...interface{})
	Logf(l LogLevel, format string, v ...interface{})
}

type emptyLogger struct{}

func(*emptyLogger) Log(l LogLevel, v ...interface{}) {

}

func(*emptyLogger) Logf(l LogLevel, format string, v ...interface{}) {

}

var logger Logger = &emptyLogger{}

func SetLogger(l Logger) {
	logger = l
}


//Handler is the event callback
//dont't block in event handler
type Handler interface {
	//OnAccept mean server accept a new connect
	OnAccept(*Conn)
	//OnConnect mean client connecte to a server
	OnConnect(*Conn)
	//OnRecv mean conn recv a packet
	OnRecv(*Conn, Packet)
	//OnUnPackErr mean failed to unpack recved data
	OnUnpackErr(*Conn, []byte, error)
	//OnClose mean conn is closed
	OnClose(*Conn)
}

type Packet interface{
	fmt.Stringer
}

type Protocol interface {
	// PackSize return the size need for pack the Packet
	PackSize(p Packet) int
	PackTo(p Packet, w io.Writer) (int, error)
	Pack(p Packet) ([]byte, error)
	UnPack(buf []byte) (Packet, int, error)
}

type Options struct {
	Handler   Handler
	Protocol  Protocol
	RecvBufSize  int
	SendBufSize  int
	AsyncWrite   bool
	NoDelay      bool
	KeepAlive    bool
	KeepAlivePeriod time.Duration
	ReadDeadline    time.Duration
	WriteDeadline   time.Duration
}

func NewOpts( h Handler, p Protocol) *Options {
	if h == nil || p == nil {
		panic("tcp.NewOpts: nil Handler or Protocol")
	}
	return &Options{
		Handler: h,
		Protocol: p,
		RecvBufSize: DefaultRecvBufSize,
		SendBufSize: DefaultSendBufSize,
		AsyncWrite: DefaultAsyncWrite,
		NoDelay: true,
		KeepAlive: false,
	}
}