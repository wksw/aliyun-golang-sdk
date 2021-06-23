package contentsecurity

const (
	// TEXT_API_PATH 文本内容检测API路径
	TEXT_API_PATH = "/green/text/scan"
	// IMAGE_SYNC_API_PATH 图片内容检测API路径
	IMAGE_SYNC_API_PATH = "/green/image/scan"
	// IMAGE_ASYNC_API_PATH 图片异步检测API路径
	IMAGE_ASYNC_API_PATH = "/green/image/asyncscan"
	// IMAGE_ASYNC_RESULT_API_PATH 图片异步检测结果API路径
	IMAGE_ASYNC_RESULT_API_PATH = "/green/image/results"
	// VIDEO_SYNC_API_PATH 视频同步检测API路径
	VIDEO_SYNC_API_PATH = "/green/video/syncscan"
	// VIDEO_ASYNC_API_PATH 视频异步检测API路径
	VIDEO_ASYNC_API_PATH = "/green/video/asyncscan"
	// VIDEO_CANCEL_API_PATH 取消视频检测
	VIDEO_CANCEL_API_PATH = "/green/video/cancelscan"
	// VIDEO_RESULT_API_PATH 视频检测结果
	VIDEO_RESULT_API_PATH = "/green/video/results"
	// WEBPAGE_SYNC_API_PATH 网页内容同步检测api路径
	WEBPAGE_SYNC_API_PATH = "/green/webpage/scan"
	// WEBPAGE_ASYNC_API_PATH 网页内容异步检测api路径
	WEBPAGE_ASYNC_API_PATH = "/green/webpage/asyncscan"
	// WEBPAGE_RESULT_API_PATH 网页内容检测结果
	WEBPAGE_RESULT_API_PATH = "/green/webpage/results"
)

const (
	// UserOther 其他用户
	UserOther string = "others"
	// UserTaobao 淘宝用户
	UserTaobao string = "taobao"
)

const (
	// TextSceneAntispam 文本检测反垃圾场景
	TextSceneAntispam string = "antispam"
	// ImageScenePorn 智能鉴黄
	ImageScenePorn string = "porn"
	// ImageSceneTerrorism 暴恐涉政
	ImageSceneTerrorism string = "terrorism"
	// ImageSceneAd 图文违规
	ImageSceneAd string = "ad"
	// ImageSceneQrcode 二维码
	ImageSceneQrcode string = "qrcode"
	// ImageSceneLive 不良场景
	ImageSceneLive string = "live"
	// ImageSceneLogo 图片log
	ImageSceneLogo string = "logo"
)

// ClientInfo 客户端详情
// 参考https://help.aliyun.com/document_detail/53413.html?spm=a2c4g.11186623.6.622.78be2ba1jPje6N
type ClientInfo struct {
	// SDK 版本号
	SDKVersion string `json:"sdkVersion,omitempty"`
	// 配置信息版本
	CfgVersion string `json:"cfgVersion,omitempty"`
	// 用户类型
	UserType string `json:"userType,omitempty"`
	// 用户唯一标识符
	UserId string `json:"userId,omitempty"`
	// 用户昵称
	UserNick string `json:"userNick,omitempty"`
	// 用户头像
	Avatar string `json:"avatar,omitempty"`
	// 硬件设备码
	Imei string `json:"imei,omitempty"`
	// 运营商设备码
	Imsi string `json:"imsi,omitempty"`
	// 设备指纹
	Umid string `json:"umid,omitempty"`
	// 客户端公网IP地址
	IP string `json:"ip,omitempty"`
	// 设备操作系统
	OS string `json:"os,omitempty"`
	// 渠道号
	Channel string `json:"channel,omitempty"`
	// 宿主应用名称
	HostAppName string `json:"hostAppName,omitempty"`
	// 宿主应用包
	HostPackage string `json:"hostPackage,omitempty"`
	// 宿主应用版本
	HostVersion string `json:"hostVersion,omitempty"`
}

// ScanCommonReq 检测公共请求
type ScanCommonReq struct {
	// 业务场景
	BizType string `json:"bizType"`
	// 业务场景
	Scenes []string `json:"scenes,omitempty" validate:"required"`
}

// ScanCommonDataResp 检测公共返回
type ScanCommonDataResp struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	DataId  string `json:"dataId"`
	TaskId  string `json:"taskId"`
}

// ScanCommonResultResp 检测公共结果
type ScanCommonResultResp struct {
	Scene      string  `json:"scene"`
	Suggestion string  `json:"suggestion"`
	Label      string  `json:"label"`
	SubLabel   string  `json:"sublabel,omitempty"`
	Rate       float32 `json:"rate"`
}

// ScanExtras 附加信息
type ScanExtras struct {
	HitLibInfo []struct {
		Context string `json:"context"`
		LibCode string `json:"libCode"`
		LibName string `json:"libName"`
	} `json:"hitLibInfo,omitempty"`
	NewFramePrefix string `json:"newFramePrefix,omitempty"`
	NewFrames      string `json:"newFrames,omitempty"`
}

// ScanCommonAsyncReq 异步请求
type ScanCommonAsyncReq struct {
	Callback  string `json:"callback"`
	Seed      string `json:"seed"`
	CryptType string `json:"cryptType"`
}
