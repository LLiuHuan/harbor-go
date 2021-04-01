package client

import (
	"context"
	"encoding/json"
	"github.com/lliuhuan/harbor-go/schema"
	"net/url"
)

const (
	PATH_FMT_LIST_PROJECT = "/projects"
)

func (cli *Client) ListProjects(ctx context.Context, options schema.ProjectListOptions) ([]schema.Project, error) {
	var projects []schema.Project

	return projects, nil
}
