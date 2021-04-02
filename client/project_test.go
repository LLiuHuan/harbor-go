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
