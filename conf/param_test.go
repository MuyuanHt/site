package conf

import (
	"testing"
)

func TestLoadParamConfig(t *testing.T) {
	err := LoadParamConf()
	if err != nil {
		t.Errorf("LoadParamConfig failed: %v", err)
	}
	t.Logf("ParamConfig: %v", ParamConfig)
}
