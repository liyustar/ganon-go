package securityexpl_test

import (
	"fmt"
	"testing"

	"github.com/liyustar/nuts/securityexpl"
)

func TestDes(t *testing.T) {
	var key, origData = []byte("lyx!hell"), []byte("hello!")
		fmt.Println(key, origData)

	crypted, err := securityexpl.DesEncrypt(origData, key)
	if err != nil {
		t.Fatalf("加密出错，%v", err)
	}

	decrypt, err := securityexpl.DesDecrypt(crypted, key)
	if err != nil {
		t.Fatalf("解密出错，%v", err)
	}

	if string(origData) != string(decrypt) {
		t.Fatal("解密后与明文不一致")
	}
}

func TestTripleDes(t *testing.T) {
	var key, origData = []byte("lyx!hell" + "lyx!hell" + "lyx!hell"), []byte("hello!")
	fmt.Println(key, origData)

	crypted, err := securityexpl.TripleDesEncrypt(origData, key)
	if err != nil {
		t.Fatalf("加密出错，%v", err)
	}

	decrypt, err := securityexpl.TripleDesDecrypt(crypted, key)
	if err != nil {
		t.Fatalf("解密出错，%v", err)
	}

	if string(origData) != string(decrypt) {
		t.Fatal("解密后与明文不一致")
	}
}
