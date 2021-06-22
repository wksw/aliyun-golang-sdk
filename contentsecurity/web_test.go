package contentsecurity

import (
	"encoding/json"
	"testing"
)

func TestScanWebpageSync(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanWebpageSync(&ScanWebpageReq{
		TextScenes:  []string{"antispam"},
		ImageScenes: []string{"porn"},
		Tasks: []ScanWebpageTask{
			{
				DataId: "123",
				Url:    "https://ziel.cn/about/",
			},
		},
	})
	b, _ := json.Marshal(resp)
	t.Log(string(b))
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log(resp)
	if resp.Code != 200 {
		t.Errorf("response code not 200")
	}
}
