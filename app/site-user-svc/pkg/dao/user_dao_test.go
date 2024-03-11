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
