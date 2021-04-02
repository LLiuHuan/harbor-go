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

type RepositoryListOptions struct {
	ProjectName string
	Q           string
	Sort        string
	Page        string
	PageSize    string
}

type RepositoryGetOptions struct {
	ProjectName    string
	RepositoryName string
}
