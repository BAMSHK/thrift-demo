// system
#include <iostream>

// lib
#include <thrift/protocol/TBinaryProtocol.h>
// 新增 TMultiplexedProtocol --zxh
#include <thrift/protocol/TMultiplexedProtocol.h>
#include <thrift/transport/TSocket.h>
#include <thrift/transport/TTransportUtils.h>
#include <boost/shared_ptr.hpp>
using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;
using boost::shared_ptr;

// project
#include "../../../gen-cpp/timeServe.h"

// 连接单服务
// int main() {
//   std::shared_ptr<TTransport> socket(new TSocket("10.2.238.171", 9090));
//   std::shared_ptr<TTransport> transport(new TBufferedTransport(socket));
//   std::shared_ptr<TProtocol> protocol(new TBinaryProtocol(transport));
//   // open connect
//   transport->open();
//   // 创建对象
//   timeServeClient client(protocol);
//   auto timeNow = client.getCurrtentTime();
//   std::cout << timeNow << std::endl;
//   transport->close();
//   return 0;
// }
// 连接多服务
int main() {
  std::shared_ptr<TTransport> socket(new TSocket("10.2.238.171", 9090));
  std::shared_ptr<TTransport> transport(new TBufferedTransport(socket));
  std::shared_ptr<TProtocol> protocol(new TBinaryProtocol(transport));
  // 新增 TMultiplexedProtocol   --zxh
  std::shared_ptr<TMultiplexedProtocol> mp (new TMultiplexedProtocol(protocol, "timeService"));
  // open connect
  transport->open();
  // 传入的protocol 为 TMultiplexedProtocol --zxh
  timeServeClient client(mp);
  auto timeNow = client.getCurrtentTime();
  std::cout << timeNow << std::endl;
  transport->close();
  return 0;
}
