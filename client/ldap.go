package client

import (
	"context"
	"encoding/json"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_GET_LDAP_GROUP_SEARCH = "/ldap/groups/search"
)

/**
GetLdapGroupSearch 该端点根据相关的配置参数搜索可用的ldap组。支持按groupname或groupdn搜索。
url: /ldap/groups/search
*/
// TODO: 未测试
func (cli *Client) GetLdapGroupSearch(ctx context.Context, options schema.GetLdapGroupSearchOptions) (resp []schema.GetLdapInfo, err error) {
	query := StructQuery(options)

	serverResp, err := cli.get(ctx, PATH_GET_LDAP_GROUP_SEARCH, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return resp, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&resp)

	return resp, nil
}
