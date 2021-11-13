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

type GetProjectArtifactsReference struct {
	// header
	XRequestID             string `json:"X-Request-Id" header:"true" default:""`
	XAcceptVulnerabilities string `json:"x_accept_vulnerabilities" header:"true" default:""`

	// path
	ProjectName    string `json:"project_name" path:"true"`
	RepositoryName string `json:"repository_name" path:"true"`
	Reference      string `json:"reference" path:"true"`

	// query
	Page                int64 `json:"page" query:"true"`
	PageSize            int64 `json:"page_size" query:"true"`
	WithTag             bool  `json:"with_tag" query:"true"`
	WithLabel           bool  `json:"with_label" query:"true"`
	WithScanOverview    bool  `json:"with_scan_overview" query:"true"`
	WithSignature       bool  `json:"with_signature" query:"true"`
	WithImmutableStatus bool  `json:"with_immutable_status" query:"true"`
}

type ProjectArtifactsReference struct {
	Size         int       `json:"size"`
	PushTime     time.Time `json:"push_time"`
	ScanOverview struct {
		AdditionalProp1 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  time.Time `json:"start_time"`
			ScanStatus string    `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int       `json:"complete_percent"`
			EndTime         time.Time `json:"end_time"`
			Duration        int       `json:"duration"`
			ReportId        string    `json:"report_id"`
			Severity        string    `json:"severity"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  time.Time `json:"start_time"`
			ScanStatus string    `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int       `json:"complete_percent"`
			EndTime         time.Time `json:"end_time"`
			Duration        int       `json:"duration"`
			ReportId        string    `json:"report_id"`
			Severity        string    `json:"severity"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  time.Time `json:"start_time"`
			ScanStatus string    `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int       `json:"complete_percent"`
			EndTime         time.Time `json:"end_time"`
			Duration        int       `json:"duration"`
			ReportId        string    `json:"report_id"`
			Severity        string    `json:"severity"`
		} `json:"additionalProp3"`
	} `json:"scan_overview"`
	Tags []struct {
		RepositoryId int       `json:"repository_id"`
		Name         string    `json:"name"`
		PushTime     time.Time `json:"push_time"`
		PullTime     time.Time `json:"pull_time"`
		Signed       bool      `json:"signed"`
		Id           int       `json:"id"`
		Immutable    bool      `json:"immutable"`
		ArtifactId   int       `json:"artifact_id"`
	} `json:"tags"`
	PullTime time.Time `json:"pull_time"`
	Labels   []struct {
		UpdateTime   time.Time `json:"update_time"`
		Description  string    `json:"description"`
		Color        string    `json:"color"`
		CreationTime time.Time `json:"creation_time"`
		Deleted      bool      `json:"deleted"`
		Scope        string    `json:"scope"`
		ProjectId    int       `json:"project_id"`
		Id           int       `json:"id"`
		Name         string    `json:"name"`
	} `json:"labels"`
	References []struct {
		Platform struct {
			Os           string   `json:"os"`
			Variant      string   `json:"variant"`
			Architecture string   `json:"architecture"`
			OsFeatures   []string `json:"'os.features'"`
			OsVersion    string   `json:"'os.version'"`
		} `json:"platform"`
		ChildDigest string   `json:"child_digest"`
		Urls        []string `json:"urls"`
		ParentId    int      `json:"parent_id"`
		ChildId     int      `json:"child_id"`
		Annotations struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"annotations"`
	} `json:"references"`
	ManifestMediaType string `json:"manifest_media_type"`
	ExtraAttrs        struct {
		AdditionalProp1 struct {
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
		} `json:"additionalProp3"`
	} `json:"extra_attrs"`
	Id            int    `json:"id"`
	Digest        string `json:"digest"`
	Icon          string `json:"icon"`
	RepositoryId  int    `json:"repository_id"`
	AdditionLinks struct {
		AdditionalProp1 struct {
			Href     string `json:"href"`
			Absolute bool   `json:"absolute"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Href     string `json:"href"`
			Absolute bool   `json:"absolute"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Href     string `json:"href"`
			Absolute bool   `json:"absolute"`
		} `json:"additionalProp3"`
	} `json:"addition_links"`
	MediaType   string `json:"media_type"`
	ProjectId   int    `json:"project_id"`
	Type        string `json:"type"`
	Annotations struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"annotations"`
}
