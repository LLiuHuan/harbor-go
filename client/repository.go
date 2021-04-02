package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_REPOSITORY_LIST = "/projects/%s/repositories"    // /projects/{project_name}/repositories
	PATH_REPOSITORY_GET  = "/projects/%s/repositories/%s" // /projects/{project_name}/repositories/{repository_name}
)

func (cli *Client) ListRepository(ctx context.Context, options schema.RepositoryListOptions) ([]schema.Repository, error) {
	var repository []schema.Repository

	query := url.Values{}

	if v := options.ProjectName; v != "" {
		query.Set("project_name", v)
	}

	if v := options.Q; v != "" {
		query.Set("q", v)
	}

	if v := options.Sort; v != "" {
		query.Set("sort", v)
	}

	if v := options.Page; v != "" {
		query.Set("page", v)
	}

	if v := options.PageSize; v != "" {
		query.Set("page_size", v)
	}

	//headers := make(map[string][]string)
	//headers["Authorization"] = []string{"Basic bGg6TGl1SHVhbjEyMw=="}
	//req.SetBasicAuth()
	path := fmt.Sprintf(PATH_REPOSITORY_LIST, options.ProjectName)
	serverResp, err := cli.get(ctx, path, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return repository, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&repository)

	return repository, nil
}

func (cli *Client) GetRepository(ctx context.Context, options schema.RepositoryGetOptions) ([]schema.Repository, error) {
	var repository []schema.Repository

	query := url.Values{}

	if v := options.ProjectName; v != "" {
		query.Set("project_name", v)
	}

	if v := options.RepositoryName; v != "" {
		query.Set("repository_name", v)
	}

	path := fmt.Sprintf(PATH_REPOSITORY_GET, options.ProjectName, options.RepositoryName)
	serverResp, err := cli.get(ctx, path, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return repository, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&repository)

	return repository, nil
}
