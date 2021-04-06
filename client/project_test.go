package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/lliuhuan/harbor-go/schema"
)

func TestListProjects(t *testing.T) {
	var c, err = NewClientWithOpts(
		WithHost("http://10.0.88.69:8080"),
		WithBasicAuth("admin", "Harbor12345"),
	)
	if err != nil {
		fmt.Println(err)
	}
	projects, err := c.ListProjects(context.Background(), schema.ProjectListOptions{})
	for k, v := range projects {
		fmt.Println(k, v)
	}
}

func TestPostProjectMetadata(t *testing.T) {
	var c, err = NewClientWithOpts(
		WithHost("http://10.0.88.69:8080"),
		WithBasicAuth("admin", "Harbor12345"),
	)
	if err != nil {
		fmt.Println(err)
	}

	projects, err := c.PostProjectMetadata(context.Background(), schema.PostProjectMetadataOptions{
		ProjectID: 2,
		Metadata:  schema.ProjectMetadataOptions{EnableContentTrust: "false", AutoScan: "false", Severity: "low", Public: "true", ReuseSysCveAllowlist: "true", PreventVul: "false", RetentionId: ""},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(projects)
	}
}

func TestGetProjectMetadata(t *testing.T) {
	var c, err = NewClientWithOpts(
		WithHost("http://10.0.88.69:8080"),
		WithBasicAuth("admin", "Harbor12345"),
	)
	if err != nil {
		fmt.Println(err)
	}

	projects, err := c.GetProjectMetadata(context.Background(), schema.GetProjectMetadataOptions{ProjectID: 2})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(projects)
	fmt.Printf("%#v \n", projects)
}
