package utils

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Interface2Anypb -
func Interface2Anypb(v interface{}) (*anypb.Any, error) {
	var (
		wrapper proto.Message
		any     *anypb.Any
		err     error
	)

	switch v.(type) {
	case string:
		wrapper = wrapperspb.String(v.(string))
	case []byte:
		wrapper = wrapperspb.Bytes(v.([]byte))
	case int, int32:
		wrapper = wrapperspb.Int32(v.(int32))
	case int64:
		wrapper = wrapperspb.Int64(v.(int64))
	case bool:
		wrapper = wrapperspb.Bool(v.(bool))
	}
	any, err = anypb.New(wrapper)
	if err != nil {
		return nil, err
	}

	return any, err
}
