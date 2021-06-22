package contentsecurity

import (
	"testing"
)

func TestRandString(t *testing.T) {

}

func TestMd5(t *testing.T) {
	request := []byte("abc123")
	expect := "6ZoYxCjLONXyYIU2eJIuAw=="

	result := newBase64Md5(request)
	if result != expect {
		t.Error("md5 not expect")
	}
}

func TestHmacSha1(t *testing.T) {
	request := `POST
application/json
C+5Y0crpO4sYgC2DNjycug==
application/json
Tue, 14 Mar 2017 06:29:50 GMT
x-acs-signature-method:HMAC-SHA1
x-acs-signature-nonce:339497c2-d91f-4c17-a0a3-1192ee9e2202
x-acs-signature-version:1.0
x-acs-version:2018-05-09
/green/image/scan?clientInfo={"ip":"127.xxx.xxx.2","userId":"12023xxxx","userNick":"Mike","userType":"others"}`
	expect := "4n2UfyR1BSAJ//O+yGraLUobI/M="
	key := "abc1234567"
	result, err := newBase64HmacSha1(key, []byte(request))
	if err != nil {
		t.Errorf("base64 hmac sha1 got error[%s]", err.Error())
	}
	if result != expect {
		t.Error("base64 hmac sha1 not expect")
	}
}

func BenchmarkRand(b *testing.B) {
	repeatMap := make(map[string]bool)
	for i := 0; i < b.N; i++ {
		result := ranString(10)
		if _, ok := repeatMap[result]; ok {
			b.Error("rand repeat")
		}
		repeatMap[result] = true
	}
}

func TestRand(t *testing.T) {
	expect := 10
	result := ranString(expect)
	if len(result) != expect {
		t.Errorf("rand length not expect")
	}
}
