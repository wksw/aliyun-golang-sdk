// Package contentsecurity 图片内容安全检测
package contentsecurity

import (
	"net/http"
)

// ScanImageSyncReq 图片同步检测请求
// 参考https://help.aliyun.com/document_detail/70292.html?spm=a2c4g.11186623.6.628.5234aba5r5qacM
type ScanImageSyncReq struct {
	ScanCommonReq
	Tasks []ScanImageTask `json:"tasks"`
}

// ScanImageAsyncReq 图片异步检测请求
// 参考https://help.aliyun.com/document_detail/70430.html?spm=a2c4g.11186623.6.629.21ae4caccVYt0l
type ScanImageAsyncReq struct {
	ScanImageSyncReq
	ScanCommonAsyncReq
	Offline bool `json:"offline"`
}

// ScanImageTask 图片检测任务
type ScanImageTask struct {
	ClientInfo ClientInfo   `json:"clientInfo"`
	DataId     string       `json:"dataId" validate:"max=128"`
	Url        string       `json:"url" validate:"required"`
	Extras     []ScanExtras `json:"extras"`
	Interval   int          `json:"interval,omitempty"`
	MaxFrames  int          `json:"maxFrames,omitempty"`
}

// ScanImageAsyncResp 图片异步检测返回
type ScanImageAsyncResp struct {
	ContentSecurityCommonResp
	Data []struct {
		ScanCommonDataResp
		Url string `json:"url"`
	} `json:"data"`
}

// ScanImageData 图片检测返回
type ScanImageData struct {
	ScanCommonDataResp
	Extras   ScanExtras `json:"extras,omitempty"`
	Url      string     `json:"url"`
	StoreUrl string     `json:"storeUrl"`
	Results  []struct {
		ScanCommonResultResp
		// 截断后的每一帧图像的临时访问地址
		Frames []struct {
			Rate float32 `json:"rate"`
			Url  string  `json:"url" validate:"required"`
		} `json:"frams"`
		// 图片中含有广告或文字违规信息时，返回图片中广告文字命中的风险关键词信息。
		HintWordsInfo []struct {
			Context string `json:"context"`
		} `json:"hintWordsInfo,omitempty"`
		QrcodeData      []string `json:"qrcodeData,omitempty"`
		QrcodeLocations []struct {
			X      float32 `json:"x"`
			Y      float32 `json:"y"`
			W      float32 `json:"w"`
			H      float32 `json:"h"`
			Qrcode string  `json:"qrcode"`
		} `json:"qrcodeLocations,omitempty"`
		ProgramCodeData []struct {
			X float32 `json:"x"`
			Y float32 `json:"y"`
			W float32 `json:"w"`
			H float32 `json:"h"`
		} `json:"programCodeData,omitempty"`
		LogoData []struct {
			Type string  `json:"type"`
			Name string  `json:"name"`
			X    float32 `json:"x"`
			Y    float32 `json:"y"`
			W    float32 `json:"w"`
			H    float32 `json:"h"`
		} `json:"logoData,omitempty"`
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
		} `json:"sfaceData,omitempty"`
		OcrData []string `json:"ocrData,omitempty"`
	} `json:"results"`
}

// ScanImageResult  图片检测结果
// type ScanImageResult struct {
// 	ScanCommonResultResp
// 	// 截断后的每一帧图像的临时访问地址
// 	Frames []struct {
// 		Rate float32 `json:"rate"`
// 		Url  string  `json:"url" validate:"required"`
// 	} `json:"frams"`
// 	// 图片中含有广告或文字违规信息时，返回图片中广告文字命中的风险关键词信息。
// 	HintWordsInfo []struct {
// 		Context string `json:"context"`
// 	} `json:"hintWordsInfo,omitempty"`
// 	QrcodeData      []string `json:"qrcodeData,omitempty"`
// 	QrcodeLocations []struct {
// 		X      float32 `json:"x"`
// 		Y      float32 `json:"y"`
// 		W      float32 `json:"w"`
// 		H      float32 `json:"h"`
// 		Qrcode string  `json:"qrcode"`
// 	} `json:"qrcodeLocations,omitempty"`
// 	ProgramCodeData []struct {
// 		X float32 `json:"x"`
// 		Y float32 `json:"y"`
// 		W float32 `json:"w"`
// 		H float32 `json:"h"`
// 	} `json:"programCodeData,omitempty"`
// 	LogoData []struct {
// 		Type string  `json:"type"`
// 		Name string  `json:"name"`
// 		X    float32 `json:"x"`
// 		Y    float32 `json:"y"`
// 		W    float32 `json:"w"`
// 		H    float32 `json:"h"`
// 	} `json:"logoData,omitempty"`
// 	SfaceData []struct {
// 		X     float32 `json:"x"`
// 		Y     float32 `json:"y"`
// 		W     float32 `json:"w"`
// 		H     float32 `json:"h"`
// 		Faces []struct {
// 			Id   string `json:"id"`
// 			Name string `json:"name"`
// 			Rate string `json:"rate"`
// 		} `json:"faces"`
// 	} `json:"sfaceData,omitempty"`
// 	OcrData []string `json:"ocrData,omitempty"`
// }

// ScanImageResp 图片检测返回
type ScanImageResp struct {
	ContentSecurityCommonResp
	Data []ScanImageData `json:"data"`
}

// ScanImageSync 同步检测图片
func (c Client) ScanImageSync(in *ScanImageSyncReq) (result *ScanImageResp, err error) {
	resp, err := c.Do(http.MethodPost, IMAGE_SYNC_API_PATH, in)
	if err != nil {
		return nil, err
	}
	result = &ScanImageResp{}
	return result, interfaceConvert(resp, result)
}

// ScanImageAsync 异步检测图片
func (c Client) ScanImageAsync(in *ScanImageAsyncReq) (result *ScanImageAsyncResp, err error) {
	resp, err := c.Do(http.MethodPost, IMAGE_ASYNC_API_PATH, in)
	if err != nil {
		return nil, err
	}
	result = &ScanImageAsyncResp{}
	return result, interfaceConvert(resp, result)
}

// ScanImageAsyncResult 图片异步检测结果查询
func (c Client) ScanImageAsyncResult(taskIds []string) (result *ScanImageResp, err error) {
	resp, err := c.Do(http.MethodPost, IMAGE_ASYNC_RESULT_API_PATH, taskIds)
	if err != nil {
		return nil, err
	}
	result = &ScanImageResp{}
	return result, interfaceConvert(resp, result)
}
