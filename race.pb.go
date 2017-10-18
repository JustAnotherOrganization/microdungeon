// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: race.proto

package microdungeon

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Race is a playable D&D race.
type Race struct {
	Name           string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Genus          string   `protobuf:"bytes,2,opt,name=genus,proto3" json:"genus,omitempty"`
	Description    string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Score          *Score   `protobuf:"bytes,4,opt,name=score" json:"score,omitempty"`
	NightVision    bool     `protobuf:"varint,5,opt,name=night_vision,json=nightVision,proto3" json:"night_vision,omitempty"`
	KnownLanguages []string `protobuf:"bytes,6,rep,name=known_languages,json=knownLanguages" json:"known_languages,omitempty"`
	Undead         bool     `protobuf:"varint,7,opt,name=undead,proto3" json:"undead,omitempty"`
}

func (m *Race) Reset()                    { *m = Race{} }
func (m *Race) String() string            { return proto.CompactTextString(m) }
func (*Race) ProtoMessage()               {}
func (*Race) Descriptor() ([]byte, []int) { return fileDescriptorRace, []int{0} }

func (m *Race) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Race) GetGenus() string {
	if m != nil {
		return m.Genus
	}
	return ""
}

func (m *Race) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Race) GetScore() *Score {
	if m != nil {
		return m.Score
	}
	return nil
}

func (m *Race) GetNightVision() bool {
	if m != nil {
		return m.NightVision
	}
	return false
}

func (m *Race) GetKnownLanguages() []string {
	if m != nil {
		return m.KnownLanguages
	}
	return nil
}

func (m *Race) GetUndead() bool {
	if m != nil {
		return m.Undead
	}
	return false
}

func init() {
	proto.RegisterType((*Race)(nil), "microdungeon.Race")
}

func init() { proto.RegisterFile("race.proto", fileDescriptorRace) }

var fileDescriptorRace = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0x9a, 0x04, 0x7a, 0xa9, 0x8a, 0x64, 0x2a, 0x64, 0x31, 0x19, 0x16, 0xc2, 0x92,
	0x01, 0xfe, 0x01, 0x33, 0x93, 0x91, 0x18, 0x58, 0x22, 0xd7, 0x39, 0x19, 0x8b, 0xf6, 0x5c, 0xd9,
	0x09, 0x88, 0x9f, 0xcb, 0x3f, 0x41, 0xbd, 0x80, 0x14, 0x36, 0xbf, 0xef, 0x3d, 0x3d, 0xdf, 0x1d,
	0x40, 0xb2, 0x0e, 0xdb, 0x43, 0x8a, 0x43, 0x94, 0xab, 0x7d, 0x70, 0x29, 0xf6, 0x23, 0x79, 0x8c,
	0x74, 0xb5, 0xb1, 0xdb, 0xb0, 0x0b, 0xc3, 0x57, 0x97, 0x5d, 0x4c, 0x98, 0xa7, 0xcc, 0xcd, 0xb7,
	0x80, 0xc2, 0x58, 0x87, 0x52, 0x42, 0x41, 0x76, 0x8f, 0x4a, 0x68, 0xd1, 0x2c, 0x0d, 0xbf, 0xe5,
	0x06, 0x4a, 0x8f, 0x34, 0x66, 0x75, 0xc2, 0x70, 0x12, 0x52, 0x43, 0xdd, 0x63, 0x76, 0x29, 0x1c,
	0x86, 0x10, 0x49, 0x2d, 0xd8, 0x9b, 0x23, 0x79, 0x07, 0x25, 0x7f, 0xa2, 0x0a, 0x2d, 0x9a, 0xfa,
	0xfe, 0xa2, 0x9d, 0x0f, 0xd2, 0x3e, 0x1f, 0x2d, 0x33, 0x25, 0xe4, 0x35, 0xac, 0x28, 0xf8, 0xb7,
	0xa1, 0xfb, 0x08, 0xf9, 0xd8, 0x56, 0x6a, 0xd1, 0x9c, 0x99, 0x9a, 0xd9, 0x0b, 0x23, 0x79, 0x0b,
	0xe7, 0xef, 0x14, 0x3f, 0xa9, 0xdb, 0x59, 0xf2, 0xa3, 0xf5, 0x98, 0x55, 0xa5, 0x17, 0xcd, 0xd2,
	0xac, 0x19, 0x3f, 0xfd, 0x51, 0x79, 0x09, 0xd5, 0x48, 0x3d, 0xda, 0x5e, 0x9d, 0x72, 0xcb, 0xaf,
	0x7a, 0x5c, 0xbf, 0xfe, 0xbb, 0xc4, 0xb6, 0xe2, 0xd5, 0x1f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0x8f, 0x46, 0x24, 0x2c, 0x01, 0x00, 0x00,
}
