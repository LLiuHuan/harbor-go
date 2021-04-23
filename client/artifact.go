package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_GET_PROJECT_ARTIFACTS_LIST = "/projects/%s/repositories/%s/artifacts"
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
