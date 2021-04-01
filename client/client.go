package client

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// BasicAuth
//  Username 用户名
//	Password 密码
type BasicAuth struct {
	Username string
	Password string
}

//	Client 客户端结构体
//	client    http客户端
//	host      Ip
//	hostname  Ip/域名 名称
//	basicAuth *BasicAuth  用户名密码
//	version   接口版本  目前只实现了v2
//	scheme    http/https
//	basePath  默认接口地址
type Client struct {
	client    *http.Client
	host      string
	hostname  string
	basicAuth *BasicAuth
	version   string
	scheme    string
	basePath  string
}

// NewClientWithOpts 初始化客户端
func NewClientWithOpts(opts ...Opt) (*Client, error) {
	c := &Client{
		client:    http.DefaultClient,
		host:      DefaultHarborHost,
		basicAuth: nil,
		version:   Version,
		scheme:    "http",
		basePath:  "/api",
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// 返回客户端副本
func (cli *Client) HTTPClient() *http.Client {
	return &*cli.client
}

// 返回服务器地址
func (cli *Client) Host() string {
	return cli.host
}

// 返回http/https
func (cli *Client) Scheme() string {
	return cli.scheme
}

// 返回接口版本
func (cli *Client) Version() string {
	return cli.version
}

// 拼接请求地址
func (cli *Client) getAPIPath(ctx context.Context, p string, query url.Values) string {
	var apiPath string
	if cli.version != "" {
		v := strings.TrimPrefix(cli.version, "v")
		apiPath = path.Join(cli.basePath, "/v"+v, p)
	} else {
		apiPath = path.Join(cli.basePath, p)
	}
	return (&url.URL{Path: apiPath, RawQuery: query.Encode()}).String()
}
