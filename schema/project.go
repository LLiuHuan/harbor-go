package schema

import "time"

type Project struct {
	ProjectID    int64             `json:"project_id"`
	OwnerID      int               `json:"owner_id"`
	Name         string            `json:"name"`
	CreationTime time.Time         `json:"creation_time"`
	UpdateTime   time.Time         `json:"update_time"`
	Deleted      bool              `json:"deleted"`
	OwnerName    string            `json:"owner_name"`
	Role         int               `json:"current_user_role_id"`
	RoleList     []int             `json:"current_user_role_ids"`
	RepoCount    int64             `json:"repo_count"`
	ChartCount   uint64            `json:"chart_count"`
	Metadata     map[string]string `json:"metadata"`
	CVEAllowlist CVEAllowlist      `json:"cve_allowlist"`
	RegistryID   int64             `json:"registry_id"`
}

type ProjectListOptions struct {
	Page     string `json:"page"`
	PageSize string `json:"page_size"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Public   *bool  `json:"public" not:"true"`
}
