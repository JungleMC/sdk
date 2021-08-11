package messages

import (
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
)

func UnmarshalMessage(in *redis.Message, out proto.Message) error {
	err := proto.Unmarshal([]byte(in.Payload), out)
	if err != nil {
		return err
	}
	return nil
}
