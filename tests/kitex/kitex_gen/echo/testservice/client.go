// Code generated by Kitex v0.7.0. DO NOT EDIT.

package testservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	dubbo "github.com/kitex-contrib/codec-dubbo/pkg"
	echo "github.com/kitex-contrib/codec-dubbo/tests/kitex/kitex_gen/echo"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	EchoInt(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error)
	EchoBool(ctx context.Context, req bool, callOptions ...callopt.Option) (r bool, err error)
	EchoByte(ctx context.Context, req int8, callOptions ...callopt.Option) (r int8, err error)
	EchoInt16(ctx context.Context, req int16, callOptions ...callopt.Option) (r int16, err error)
	EchoInt32(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error)
	EchoInt64(ctx context.Context, req int64, callOptions ...callopt.Option) (r int64, err error)
	EchoDouble(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error)
	EchoString(ctx context.Context, req string, callOptions ...callopt.Option) (r string, err error)
	EchoBinary(ctx context.Context, req []byte, callOptions ...callopt.Option) (r []byte, err error)
	Echo(ctx context.Context, req *echo.EchoRequest, callOptions ...callopt.Option) (r *echo.EchoResponse, err error)
	EchoBoolList(ctx context.Context, req []bool, callOptions ...callopt.Option) (r []bool, err error)
	EchoByteList(ctx context.Context, req []int8, callOptions ...callopt.Option) (r []int8, err error)
	EchoInt16List(ctx context.Context, req []int16, callOptions ...callopt.Option) (r []int16, err error)
	EchoInt32List(ctx context.Context, req []int32, callOptions ...callopt.Option) (r []int32, err error)
	EchoInt64List(ctx context.Context, req []int64, callOptions ...callopt.Option) (r []int64, err error)
	EchoDoubleList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error)
	EchoStringList(ctx context.Context, req []string, callOptions ...callopt.Option) (r []string, err error)
	EchoBinaryList(ctx context.Context, req [][]byte, callOptions ...callopt.Option) (r [][]byte, err error)
	EchoBool2BoolMap(ctx context.Context, req map[bool]bool, callOptions ...callopt.Option) (r map[bool]bool, err error)
	EchoBool2ByteMap(ctx context.Context, req map[bool]int8, callOptions ...callopt.Option) (r map[bool]int8, err error)
	EchoBool2Int16Map(ctx context.Context, req map[bool]int16, callOptions ...callopt.Option) (r map[bool]int16, err error)
	EchoBool2Int32Map(ctx context.Context, req map[bool]int32, callOptions ...callopt.Option) (r map[bool]int32, err error)
	EchoBool2Int64Map(ctx context.Context, req map[bool]int64, callOptions ...callopt.Option) (r map[bool]int64, err error)
	EchoBool2DoubleMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error)
	EchoBool2StringMap(ctx context.Context, req map[bool]string, callOptions ...callopt.Option) (r map[bool]string, err error)
	EchoBool2BinaryMap(ctx context.Context, req map[bool][]byte, callOptions ...callopt.Option) (r map[bool][]byte, err error)
	EchoBool2BoolListMap(ctx context.Context, req map[bool][]bool, callOptions ...callopt.Option) (r map[bool][]bool, err error)
	EchoBool2ByteListMap(ctx context.Context, req map[bool][]int8, callOptions ...callopt.Option) (r map[bool][]int8, err error)
	EchoBool2Int16ListMap(ctx context.Context, req map[bool][]int16, callOptions ...callopt.Option) (r map[bool][]int16, err error)
	EchoBool2Int32ListMap(ctx context.Context, req map[bool][]int32, callOptions ...callopt.Option) (r map[bool][]int32, err error)
	EchoBool2Int64ListMap(ctx context.Context, req map[bool][]int64, callOptions ...callopt.Option) (r map[bool][]int64, err error)
	EchoBool2DoubleListMap(ctx context.Context, req map[bool][]float64, callOptions ...callopt.Option) (r map[bool][]float64, err error)
	EchoBool2StringListMap(ctx context.Context, req map[bool][]string, callOptions ...callopt.Option) (r map[bool][]string, err error)
	EchoBool2BinaryListMap(ctx context.Context, req map[bool][][]byte, callOptions ...callopt.Option) (r map[bool][][]byte, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithCodec(dubbo.NewDubboCodec()))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTestServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTestServiceClient struct {
	*kClient
}

func (p *kTestServiceClient) EchoInt(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt(ctx, req)
}

func (p *kTestServiceClient) EchoBool(ctx context.Context, req bool, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool(ctx, req)
}

func (p *kTestServiceClient) EchoByte(ctx context.Context, req int8, callOptions ...callopt.Option) (r int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoByte(ctx, req)
}

func (p *kTestServiceClient) EchoInt16(ctx context.Context, req int16, callOptions ...callopt.Option) (r int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt16(ctx, req)
}

func (p *kTestServiceClient) EchoInt32(ctx context.Context, req int32, callOptions ...callopt.Option) (r int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt32(ctx, req)
}

func (p *kTestServiceClient) EchoInt64(ctx context.Context, req int64, callOptions ...callopt.Option) (r int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt64(ctx, req)
}

func (p *kTestServiceClient) EchoDouble(ctx context.Context, req float64, callOptions ...callopt.Option) (r float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoDouble(ctx, req)
}

func (p *kTestServiceClient) EchoString(ctx context.Context, req string, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoString(ctx, req)
}

func (p *kTestServiceClient) EchoBinary(ctx context.Context, req []byte, callOptions ...callopt.Option) (r []byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBinary(ctx, req)
}

func (p *kTestServiceClient) Echo(ctx context.Context, req *echo.EchoRequest, callOptions ...callopt.Option) (r *echo.EchoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Echo(ctx, req)
}

func (p *kTestServiceClient) EchoBoolList(ctx context.Context, req []bool, callOptions ...callopt.Option) (r []bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBoolList(ctx, req)
}

func (p *kTestServiceClient) EchoByteList(ctx context.Context, req []int8, callOptions ...callopt.Option) (r []int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoByteList(ctx, req)
}

func (p *kTestServiceClient) EchoInt16List(ctx context.Context, req []int16, callOptions ...callopt.Option) (r []int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt16List(ctx, req)
}

func (p *kTestServiceClient) EchoInt32List(ctx context.Context, req []int32, callOptions ...callopt.Option) (r []int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt32List(ctx, req)
}

func (p *kTestServiceClient) EchoInt64List(ctx context.Context, req []int64, callOptions ...callopt.Option) (r []int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoInt64List(ctx, req)
}

func (p *kTestServiceClient) EchoDoubleList(ctx context.Context, req []float64, callOptions ...callopt.Option) (r []float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoDoubleList(ctx, req)
}

func (p *kTestServiceClient) EchoStringList(ctx context.Context, req []string, callOptions ...callopt.Option) (r []string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoStringList(ctx, req)
}

func (p *kTestServiceClient) EchoBinaryList(ctx context.Context, req [][]byte, callOptions ...callopt.Option) (r [][]byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBinaryList(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BoolMap(ctx context.Context, req map[bool]bool, callOptions ...callopt.Option) (r map[bool]bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BoolMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2ByteMap(ctx context.Context, req map[bool]int8, callOptions ...callopt.Option) (r map[bool]int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2ByteMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int16Map(ctx context.Context, req map[bool]int16, callOptions ...callopt.Option) (r map[bool]int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int16Map(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int32Map(ctx context.Context, req map[bool]int32, callOptions ...callopt.Option) (r map[bool]int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int32Map(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int64Map(ctx context.Context, req map[bool]int64, callOptions ...callopt.Option) (r map[bool]int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int64Map(ctx, req)
}

func (p *kTestServiceClient) EchoBool2DoubleMap(ctx context.Context, req map[bool]float64, callOptions ...callopt.Option) (r map[bool]float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2DoubleMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2StringMap(ctx context.Context, req map[bool]string, callOptions ...callopt.Option) (r map[bool]string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2StringMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BinaryMap(ctx context.Context, req map[bool][]byte, callOptions ...callopt.Option) (r map[bool][]byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BinaryMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BoolListMap(ctx context.Context, req map[bool][]bool, callOptions ...callopt.Option) (r map[bool][]bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BoolListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2ByteListMap(ctx context.Context, req map[bool][]int8, callOptions ...callopt.Option) (r map[bool][]int8, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2ByteListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int16ListMap(ctx context.Context, req map[bool][]int16, callOptions ...callopt.Option) (r map[bool][]int16, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int16ListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int32ListMap(ctx context.Context, req map[bool][]int32, callOptions ...callopt.Option) (r map[bool][]int32, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int32ListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2Int64ListMap(ctx context.Context, req map[bool][]int64, callOptions ...callopt.Option) (r map[bool][]int64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2Int64ListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2DoubleListMap(ctx context.Context, req map[bool][]float64, callOptions ...callopt.Option) (r map[bool][]float64, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2DoubleListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2StringListMap(ctx context.Context, req map[bool][]string, callOptions ...callopt.Option) (r map[bool][]string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2StringListMap(ctx, req)
}

func (p *kTestServiceClient) EchoBool2BinaryListMap(ctx context.Context, req map[bool][][]byte, callOptions ...callopt.Option) (r map[bool][][]byte, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoBool2BinaryListMap(ctx, req)
}
