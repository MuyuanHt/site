package dao

import (
	"site/protocol/shared"
	"testing"
)

func TestUpdateUserInfo(t *testing.T) {
	d := initTestDao()
	res, err := d.FindOneUserByAccount(shared.AccountTypePhone, "17334394823")
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}
	info := res.UserInfo
	info.Age = 20
	info.Gender = 1
	err = d.UpdateUserInfo(int64(res.UserId), &info)
	if err != nil {
		t.Errorf("Error updating user: %v", err)
	}
}

func TestDao_FindUserInfosLikeName(t *testing.T) {
	d := initTestDao()
	name := "u"
	infos, err := d.FindUserInfosLikeName(name)
	if err != nil {
		t.Errorf("Error getting users: %v", err)
	}
	t.Logf("len: %d", len(infos))
	for _, info := range infos {
		t.Logf("id: %v, name: %v", info.ID, info.Username)
	}
}
