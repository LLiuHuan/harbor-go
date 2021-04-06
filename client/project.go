package client

import (
	"context"
	"encoding/json"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_LIST_PROJECT = "/projects"
)

func (cli *Client) ListProjects(ctx context.Context, options schema.ProjectListOptions) ([]schema.Project, error) {
	var projects []schema.Project

	query := StructQuery(options)
	if options.Public != nil {
		if *options.Public {
			query.Set("public", "1")
		} else {
			query.Set("public", "0")
		}
	}

	serverResp, err := cli.get(ctx, PATH_LIST_PROJECT, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return projects, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&projects)

	return projects, nil
}
