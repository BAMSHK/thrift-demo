// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package timerpc

import(
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type TimeServe interface {
  GetCurrtentTime(ctx context.Context) (r int32, err error)
}

type TimeServeClient struct {
  c thrift.TClient
}

func NewTimeServeClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TimeServeClient {
  return &TimeServeClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewTimeServeClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TimeServeClient {
  return &TimeServeClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewTimeServeClient(c thrift.TClient) *TimeServeClient {
  return &TimeServeClient{
    c: c,
  }
}

func (p *TimeServeClient) Client_() thrift.TClient {
  return p.c
}
func (p *TimeServeClient) GetCurrtentTime(ctx context.Context) (r int32, err error) {
  var _args0 TimeServeGetCurrtentTimeArgs
  var _result1 TimeServeGetCurrtentTimeResult
  if err = p.Client_().Call(ctx, "getCurrtentTime", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type TimeServeProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler TimeServe
}

func (p *TimeServeProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *TimeServeProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *TimeServeProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewTimeServeProcessor(handler TimeServe) *TimeServeProcessor {

  self2 := &TimeServeProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["getCurrtentTime"] = &timeServeProcessorGetCurrtentTime{handler:handler}
return self2
}

func (p *TimeServeProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x3

}

type timeServeProcessorGetCurrtentTime struct {
  handler TimeServe
}

func (p *timeServeProcessorGetCurrtentTime) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TimeServeGetCurrtentTimeArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getCurrtentTime", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TimeServeGetCurrtentTimeResult{}
var retval int32
  var err2 error
  if retval, err2 = p.handler.GetCurrtentTime(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getCurrtentTime: " + err2.Error())
    oprot.WriteMessageBegin("getCurrtentTime", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("getCurrtentTime", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

type TimeServeGetCurrtentTimeArgs struct {
}

func NewTimeServeGetCurrtentTimeArgs() *TimeServeGetCurrtentTimeArgs {
  return &TimeServeGetCurrtentTimeArgs{}
}

func (p *TimeServeGetCurrtentTimeArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TimeServeGetCurrtentTimeArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getCurrtentTime_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TimeServeGetCurrtentTimeArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TimeServeGetCurrtentTimeArgs(%+v)", *p)
}

// Attributes:
//  - Success
type TimeServeGetCurrtentTimeResult struct {
  Success *int32 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewTimeServeGetCurrtentTimeResult() *TimeServeGetCurrtentTimeResult {
  return &TimeServeGetCurrtentTimeResult{}
}

var TimeServeGetCurrtentTimeResult_Success_DEFAULT int32
func (p *TimeServeGetCurrtentTimeResult) GetSuccess() int32 {
  if !p.IsSetSuccess() {
    return TimeServeGetCurrtentTimeResult_Success_DEFAULT
  }
return *p.Success
}
func (p *TimeServeGetCurrtentTimeResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TimeServeGetCurrtentTimeResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TimeServeGetCurrtentTimeResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *TimeServeGetCurrtentTimeResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getCurrtentTime_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TimeServeGetCurrtentTimeResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteI32(int32(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TimeServeGetCurrtentTimeResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TimeServeGetCurrtentTimeResult(%+v)", *p)
}


