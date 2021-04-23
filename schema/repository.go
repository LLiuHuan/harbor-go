package schema

import "time"

type Repository struct {
	ID            int64     `json:"id"`
	ProjectID     int64     `json:"project_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	ArtifactCount int64     `json:"artifact_count"`
	PullCount     int64     `json:"pull_count"`
	CreationTime  time.Time `json:"creation_time"`
	UpdateTime    time.Time `json:"update_time"`
}

type GetRepositoryListOptions struct {
	ProjectName string `json:"project_name" path:"true"`
	Q           string `json:"q" query:"true"`
	Sort        string `json:"sort" query:"true"`
	Page        string `json:"page" query:"true"`
	PageSize    string `json:"page_size" query:"true"`
}

type GetRepositoryByNameOptions struct {
	ProjectName    string `json:"project_name" path:"true"`
	RepositoryName string `json:"repository_name" path:"true"`
}

type PutRepositoryByNameOptions struct {
	ProjectName    string      `json:"project_name" path:"true"`
	RepositoryName string      `json:"repository_name" path:"true"`
	Repository     *Repository `json:"repository" query:"true"`
}

type DelRepositoryByNameOptions struct {
	ProjectName    string `json:"project_name" path:"true"`
	RepositoryName string `json:"repository_name" path:"true"`
}
