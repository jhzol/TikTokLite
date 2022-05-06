package ProtoMessage

import (
	"TikTokLite/proto/pkg"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestProto(t *testing.T) {
	user := &message.User{
		Id:       new(int64),
		Name:     new(string),
		IsFollow: new(bool),
	}
	*user.Id = 123
	*user.Name = "someName"
	*user.IsFollow = false
	data, err := proto.Marshal(user)
	if err != nil {
		t.Errorf("Marshal error\n")
	}
	newUser := &message.User{}
	err = proto.Unmarshal(data, newUser)
	if err != nil {
		t.Errorf("Unmarshal error\n")
	}
	if user.GetId() != newUser.GetId() {
		t.Errorf("user:%+v,newUser:%+v\n", user, newUser)
	}
}
