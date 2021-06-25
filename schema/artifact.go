package schema

import "time"

type GetProjectArtifactsListOptions struct {
	// header
	XRequestID             string `json:"X-Request-Id" header:"true" default:""`
	XAcceptVulnerabilities string `json:"X-Accept-Vulnerabilities" header:"true" default:"application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0"`
	// path
	ProjectName    string `json:"project_name" path:"true"`
	RepositoryName string `json:"repository_name" path:"true"`
	// query
	Q                   string `json:"q" query:"true" default:""`
	Page                int64  `json:"page" query:"true" default:"1"`
	PageSize            int64  `json:"page_size" query:"true" default:"10"`
	WithTag             bool   `json:"with_tag" query:"true" default:"true"`
	WithLabel           bool   `json:"with_label" query:"true" default:"false"`
	WithScanOverview    bool   `json:"with_scan_overview" query:"true" default:"false"`
	WithSignature       bool   `json:"with_signature" query:"true" default:"false"`
	WithImmutableStatus bool   `json:"with_immutable_status" query:"true" default:"false"`
}

type GetProjectArtifactsList struct {
	AdditionLinks struct {
		BuildHistory struct {
			Absolute bool   `json:"absolute"`
			Href     string `json:"href"`
		} `json:"build_history"`
	} `json:"addition_links"`
	Digest     string `json:"digest"`
	ExtraAttrs struct {
		Architecture string `json:"architecture"`
		Author       string `json:"author"`
		Config       struct {
			Cmd          []string `json:"Cmd"`
			Entrypoint   []string `json:"Entrypoint"`
			Env          []string `json:"Env"`
			ExposedPorts struct {
				Tcp struct {
				} `json:"80/tcp"`
			} `json:"ExposedPorts"`
			Labels struct {
				Maintainer string `json:"maintainer"`
			} `json:"Labels"`
			StopSignal string `json:"StopSignal"`
		} `json:"config"`
		Created time.Time `json:"created"`
		Os      string    `json:"os"`
	} `json:"extra_attrs"`
	Icon              string      `json:"icon"`
	Id                int         `json:"id"`
	Labels            interface{} `json:"labels"`
	ManifestMediaType string      `json:"manifest_media_type"`
	MediaType         string      `json:"media_type"`
	ProjectId         int         `json:"project_id"`
	PullTime          time.Time   `json:"pull_time"`
	PushTime          time.Time   `json:"push_time"`
	References        interface{} `json:"references"`
	RepositoryId      int         `json:"repository_id"`
	Size              int         `json:"size"`
	Tags              []struct {
		ArtifactId   int       `json:"artifact_id"`
		Id           int       `json:"id"`
		Immutable    bool      `json:"immutable"`
		Name         string    `json:"name"`
		PullTime     time.Time `json:"pull_time"`
		PushTime     time.Time `json:"push_time"`
		RepositoryId int       `json:"repository_id"`
		Signed       bool      `json:"signed"`
	} `json:"tags"`
	Type string `json:"type"`
}

type GetProjectArtifactsTagsOptions struct {
	// header
	XRequestID             string `json:"X-Request-Id" header:"true" default:""`
	XAcceptVulnerabilities string `json:"X-Accept-Vulnerabilities" header:"true" default:"application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0"`
	// path
	ProjectName    string `json:"project_name" path:"true"`
	RepositoryName string `json:"repository_name" path:"true"`
	Reference      string `json:"reference" path:"true"`
	// query
	Page                int64 `json:"page" query:"true" default:"1"`
	PageSize            int64 `json:"page_size" query:"true" default:"10"`
	WithTag             bool  `json:"with_tag" query:"true" default:"true"`
	WithLabel           bool  `json:"with_label" query:"true" default:"false"`
	WithScanOverview    bool  `json:"with_scan_overview" query:"true" default:"false"`
	WithSignature       bool  `json:"with_signature" query:"true" default:"false"`
	WithImmutableStatus bool  `json:"with_immutable_status" query:"true" default:"false"`
}

type GetProjectArtifactsTags struct {
	RepositoryId int       `json:"repository_id"`
	Name         string    `json:"name"`
	PushTime     time.Time `json:"push_time"`
	PullTime     time.Time `json:"pull_time"`
	Signed       bool      `json:"signed"`
	Id           int       `json:"id"`
	Immutable    bool      `json:"immutable"`
	ArtifactId   int       `json:"artifact_id"`
}

type GetProjectArtifactsInfoOptions struct {
	// header
	XRequestID string `json:"X-Request-Id" header:"true" default:""`
	// path
	ProjectName    string `json:"project_name" path:"true"`
	RepositoryName string `json:"repository_name" path:"true"`
	Reference      string `json:"reference" path:"true"`
	// query
	Q                   string `json:"q" query:"true" default:""`
	Page                int64  `json:"page" query:"true" default:"1"`
	PageSize            int64  `json:"page_size" query:"true" default:"10"`
	WithSignature       bool   `json:"with_signature" query:"true" default:"false"`
	WithImmutableStatus bool   `json:"with_immutable_status" query:"true" default:"false"`
}

type GetProjectArtifactsInfo struct {
	Size         int       `json:"size"`
	PushTime     time.Time `json:"push_time"`
	ScanOverview struct {
		AdditionalProp1 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  string `json:"start_time"`
			ScanStatus string `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int    `json:"complete_percent"`
			EndTime         string `json:"end_time"`
			Duration        int    `json:"duration"`
			ReportId        string `json:"report_id"`
			Severity        string `json:"severity"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  string `json:"start_time"`
			ScanStatus string `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int    `json:"complete_percent"`
			EndTime         string `json:"end_time"`
			Duration        int    `json:"duration"`
			ReportId        string `json:"report_id"`
			Severity        string `json:"severity"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Scanner struct {
				Version string `json:"version"`
				Vendor  string `json:"vendor"`
				Name    string `json:"name"`
			} `json:"scanner"`
			StartTime  string `json:"start_time"`
			ScanStatus string `json:"scan_status"`
			Summary    struct {
				Fixable int `json:"fixable"`
				Total   int `json:"total"`
				Summary struct {
					High     int `json:"High"`
					Critical int `json:"Critical"`
				} `json:"summary"`
			} `json:"summary"`
			CompletePercent int    `json:"complete_percent"`
			EndTime         string `json:"end_time"`
			Duration        int    `json:"duration"`
			ReportId        string `json:"report_id"`
			Severity        string `json:"severity"`
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

type GetProjectArtifactsAdditionsOptions struct {
	// header
	XRequestID string `json:"X-Request-Id" header:"true" default:""`
	// path
	ProjectName    string `json:"project_name" path:"true"`
	RepositoryName string `json:"repository_name" path:"true"`
	Reference      string `json:"reference" path:"true"`
	Addition       string `json:"addition" path:"true"`
}

type GetProjectArtifactsAdditionsBuildHistory struct {
	Created    time.Time `json:"created"`
	CreatedBy  string    `json:"created_by"`
	EmptyLayer bool      `json:"empty_layer,omitempty"`
}
