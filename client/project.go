package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_LIST_PROJECT = "/projects"
)

func (cli *Client) ListProjects(ctx context.Context, options schema.ProjectListOptions) ([]schema.Project, error) {
	var projects []schema.Project

	query := url.Values{}

	if options.Public != nil {
		if *options.Public {
			query.Set("public", "1")
		} else {
			query.Set("public", "0")
		}
	}

	if v := options.Name; v != "" {
		query.Set("name", v)
	}

	if v := options.Owner; v != "" {
		query.Set("owner", v)
	}

	if v := options.Page; v != "" {
		query.Set("page", v)
	}

	if v := options.PageSize; v != "" {
		query.Set("page_size", v)
	}

	serverResp, err := cli.get(ctx, PATH_LIST_PROJECT, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return projects, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&projects)

	return projects, nil
}
