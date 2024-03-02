package conf

import "testing"

func TestLoadKey(t *testing.T) {
	priKey, err := LoadPrivateKey()
	if err != nil {
		t.Errorf("LoadPriKey error: %v", err)
	}
	pubKey, err := LoadPublicKey()
	if err != nil {
		t.Errorf("LoadPubKey error: %v", err)
	}
	t.Logf("priKey: %v", string(priKey))
	t.Logf("pubKey: %v", string(pubKey))
}
