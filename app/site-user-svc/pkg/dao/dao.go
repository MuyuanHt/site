package dao

import (
	"gorm.io/gorm"
	"site/app/site-user-svc/pkg/models"
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
	err := d.DB.AutoMigrate(&models.UserInfo{})
	if err != nil {
		return nil, err
	}
	err = d.DB.AutoMigrate(&models.UserRelation{})
	if err != nil {
		return nil, err
	}
	err = d.DB.AutoMigrate(&models.Account{})
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Close 关闭数据库连接
func (d *Dao) Close() error {
	return dbs.CloseMysql(d.DB)
}
