package contentsecurity

import (
	"encoding/json"
	"testing"
	"time"
)

func TestScanImageSync(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanImageSync(&ScanImageSyncReq{
		ScanCommonReq: ScanCommonReq{
			Scenes: []string{ImageScenePorn},
		},
		Tasks: []ScanImageTask{
			{
				ClientInfo: ClientInfo{
					UserId:   "abc",
					UserType: UserOther,
				},
				DataId: "123",
				Url:    "https://tcloud-public.oss-cn-hongkong.aliyuncs.com/avatar/1384693247049883648/6f693c25-13ad-4979-941f-5966d9dc92cd.jpg?X_PP_Audience%3D1384693247049883648%26X_PP_ExpiredAt%3D1620465261%26X_PP_GrantedAt%3D1620465261%26X_PP_Method%3D%2A%26X_PP_ObjectName%3Davatar%2F1384693247049883648%2F6f693c25-13ad-4979-941f-5966d9dc92cd.jpg%26X_PP_Owner%3D1384693247049883648%26X_PP_ResourceType%3Davatar%26X_PP_Signature%3D12ee0717ba6cd2dc0b9c8e4a72a4d9e4",
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

func TestScanImageAsync(t *testing.T) {
	client := New(testEndpoint, testAccessKey, testSecretKey)
	resp, err := client.ScanImageAsync(&ScanImageAsyncReq{
		ScanImageSyncReq: ScanImageSyncReq{
			ScanCommonReq: ScanCommonReq{
				Scenes: []string{ImageScenePorn},
			},
			Tasks: []ScanImageTask{
				{
					ClientInfo: ClientInfo{
						UserId:   "abc",
						UserType: UserOther,
					},
					DataId: "123",
					Url:    "https://tcloud-public.oss-cn-hongkong.aliyuncs.com/avatar/1384693247049883648/6f693c25-13ad-4979-941f-5966d9dc92cd.jpg?X_PP_Audience%3D1384693247049883648%26X_PP_ExpiredAt%3D1620465261%26X_PP_GrantedAt%3D1620465261%26X_PP_Method%3D%2A%26X_PP_ObjectName%3Davatar%2F1384693247049883648%2F6f693c25-13ad-4979-941f-5966d9dc92cd.jpg%26X_PP_Owner%3D1384693247049883648%26X_PP_ResourceType%3Davatar%26X_PP_Signature%3D12ee0717ba6cd2dc0b9c8e4a72a4d9e4",
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
	if resp.Code != 200 {
		t.Errorf("response code not 200[%s]", string(b))
		t.FailNow()
	}

	var taskIds []string
	for _, data := range resp.Data {
		taskIds = append(taskIds, data.TaskId)
	}
	time.Sleep(10 * time.Second)
	asyncResp, err := client.ScanImageAsyncResult(taskIds)
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
