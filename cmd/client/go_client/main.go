package main

// 测试thrift客户端连接 使用TMultiplexedProtocol注册服务的 thrift服务
import (
	"context"
	"fmt"
	"sample/gen-go/timerpc"

	"github.com/apache/thrift/lib/go/thrift"
)

var ctx = context.Background()

//  获取TMultiplexedProtocol
func GetMultiplexedProtocol(serviceName string) *thrift.TMultiplexedProtocol {
	addr := "10.2.238.171:9090"
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	//Binary protocol
	var protocolFactory thrift.TProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	//no buffered
	var transportFactory thrift.TTransportFactory = thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error get transport :", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error get transport :", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	prot := thrift.NewTMultiplexedProtocol(iprot, serviceName)

	return prot
}

// 获取TimeService客户端
func GetTimeServiceClient(prot *thrift.TMultiplexedProtocol) *timerpc.TimeServeClient {
	return timerpc.NewTimeServeClient(thrift.NewTStandardClient(prot, prot))
}

func main() {
	prot := GetMultiplexedProtocol("timeService")
	client := GetTimeServiceClient(prot)
	// 调用TimeService中的GetCurrtentTime(ctx)方法
	timeDate, err := client.GetCurrtentTime(ctx)
	if err != nil {
		fmt.Printf("err is :%v\n", err)
	}
	fmt.Printf("currentTime is %d clock\n", timeDate)
}
