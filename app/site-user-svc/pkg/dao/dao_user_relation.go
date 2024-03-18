package dao

import (
	"errors"
	"site/app/site-user-svc/pkg/models"
	"site/protocol/shared"
)

// CreateRelation 插入用户关系信息
func (d *Dao) CreateRelation(relation *models.UserRelation) error {
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	result := tx.Model(models.UserRelation{}).Create(relation)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// FindRelationById 通过 id 查询用户关系
func (d *Dao) FindRelationById(id int64) (*models.UserRelation, error) {
	var r *models.UserRelation
	result := d.DB.Model(models.UserRelation{}).First(&r, uint(id))
	if result.Error != nil {
		return nil, result.Error
	}
	if r.ID == 0 {
		return nil, errors.New(shared.CodeMessageIgnoreCode(shared.RecordNotFound))
	}
	return r, nil
}

// AddFriendById 添加或删除好友同时修改两条记录 felFlag 为删除标识
func (d *Dao) AddFriendById(id1 int64, id2 int64, delFlag bool) error {
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	r1, err := d.FindRelationById(id1)
	if err != nil {
		return err
	}
	r2, err := d.FindRelationById(id2)
	if err != nil {
		return err
	}
	if delFlag {
		r1.FriendNum -= 1
		r2.FriendNum -= 1
	} else {
		r1.FriendNum += 1
		r2.FriendNum += 1
	}
	result := d.DB.Save(r1)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = d.DB.Save(r2)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

// UpdateUserRelation 修改用户关系信息
func (d *Dao) UpdateUserRelation(id int64, relation *models.UserRelation) error {
	r, err := d.FindRelationById(id)
	if err != nil {
		return err
	}
	r.SearchLimit = relation.SearchLimit
	r.VisitLimit = relation.VisitLimit
	r.AddLimit = relation.AddLimit
	r.FriendNum = relation.FriendNum
	r.TopFriendNum = relation.TopFriendNum
	r.BlackFriendNum = relation.BlackFriendNum
	r.GroupNum = relation.GroupNum
	r.OwnerGroupNum = relation.OwnerGroupNum
	r.AdminGroupNum = relation.AdminGroupNum
	r.TopGroupNum = relation.TopGroupNum
	r.BlackGroupNum = relation.BlackGroupNum
	result := d.DB.Save(r)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SagaAdjustUserRelation DTM 处理 saga 事务修改用户关系信息
// 仅修改用户好友等相关数量 不修改隐私权限 传入参数为偏移量
func (d *Dao) SagaAdjustUserRelation(id int64, deltaRelation *models.UserRelation) error {
	r, err := d.FindRelationById(id)
	if err != nil {
		return err
	}
	r.FriendNum += deltaRelation.FriendNum
	r.TopFriendNum += deltaRelation.TopFriendNum
	r.BlackFriendNum += deltaRelation.BlackFriendNum
	r.GroupNum += deltaRelation.GroupNum
	r.OwnerGroupNum += deltaRelation.OwnerGroupNum
	r.AdminGroupNum += deltaRelation.AdminGroupNum
	r.TopGroupNum += deltaRelation.TopGroupNum
	r.BlackGroupNum += deltaRelation.BlackGroupNum
	result := d.DB.Save(r)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
