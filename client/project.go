package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_LIST_PROJECT           = "/projects"
	PATH_POST_PROJECT_METADATAS = "/projects/%d/metadatas"
	PATH_GET_PROJECT_METADATAS  = "/projects/%d/metadatas"
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

// PostAddProjectMetadatas This endpoint is aimed to add metadata of a project.
// url: /projects/{project_id}/metadatas
// TODO: 怎么测试都不好使，一直提示权限问题 已经使用了最高级权限，还是无法调用成功
func (cli *Client) PostProjectMetadata(ctx context.Context, options schema.PostProjectMetadataOptions) (resp schema.Response, err error) {
	query := url.Values{}
	path := fmt.Sprintf(PATH_POST_PROJECT_METADATAS, options.ProjectID)
	serverResp, err := cli.post(ctx, path, query, options, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return
	}

	var message string
	switch serverResp.statusCode {
	case 200:
		message = "Add metadata successfully."
	case 400:
		message = "Invalid request."
	case 401:
		message = "User need to log in first."
	case 403:
		message = "User does not have permission to the project."
	case 404:
		message = "Project ID does not exist."
	case 415:
		message = "The Media Type of the request is not supported, it has to be \"application/json\""
	case 500:
		message = "Internal server errors."
	}

	resp = schema.OkWithDetailed(serverResp.statusCode, message, serverResp.body)
	return
}

// GetProjectMetadata This endpoint returns metadata of the project specified by project ID.
// url: /projects/{project_id}/metadatas
func (cli *Client) GetProjectMetadata(ctx context.Context, options schema.GetProjectMetadataOptions) (res schema.ProjectMetadataOptions, err error) {
	query := StructQuery(options)
	path := fmt.Sprintf(PATH_GET_PROJECT_METADATAS, options.ProjectID)
	serverResp, err := cli.get(ctx, path, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return
	}

	err = json.NewDecoder(serverResp.body).Decode(&res)
	return
}
