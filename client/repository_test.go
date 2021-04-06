package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/lliuhuan/harbor-go/schema"
)

func TestGetRepositoryList(t *testing.T) {
	var c, err = NewClientWithOpts(
		WithHost("http://10.0.88.69:8080"),
		WithBasicAuth("admin", "Harbor12345"),
	)
	if err != nil {
		fmt.Println(err)
	}
	repository, err := c.GetRepositoryList(context.Background(), schema.GetRepositoryListOptions{ProjectName: "test"})
	for k, v := range repository {
		fmt.Println(k, v)
	}
}

func TestGetRepository(t *testing.T) {
	var c, err = NewClientWithOpts(
		WithHost("http://10.0.88.69:8080"),
		WithBasicAuth("admin", "Harbor12345"),
	)
	if err != nil {
		fmt.Println(err)
	}
	repository, err := c.GetRepositoryByName(context.Background(), schema.GetRepositoryByNameOptions{ProjectName: "test", RepositoryName: "test/nginx"})
	fmt.Println(repository)
	fmt.Println(err)
	for k, v := range repository {
		fmt.Println(k, v)
	}
}
