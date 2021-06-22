package contentsecurity

import (
	"encoding/json"
	"testing"
	"time"
)

func TestScanVideoAsync(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanVideoAsync(&ScanVideoAsyncReq{
		ScanCommonReq: ScanCommonReq{
			Scenes: []string{"porn"},
		},
		Tasks: []ScanVideoTask{
			{
				ClientInfo: ClientInfo{
					UserId:   "abc",
					UserType: UserOther,
				},
				DataId:    "123",
				Url:       "https://zhiou-index-video.oss-cn-beijing.aliyuncs.com/index_video1.mp4",
				Interval:  5,
				MaxFrames: 100,
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
		t.Errorf("response code not 200[%s]", string(b))
		t.FailNow()
	}

	var taskIds []string
	for _, data := range resp.Data {
		taskIds = append(taskIds, data.TaskId)
	}
	time.Sleep(time.Minute)
	asyncResp, err := client.ScanVideoResult(taskIds)
	ab, _ := json.Marshal(asyncResp)
	t.Log(string(ab))
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if asyncResp.Code/100 != 2 {
		t.Errorf("response code not 200[%s]", string(ab))
		t.FailNow()
	}
}

func TestScanVideoResult(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanVideoResult([]string{"vi1cTYTYXVbFi5jj9RSZoQzG-1uzab2", "123"})
	b, _ := json.Marshal(resp)
	t.Log(string(b))
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log(resp)
	if resp.Code != 200 {
		t.Errorf("response code not 200[%s]", string(b))
		t.FailNow()
	}

}

func TestScanVideoCancel(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanVideoCancel([]string{"vi1cTYTYXVbFi5jj9RSZoQzG-1uzab2", "123"})
	b, _ := json.Marshal(resp)
	t.Log(string(b))
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log(resp)
	if resp.Code != 200 {
		t.Errorf("response code not 200[%s]", string(b))
		t.FailNow()
	}

}
