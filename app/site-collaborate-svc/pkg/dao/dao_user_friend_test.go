package dao

import (
	"site/app/site-collaborate-svc/pkg/models"
	"site/common/dbs"
	"site/conf"
	"site/protocol/shared"
	"testing"
)

// 测试时连接数据库使用
func initTestFriendDao() *Dao {
	c, err := conf.LoadAppConf()
	if err != nil {
		panic(err)
	}
	dsn := c.GetServiceConf("team").Mysql.GetMysqlDsn()
	d := &Dao{
		DB: dbs.InitMysql(dsn),
	}
	// 自动迁移
	_ = d.DB.AutoMigrate(&models.UserFriend{})
	return d
}

// 关闭测试数据库连接
func closeTestFriendDao(d *Dao) {
	_ = dbs.CloseMysql(d.DB)
}

// 测试时为了方便先初始化后根据需要修改字段值
func initTestTempFriend() *models.UserFriend {
	return &models.UserFriend{
		IsTop:    false,
		IsBlack:  false,
		UserId:   12345,
		FriendId: 23456,
		Label:    "text",
	}
}

func TestDao_CreateBecomeFriend(t *testing.T) {
	d := initTestFriendDao()
	defer closeTestFriendDao(d)
	u1 := initTestTempFriend()
	u1.UserId = 123
	u1.FriendId = 237
	u2 := initTestTempFriend()
	u2.UserId = u1.FriendId
	u2.FriendId = u1.UserId
	err := d.CreateBecomeFriend(u1, u2)
	if err != nil {
		t.Errorf("CreateBecomeFriend failed: %v", err)
	}
}

func TestDao_FindFriendRecord(t *testing.T) {
	d := initTestFriendDao()
	defer closeTestFriendDao(d)
	uid1 := int64(123)
	uid2 := int64(234)
	f, err := d.FindFriendRecord(uid1, uid2)
	if err != nil {
		t.Errorf("FindRecord err: %v", err)
		return
	}
	t.Log(f.ID)
}

func TestDao_FindIsFriend(t *testing.T) {
	d := initTestFriendDao()
	defer closeTestFriendDao(d)
	uid1 := int64(123)
	uid2 := int64(234)
	f1, f2, err := d.FindIsFriend(uid1, uid2)
	if err != nil {
		t.Errorf("FindRecord err: %v", err)
		return
	}
	t.Log(f1.ID)
	t.Log(f2.ID)
}

func TestDao_DeleteFriendRecord(t *testing.T) {
	d := initTestFriendDao()
	defer closeTestFriendDao(d)
	uid1 := int64(123)
	uid2 := int64(234)
	err := d.DeleteFriendRecord(uid1, uid2)
	if err != nil {
		t.Errorf("Delete friend record err: %v", err)
	}
}

func TestDao_UpdateFriendInfo(t *testing.T) {
	d := initTestFriendDao()
	defer closeTestFriendDao(d)
	info := &models.UserFriend{
		UserId:   123,
		FriendId: 236,
		IsTop:    true,
		IsBlack:  false,
		Label:    "friend1",
	}
	err := d.UpdateFriendInfo(info)
	if err != nil {
		t.Errorf("Update friend info err: %v", err)
	}
}

func TestDao_FindUserAllFriend(t *testing.T) {
	d := initTestFriendDao()
	defer closeTestFriendDao(d)
	uid := int64(123)
	opt := shared.FindAllOpt
	friends, err := d.FindUserAllFriends(uid, opt)
	if err != nil {
		t.Errorf("Find all friends err: %v", err)
		return
	}
	t.Log(len(friends))
	for _, v := range friends {
		t.Logf("id: %v, friendId: %v, label: %v", v.ID, v.FriendId, v.Label)
	}
}
