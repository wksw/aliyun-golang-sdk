package contentsecurity

import (
	"encoding/json"
	"testing"
)

func TestRequest(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	var out ScanTextResp
	if err := client.Request(TEXT_API_PATH, H{
		"scenes": []string{TextSceneAntispam},
		"tasks": []H{
			{
				"dataId":  "123",
				"content": "本校小额贷款，安全、快捷、方便、无抵押，随机随贷，当天放款，上门服务。",
			},
		},
	}, &out); err != nil {
		t.Error(err)
		t.FailNow()
	}
	b, _ := json.Marshal(&out)
	t.Log(string(b))
}
