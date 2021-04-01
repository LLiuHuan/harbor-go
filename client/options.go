package client

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	Version           = "v2.0"
	DefaultHarborHost = "http://localhost"
)

type Opt func(*Client) error

// WithHost 写入Host地址
func WithHost(host string) Opt {
	return func(c *Client) error {
		if strings.HasSuffix(host, "/") {
			host = host[:len(host)-1]
		}
		//c.host = host

		hostURL, err := ParseHostURL(host)
		if err != nil {
			return err
		}
		c.host = hostURL.Host
		c.scheme = hostURL.Scheme

		//c.basePath = host
		return nil
	}
}

// ParseHostURL 解析url字符串，验证该字符串是否为主机url 返回解析后的URL
func ParseHostURL(host string) (*url.URL, error) {
	protoAddrParts := strings.SplitN(host, "://", 2)
	if len(protoAddrParts) == 1 {
		return nil, fmt.Errorf("unable to parse harbor host `%s`", host)
	}

	proto, addr := protoAddrParts[0], protoAddrParts[1]
	return &url.URL{
		Scheme: proto,
		Host:   addr,
	}, nil
}

// WithScheme 修改 http/https
func WithScheme(scheme string) Opt {
	return func(c *Client) error {
		c.scheme = scheme
		return nil
	}
}

// WithBasicAuth 设置用户名密码
func WithBasicAuth(username, password string) Opt {
	return func(c *Client) error {
		c.basicAuth = &BasicAuth{Username: username, Password: password}
		return nil
	}
}
