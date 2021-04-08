package schema

type PutUserCliOptions struct {
	UserID      int         `json:"user_id"`
	InputSecret InputSecret `json:"input_secret"`
}
type InputSecret struct {
	Secret string `json:"secret"`
}

type GetUserPermissionOptions struct {
	Scope    string `json:"scope"`
	Relative bool   `json:"relative"`
}

type GetUserPermission struct {
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

type PutUserInfoOptions struct {
	UserID  int         `json:"user_id"`
	Profile UserProfile `json:"profile"`
}

type UserProfile struct {
	Comment  string `json:"comment"`
	Email    string `json:"email"`
	RealName string `json:"realname"`
}

type GetUserInfoOptions struct {
	UserID int `json:"user_id"`
}

type GetUserInfo struct {
	UserName        string `json:"username"`
	Comment         string `json:"comment"`
	UpdateTime      string `json:"update_time"`
	PassWord        string `json:"password"`
	UserID          int    `json:"user_id"`
	RealName        string `json:"realname"`
	Deleted         bool   `json:"deleted"`
	CreateTime      string `json:"create_time"`
	AdminRoleInAuth bool   `json:"admin_role_in_auth"`
	RoleID          int    `json:"role_id"`
	SysAdminFlag    bool   `json:"sysadmin_flag"`
	RoleName        string `json:"role_name"`
	ResetUUID       string `json:"reset_uuid"`
	Salt            string `json:"Salt"`
	Email           string `json:"email"`
}

type DelUserInfoOptions struct {
	UserID int `json:"user_id"`
}

type GetUserSearchByNameOptions struct {
	UserName string `json:"username"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"page_size"`
}

type GetUserSearchByName struct {
	UserName string `json:"username"`
	UserID   int    `json:"user_id"`
}
