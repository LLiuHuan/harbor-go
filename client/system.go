package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_GET_SYSTEM_INFO = "/systeminfo"
)

// GetSystemInfo Get general system info
// This API is for retrieving general system info, this can be called by anonymous request. Some attributes will be omitted in the response when this API is called by anonymous request.
// url: /systeminfo
func (cli *Client) GetSystemInfo(ctx context.Context) (res schema.GetSystemInfo, err error) {
	serverResp, err := cli.get(ctx, PATH_GET_SYSTEM_INFO, nil, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewDecoder(serverResp.body).Decode(&res)
	return
}
