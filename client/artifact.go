package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_GET_PROJECT_ARTIFACTS_LIST      = "/projects/%s/repositories/%s/artifacts"
	PATH_GET_PROJECT_ARTIFACTS_TAGS      = "/projects/%s/repositories/%s/artifacts/%s/tags"
	PATH_GET_PROJECT_ARTIFACTS_INFO      = "/projects/%s/repositories/%s/artifacts/%s"
	PATH_GET_PROJECT_ARTIFACTS_ADDITIONS = "/projects/%s/repositories/%s/artifacts/%s/additions/%s"
)

// GetProjectArtifactsList List artifacts
// List artifacts under the specific project and repository. Except the basic properties, the other supported queries in "q" includes "tags=*" to list only tagged artifacts, "tags=nil" to list only untagged artifacts, "tags=~v" to list artifacts whose tag fuzzy matches "v", "tags=v" to list artifact whose tag exactly matches "v", "labels=(id1, id2)" to list artifacts that both labels with id1 and id2 are added to
// url: /projects/{project_name}/repositories/{repository_name}/artifacts
func (cli *Client) GetProjectArtifactsList(ctx context.Context, options schema.GetProjectArtifactsListOptions) (res []schema.GetProjectArtifactsList, err error) {
	query := StructQuery(options)
	header := StructHeader(options)
	path := fmt.Sprintf(PATH_GET_PROJECT_ARTIFACTS_LIST, options.ProjectName, options.RepositoryName)
	serverResp, err := cli.get(ctx, path, query, header)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(serverResp.body).Decode(&res)
	return
}

// GetProjectArtifactsTags List tags
// List tags of the specific artifact
// url: /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/tags
func (cli *Client) GetProjectArtifactsTags(ctx context.Context, options schema.GetProjectArtifactsTagsOptions) (res []schema.GetProjectArtifactsTags, err error) {
	query := StructQuery(options)
	header := StructHeader(options)
	path := fmt.Sprintf(PATH_GET_PROJECT_ARTIFACTS_TAGS, options.ProjectName, options.RepositoryName, options.Reference)
	serverResp, err := cli.get(ctx, path, query, header)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(serverResp.body).Decode(&res)
	return
}

// GetProjectArtifactsInfo Get the specific artifact
// Get the artifact specified by the reference under the project and repository. The reference can be digest or tag.
// url: /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}
func (cli *Client) GetProjectArtifactsInfo(ctx context.Context, options schema.GetProjectArtifactsTagsOptions) (res schema.GetProjectArtifactsInfo, err error) {
	query := StructQuery(options)
	header := StructHeader(options)
	path := fmt.Sprintf(PATH_GET_PROJECT_ARTIFACTS_INFO, options.ProjectName, options.RepositoryName, options.Reference)
	serverResp, err := cli.get(ctx, path, query, header)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(serverResp.body).Decode(&res)
	return
}

// GetProjectArtifactsAdditions Get the addition of the specific artifact
// Get the addition of the artifact specified by the reference under the project and repository.
// url: /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition}
// addition = [build_history, values.yaml, readme.md, dependencies]
func (cli *Client) GetProjectArtifactsAdditions(ctx context.Context, options schema.GetProjectArtifactsAdditionsOptions) (res interface{}, err error) {
	query := StructQuery(options)
	header := StructHeader(options)
	path := fmt.Sprintf(PATH_GET_PROJECT_ARTIFACTS_ADDITIONS, options.ProjectName, options.RepositoryName, options.Reference, options.Addition)
	serverResp, err := cli.get(ctx, path, query, header)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch options.Addition {
	case "build_history":
		res = make([]schema.GetProjectArtifactsAdditionsBuildHistory, 0)
	default:
		res = ""
	}
	err = json.NewDecoder(serverResp.body).Decode(&res)
	return
}
