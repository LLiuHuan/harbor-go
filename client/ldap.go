package client

import (
	"context"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_GET_LDAP_GROUP_SEARCH = "/ldap/groups/search"
)

func (cli *Client) GetLdapGroupSearch(ctx context.Context, options schema.GetLdapGroupSearchOptions) {

}
