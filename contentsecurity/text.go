// Package contentsecurity 图片内容安全检测
package contentsecurity

import (
	"net/http"
)

// ScanTextReq 文本内容检测请求
// 参考https://help.aliyun.com/document_detail/70439.html?spm=a2c4g.11186623.6.701.7cae3860tgGsFO
type ScanTextReq struct {
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

// ScanTextResp 文本内容检测返回
type ScanTextResp struct {
	ContentSecurityCommonResp
	Data []ScanTextData `json:"data"`
}

// ScanTextData 文本内容检测返回数据
type ScanTextData struct {
	ScanCommonDataResp
	Content         string `json:"content"`
	FilteredContent string `json:"filteredContent"`
	Results         []struct {
		ScanCommonResultResp
		Extras  map[string]string `json:"extras,omitempty"`
		Details []struct {
			Label    string `json:"label,omitempty"`
			Contexts []struct {
				Context   string           `json:"context,omitempty"`
				Positions []map[string]int `json:"positions,omitempty"`
				LibName   string           `json:"libName,omitempty"`
				LibCode   string           `json:"libCode,omitempty"`
				RuleType  string           `json:"ruleType,omitempty"`
			} `json:"contexts,omitempty"`
		}
	} `json:"results,omitempty"`
}

// // ScanTextResult 文本检测结果
// type ScanTextResult struct {
// 	ScanCommonResultResp
// 	Extras  map[string]string `json:"extras,omitempty"`
// 	Details []struct {
// 		Label    string `json:"label,omitempty"`
// 		Contexts []struct {
// 			Context   string           `json:"context,omitempty"`
// 			Positions []map[string]int `json:"positions,omitempty"`
// 			LibName   string           `json:"libName,omitempty"`
// 			LibCode   string           `json:"libCode,omitempty"`
// 			RuleType  string           `json:"ruleType,omitempty"`
// 		} `json:"contexts,omitempty"`
// 	} `json:"details,omitempty"`
// }

// ScanText 文本内容安全检测
func (c Client) ScanText(in *ScanTextReq) (result *ScanTextResp, err error) {
	resp, err := c.Do(http.MethodPost, TEXT_API_PATH, in)
	if err != nil {
		return nil, err
	}
	result = &ScanTextResp{}
	return result, interfaceConvert(resp, result)
}
