// Code generated by thriftgo (0.3.0). DO NOT EDIT.

package echo

import (
	"context"
	"fmt"
)

type EchoRequest struct {
	Int32 int32 `thrift:"int32,1,required" frugal:"1,required,i32" json:"int32"`
}

func NewEchoRequest() *EchoRequest {
	return &EchoRequest{}
}

func (p *EchoRequest) GetInt32() (v int32) {
	return p.Int32
}
func (p *EchoRequest) SetInt32(val int32) {
	p.Int32 = val
}

func (p *EchoRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoRequest(%+v)", *p)
}

type EchoResponse struct {
	Int32 int32 `thrift:"int32,1,required" frugal:"1,required,i32" json:"int32"`
}

func NewEchoResponse() *EchoResponse {
	return &EchoResponse{}
}

func (p *EchoResponse) GetInt32() (v int32) {
	return p.Int32
}
func (p *EchoResponse) SetInt32(val int32) {
	p.Int32 = val
}

func (p *EchoResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoResponse(%+v)", *p)
}

type EchoMultiBoolResponse struct {
	BaseResp bool          `thrift:"baseResp,1,required" frugal:"1,required,bool" json:"baseResp"`
	ListResp []bool        `thrift:"listResp,2,required" frugal:"2,required,list<bool>" json:"listResp"`
	MapResp  map[bool]bool `thrift:"mapResp,3,required" frugal:"3,required,map<bool:bool>" json:"mapResp"`
}

func NewEchoMultiBoolResponse() *EchoMultiBoolResponse {
	return &EchoMultiBoolResponse{}
}

func (p *EchoMultiBoolResponse) GetBaseResp() (v bool) {
	return p.BaseResp
}

func (p *EchoMultiBoolResponse) GetListResp() (v []bool) {
	return p.ListResp
}

func (p *EchoMultiBoolResponse) GetMapResp() (v map[bool]bool) {
	return p.MapResp
}
func (p *EchoMultiBoolResponse) SetBaseResp(val bool) {
	p.BaseResp = val
}
func (p *EchoMultiBoolResponse) SetListResp(val []bool) {
	p.ListResp = val
}
func (p *EchoMultiBoolResponse) SetMapResp(val map[bool]bool) {
	p.MapResp = val
}

func (p *EchoMultiBoolResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoMultiBoolResponse(%+v)", *p)
}

type EchoMultiByteResponse struct {
	BaseResp int8          `thrift:"baseResp,1,required" frugal:"1,required,byte" json:"baseResp"`
	ListResp []int8        `thrift:"listResp,2,required" frugal:"2,required,list<byte>" json:"listResp"`
	MapResp  map[int8]int8 `thrift:"mapResp,3,required" frugal:"3,required,map<byte:byte>" json:"mapResp"`
}

func NewEchoMultiByteResponse() *EchoMultiByteResponse {
	return &EchoMultiByteResponse{}
}

func (p *EchoMultiByteResponse) GetBaseResp() (v int8) {
	return p.BaseResp
}

func (p *EchoMultiByteResponse) GetListResp() (v []int8) {
	return p.ListResp
}

func (p *EchoMultiByteResponse) GetMapResp() (v map[int8]int8) {
	return p.MapResp
}
func (p *EchoMultiByteResponse) SetBaseResp(val int8) {
	p.BaseResp = val
}
func (p *EchoMultiByteResponse) SetListResp(val []int8) {
	p.ListResp = val
}
func (p *EchoMultiByteResponse) SetMapResp(val map[int8]int8) {
	p.MapResp = val
}

func (p *EchoMultiByteResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoMultiByteResponse(%+v)", *p)
}

type EchoMultiInt16Response struct {
	BaseResp int16           `thrift:"baseResp,1,required" frugal:"1,required,i16" json:"baseResp"`
	ListResp []int16         `thrift:"listResp,2,required" frugal:"2,required,list<i16>" json:"listResp"`
	MapResp  map[int16]int16 `thrift:"mapResp,3,required" frugal:"3,required,map<i16:i16>" json:"mapResp"`
}

func NewEchoMultiInt16Response() *EchoMultiInt16Response {
	return &EchoMultiInt16Response{}
}

func (p *EchoMultiInt16Response) GetBaseResp() (v int16) {
	return p.BaseResp
}

func (p *EchoMultiInt16Response) GetListResp() (v []int16) {
	return p.ListResp
}

func (p *EchoMultiInt16Response) GetMapResp() (v map[int16]int16) {
	return p.MapResp
}
func (p *EchoMultiInt16Response) SetBaseResp(val int16) {
	p.BaseResp = val
}
func (p *EchoMultiInt16Response) SetListResp(val []int16) {
	p.ListResp = val
}
func (p *EchoMultiInt16Response) SetMapResp(val map[int16]int16) {
	p.MapResp = val
}

func (p *EchoMultiInt16Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoMultiInt16Response(%+v)", *p)
}

type EchoMultiInt32Response struct {
	BaseResp int32           `thrift:"baseResp,1,required" frugal:"1,required,i32" json:"baseResp"`
	ListResp []int32         `thrift:"listResp,2,required" frugal:"2,required,list<i32>" json:"listResp"`
	MapResp  map[int32]int32 `thrift:"mapResp,3,required" frugal:"3,required,map<i32:i32>" json:"mapResp"`
}

func NewEchoMultiInt32Response() *EchoMultiInt32Response {
	return &EchoMultiInt32Response{}
}

func (p *EchoMultiInt32Response) GetBaseResp() (v int32) {
	return p.BaseResp
}

func (p *EchoMultiInt32Response) GetListResp() (v []int32) {
	return p.ListResp
}

func (p *EchoMultiInt32Response) GetMapResp() (v map[int32]int32) {
	return p.MapResp
}
func (p *EchoMultiInt32Response) SetBaseResp(val int32) {
	p.BaseResp = val
}
func (p *EchoMultiInt32Response) SetListResp(val []int32) {
	p.ListResp = val
}
func (p *EchoMultiInt32Response) SetMapResp(val map[int32]int32) {
	p.MapResp = val
}

func (p *EchoMultiInt32Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoMultiInt32Response(%+v)", *p)
}

type EchoMultiInt64Response struct {
	BaseResp int64           `thrift:"baseResp,1,required" frugal:"1,required,i64" json:"baseResp"`
	ListResp []int64         `thrift:"listResp,2,required" frugal:"2,required,list<i64>" json:"listResp"`
	MapResp  map[int64]int64 `thrift:"mapResp,3,required" frugal:"3,required,map<i64:i64>" json:"mapResp"`
}

func NewEchoMultiInt64Response() *EchoMultiInt64Response {
	return &EchoMultiInt64Response{}
}

func (p *EchoMultiInt64Response) GetBaseResp() (v int64) {
	return p.BaseResp
}

func (p *EchoMultiInt64Response) GetListResp() (v []int64) {
	return p.ListResp
}

func (p *EchoMultiInt64Response) GetMapResp() (v map[int64]int64) {
	return p.MapResp
}
func (p *EchoMultiInt64Response) SetBaseResp(val int64) {
	p.BaseResp = val
}
func (p *EchoMultiInt64Response) SetListResp(val []int64) {
	p.ListResp = val
}
func (p *EchoMultiInt64Response) SetMapResp(val map[int64]int64) {
	p.MapResp = val
}

func (p *EchoMultiInt64Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoMultiInt64Response(%+v)", *p)
}

type EchoMultiDoubleResponse struct {
	BaseResp float64             `thrift:"baseResp,1,required" frugal:"1,required,double" json:"baseResp"`
	ListResp []float64           `thrift:"listResp,2,required" frugal:"2,required,list<double>" json:"listResp"`
	MapResp  map[float64]float64 `thrift:"mapResp,3,required" frugal:"3,required,map<double:double>" json:"mapResp"`
}

func NewEchoMultiDoubleResponse() *EchoMultiDoubleResponse {
	return &EchoMultiDoubleResponse{}
}

func (p *EchoMultiDoubleResponse) GetBaseResp() (v float64) {
	return p.BaseResp
}

func (p *EchoMultiDoubleResponse) GetListResp() (v []float64) {
	return p.ListResp
}

func (p *EchoMultiDoubleResponse) GetMapResp() (v map[float64]float64) {
	return p.MapResp
}
func (p *EchoMultiDoubleResponse) SetBaseResp(val float64) {
	p.BaseResp = val
}
func (p *EchoMultiDoubleResponse) SetListResp(val []float64) {
	p.ListResp = val
}
func (p *EchoMultiDoubleResponse) SetMapResp(val map[float64]float64) {
	p.MapResp = val
}

func (p *EchoMultiDoubleResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoMultiDoubleResponse(%+v)", *p)
}

type EchoMultiStringResponse struct {
	BaseResp string            `thrift:"baseResp,1,required" frugal:"1,required,string" json:"baseResp"`
	ListResp []string          `thrift:"listResp,2,required" frugal:"2,required,list<string>" json:"listResp"`
	MapResp  map[string]string `thrift:"mapResp,3,required" frugal:"3,required,map<string:string>" json:"mapResp"`
}

func NewEchoMultiStringResponse() *EchoMultiStringResponse {
	return &EchoMultiStringResponse{}
}

func (p *EchoMultiStringResponse) GetBaseResp() (v string) {
	return p.BaseResp
}

func (p *EchoMultiStringResponse) GetListResp() (v []string) {
	return p.ListResp
}

func (p *EchoMultiStringResponse) GetMapResp() (v map[string]string) {
	return p.MapResp
}
func (p *EchoMultiStringResponse) SetBaseResp(val string) {
	p.BaseResp = val
}
func (p *EchoMultiStringResponse) SetListResp(val []string) {
	p.ListResp = val
}
func (p *EchoMultiStringResponse) SetMapResp(val map[string]string) {
	p.MapResp = val
}

func (p *EchoMultiStringResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoMultiStringResponse(%+v)", *p)
}

type TestService interface {
	EchoInt(ctx context.Context, req int32) (r int32, err error)

	EchoBool(ctx context.Context, req bool) (r bool, err error)

	EchoByte(ctx context.Context, req int8) (r int8, err error)

	EchoInt16(ctx context.Context, req int16) (r int16, err error)

	EchoInt32(ctx context.Context, req int32) (r int32, err error)

	EchoInt64(ctx context.Context, req int64) (r int64, err error)

	EchoDouble(ctx context.Context, req float64) (r float64, err error)

	EchoString(ctx context.Context, req string) (r string, err error)

	EchoBinary(ctx context.Context, req []byte) (r []byte, err error)

	Echo(ctx context.Context, req *EchoRequest) (r *EchoResponse, err error)

	EchoBoolList(ctx context.Context, req []bool) (r []bool, err error)

	EchoByteList(ctx context.Context, req []int8) (r []int8, err error)

	EchoInt16List(ctx context.Context, req []int16) (r []int16, err error)

	EchoInt32List(ctx context.Context, req []int32) (r []int32, err error)

	EchoInt64List(ctx context.Context, req []int64) (r []int64, err error)

	EchoDoubleList(ctx context.Context, req []float64) (r []float64, err error)

	EchoStringList(ctx context.Context, req []string) (r []string, err error)

	EchoBinaryList(ctx context.Context, req [][]byte) (r [][]byte, err error)

	EchoBool2BoolMap(ctx context.Context, req map[bool]bool) (r map[bool]bool, err error)

	EchoBool2ByteMap(ctx context.Context, req map[bool]int8) (r map[bool]int8, err error)

	EchoBool2Int16Map(ctx context.Context, req map[bool]int16) (r map[bool]int16, err error)

	EchoBool2Int32Map(ctx context.Context, req map[bool]int32) (r map[bool]int32, err error)

	EchoBool2Int64Map(ctx context.Context, req map[bool]int64) (r map[bool]int64, err error)

	EchoBool2DoubleMap(ctx context.Context, req map[bool]float64) (r map[bool]float64, err error)

	EchoBool2StringMap(ctx context.Context, req map[bool]string) (r map[bool]string, err error)

	EchoBool2BinaryMap(ctx context.Context, req map[bool][]byte) (r map[bool][]byte, err error)

	EchoBool2BoolListMap(ctx context.Context, req map[bool][]bool) (r map[bool][]bool, err error)

	EchoBool2ByteListMap(ctx context.Context, req map[bool][]int8) (r map[bool][]int8, err error)

	EchoBool2Int16ListMap(ctx context.Context, req map[bool][]int16) (r map[bool][]int16, err error)

	EchoBool2Int32ListMap(ctx context.Context, req map[bool][]int32) (r map[bool][]int32, err error)

	EchoBool2Int64ListMap(ctx context.Context, req map[bool][]int64) (r map[bool][]int64, err error)

	EchoBool2DoubleListMap(ctx context.Context, req map[bool][]float64) (r map[bool][]float64, err error)

	EchoBool2StringListMap(ctx context.Context, req map[bool][]string) (r map[bool][]string, err error)

	EchoBool2BinaryListMap(ctx context.Context, req map[bool][][]byte) (r map[bool][][]byte, err error)

	EchoMultiBool(ctx context.Context, baseReq bool, listReq []bool, mapReq map[bool]bool) (r *EchoMultiBoolResponse, err error)

	EchoMultiByte(ctx context.Context, baseReq int8, listReq []int8, mapReq map[int8]int8) (r *EchoMultiByteResponse, err error)

	EchoMultiInt16(ctx context.Context, baseReq int16, listReq []int16, mapReq map[int16]int16) (r *EchoMultiInt16Response, err error)

	EchoMultiInt32(ctx context.Context, baseReq int32, listReq []int32, mapReq map[int32]int32) (r *EchoMultiInt32Response, err error)

	EchoMultiInt64(ctx context.Context, baseReq int64, listReq []int64, mapReq map[int64]int64) (r *EchoMultiInt64Response, err error)

	EchoMultiDouble(ctx context.Context, baseReq float64, listReq []float64, mapReq map[float64]float64) (r *EchoMultiDoubleResponse, err error)

	EchoMultiString(ctx context.Context, baseReq string, listReq []string, mapReq map[string]string) (r *EchoMultiStringResponse, err error)
}
