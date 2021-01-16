package rpc

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

type Options struct {
	Protocol string
	Framed   bool
	Buffered bool
	Ip       string
	Port     int
}
type Option func(*Options)

func newOptions(opts ...Option) *Options {
	options := &Options{
		Protocol: "binary",
		Framed:   false,
		Buffered: false,
		Ip:       "127.0.0.1,10.2.238.171",
		Port:     9090,
	}
	for _, o := range opts {
		o(options)
	}
	return options
}

type ThriftService interface {
	Run(handlers []ThriftHandlers)
	runServer(addr string, handlers []ThriftHandlers) error
}
type thriftService struct {
	opts *Options
}

type ThriftHandlers struct {
	ServiceName string
	Processor   thrift.TProcessor
}

func NewServer(opts ...Option) ThriftService {
	return newThriftService(opts...)
}

func newThriftService(opts ...Option) ThriftService {
	options := newOptions(opts...)
	return &thriftService{
		opts: options,
	}
}

func (s *thriftService) Run(handlers []ThriftHandlers) {
	ips := strings.Split(s.opts.Ip, ",")
	if len(ips) > 0 {
		for _, v := range ips {
			ip := v
			g.Go(func() error {
				return s.runServer(ip+":"+strconv.Itoa(s.opts.Port), handlers)
			})
		}
	}
	if err := g.Wait(); err != nil {
		log.Printf("ThriftService false,%v\n", err)
	}
}

func (s *thriftService) runServer(addr string, handlers []ThriftHandlers) error {

	var protocolFactory thrift.TProtocolFactory
	switch s.opts.Protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		log.Printf("Invalid protocol specified,%s\n", s.opts.Protocol)
		os.Exit(1)
	}
	var transportFactory thrift.TTransportFactory
	if s.opts.Buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}
	if s.opts.Framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}
	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		log.Printf("transport generate  false %v\n", err)
		return err
	}
	log.Printf("【thrift-server】%T", transport)
	//服务注册
	multiProcessor := thrift.NewTMultiplexedProcessor()
	// 给每个service起一个名字
	if len(handlers) > 0 {
		for _, handler := range handlers {
			multiProcessor.RegisterProcessor(handler.ServiceName, handler.Processor)
		}
	}
	server := thrift.NewTSimpleServer4(multiProcessor, transport, transportFactory, protocolFactory)
	log.Println("【thrift-server】Starting the simple server... on ", addr)
	return server.Serve()
	// ===========
	// // 下面为单个服务注册得到代码
	// server := thrift.NewTSimpleServer4(handlers[0].Processor, transport, transportFactory, protocolFactory)
	// return server.Serve()
}
