package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lliuhuan/harbor-go/schema"
)

const (
	PATH_PUT_USER_CLI        = "/users/%d/cli_secret"
	PATH_GET_USER_PERMISSION = "/users/current/permissions"
	PATH_PUT_USERINFO        = "/users/%d"
	PATH_GET_USERINFO        = "/users/%d"
	PATH_DEL_USERINFO        = "/users/%d"
	PATH_GET_USER            = "/users/search"
)

// PutUserCli Set CLI secret for a user.
// This endpoint let user generate a new CLI secret for himself. This API only works when auth mode is set to 'OIDC'.
// Once this API returns with successful status, the old secret will be invalid, as there will be only one CLI secret
// for a user.
// url: /users/{user_id}/cli_secret
// TODO: 未测试
func (cli *Client) PutUserCli(ctx context.Context, options schema.PutUserCliOptions) (resp schema.Response, err error) {
	query := StructQuery(options)
	path := fmt.Sprintf(PATH_PUT_USER_CLI, options.UserID)
	serverResp, err := cli.put(ctx, path, query, options, nil)
	if err != nil {
		return
	}
	defer ensureReaderClosed(serverResp)

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

/*
GetUserPermission 获取当前用户权限。
url: /users/current/permissions
*/
// TODO: 未测试
func (cli *Client) GetUserPermission(ctx context.Context, options schema.GetUserPermissionOptions) (resp []schema.GetUserPermission, err error) {
	query := StructQuery(options)

	serverResp, err := cli.get(ctx, PATH_GET_USER_PERMISSION, query, nil)
	if err != nil {
		return resp, err
	}
	defer ensureReaderClosed(serverResp)

	err = json.NewDecoder(serverResp.body).Decode(&resp)

	return resp, nil
}

/*
PutUserInfoById 修改用户个人资料。
url: /users/{user_id}
*/
// TODO: 未测试
func (cli *Client) PutUserInfoById(ctx context.Context, options schema.PutUserInfoOptions) (resp schema.Response, err error) {
	query := StructQuery(options)
	path := fmt.Sprintf(PATH_PUT_USERINFO, options.UserID)
	serverResp, err := cli.put(ctx, path, query, options, nil)
	if err != nil {
		return
	}
	defer ensureReaderClosed(serverResp)

	var message string
	switch serverResp.statusCode {
	case 200:
		message = "Updated user's profile successfully."
	case 400:
		message = "Invalid user ID."
	case 401:
		message = "User need to log in first."
	case 403:
		message = "User does not have permission of admin role."
	case 404:
		message = "User ID does not exist."
	case 500:
		message = "Unexpected internal errors."
	}

	resp = schema.OkWithDetailed(serverResp.statusCode, message, serverResp.body)
	return
}

/*
GetUserInfoById 获取用户个人资料。
url: /users/{user_id}
*/
// TODO: 未测试
func (cli *Client) GetUserInfoById(ctx context.Context, options schema.GetUserInfoOptions) (resp schema.GetUserInfo, err error) {
	query := StructQuery(options)
	path := fmt.Sprintf(PATH_GET_USERINFO, options.UserID)
	serverResp, err := cli.get(ctx, path, query, nil)
	if err != nil {
		return
	}
	defer ensureReaderClosed(serverResp)

	err = json.NewDecoder(serverResp.body).Decode(&resp)
	return
}

/*
DelUserInfoById 将用户标记为删除，不会从数据库删除。
url: /users/{user_id}
*/
// TODO: 未测试
func (cli *Client) DelUserInfoById(ctx context.Context, options schema.DelUserInfoOptions) (resp schema.Response, err error) {
	query := StructQuery(options)
	path := fmt.Sprintf(PATH_DEL_USERINFO, options.UserID)
	serverResp, err := cli.put(ctx, path, query, options, nil)
	if err != nil {
		return
	}
	defer ensureReaderClosed(serverResp)

	var message string
	switch serverResp.statusCode {
	case 200:
		message = "Marked user as be removed successfully."
	case 400:
		message = "Invalid user ID."
	case 401:
		message = "User need to log in first."
	case 403:
		message = "User does not have permission of admin role."
	case 404:
		message = "User ID does not exist."
	case 500:
		message = "Unexpected internal errors."
	}

	resp = schema.OkWithDetailed(serverResp.statusCode, message, serverResp.body)
	return
}

/*
GetUserSearchByName 根据名称搜索用户
url: /users/search
*/
// TODO: 未测试
func (cli *Client) GetUserSearchByName(ctx context.Context, options schema.GetUserSearchByNameOptions) (resp []schema.GetUserSearchByName, err error) {
	query := StructQuery(options)
	serverResp, err := cli.get(ctx, PATH_GET_USER, query, nil)
	if err != nil {
		return nil, err
	}
	defer ensureReaderClosed(serverResp)
	err = json.NewDecoder(serverResp.body).Decode(&resp)
	return
}
