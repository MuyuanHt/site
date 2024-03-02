package conf

import (
	"reflect"
	"testing"
)

func TestAppAddrConfig_GetServiceAddr(t *testing.T) {
	want := &ServiceConf{
		Name: "user-service",
		Host: "localhost",
		Port: "50011",
	}
	c, _ := LoadAppConf()
	got := c.GetServiceConf("user")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v \n", got, want)
	}
}

func TestFilePath(t *testing.T) {
	want := "./../conf/yaml/app.yaml"
	got := filePath("app.yaml", "yaml")
	if got != want {
		t.Errorf("got: %v, want %v", got, want)
	}
}

func TestGetMysqlDsn(t *testing.T) {
	want := "root:my0309510.@tcp(localhost:3306)/site_user_svs?charset=utf8mb4&parseTime=True&loc=Local"
	c, _ := LoadAppConf()
	got := c.GetServiceConf("user").Mysql.GetMysqlDsn()
	if want != got {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
