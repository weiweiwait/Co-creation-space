package domain

import (
	"context"
	"my_project/project-grpc/user/login"
	"my_project/project-project/internal/rpc"
)

type UserRpcDomain struct {
	lc login.LoginServiceClient
}

func NewUserRpcDomain() *UserRpcDomain {
	return &UserRpcDomain{
		lc: rpc.LoginServiceClient,
	}
}
func (d *UserRpcDomain) MemberList(c context.Context, mIdList []int64) ([]*login.MemberMessage, map[int64]*login.MemberMessage, error) {
	messageList, err := d.lc.FindMemInfoByIds(c, &login.UserMessage{MIds: mIdList})
	mMap := make(map[int64]*login.MemberMessage)
	for _, v := range messageList.List {
		mMap[v.Id] = v
	}
	return messageList.List, mMap, err
}
