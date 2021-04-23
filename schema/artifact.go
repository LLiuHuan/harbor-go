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
