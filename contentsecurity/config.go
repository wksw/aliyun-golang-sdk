package contentsecurity

import "time"

// ClientConfig 客户端配置
type ClientConfig struct {
	// 接口地址
	endpoint string
	// 阿里云accessKey,用于接口调用身份识别, 联系管理员获取
	accessKey string
	// 阿里云accessSecretKey, 用于接口签名, 联系管理员获取
	secretKey string
	// 代理地址
	proxy string
	// 代理用户名
	proxyUser string
	// 代理密码
	proxyPassword string
	// 超时时间, 默认10秒
	timeout time.Duration
}
