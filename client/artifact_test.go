package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/lliuhuan/harbor-go/schema"
)

func TestGetProjectArtifactsList(t *testing.T) {
	var c, err = NewClientWithOpts(
		WithHost("http://49.233.178.154:8082"),
		WithBasicAuth("admin", "Harbor12345"),
	)
	if err != nil {
		fmt.Println(err)
	}
	projects, err := c.GetProjectArtifactsList(context.Background(), schema.GetProjectArtifactsListOptions{ProjectName: "fairman", RepositoryName: "nginx"})
	for k, v := range projects {
		fmt.Println(k, v.Tags)
	}
}
