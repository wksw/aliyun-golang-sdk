package contentsecurity

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randSource = rand.NewSource(time.Now().UnixNano())

// md5加密
func newBase64Md5(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// 生成随机数
func ranString(n int) string {
	b := make([]byte, n)
	// A randSource.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, randSource.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSource.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// hmac-sha1签名
func newBase64HmacSha1(key string, data []byte) (string, error) {
	h := hmac.New(sha1.New, []byte(key))
	if _, err := h.Write([]byte(data)); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil

}

// 用json反序列化、序列化把interface转成指定struct类型
func interfaceConvert(in, out interface{}) error {
	d, err := json.Marshal(in)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(d, out); err != nil {
		return err
	}
	return nil
}
