package dao

import (
	"site/app/site-user-svc/pkg/models"
	"site/common/db"
	"site/conf"
	"site/protocol/shared"
	"testing"
)

// 测试时连接数据库使用
func initTestDao() *Dao {
	c, err := conf.LoadAppConf()
	if err != nil {
		panic(err)
	}
	dsn := c.GetServiceConf("user").Mysql.GetMysqlDsn()
	d := &Dao{
		DB: db.InitMysql(dsn),
	}
	// 自动迁移
	_ = d.DB.AutoMigrate(&models.UserInfo{})
	_ = d.DB.AutoMigrate(&models.Account{})
	return d
}

// 测试时为了方便先初始化后根据需要修改字段值
func initTestTempAccount() *models.Account {
	return &models.Account{
		Uid:      100001,
		Phone:    "12345678901",
		Password: "password_password",
		Email:    "email@123.com",
		UserInfo: models.UserInfo{
			Username:    "username",
			Age:         20,
			Gender:      1,
			Region:      "sc",
			Icon:        "temp://xxx.xxx.xxx/xxx.png",
			Description: "description text",
		},
	}
}

func TestAccountDao_CreateAccount(t *testing.T) {
	d := initTestDao()
	a := initTestTempAccount()
	a.Phone = "12345678902"
	d.CreateAccount(a)
}

func TestAccountDao_FindOneAccountByAccount(t *testing.T) {
	d := initTestDao()
	a := initTestTempAccount()
	a.Phone = "12345678903"
	a.Email = "123@123.com"
	d.CreateAccount(a)
	res, _ := d.FindOneAccountByAccount(shared.AccountTypePhone, a.Phone)
	want := a.Email
	got := res.Email
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestAccountDao_FindOneUserByAccount(t *testing.T) {
	d := initTestDao()
	a := initTestTempAccount()
	a.Phone = "12345678916"
	a.UserInfo.Username = "test_name15"
	d.CreateAccount(a)
	res, _ := d.FindOneUserByAccount(shared.AccountTypePhone, a.Phone)
	// 看一下 userinfo 有没有查成功
	want := a.UserInfo.Username
	got := res.UserInfo.Username
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestAccountDao_UnsubscribeById(t *testing.T) {
	d := initTestDao()
	uid := int64(551428958831517696)
	err := d.UnsubscribeByUId(uid)
	if err != nil {
		t.Errorf("unsubscribe failed: %v", err)
	}
}

func TestAccountDao_FindUsersLikePhone(t *testing.T) {
	d := initTestDao()
	param := "848"
	accounts, err := d.FindUsersLikePhone(param)
	if err != nil {
		t.Errorf("find users err: %v", err)
	}
	t.Logf("len %v", len(accounts))
	for _, v := range accounts {
		t.Logf("phone: %v, name: %v", v.Phone, v.Username)
	}
}

func TestAccountDao_FindUsersLikeEmail(t *testing.T) {
	d := initTestDao()
	param := "bc"
	accounts, err := d.FindUsersLikeEmail(param)
	if err != nil {
		t.Errorf("find users err: %v", err)
	}
	t.Logf("len %v", len(accounts))
	for _, v := range accounts {
		t.Logf("phone: %v, name: %v", v.Email, v.Username)
	}
}

func TestAccountDao_FindUsersLikeName(t *testing.T) {
	d := initTestDao()
	param := "u"
	accounts, err := d.FindUsersLikeName(param)
	if err != nil {
		t.Errorf("find users err: %v", err)
	}
	t.Logf("len %v", len(accounts))
	for _, v := range accounts {
		t.Logf("phone: %v, name: %v", v.Phone, v.Username)
	}
}
