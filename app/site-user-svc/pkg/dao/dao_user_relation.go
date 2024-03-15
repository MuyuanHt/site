package dao

import "site/app/site-user-svc/pkg/models"

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
	return r, nil
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
	result := d.DB.Save(r)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
