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
	CVEAllowList CVEAllowlist      `json:"cve_allowlist"`
	RegistryID   int64             `json:"registry_id"`
}

type ProjectListOptions struct {
	Page     string `json:"page"`
	PageSize string `json:"page_size"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Public   *bool  `json:"public" query:"true"`
}

type PostProjectMetadataOptions struct {
	ProjectID int64                  `json:"project_id" query:"false"`
	Metadata  ProjectMetadataOptions `json:"metadata"`
}

type ProjectMetadataOptions struct {
	EnableContentTrust   string `json:"enable_content_trust"`
	AutoScan             string `json:"auto_scan"`
	Severity             string `json:"severity"`
	Public               string `json:"public"`
	ReuseSysCveAllowList string `json:"reuse_sys_cve_allowlist"`
	PreventVul           string `json:"prevent_vul"`
	RetentionId          string `json:"retention_id"`
}

type GetProjectMetadataOptions struct {
	ProjectID int64 `json:"project_id"`
}

type HeadProjects struct {
	// header
	XRequestID string `json:"X-Request-Id" header:"true" default:""`

	// query
	ProjectName string `json:"project_name" query:"true"`
}

type Metadata struct {
	EnableContentTrust   string `json:"enable_content_trust"`
	AutoScan             string `json:"auto_scan"`
	Severity             string `json:"severity"`
	Public               string `json:"public" body:"true"`
	ReuseSysCveAllowList string `json:"reuse_sys_cve_allowlist"`
	PreventVul           string `json:"prevent_vul"`
	RetentionId          string `json:"retention_id"`
}

type CveAllowListItems struct {
	CveId string `json:"cve_id" query:"true"`
}

type CveAllowList struct {
	Items        []CveAllowListItems `json:"items"`
	ProjectId    int                 `json:"project_id"`
	Id           int                 `json:"id"`
	ExpiresAt    int                 `json:"expires_at"`
	UpdateTime   time.Time           `json:"update_time"`
	CreationTime time.Time           `json:"creation_time"`
}

type PostProjects struct {
	// header
	XRequestID              string `json:"X-Request-Id" header:"true" default:""`
	XResourceNameInLocation string `json:"X-Resource-Name-In-Location" header:"true" default:"false"`

	// body
	ProjectName  string       `json:"project_name" body:"true"`
	CveAllowList CveAllowList `json:"cve_allowlist"`
	CountLimit   int          `json:"count_limit"`
	RegistryId   int          `json:"registry_id"`
	StorageLimit int          `json:"storage_limit" body:"true"`
	Metadata     Metadata     `json:"metadata" body:"true"`
	Public       bool         `json:"public" body:"true"`
}

type DeleteProjects struct {
	// header
	XRequestID              string `json:"X-Request-Id" header:"true" default:""`
	XResourceNameInLocation string `json:"X-Resource-Name-In-Location" header:"true" default:"false"`

	// path
	ProjectNameOrId string `json:"project_name_or_id" path:"true"`
}
