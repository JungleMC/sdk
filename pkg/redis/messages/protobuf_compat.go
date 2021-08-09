package messages

import "google.golang.org/protobuf/proto"

func (x *JavaLoginBegin) MarshalBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *JavaLoginBegin) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, x)
}

func (x *JavaLoginBegin) UnmarshalBinaryString(data string) error {
	return proto.Unmarshal([]byte(data), x)
}

func (x *Player) MarshalBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *Player) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, x)
}

func (x *Player) UnmarshalBinaryString(data string) error {
	return proto.Unmarshal([]byte(data), x)
}
