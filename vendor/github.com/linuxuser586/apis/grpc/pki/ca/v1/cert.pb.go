// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/pki/ca/v1/cert.proto

package v1 // import "github.com/linuxuser586/apis/grpc/pki/ca/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CertRequest is currently empty
type CertRequest struct {
}

func (m *CertRequest) Reset()         { *m = CertRequest{} }
func (m *CertRequest) String() string { return proto.CompactTextString(m) }
func (*CertRequest) ProtoMessage()    {}
func (*CertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cert_02e2a7417f02f541, []int{0}
}
func (m *CertRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertRequest.Merge(dst, src)
}
func (m *CertRequest) XXX_Size() int {
	return m.Size()
}
func (m *CertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertRequest proto.InternalMessageInfo

// CertResponse is the CertService response
type CertResponse struct {
	Cert string `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (m *CertResponse) Reset()         { *m = CertResponse{} }
func (m *CertResponse) String() string { return proto.CompactTextString(m) }
func (*CertResponse) ProtoMessage()    {}
func (*CertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cert_02e2a7417f02f541, []int{1}
}
func (m *CertResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertResponse.Merge(dst, src)
}
func (m *CertResponse) XXX_Size() int {
	return m.Size()
}
func (m *CertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertResponse proto.InternalMessageInfo

func (m *CertResponse) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func init() {
	proto.RegisterType((*CertRequest)(nil), "linuxuser586.pki.ca.v1.CertRequest")
	proto.RegisterType((*CertResponse)(nil), "linuxuser586.pki.ca.v1.CertResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertServiceClient is the client API for CertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertServiceClient interface {
	// Get the CA certificate
	Get(ctx context.Context, in *CertRequest, opts ...grpc.CallOption) (*CertResponse, error)
}

type certServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertServiceClient(cc *grpc.ClientConn) CertServiceClient {
	return &certServiceClient{cc}
}

func (c *certServiceClient) Get(ctx context.Context, in *CertRequest, opts ...grpc.CallOption) (*CertResponse, error) {
	out := new(CertResponse)
	err := c.cc.Invoke(ctx, "/linuxuser586.pki.ca.v1.CertService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertServiceServer is the server API for CertService service.
type CertServiceServer interface {
	// Get the CA certificate
	Get(context.Context, *CertRequest) (*CertResponse, error)
}

func RegisterCertServiceServer(s *grpc.Server, srv CertServiceServer) {
	s.RegisterService(&_CertService_serviceDesc, srv)
}

func _CertService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linuxuser586.pki.ca.v1.CertService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServiceServer).Get(ctx, req.(*CertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "linuxuser586.pki.ca.v1.CertService",
	HandlerType: (*CertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CertService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/pki/ca/v1/cert.proto",
}

func (m *CertRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *CertResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cert) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCert(dAtA, i, uint64(len(m.Cert)))
		i += copy(dAtA[i:], m.Cert)
	}
	return i, nil
}

func encodeVarintCert(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CertRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *CertResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cert)
	if l > 0 {
		n += 1 + l + sovCert(uint64(l))
	}
	return n
}

func sovCert(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCert(x uint64) (n int) {
	return sovCert(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CertRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCert
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCert(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCert
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CertResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCert
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCert
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCert
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCert(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCert
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCert(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCert
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCert
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCert
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCert
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCert
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCert(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCert = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCert   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protos/pki/ca/v1/cert.proto", fileDescriptor_cert_02e2a7417f02f541) }

var fileDescriptor_cert_02e2a7417f02f541 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0xc8, 0xce, 0xd4, 0x4f, 0x4e, 0xd4, 0x2f, 0x33, 0xd4, 0x4f, 0x4e, 0x2d,
	0x2a, 0xd1, 0x03, 0x8b, 0x0a, 0x89, 0xe5, 0x64, 0xe6, 0x95, 0x56, 0x94, 0x16, 0xa7, 0x16, 0x99,
	0x5a, 0x98, 0xe9, 0x15, 0x64, 0x67, 0xea, 0x25, 0x27, 0xea, 0x95, 0x19, 0x2a, 0xf1, 0x72, 0x71,
	0x3b, 0xa7, 0x16, 0x95, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x29, 0x71, 0xf1, 0x40,
	0xb8, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0x20, 0x43, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa3, 0x44, 0x88, 0x96, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4,
	0x54, 0xa1, 0x20, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x65, 0x3d, 0xec, 0x36, 0xe8, 0x21, 0x19,
	0x2f, 0xa5, 0x82, 0x5f, 0x11, 0xc4, 0x52, 0x25, 0x06, 0x27, 0xd7, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x47, 0x36, 0x4b, 0x3f, 0xb1, 0x20, 0xb3, 0x58, 0x3f, 0xbd, 0xa8, 0x20, 0x19, 0xe1, 0xff,
	0x24, 0x36, 0xb0, 0xdf, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x27, 0xa2, 0x9c, 0x1a,
	0x01, 0x00, 0x00,
}