package messages

import "google.golang.org/protobuf/proto"

func (x *LoginBegin) MarshalBinary() ([]byte, error) {
	return proto.Marshal(x)
}

func (x *LoginBegin) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, x)
}
