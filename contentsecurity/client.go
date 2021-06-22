package contentsecurity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"sort"
	"strings"
	"time"
)

const (
	// NEED_SIGN_HEAD_PREFIX 需要签名的请求头前缀
	NEED_SIGN_HEAD_PREFIX = "x-acs-"
	// CONTENT_TYPE 请求数据格式
	CONTENT_TYPE = "application/json"
	// SIGN_NEW_LINE 签名换行符号
	SIGN_NEW_LINE = "\n"
	// SDK_VERSION sdk版本
	SDK_VERSION = "1.0"
	// API_VERSION 接口版本
	API_VERSION = "2018-05-09"
	// SIGNATURE_VERSION 签名版本
	SIGNATURE_VERSION = "1.0"
	// SIGNATURE_METHOD 签名方式
	SIGNATURE_METHOD = "HMAC-SHA1"
)

// Client Http请求客户端
type Client struct {
	client     *http.Client
	clientInfo *ClientInfo
	config     *ClientConfig
}

// ContentSecurityCommonResp 内容安全公共返回返回
// 参考https://help.aliyun.com/document_detail/53414.html?spm=a2c4g.11186623.6.623.5093245eypvnRE
type ContentSecurityCommonResp struct {
	// 时间、签名、md5值错误才会出现这个值
	Code1 string `json:"Code,omitempty"`
	// 错误码
	Code int `json:"code"`
	// 错误描述
	Message string `json:"msg,omitempty"`
	// 和Code1功能保持一致
	Message1 string `json:"Message,omitempty"`
	// 请求唯一标识符
	RequestId string `json:"requestId,omitempty"`
}

// ContentSecurityResp 结果返回
type ContentSecurityResp struct {
	ContentSecurityCommonResp
	// 返回内容
	Data interface{}
}

// New 创建http请求客户端
func New(endpoint, accessKey, secretKey string) Client {
	return Client{
		client: http.DefaultClient,
		config: &ClientConfig{
			endpoint:  endpoint,
			accessKey: accessKey,
			secretKey: secretKey,
			timeout:   10 * time.Second,
		},
		clientInfo: &ClientInfo{
			SDKVersion: SDK_VERSION,
			OS:         runtime.GOOS,
		},
	}
}

// NewWithConfig 根据配置创建http请求客户端
func NewWithConfig(config *ClientConfig) Client {
	return Client{
		config: config,
	}
}

// WithProxy 设置代理
func (c *Client) WithProxy(proxy, user, password string) error {
	c.config.proxy = proxy
	c.config.proxyUser = user
	c.config.proxyPassword = password
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return fmt.Errorf("parse proxy host fail[%s]", err.Error())
	}
	proxyURL.User = url.UserPassword(user, password)
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	c.client.Transport = transport
	return nil
}

// WithTimeout 设置超时时间
func (c *Client) WithTimeout(duration time.Duration) {
	c.config.timeout = duration
	c.client.Timeout = duration
}

// WithEndpoint 设置接口地址
func (c *Client) WithEndpoint(endpoint string) {
	c.config.endpoint = endpoint
}

// WithAkSk 设置aksk
func (c *Client) WithAkSk(accesskey, secretkey string) {
	c.config.accessKey = accesskey
	c.config.secretKey = secretkey
}

// Config 获取客户端配置
func (c *Client) Config() *ClientConfig {
	return c.config
}

// Do Http请求发送
func (c *Client) Do(method, path string, body interface{}) (*ContentSecurityResp, error) {
	bodyByte, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal request body fail[%s]", err.Error())
	}
	uri, err := url.ParseRequestURI(c.config.endpoint + path)
	if err != nil {
		return nil, fmt.Errorf("parse request host and request path fail[%s]", err.Error())
	}
	uri.RawQuery = ""
	clientInfo, err := json.Marshal(c.clientInfo)
	if err != nil {
		return nil, fmt.Errorf("marshal clientInfo fail[%s]", err.Error())
	}

	req, err := http.NewRequest(method, uri.String()+"?clientInfo="+url.QueryEscape(string(clientInfo)), bytes.NewReader(bodyByte))
	if err != nil {
		return nil, fmt.Errorf("create request fail[%s]", err.Error())
	}
	req.Header.Add("Accept", CONTENT_TYPE)
	req.Header.Add("Content-Type", CONTENT_TYPE)
	req.Header.Add("Content-Md5", newBase64Md5(bodyByte))
	req.Header.Add("Date", time.Now().UTC().Format(http.TimeFormat))
	req.Header.Add("x-acs-version", API_VERSION)
	req.Header.Add("x-acs-signature-nonce", ranString(10))
	req.Header.Add("x-acs-signature-version", SIGNATURE_VERSION)
	req.Header.Add("x-acs-signature-method", SIGNATURE_METHOD)
	// 请求签名
	signStr, err := c.signature(req)
	if err != nil {
		return nil, fmt.Errorf("signature request fail[%s]", err.Error())
	}
	// 签名按照指定格式放在请求头中
	req.Header.Add("Authorization", "acs"+" "+c.config.accessKey+":"+signStr)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request remote server fail[%s]", err.Error())
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body fail[%s]", err.Error())
	}
	var respData ContentSecurityResp
	if err := json.Unmarshal(respBody, &respData); err != nil {
		return nil, fmt.Errorf("unmarshal response body fail[%s]", err.Error())
	}
	// 签名之类的错误会出现这个错误码
	if respData.Code1 != "" {
		err = fmt.Errorf("(%s)%s", respData.Code1, respData.Message1)
	}
	// 非2xx状态码
	if err == nil && respData.Code/100 != 2 {
		err = fmt.Errorf("(%d)%s", respData.Code, respData.Message)
	}
	return &respData, err
}

// 请求签名
// 参考https://help.aliyun.com/document_detail/53415.html?spm=a2c4g.11186623.6.624.4fa120d8muS0Ng
func (c *Client) signature(request *http.Request) (string, error) {
	// 序列化请求头
	var needSignHeadKeys []string
	for key := range request.Header {
		key = strings.ToLower(key)
		if strings.HasPrefix(key, NEED_SIGN_HEAD_PREFIX) {
			needSignHeadKeys = append(needSignHeadKeys, key)
		}
	}
	// 排序
	sort.Strings(needSignHeadKeys)

	/*
		"POST\n" +
		"application/json\n" +
		"HTTP头Content-MD5的值" + "\n" +
		"application/json" + "\n" +
		"HTTP头Date的值" + "\n" +
		"序列化请求头" +
		"序列化uri和query参数"

	*/
	var signBuffer bytes.Buffer
	signBuffer.WriteString(request.Method)
	signBuffer.WriteString(SIGN_NEW_LINE)
	signBuffer.WriteString(CONTENT_TYPE)
	signBuffer.WriteString(SIGN_NEW_LINE)
	signBuffer.WriteString(request.Header.Get("Content-MD5"))
	signBuffer.WriteString(SIGN_NEW_LINE)
	signBuffer.WriteString(CONTENT_TYPE)
	signBuffer.WriteString(SIGN_NEW_LINE)
	signBuffer.WriteString(request.Header.Get("Date"))
	signBuffer.WriteString(SIGN_NEW_LINE)
	// 待签名字符串添加头域
	for _, key := range needSignHeadKeys {
		signBuffer.WriteString(key + ":" + strings.TrimSpace(request.Header.Get(key)))
		signBuffer.WriteString(SIGN_NEW_LINE)
	}

	signBuffer.WriteString(request.URL.Path + "?clientInfo=" + request.URL.Query().Get("clientInfo"))
	signStr, err := newBase64HmacSha1(c.config.secretKey, signBuffer.Bytes())
	return signStr, err
}
