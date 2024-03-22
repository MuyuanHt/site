package dao

import (
	"gorm.io/gorm"
	"site/app/site-collaborate-svc/pkg/models"
	"site/common/dbs"
)

type Dao struct {
	DB *gorm.DB
}

// InitDao 初始化
func InitDao(mysqlDsn string) (*Dao, error) {
	d := &Dao{
		DB: dbs.InitMysql(mysqlDsn),
	}
	// 数据库自动迁移
	err := d.DB.AutoMigrate(&models.UserFriend{})
	if err != nil {
		return nil, err
	}
	//err = d.DB.AutoMigrate(&models.GroupInfo{})
	//if err != nil {
	//	return nil, err
	//}
	//err = d.DB.AutoMigrate(&models.GroupMember{})
	//if err != nil {
	//	return nil, err
	//}
	//err = d.DB.AutoMigrate(&models.ItemInfo{})
	//if err != nil {
	//	return nil, err
	//}
	//err = d.DB.AutoMigrate(&models.ItemNode{})
	//if err != nil {
	//	return nil, err
	//}
	//err = d.DB.AutoMigrate(&models.TeamInfo{})
	//if err != nil {
	//	return nil, err
	//}
	//err = d.DB.AutoMigrate(&models.TeamMember{})
	//if err != nil {
	//	return nil, err
	//}
	return d, nil
}

// Close 关闭数据库连接
func (d *Dao) Close() error {
	return dbs.CloseMysql(d.DB)
}
