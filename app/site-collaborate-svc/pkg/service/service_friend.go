package service

import (
	"errors"
	"site/app/site-collaborate-svc/pkg/client"
	"site/app/site-collaborate-svc/pkg/dao"
	"site/app/site-collaborate-svc/pkg/models"
	"site/common/logs"
	"site/protocol/shared"
)

type FriendService struct {
	D       *dao.Dao
	UserSvc client.UserServiceClient
}

// InitFriendService 初始化 FriendService
func InitFriendService(d *dao.Dao, us client.UserServiceClient) FriendService {
	return FriendService{
		D:       d,
		UserSvc: us,
	}
}

// FindUserOneFriend 查询一个用户的某个好友
func (s *FriendService) FindUserOneFriend(userId int64, friendId int64) (*models.UserFriend, error) {
	f, err := s.D.FindFriendRecord(userId, friendId)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// AddFriend 添加好友
func (s *FriendService) AddFriend(userId int64, friendId int64, userLabel string, friendLabel string) error {
	// 无法添加自己
	if userId == friendId {
		return errors.New(shared.CodeMessageIgnoreCode(shared.AddYourselfError))
	}
	// 创建记录前先检查有无记录
	_, _, err := s.D.FindIsFriend(userId, friendId)
	if err == nil {
		return errors.New(shared.CodeMessageIgnoreCode(shared.ExistsFriendOrGroup))
	}
	userInfo, err := s.UserSvc.FindOneUserByUid(userId)
	if err != nil {
		return err
	}
	if userInfo.Msg.Error != "" || userInfo.Msg.Status != 200 {
		return errors.New(userInfo.Msg.Error)
	}
	friendInfo, err := s.UserSvc.FindOneUserByUid(friendId)
	if err != nil {
		return err
	}
	if friendInfo.Msg.Error != "" || friendInfo.Msg.Status != 200 {
		return errors.New(userInfo.Msg.Error)
	}
	// userLabel 是用户对好友的备注
	if userLabel == "" {
		userLabel = friendInfo.Data.Username
	}
	// friendLabel 是好友对用户的备注
	if friendLabel == "" {
		friendLabel = userInfo.Data.Username
	}
	err = s.D.CreateBecomeFriendById(userId, friendId, userLabel, friendLabel)
	if err != nil {
		return err
	}
	// 添加好友后修改用户关系信息
	resp, err := s.UserSvc.UpdateUserFriendNum(userId, friendId, false)
	// 处理失败进行补偿操作
	if err != nil || resp.Msg.Error != "" || resp.Msg.Status != 200 {
		err1 := s.D.DeleteFriendRecord(userId, friendId)
		if err1 != nil {
			logs.SugarLogger.Errorf("UserId[%v] FriendId[%v] Add friend record and compensate failed: err[%v], err1[%v], msg.error[%v]", userId, friendId, err, err1, resp.Msg.Error)
			return errors.New(shared.CodeMessageIgnoreCode(shared.AddError))
		}
		logs.SugarLogger.Warnf("UserId[%v] FriendId[%v] Add friend record failed: %v, and compensate success!", userId, friendId, err)
		return errors.New(shared.CodeMessageIgnoreCode(shared.AddError))
	}
	return nil
}

// DeleteFriend 删除好友
func (s *FriendService) DeleteFriend(userId int64, friendId int64) error {
	// 修改用户前服务用户关系信息
	resp, err := s.UserSvc.UpdateUserFriendNum(userId, friendId, true)
	// 处理失败时直接 return
	if err != nil || resp.Msg.Error != "" || resp.Msg.Status != 200 {
		return errors.New(shared.CodeMessageIgnoreCode(shared.DeleteError))
	}
	err = s.D.DeleteFriendRecord(userId, friendId)
	if err != nil {
		// 删除失败时对用户关系信息进行补偿操作
		resp2, err1 := s.UserSvc.UpdateUserFriendNum(userId, friendId, false)
		if err1 != nil || resp2.Msg.Error != "" || resp2.Msg.Status != 200 {
			logs.SugarLogger.Errorf("UserId[%v] FriendId[%v] Delete friend record failed and compensate failed err: err[%v], err1[%v], msg.error[%v]", userId, friendId, err, err1, resp.Msg.Error)
			return errors.New(shared.CodeMessageIgnoreCode(shared.DeleteError))
		}
		logs.SugarLogger.Warnf("UserId[%v] FriendId[%v] Delete friend record failed: %v, and compensate success!", userId, friendId, err)
		return errors.New(shared.CodeMessageIgnoreCode(shared.DeleteError))
	}
	if err != nil {
		return err
	}
	return nil
}

// UpdateFriendInfo 修改好友信息
func (s *FriendService) UpdateFriendInfo(friendInfo *models.UserFriend) (*models.UserFriend, error) {
	if friendInfo == nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	err := s.D.UpdateFriendInfo(friendInfo)
	if err != nil {
		return nil, err
	}
	f, err := s.FindUserOneFriend(friendInfo.UserId, friendInfo.FriendId)
	if err != nil {
		logs.SugarLogger.Warnf("UserId[%v] FriendId[%v] Update friend info but find failed: %v", friendInfo.UserId, friendInfo.FriendId, err)
		return nil, err
	}
	return f, nil
}

// FindAllFriends 根据筛选条件查询用户好友
func (s *FriendService) FindAllFriends(uid int64, opt int) ([]*models.UserFriend, error) {
	friends, err := s.D.FindUserAllFriends(uid, opt)
	if err != nil {
		return make([]*models.UserFriend, 0), err
	}
	return friends, nil
}
