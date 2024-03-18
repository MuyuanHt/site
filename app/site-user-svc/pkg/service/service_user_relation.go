package service

import (
	"errors"
	"site/app/site-user-svc/pkg/models"
	"site/common/tools"
	"site/protocol/shared"
)

// AddFriendByUid 添加或删除好友
func (s *UserService) AddFriendByUid(uid1 int64, uid2 int64, delFlag bool) error {
	account1, err := s.FindOneUserByUid(uid1)
	if err != nil {
		return err
	}
	account2, err := s.FindOneUserByUid(uid2)
	if err != nil {
		return err
	}
	return s.D.AddFriendById(int64(account1.RelationId), int64(account2.RelationId), delFlag)
}

// UpdateUserLimit 修改用户隐私权限
func (s *UserService) UpdateUserLimit(uid int64, limit *models.UserLimit) (*models.Account, error) {
	if limit == nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	result, err := s.D.FindOneUserByUid(uid)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UserNotFound))
	}
	if !tools.IsSearchLimit(int(limit.SearchLimit)) || !tools.IsVisitLimit(int(limit.VisitLimit)) || !tools.IsAddLimit(int(limit.AddLimit)) {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	newRelation := &models.UserRelation{
		SearchLimit:    limit.SearchLimit,
		VisitLimit:     limit.VisitLimit,
		AddLimit:       limit.AddLimit,
		FriendNum:      result.UserRelation.FriendNum,
		TopFriendNum:   result.UserRelation.TopFriendNum,
		BlackFriendNum: result.UserRelation.BlackFriendNum,
		GroupNum:       result.UserRelation.GroupNum,
		OwnerGroupNum:  result.UserRelation.OwnerGroupNum,
		AdminGroupNum:  result.UserRelation.AdminGroupNum,
		TopGroupNum:    result.UserRelation.TopGroupNum,
		BlackGroupNum:  result.UserRelation.BlackGroupNum,
	}
	err = s.D.UpdateUserRelation(int64(result.RelationId), newRelation)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UpdateUserRelationError))
	}
	resAccount, err := s.D.FindOneUserByUid(uid)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UserNotFound))
	}
	return resAccount, nil
}

// UpdateUserRelationNum 修改用户关系数量 传入参数为变化量
func (s *UserService) UpdateUserRelationNum(uid int64, delta *models.UserRelationNum) (*models.Account, error) {
	if delta == nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	result, err := s.D.FindOneUserByUid(uid)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UserNotFound))
	}
	friendNum := result.UserRelation.FriendNum + delta.FriendNum
	topFriendNum := result.UserRelation.TopFriendNum + delta.TopFriendNum
	blackFriendNum := result.UserRelation.BlackFriendNum + delta.BlackFriendNum
	groupNum := result.UserRelation.GroupNum + delta.GroupNum
	ownerGroupNum := result.UserRelation.OwnerGroupNum + delta.OwnerGroupNum
	adminGroupNum := result.UserRelation.AdminGroupNum + delta.AdminGroupNum
	topGroupNum := result.UserRelation.TopGroupNum + delta.TopGroupNum
	blackGroupNum := result.UserRelation.BlackGroupNum + delta.BlackGroupNum
	if friendNum < 0 || topFriendNum < 0 || blackFriendNum < 0 || groupNum < 0 || ownerGroupNum < 0 || adminGroupNum < 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	newRelation := &models.UserRelation{
		SearchLimit:    result.UserRelation.SearchLimit,
		VisitLimit:     result.UserRelation.VisitLimit,
		AddLimit:       result.UserRelation.AddLimit,
		FriendNum:      friendNum,
		TopFriendNum:   topFriendNum,
		BlackFriendNum: blackFriendNum,
		GroupNum:       groupNum,
		OwnerGroupNum:  ownerGroupNum,
		AdminGroupNum:  adminGroupNum,
		TopGroupNum:    topGroupNum,
		BlackGroupNum:  blackGroupNum,
	}
	err = s.D.UpdateUserRelation(int64(result.RelationId), newRelation)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UpdateUserRelationError))
	}
	resAccount, err := s.D.FindOneUserByUid(uid)
	if err != nil {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.UserNotFound))
	}
	return resAccount, nil
}
