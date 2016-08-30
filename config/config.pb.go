// Code generated by protoc-gen-gogo.
// source: cockroach/config/config.proto
// DO NOT EDIT!

/*
	Package config is a generated protocol buffer package.

	It is generated from these files:
		cockroach/config/config.proto

	It has these top-level messages:
		GCPolicy
		Constraint
		Constraints
		ZoneConfig
		SystemConfig
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/roachpb"
import cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

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

type Constraint_Type int32

const (
	// POSITIVE will attempt to ensure all stores the replicas are on has this
	// constraint.
	Constraint_POSITIVE Constraint_Type = 0
	// REQUIRED is like POSITIVE except replication will fail if not satisfied.
	Constraint_REQUIRED Constraint_Type = 1
	// PROHIBITED will prevent replicas from having this key, value.
	Constraint_PROHIBITED Constraint_Type = 2
)

var Constraint_Type_name = map[int32]string{
	0: "POSITIVE",
	1: "REQUIRED",
	2: "PROHIBITED",
}
var Constraint_Type_value = map[string]int32{
	"POSITIVE":   0,
	"REQUIRED":   1,
	"PROHIBITED": 2,
}

func (x Constraint_Type) Enum() *Constraint_Type {
	p := new(Constraint_Type)
	*p = x
	return p
}
func (x Constraint_Type) String() string {
	return proto.EnumName(Constraint_Type_name, int32(x))
}
func (x *Constraint_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Constraint_Type_value, data, "Constraint_Type")
	if err != nil {
		return err
	}
	*x = Constraint_Type(value)
	return nil
}
func (Constraint_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1, 0} }

// GCPolicy defines garbage collection policies which apply to MVCC
// values within a zone.
//
// TODO(spencer): flesh this out to include maximum number of values
//   as well as whether there's an intersection between max values
//   and TTL or a union.
type GCPolicy struct {
	// TTLSeconds specifies the maximum age of a value before it's
	// garbage collected. Only older versions of values are garbage
	// collected. Specifying <=0 mean older versions are never GC'd.
	TTLSeconds int32 `protobuf:"varint,1,opt,name=ttl_seconds,json=ttlSeconds" json:"ttl_seconds"`
}

func (m *GCPolicy) Reset()                    { *m = GCPolicy{} }
func (m *GCPolicy) String() string            { return proto.CompactTextString(m) }
func (*GCPolicy) ProtoMessage()               {}
func (*GCPolicy) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

// Constraint constrains the stores a replica can be stored on.
type Constraint struct {
	Type Constraint_Type `protobuf:"varint,1,opt,name=type,enum=cockroach.config.Constraint_Type" json:"type"`
	// Key is only set if this is a constraint on locality.
	Key string `protobuf:"bytes,2,opt,name=key" json:"key"`
	// Value to constrain to.
	Value string `protobuf:"bytes,3,opt,name=value" json:"value"`
}

func (m *Constraint) Reset()                    { *m = Constraint{} }
func (*Constraint) ProtoMessage()               {}
func (*Constraint) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

// Constraints is a collection of constraints.
type Constraints struct {
	Constraints []Constraint `protobuf:"bytes,6,rep,name=constraints" json:"constraints"`
}

func (m *Constraints) Reset()                    { *m = Constraints{} }
func (m *Constraints) String() string            { return proto.CompactTextString(m) }
func (*Constraints) ProtoMessage()               {}
func (*Constraints) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

// ZoneConfig holds configuration that is needed for a range of KV pairs. This
// and the conversion methods must stay in sync with ZoneConfigHuman.
type ZoneConfig struct {
	// TODO(d4l3k): Remove replica_attrs after a sufficient amount of time has passed.
	ReplicaAttrs  []cockroach_roachpb.Attributes `protobuf:"bytes,1,rep,name=replica_attrs,json=replicaAttrs" json:"replica_attrs" yaml:",omitempty"`
	RangeMinBytes int64                          `protobuf:"varint,2,opt,name=range_min_bytes,json=rangeMinBytes" json:"range_min_bytes" yaml:"range_min_bytes"`
	RangeMaxBytes int64                          `protobuf:"varint,3,opt,name=range_max_bytes,json=rangeMaxBytes" json:"range_max_bytes" yaml:"range_max_bytes"`
	// If GC policy is not set, uses the next highest, non-null policy
	// in the zone config hierarchy, up to the default policy if necessary.
	GC GCPolicy `protobuf:"bytes,4,opt,name=gc" json:"gc"`
	// NumReplicas specifies the desired number of replicas
	NumReplicas int32 `protobuf:"varint,5,opt,name=num_replicas,json=numReplicas" json:"num_replicas" yaml:"num_replicas"`
	// Constraints constrains which stores the replicas can be stored on. The
	// order in which the constraints are stored is arbitrary and may change.
	// https://github.com/cockroachdb/cockroach/blob/master/docs/RFCS/expressive_zone_config.md#constraint-system
	Constraints Constraints `protobuf:"bytes,6,opt,name=constraints" json:"constraints" yaml:"constraints,flow"`
}

func (m *ZoneConfig) Reset()                    { *m = ZoneConfig{} }
func (m *ZoneConfig) String() string            { return proto.CompactTextString(m) }
func (*ZoneConfig) ProtoMessage()               {}
func (*ZoneConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

type SystemConfig struct {
	Values []cockroach_roachpb1.KeyValue `protobuf:"bytes,1,rep,name=values" json:"values"`
}

func (m *SystemConfig) Reset()                    { *m = SystemConfig{} }
func (m *SystemConfig) String() string            { return proto.CompactTextString(m) }
func (*SystemConfig) ProtoMessage()               {}
func (*SystemConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func init() {
	proto.RegisterType((*GCPolicy)(nil), "cockroach.config.GCPolicy")
	proto.RegisterType((*Constraint)(nil), "cockroach.config.Constraint")
	proto.RegisterType((*Constraints)(nil), "cockroach.config.Constraints")
	proto.RegisterType((*ZoneConfig)(nil), "cockroach.config.ZoneConfig")
	proto.RegisterType((*SystemConfig)(nil), "cockroach.config.SystemConfig")
	proto.RegisterEnum("cockroach.config.Constraint_Type", Constraint_Type_name, Constraint_Type_value)
}
func (m *GCPolicy) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GCPolicy) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintConfig(data, i, uint64(m.TTLSeconds))
	return i, nil
}

func (m *Constraint) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Constraint) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintConfig(data, i, uint64(m.Type))
	data[i] = 0x12
	i++
	i = encodeVarintConfig(data, i, uint64(len(m.Key)))
	i += copy(data[i:], m.Key)
	data[i] = 0x1a
	i++
	i = encodeVarintConfig(data, i, uint64(len(m.Value)))
	i += copy(data[i:], m.Value)
	return i, nil
}

func (m *Constraints) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Constraints) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Constraints) > 0 {
		for _, msg := range m.Constraints {
			data[i] = 0x32
			i++
			i = encodeVarintConfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ZoneConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ZoneConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ReplicaAttrs) > 0 {
		for _, msg := range m.ReplicaAttrs {
			data[i] = 0xa
			i++
			i = encodeVarintConfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x10
	i++
	i = encodeVarintConfig(data, i, uint64(m.RangeMinBytes))
	data[i] = 0x18
	i++
	i = encodeVarintConfig(data, i, uint64(m.RangeMaxBytes))
	data[i] = 0x22
	i++
	i = encodeVarintConfig(data, i, uint64(m.GC.Size()))
	n1, err := m.GC.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x28
	i++
	i = encodeVarintConfig(data, i, uint64(m.NumReplicas))
	data[i] = 0x32
	i++
	i = encodeVarintConfig(data, i, uint64(m.Constraints.Size()))
	n2, err := m.Constraints.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *SystemConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SystemConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			data[i] = 0xa
			i++
			i = encodeVarintConfig(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Config(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Config(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConfig(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *GCPolicy) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovConfig(uint64(m.TTLSeconds))
	return n
}

func (m *Constraint) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovConfig(uint64(m.Type))
	l = len(m.Key)
	n += 1 + l + sovConfig(uint64(l))
	l = len(m.Value)
	n += 1 + l + sovConfig(uint64(l))
	return n
}

func (m *Constraints) Size() (n int) {
	var l int
	_ = l
	if len(m.Constraints) > 0 {
		for _, e := range m.Constraints {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *ZoneConfig) Size() (n int) {
	var l int
	_ = l
	if len(m.ReplicaAttrs) > 0 {
		for _, e := range m.ReplicaAttrs {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	n += 1 + sovConfig(uint64(m.RangeMinBytes))
	n += 1 + sovConfig(uint64(m.RangeMaxBytes))
	l = m.GC.Size()
	n += 1 + l + sovConfig(uint64(l))
	n += 1 + sovConfig(uint64(m.NumReplicas))
	l = m.Constraints.Size()
	n += 1 + l + sovConfig(uint64(l))
	return n
}

func (m *SystemConfig) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GCPolicy) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GCPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GCPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTLSeconds", wireType)
			}
			m.TTLSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TTLSeconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func (m *Constraint) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Constraint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Constraint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Type |= (Constraint_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func (m *Constraints) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Constraints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Constraints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constraints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Constraints = append(m.Constraints, Constraint{})
			if err := m.Constraints[len(m.Constraints)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func (m *ZoneConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ZoneConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ZoneConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaAttrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReplicaAttrs = append(m.ReplicaAttrs, cockroach_roachpb.Attributes{})
			if err := m.ReplicaAttrs[len(m.ReplicaAttrs)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeMinBytes", wireType)
			}
			m.RangeMinBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeMinBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeMaxBytes", wireType)
			}
			m.RangeMaxBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeMaxBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GC", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GC.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumReplicas", wireType)
			}
			m.NumReplicas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NumReplicas |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constraints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Constraints.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func (m *SystemConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SystemConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SystemConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, cockroach_roachpb1.KeyValue{})
			if err := m.Values[len(m.Values)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipConfig(data[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x71, 0x5a, 0x95, 0x71, 0x5a, 0xc2, 0x82, 0x8a, 0x95, 0xb6, 0x8e, 0xf1, 0x29,
	0x87, 0xca, 0x95, 0x82, 0x84, 0x44, 0x91, 0x40, 0x38, 0x0d, 0xc5, 0x02, 0xd4, 0xe2, 0x84, 0x0a,
	0xf5, 0x62, 0xb6, 0xee, 0xd6, 0x58, 0xb5, 0xbd, 0x96, 0xbd, 0x81, 0xfa, 0x2d, 0x38, 0x72, 0xe4,
	0x2d, 0x78, 0x85, 0x1e, 0xb9, 0xc1, 0xa9, 0x82, 0xf0, 0x06, 0x3c, 0x01, 0xb2, 0xbd, 0x69, 0x4c,
	0x8a, 0x7a, 0x49, 0x76, 0x66, 0xfe, 0xf9, 0xbc, 0x3b, 0xff, 0xc0, 0x86, 0xcb, 0xdc, 0xd3, 0x84,
	0x11, 0xf7, 0xfd, 0x96, 0xcb, 0xa2, 0x13, 0xdf, 0x13, 0x7f, 0x46, 0x9c, 0x30, 0xce, 0x70, 0xeb,
	0xb2, 0x6c, 0x94, 0xf9, 0xf6, 0xfa, 0xac, 0xa1, 0xf8, 0x8d, 0x8f, 0xb6, 0x8e, 0x09, 0x27, 0xa5,
	0xbe, 0xad, 0x5d, 0xad, 0x86, 0x94, 0x93, 0x8a, 0xe2, 0x8e, 0xc7, 0x3c, 0x56, 0x1c, 0xb7, 0xf2,
	0x53, 0x99, 0xd5, 0x9f, 0xc0, 0xd2, 0x6e, 0x7f, 0x9f, 0x05, 0xbe, 0x9b, 0xe1, 0xfb, 0x20, 0x73,
	0x1e, 0x38, 0x29, 0x75, 0x59, 0x74, 0x9c, 0x2a, 0x48, 0x43, 0xdd, 0x05, 0x13, 0x9f, 0x5f, 0x74,
	0x6a, 0x93, 0x8b, 0x0e, 0x8c, 0x46, 0x2f, 0x87, 0x65, 0xc5, 0x06, 0xce, 0x03, 0x71, 0xd6, 0xbf,
	0x22, 0x80, 0x3e, 0x8b, 0x52, 0x9e, 0x10, 0x3f, 0xe2, 0xf8, 0x11, 0x34, 0x78, 0x16, 0xd3, 0xa2,
	0x79, 0xa5, 0x77, 0xcf, 0x98, 0x7f, 0x86, 0x31, 0xd3, 0x1a, 0xa3, 0x2c, 0xa6, 0x66, 0x23, 0xe7,
	0xdb, 0x45, 0x13, 0x5e, 0x05, 0xe9, 0x94, 0x66, 0x4a, 0x5d, 0x43, 0xdd, 0x1b, 0xa2, 0x90, 0x27,
	0x70, 0x1b, 0x16, 0x3e, 0x90, 0x60, 0x4c, 0x15, 0xa9, 0x52, 0x29, 0x53, 0x7a, 0x0f, 0x1a, 0x39,
	0x07, 0x37, 0x61, 0x69, 0x7f, 0x6f, 0x68, 0x8d, 0xac, 0x83, 0x41, 0xab, 0x96, 0x47, 0xf6, 0xe0,
	0xf5, 0x1b, 0xcb, 0x1e, 0xec, 0xb4, 0x10, 0x5e, 0x01, 0xd8, 0xb7, 0xf7, 0x9e, 0x5b, 0xa6, 0x35,
	0x1a, 0xec, 0xb4, 0xea, 0xdb, 0x8d, 0xcf, 0x5f, 0x3a, 0x35, 0x7d, 0x08, 0xf2, 0xec, 0x32, 0x29,
	0xde, 0x01, 0xd9, 0x9d, 0x85, 0xca, 0xa2, 0x26, 0x75, 0xe5, 0xde, 0xfa, 0x75, 0x0f, 0x10, 0x17,
	0xa9, 0xb6, 0xe9, 0xdf, 0x25, 0x80, 0x43, 0x16, 0xd1, 0x7e, 0x21, 0xc6, 0x0e, 0x2c, 0x27, 0x34,
	0x0e, 0x7c, 0x97, 0x38, 0x84, 0xf3, 0x24, 0x1f, 0x6a, 0x8e, 0xdd, 0xa8, 0x60, 0x85, 0x5d, 0xc6,
	0x53, 0xce, 0x13, 0xff, 0x68, 0xcc, 0x69, 0x6a, 0xae, 0xe5, 0xdc, 0x3f, 0x17, 0x9d, 0x5b, 0x19,
	0x09, 0x83, 0x6d, 0x7d, 0x93, 0x85, 0x3e, 0xa7, 0x61, 0xcc, 0x33, 0x5d, 0x41, 0x76, 0x53, 0x00,
	0x73, 0x7d, 0x8a, 0x9f, 0xc1, 0xcd, 0x84, 0x44, 0x1e, 0x75, 0x42, 0x3f, 0x72, 0x8e, 0x32, 0x4e,
	0xd3, 0x62, 0x7c, 0x92, 0xa9, 0x0a, 0xc6, 0x6a, 0xc9, 0x98, 0x13, 0xe9, 0xf6, 0x72, 0x91, 0x79,
	0xe5, 0x47, 0x66, 0x1e, 0x57, 0x38, 0xe4, 0x4c, 0x70, 0xa4, 0x6b, 0x38, 0x53, 0xd1, 0x25, 0x87,
	0x9c, 0x95, 0x9c, 0x07, 0x50, 0xf7, 0x5c, 0xa5, 0xa1, 0xa1, 0xae, 0xdc, 0x6b, 0x5f, 0x1d, 0xde,
	0x74, 0xd7, 0x4c, 0x10, 0x6b, 0x55, 0xdf, 0xed, 0xdb, 0x75, 0xcf, 0xc5, 0x8f, 0xa1, 0x19, 0x8d,
	0x43, 0x47, 0xbc, 0x2d, 0x55, 0x16, 0x8a, 0xe5, 0x9b, 0x0e, 0xe2, 0x76, 0xf9, 0xf1, 0xaa, 0x42,
	0xb7, 0xe5, 0x68, 0x1c, 0xda, 0x22, 0xc2, 0xef, 0xe6, 0xdd, 0x43, 0x73, 0x63, 0xbe, 0xe2, 0x5e,
	0x6a, 0x76, 0x04, 0xfd, 0x6e, 0x49, 0xaf, 0xf4, 0x6f, 0x9e, 0x04, 0xec, 0xa3, 0xfe, 0xaf, 0xb3,
	0x16, 0x34, 0x87, 0x59, 0xca, 0x69, 0x28, 0xac, 0x7d, 0x08, 0x8b, 0xc5, 0x06, 0x4e, 0x3d, 0x5d,
	0xfb, 0x8f, 0xa7, 0x2f, 0x68, 0x76, 0x90, 0x6b, 0xc4, 0xa6, 0x88, 0x06, 0x53, 0x3b, 0xff, 0xa5,
	0xd6, 0xce, 0x27, 0x2a, 0xfa, 0x36, 0x51, 0xd1, 0x8f, 0x89, 0x8a, 0x7e, 0x4e, 0x54, 0xf4, 0xe9,
	0xb7, 0x5a, 0x3b, 0x5c, 0x2c, 0xaf, 0xf9, 0xb6, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x29, 0x0e,
	0x84, 0x3a, 0x20, 0x04, 0x00, 0x00,
}
