package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	consts "site/protocol/shared"
	pb "site/protocol/user"
	"strconv"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

// InitUserServiceClient 初始化用户服务客户端
func InitUserServiceClient(url string) UserServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could't connect: %v", err)
	}
	return UserServiceClient{
		Client: pb.NewUserServiceClient(cc),
	}
}

// FindOneUserByUid 查询用户信息
func (c *UserServiceClient) FindOneUserByUid(uid int64) (*pb.FindOneUserResp, error) {
	req := &pb.FindOneUserReq{
		Account:     strconv.FormatInt(uid, 10),
		AccountType: consts.AccountTypeUid,
	}
	return c.Client.FindOneUser(context.Background(), req)
}

// UpdateUserFriendNum 添加或删除好友
func (c *UserServiceClient) UpdateUserFriendNum(uid1 int64, uid2 int64, delFlag bool) (*pb.UpdateUserFriendNumResp, error) {
	req := &pb.UpdateUserFriendNumReq{
		UserId:   uid1,
		FriendId: uid2,
		DelFlag:  delFlag,
	}
	return c.Client.UpdateUserFriendNum(context.Background(), req)
}

// UpdateUserRelation 修改用户关系信息 opt 为操作选项 add 标记是增加还是减少
// TODO: 删除好友或退出群聊之前先判断是否置顶，拉黑之前在前面判断的基础上判断是否是好友或是否在群聊中
func (c *UserServiceClient) UpdateUserRelation(uid int64, opt int, add bool) (*pb.UpdateUserRelationResp, error) {
	req := &pb.UpdateUserRelationReq{
		Uid: uid,
	}
	var num int32
	if add {
		num = 1
	} else {
		num = -1
	}
	switch opt {
	case consts.AddFriendOpt:
		req.Data.FriendNum = num
	case consts.TopFriendOpt:
		req.Data.TopGroupNum = num
	case consts.BlackFriendOpt:
		req.Data.BlackGroupNum = num
	case consts.CreateGroupOpt:
		req.Data.GroupNum = num
		req.Data.OwnerGroupNum = num
		req.Data.AdminGroupNum = num
	case consts.AdminGroupOpt:
		req.Data.AdminGroupNum = num
	case consts.TopGroupOpt:
		req.Data.TopGroupNum = num
	case consts.BlackGroupOpt:
		req.Data.BlackGroupNum = num
	case consts.JoinGroupOpt:
		req.Data.GroupNum = num
	}
	return c.Client.UpdateUserRelation(context.Background(), req)
}
