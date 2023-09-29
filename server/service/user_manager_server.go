package service

import (
	"context"
	"foo/server/domain"
	"foo/server/proto"
)

var users = []*domain.User{
	{Id: 1, Name: "foo"},
	{Id: 2, Name: "bar"},
	{Id: 3, Name: "baz"},
}

type UserManagerServer struct{}

func (u *UserManagerServer) Fetch(_ context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	id := req.Id

	var found *domain.User
	// idが一致するユーザーを取得
	for _, user := range users {
		if user.Id == id {
			found = user
			break
		}
	}

	return &proto.UserResponse{
		Id:   found.Id,
		Name: found.Name,
	}, nil
}
