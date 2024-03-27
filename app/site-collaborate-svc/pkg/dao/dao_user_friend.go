package dao

import (
	"errors"
	"site/app/site-collaborate-svc/pkg/models"
	"site/protocol/shared"
)

// CreateBecomeFriend 成为好友时创建两条记录
func (d *Dao) CreateBecomeFriend(user1 *models.UserFriend, user2 *models.UserFriend) error {
	if user1 == nil || user2 == nil {
		return errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	result := tx.Model(models.UserFriend{}).Create(user1)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Model(models.UserFriend{}).Create(user2)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// CreateBecomeFriendById 创建好友记录
func (d *Dao) CreateBecomeFriendById(userId int64, friendId int64, userLabel string, friendLabel string) error {
	user := &models.UserFriend{
		UserId:   userId,
		FriendId: friendId,
		Label:    userLabel,
	}
	friend := &models.UserFriend{
		UserId:   friendId,
		FriendId: userId,
		Label:    friendLabel,
	}
	return d.CreateBecomeFriend(user, friend)
}

// FindFriendRecord 查找包含两个用户的好友记录 参数1 是用户 参数2 是该用户的好友
func (d *Dao) FindFriendRecord(ownerId int64, friendId int64) (*models.UserFriend, error) {
	var f models.UserFriend
	result := d.DB.Where(&models.UserFriend{UserId: ownerId, FriendId: friendId}).Find(&f)
	if result.Error != nil {
		return nil, result.Error
	}
	if f.ID == 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.RecordNotFound))
	}
	return &f, nil
}

// FindIsFriend 查看两个用户是否是好友 检查有无同时包含两用户的记录
func (d *Dao) FindIsFriend(uid1 int64, uid2 int64) (*models.UserFriend, *models.UserFriend, error) {
	f1, err := d.FindFriendRecord(uid1, uid2)
	if err != nil {
		return nil, nil, err
	}
	f2, err := d.FindFriendRecord(uid2, uid1)
	if err != nil {
		return nil, nil, err
	}
	return f1, f2, nil
}

// DeleteFriendRecord 删除好友 同时删除两条记录
func (d *Dao) DeleteFriendRecord(uid1 int64, uid2 int64) error {
	f1, f2, err := d.FindIsFriend(uid1, uid2)
	if err != nil {
		return err
	}
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return err
	}
	result := tx.Delete(&f1)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Delete(&f2)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// UpdateFriendInfo 更新某条好友记录
func (d *Dao) UpdateFriendInfo(friend *models.UserFriend) error {
	if friend == nil {
		return errors.New(shared.CodeMessageIgnoreCode(shared.ParamError))
	}
	f, err := d.FindFriendRecord(friend.UserId, friend.FriendId)
	if err != nil {
		return err
	}
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Error; err != nil {
		return err
	}
	f.IsTop = friend.IsTop
	f.IsBlack = friend.IsBlack
	f.Label = friend.Label
	result := tx.Save(f)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// FindUserAllFriends 查找用户所有好友 opt 选择查询条件 是否置顶/拉黑
func (d *Dao) FindUserAllFriends(uid int64, opt int) ([]*models.UserFriend, error) {
	friends := make([]*models.UserFriend, 0)
	query := d.DB.Model(&models.UserFriend{}).Where("user_id=?", uid)
	switch opt {
	case shared.FindTopOpt:
		query = query.Where("is_top=?", true)
	case shared.FindBlackOpt:
		query = query.Where("is_black=?", true)
	}
	result := query.Find(&friends)
	if result.Error != nil {
		return friends, result.Error
	}
	return friends, nil
}
