package client

import (
	"context"
	"fmt"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_PUT_USER_CLI = "/users/%d/cli_secret"
)

// PutUserCli Set CLI secret for a user.
// This endpoint let user generate a new CLI secret for himself. This API only works when auth mode is set to 'OIDC'.
// Once this API returns with successful status, the old secret will be invalid, as there will be only one CLI secret
// for a user.
// url: /users/{user_id}/cli_secret
// TODO: 未测试
func (cli *Client) PutUserCli(ctx context.Context, options schema.PutUserCli) (resp schema.Response, err error) {
	query := StructQuery(options)
	path := fmt.Sprintf(PATH_PUT_USER_CLI, options.UserID)
	serverResp, err := cli.put(ctx, path, query, options, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return
	}

	var message string
	switch serverResp.statusCode {
	case 200:
		message = "The secret is successfully updated"
	case 400:
		message = "Invalid user ID. Or user is not onboarded via OIDC authentication. Or the secret does not meet the standard."
	case 401:
		message = "User need to log in first."
	case 403:
		message = "Non-admin user can only generate the cli secret of himself."
	case 404:
		message = "User ID does not exist."
	case 412:
		message = "The auth mode of the system is not \"oidc_auth\", or the user is not onboarded via OIDC AuthN."
	case 500:
		message = "Internal server errors."
	}

	resp = schema.OkWithDetailed(serverResp.statusCode, message, serverResp.body)
	return
}
