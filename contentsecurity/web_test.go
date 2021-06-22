package contentsecurity

import (
	"encoding/json"
	"testing"
	"time"
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

func TestScanWebpageAsync(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanWebpageAsync(&ScanWebpageAsyncReq{
		ScanWebpageReq: ScanWebpageReq{
			TextScenes:  []string{"antispam"},
			ImageScenes: []string{"porn"},
			Tasks: []ScanWebpageTask{
				{
					DataId: "123",
					Url:    "https://ziel.cn/about/",
				},
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
	if resp.Code/100 != 2 {
		t.Errorf("response code not 2xx")
	}
	time.Sleep(30 * time.Second)
	var taskIds []string
	for _, data := range resp.Data {
		taskIds = append(taskIds, data.TaskId)
	}
	resp1, err := client.ScanWebpageResult(taskIds)
	bb, _ := json.Marshal(resp1)
	t.Log(string(bb))
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log(resp1)
	if resp1.Code/100 != 2 {
		t.Errorf("response code not 2xx")
	}
}

func TestScanWebpageResult(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanWebpageResult([]string{"wp19GrgsoX3g6wMHOlTYcSc-1uzc@D"})
	b, _ := json.Marshal(resp)
	t.Log(string(b))
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log(resp)
	if resp.Code/100 != 2 {
		t.Errorf("response code not 2xx")
	}
}
