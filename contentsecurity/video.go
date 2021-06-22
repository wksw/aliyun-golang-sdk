// Package contentsecurity 视频内容安全检测
package contentsecurity

import "net/http"

// ScanVideoSyncReq 视频同步检测请求
type ScanVideoSyncReq struct {
	ScanCommonReq
	Tasks []ScanVideoTask `json:"tasks"`
}

// ScanVideoAsyncReq 视频异步检测请求
type ScanVideoAsyncReq struct {
	ScanCommonReq
	AudioScenes []string        `json:"audioScenes"`
	Live        bool            `json:"live"`
	Offline     bool            `json:"offline"`
	Callback    string          `json:"callback"`
	Seed        string          `json:"seed"`
	CryptType   string          `json:"cryptType"`
	Tasks       []ScanVideoTask `json:"tasks"`
}

// ScanVideoTask 视频检测任务
type ScanVideoTask struct {
	ClientInfo ClientInfo `json:"clientInfo"`
	DataId     string     `json:"dataId" validate:"max=128"`
	LiveId     string     `json:"liveId"`
	Url        string     `json:"url"`
	Frams      []struct {
		Rate float32 `json:"rate"`
		Url  string  `json:"url" validate:"required"`
	} `json:"frams"`
	FramePrefix string `json:"framePrefix"`
	Interval    int    `json:"interval"`
	MaxFrames   int    `json:"maxFrames"`
}

// ScanVideoTaskFram 待检测视频的截帧信息
type ScanVideoTaskFram struct {
	Url    string `json:"url"`
	Offset int    `json:"offset"`
}

// ScanVideoAsyncResp 视频异步检测返回
type ScanVideoAsyncResp struct {
	ContentSecurityCommonResp
	Data []struct {
		ScanCommonDataResp
	} `json:"data"`
}

// ScanVideoData 视频检测返回
type ScanVideoData struct {
	ScanCommonDataResp
	AudioScanResults []struct {
		ScanCommonResultResp
		AudioScene string `json:"audioScene"`
		Details    []struct {
			StartTime int64  `json:"startTime"`
			EndTime   int64  `json:"endTime"`
			Text      string `json:"text"`
			Label     string `json:"label"`
			Keyword   string `json:"keyword"`
			LibName   string `json:"libName"`
		}
	} `json:"audioScanResults"`
	Results []struct {
		ScanCommonResultResp
		// 截断后的每一帧图像的临时访问地址
		Frames []struct {
			Rate   float32 `json:"rate"`
			Url    string  `json:"url" validate:"required"`
			Offset int     `json:"offset"`
			Label  string  `json:"label"`
		} `json:"frams"`
		Extras       ScanExtras `json:"extras"`
		ExtrasNewUrl string     `json:"extras.newUrl"`
		// 图片中含有广告或文字违规信息时，返回图片中广告文字命中的风险关键词信息。
		HintWordsInfo []struct {
			Context string `json:"context"`
		} `json:"hintWordsInfo"`
		LogoData []struct {
			Type string  `json:"type"`
			Name string  `json:"name"`
			X    float32 `json:"x"`
			Y    float32 `json:"y"`
			W    float32 `json:"w"`
			H    float32 `json:"h"`
		} `json:"logoData"`
		SfaceData []struct {
			X     float32 `json:"x"`
			Y     float32 `json:"y"`
			W     float32 `json:"w"`
			H     float32 `json:"h"`
			Faces []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				Rate string `json:"rate"`
			} `json:"faces"`
		} `json:"sfaceData"`
		OcrData []string `json:"ocrData,omitempty"`
	} `json:"results"`
}

// ScanVideoResp 视频检测返回
type ScanVideoResp struct {
	ContentSecurityCommonResp
	Data []ScanVideoData `json:"data"`
}

// ScanVideoSync 视频同步检测
func (c Client) ScanVideoSync(in *ScanVideoSyncReq) (*ScanVideoResp, error) {
	resp, err := c.Do(http.MethodPost, VIDEO_SYNC_API_PATH, in)
	if err != nil {
		return nil, err
	}
	result := &ScanVideoResp{}
	if err := interfaceConvert(resp, result); err != nil {
		return result, err
	}
	return result, nil
}

// ScanVideoAsync 视频异步检测
func (c Client) ScanVideoAsync(in *ScanVideoAsyncReq) (*ScanVideoAsyncResp, error) {
	resp, err := c.Do(http.MethodPost, VIDEO_ASYNC_API_PATH, in)
	if err != nil {
		return nil, err
	}
	result := &ScanVideoAsyncResp{}
	if err := interfaceConvert(resp, result); err != nil {
		return result, err
	}
	return result, nil
}

// ScanVideoResult 视频异步检测结果
func (c Client) ScanVideoResult(taskIds []string) (*ScanVideoResp, error) {
	resp, err := c.Do(http.MethodPost, VIDEO_ASYNC_API_PATH, taskIds)
	if err != nil {
		return nil, err
	}
	result := &ScanVideoResp{}
	if err := interfaceConvert(resp, result); err != nil {
		return result, err
	}
	return result, nil
}

// ScanVideoCancel 取消视频检测
func (c Client) ScanVideoCancel(taskIds []string) (*ScanCommonDataResp, error) {
	resp, err := c.Do(http.MethodPost, VIDEO_CANCEL_API_PATH, taskIds)
	if err != nil {
		return nil, err
	}
	result := &ScanCommonDataResp{}
	if err := interfaceConvert(resp, result); err != nil {
		return result, err
	}
	return result, nil
}
