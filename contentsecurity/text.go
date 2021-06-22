// Package contentsecurity 图片内容安全检测
package contentsecurity

import (
	"net/http"
)

// ScanTextRequest 文本内容检测请求
// 参考https://help.aliyun.com/document_detail/70439.html?spm=a2c4g.11186623.6.701.7cae3860tgGsFO
type ScanTextRequest struct {
	ScanCommonReq
	// 检测对象列表
	Tasks []ScanTextTask `json:"tasks,omitempty"`
}

// ScanTextTask 文本内容检测任务
type ScanTextTask struct {
	ClientInfo ClientInfo `json:"clientInfo"`
	// 检测对象对应数据ID
	DataId string `json:"dataId" validate:"max=128"`
	// 待检测文本
	Content string `json:"content" validate:"required,max=10000"`
}

// ScanTextResponse 文本内容检测返回
type ScanTextResponse struct {
	ContentSecurityCommonResp
	Data []struct {
		ScanCommonDataResp
		Content         string `json:"content"`
		FilteredContent string `json:"filteredContent"`
		Results         []struct {
			ScanCommonResultResp
			Extras  map[string]string `json:"extras"`
			Details []struct {
				Label    string `json:"label"`
				Contexts []struct {
					Context   string           `json:"context"`
					Positions []map[string]int `json:"positions"`
					LibName   string           `json:"libName"`
					LibCode   string           `json:"libCode"`
					RuleType  string           `json:"ruleType"`
				} `json:"contexts"`
			} `json:"details"`
		} `json:"results,omitempty"`
	} `json:"data"`
}

// ScanText 文本内容安全检测
func (c Client) ScanText(in *ScanTextRequest) (*ScanTextResponse, error) {
	resp, err := c.Do(http.MethodPost, TEXT_API_PATH, in)
	if err != nil {
		return nil, err
	}
	result := &ScanTextResponse{}
	if err := interfaceConvert(resp, result); err != nil {
		return result, err
	}
	return result, nil
}
