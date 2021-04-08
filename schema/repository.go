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
	ProjectName string `json:"project_name" not:"true"`
	Q           string `json:"q"`
	Sort        string `json:"sort"`
	Page        string `json:"page"`
	PageSize    string `json:"page_size"`
}

type GetRepositoryByNameOptions struct {
	ProjectName    string `json:"project_name" not:"true"`
	RepositoryName string `json:"repository_name" not:"true"`
}

type PutRepositoryByNameOptions struct {
	ProjectName    string      `json:"project_name" not:"true"`
	RepositoryName string      `json:"repository_name" not:"true"`
	Repository     *Repository `json:"repository"`
}

type DelRepositoryByNameOptions struct {
	ProjectName    string `json:"project_name" not:"true"`
	RepositoryName string `json:"repository_name" not:"true"`
}
