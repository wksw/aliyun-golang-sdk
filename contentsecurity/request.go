package contentsecurity

import "net/http"

// Request 统一的一个请求
func (c Client) Request(path string, in, out interface{}) error {
	resp, err := c.Do(http.MethodPost, path, in)
	if err != nil {
		return err
	}
	return interfaceConvert(resp, out)
}
