// Package contentsecurity 网页内容安全检测
package contentsecurity

import "net/http"

// ScanWebpageReq 网页检测请求
type ScanWebpageReq struct {
	ScanCommonReq
	TextScenes          []string          `json:"textScenes"`
	ImageScenes         []string          `json:"imageScenes"`
	ReturnHighlightHtml bool              `json:"returnHighlightHtml"`
	Tasks               []ScanWebpageTask `json:"tasks"`
}

// ScanWebpageAsyncReq 网页检测异步请求
type ScanWebpageAsyncReq struct {
	ScanWebpageReq
	ScanCommonAsyncReq
}

// ScanWebpageAsyncResp 网页检测异步请求返回
type ScanWebpageAsyncResp struct {
	ContentSecurityCommonResp
	Data []struct {
		ScanCommonDataResp
		Url string `json:"string"`
	} `json:"data"`
}

// ScanWebpageTask 网页检测任务
type ScanWebpageTask struct {
	DataId string `json:"dataId" validate:"max=128"`
	Url    string `json:"url" validate:"required"`
	// 待检测文本
	// url和content二选一
	Content string `json:"content"`
}

// ScanWebpageResp 网页检测回复
type ScanWebpageResp struct {
	ContentSecurityCommonResp
	Data []struct {
		ScanCommonDataResp
		Suggestion    string          `json:"suggestion"`
		RiskFrequency map[string]int  `json:"riskFrequency"`
		TextResults   []ScanTextData  `json:"textResults,omitempty"`
		ImageResults  []ScanImageData `json:"imageResults,omitempty"`
		HighlightHtml string          `json:"highlightHtml,omitempty"`
	} `json:"data"`
}

// ScanWebpageSync 网页内容同步检测
func (c Client) ScanWebpageSync(in *ScanWebpageReq) (result *ScanWebpageResp, err error) {
	resp, err := c.Do(http.MethodPost, WEBPAGE_SYNC_API_PATH, in)
	if err != nil {
		return nil, err
	}
	result = &ScanWebpageResp{}
	return result, interfaceConvert(resp, result)
}

// ScanWebpageAsync 网页内容同步检测
func (c Client) ScanWebpageAsync(in *ScanWebpageAsyncReq) (result *ScanWebpageAsyncResp, err error) {
	resp, err := c.Do(http.MethodPost, WEBPAGE_SYNC_API_PATH, in)
	if err != nil {
		return nil, err
	}
	result = &ScanWebpageAsyncResp{}
	return result, interfaceConvert(resp, result)
}

// ScanWebpageResult 获取网页内容异步检测结果
func (c Client) ScanWebpageResult(taskIds []string) (result *ScanWebpageResp, err error) {
	resp, err := c.Do(http.MethodPost, WEBPAGE_RESULT_API_PATH, taskIds)
	if err != nil {
		return nil, err
	}
	result = &ScanWebpageResp{}
	err = interfaceConvert(resp, result)
	return result, err
}
