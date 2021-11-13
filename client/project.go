package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_LIST_PROJECT                    = "/projects"
	PATH_POST_PROJECT_METADATAS          = "/projects/%d/metadatas"
	PATH_GET_PROJECT_METADATAS           = "/projects/%d/metadatas"
	PATH_HEAD_PROJECTS                   = "/projects"
	PATH_POST_PROJECTS                   = "/projects"
	PATH_DELETE_PROJECTS                 = "/projects/%s"
	PATH_GET_PROJECT_ARTIFACTS_REFERENCE = "/projects/%s/repositories/%s/artifacts/%s"
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

// PostProjectMetadata This endpoint is aimed to add metadata of a project.
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

	resp = schema.OkWithDetailed(serverResp.statusCode, serverResp.body, message)
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

// HeadProjects Check if the project name user provided already exists.
// url: /projects
func (cli *Client) HeadProjects(ctx context.Context, options schema.HeadProjects) (resp schema.Response, err error) {
	query := StructQuery(options)
	serverResp, err := cli.head(ctx, PATH_HEAD_PROJECTS, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return
	}

	var message string
	switch serverResp.statusCode {
	case 200:
		message = "Success"
	case 404:
		message = "Not found"
	case 500:
		message = "Internal server error"
	}

	resp = schema.OkWithDetailed(serverResp.statusCode, serverResp.body, message)
	return
}

// PostProjects Create a new project.
// url: /projects
func (cli *Client) PostProjects(ctx context.Context, opt schema.PostProjects) (resp schema.Response, err error) {
	header := StructHeader(opt)
	body := StructBody(opt)
	serverResp, err := cli.post(ctx, PATH_POST_PROJECTS, nil, body, header)
	var message string
	switch serverResp.statusCode {
	case 201:
		message = "Created"
	case 400:
		message = "Bad request"
	case 401:
		message = "Unauthorized"
	case 409:
		message = "Conflict"
	case 500:
		message = "Internal server error"
	default:
		message = err.Error()
	}

	resp = schema.OkWithDetailed(serverResp.statusCode, serverResp.body, message)
	return
}

// DeleteProjects Delete project by projectID
// url: /projects/{project_name_or_id}
func (cli *Client) DeleteProjects(ctx context.Context, opt schema.DeleteProjects) (serverRes ServerResponse, err error) {
	header := StructHeader(opt)

	path := fmt.Sprintf(PATH_DELETE_PROJECTS, opt.ProjectNameOrId)
	serverRes, err = cli.delete(ctx, path, nil, header)

	return
}

// GetProjectArtifactsReference Get the specific artifact
// url: /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}
func (cli *Client) GetProjectArtifactsReference(ctx context.Context, opt schema.GetProjectArtifactsReference) (res schema.ProjectArtifactsReference, err error) {
	query := StructQuery(opt)
	path := fmt.Sprintf(PATH_GET_PROJECT_ARTIFACTS_REFERENCE, opt.ProjectName, opt.RepositoryName, opt.Reference)
	serverResp, err := cli.get(ctx, path, query, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return res, err
	}

	if serverResp.statusCode == 200 {
		err = json.NewDecoder(serverResp.body).Decode(&res)
		return res, err
	} else {
		var resNotFound schema.NotFound
		err = json.NewDecoder(serverResp.body).Decode(&resNotFound)
		return res, errors.New(strconv.Itoa(serverResp.statusCode))
	}
}
