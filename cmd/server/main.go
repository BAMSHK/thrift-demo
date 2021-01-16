package main

import (
	"sample/gen-go/Sample"
	"sample/gen-go/timerpc"
	"sample/internal/service"
	"sample/rpc"
)

func main() {
	startTriftServer()
}
func startTriftServer() {
	handlers := make([]rpc.ThriftHandlers, 0, 2)
	handlers = append(handlers, rpc.ThriftHandlers{ServiceName: "greeterService", Processor: Sample.NewGreeterProcessor(service.NewGreeterHandle())})
	handlers = append(handlers, rpc.ThriftHandlers{ServiceName: "timeService", Processor: timerpc.NewTimeServeProcessor(service.NewTimeServerHandle())})
	//start
	serve := rpc.NewServer()
	serve.Run(handlers)

}
