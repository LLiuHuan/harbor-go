package schema

type GetLdapGroupSearchOptions struct {
	GroupName string `json:"groupname"`
	GroupDN   string `json:"groupdn"`
}

type GetLdapInfo struct {
	GroupName   string `json:"group_name"`
	LdapGroupDN string `json:"ldap_group_dn"`
	GroupType   int64  `json:"group_type"`
	ID          int64  `json:"id"`
}
