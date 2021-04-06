package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_GET_REPOSITORY_LIST       = "/projects/%s/repositories"    // /projects/{project_name}/repositories
	PATH_GET_REPOSITORY_BY_NAME    = "/projects/%s/repositories/%s" // /projects/{project_name}/repositories/{repository_name}
	PATH_PUT_REPOSITORY_BY_NAME    = "/projects/%s/repositories/%s" // /projects/{project_name}/repositories/{repository_name}
	PATH_DELETE_REPOSITORY_BY_NAME = "/projects/%s/repositories/%s" // /projects/{project_name}/repositories/{repository_name}
)

// GetRepositoryList List repositories of the specified project
// url: /projects/{project_name}/repositories
func (cli *Client) GetRepositoryList(ctx context.Context, options schema.GetRepositoryListOptions) ([]schema.Repository, error) {
	var repository []schema.Repository

	query := StructQuery(options)

	//headers := make(map[string][]string)
	//headers["Authorization"] = []string{"Basic bGg6TGl1SHVhbjEyMw=="}
	//req.SetBasicAuth()
	path := fmt.Sprintf(PATH_GET_REPOSITORY_LIST, options.ProjectName)
	serverResp, err := cli.get(ctx, path, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return repository, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&repository)

	return repository, nil
}

// GetRepositoryByName Get the repository specified by name
// url: /projects/{project_name}/repositories/{repository_name}
func (cli *Client) GetRepositoryByName(ctx context.Context, options schema.GetRepositoryByNameOptions) ([]schema.Repository, error) {
	var repository []schema.Repository

	query := StructQuery(options)

	path := fmt.Sprintf(PATH_GET_REPOSITORY_BY_NAME, options.ProjectName, options.RepositoryName)
	serverResp, err := cli.get(ctx, path, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return repository, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&repository)

	return repository, nil
}

// PutRepositoryBuName Update the repository specified by name
// url: /projects/{project_name}/repositories/{repository_name}
func (cli *Client) PutRepositoryBuName(ctx context.Context, options schema.PutRepositoryByNameOptions) (schema.CurrencyError, error) {
	var resp schema.CurrencyError
	query := StructQuery(options)

	path := fmt.Sprintf(PATH_PUT_REPOSITORY_BY_NAME, options.ProjectName, options.RepositoryName)
	serverResp, err := cli.put(ctx, path, query, nil, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return resp, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&resp)
	return resp, nil
}

// DeleteRepositoryBuName Delete the repository specified by name
// url: /projects/{project_name}/repositories/{repository_name}
func (cli *Client) DeleteRepositoryBuName(ctx context.Context, options schema.DeleteRepositoryByNameOptions) (schema.CurrencyError, error) {
	var resp schema.CurrencyError
	query := StructQuery(options)

	path := fmt.Sprintf(PATH_PUT_REPOSITORY_BY_NAME, options.ProjectName, options.RepositoryName)
	serverResp, err := cli.delete(ctx, path, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return resp, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&resp)
	return resp, nil
}
