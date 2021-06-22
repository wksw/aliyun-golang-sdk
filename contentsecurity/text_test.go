package contentsecurity

import (
	"encoding/json"
	"testing"
)

func TestTextScan(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanText(&ScanTextRequest{
		ScanCommonReq: ScanCommonReq{
			Scenes: []string{"antispam"},
		},
		Tasks: []ScanTextTask{
			{
				ClientInfo: ClientInfo{
					UserId:   "abc",
					UserType: UserOther,
				},
				DataId:  "123",
				Content: "本校小额贷款，安全、快捷、方便、无抵押，随机随贷，当天放款，上门服务。",
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
