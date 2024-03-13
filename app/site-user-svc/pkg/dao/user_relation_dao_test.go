package dao

import "testing"

func TestDao_FindRelationById(t *testing.T) {
	d := initTestDao()
	res, err := d.FindRelationById(3)
	if err != nil {
		t.Errorf("FindRelationById failed: %v", err)
		return
	}
	t.Log(res.SearchLimit)
	t.Log(res.FriendNum)
}
