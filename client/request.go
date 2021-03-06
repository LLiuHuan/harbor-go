package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	errdefs "github.com/lliuhuan/harbor-go/errdefs"
	"github.com/pkg/errors"
)

type headers []map[string]string

// ServerResponse is a wrapper for http API responses.
type ServerResponse struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
	reqURL     *url.URL
}

// head sends an http request to the docker API using the method HEAD.
func (cli *Client) head(ctx context.Context, path string, query url.Values, headers []map[string]string) (ServerResponse, error) {
	return cli.sendRequest(ctx, "HEAD", path, query, nil, headers)
}

// get sends an http request to the docker API using the method GET with a specific Go context.
func (cli *Client) get(ctx context.Context, path string, query url.Values, headers []map[string]string) (ServerResponse, error) {
	return cli.sendRequest(ctx, "GET", path, query, nil, headers)
}

// post sends an http request to the docker API using the method POST with a specific Go context.
func (cli *Client) post(ctx context.Context, path string, query url.Values, obj interface{}, headers []map[string]string) (ServerResponse, error) {
	body, headers, err := encodeBody(obj, headers)
	if err != nil {
		return ServerResponse{}, err
	}
	return cli.sendRequest(ctx, "POST", path, query, body, headers)
}

func (cli *Client) postRaw(ctx context.Context, path string, query url.Values, body io.Reader, headers []map[string]string) (ServerResponse, error) {
	return cli.sendRequest(ctx, "POST", path, query, body, headers)
}

// put sends an http request to the docker API using the method PUT.
func (cli *Client) put(ctx context.Context, path string, query url.Values, obj interface{}, headers []map[string]string) (ServerResponse, error) {
	body, headers, err := encodeBody(obj, headers)
	if err != nil {
		return ServerResponse{}, err
	}
	return cli.sendRequest(ctx, "PUT", path, query, body, headers)
}

// putRaw sends an http request to the docker API using the method PUT.
func (cli *Client) putRaw(ctx context.Context, path string, query url.Values, body io.Reader, headers []map[string]string) (ServerResponse, error) {
	return cli.sendRequest(ctx, "PUT", path, query, body, headers)
}

// delete sends an http request to the docker API using the method DELETE.
func (cli *Client) delete(ctx context.Context, path string, query url.Values, headers []map[string]string) (ServerResponse, error) {
	return cli.sendRequest(ctx, "DELETE", path, query, nil, headers)
}

func (cli *Client) sendRequest(ctx context.Context, method, path string, query url.Values, body io.Reader, headers headers) (ServerResponse, error) {
	req, err := cli.buildRequest(method, cli.getAPIPath(ctx, path, query), body, headers)
	if err != nil {
		return ServerResponse{}, err
	}
	resp, err := cli.doRequest(ctx, req)
	if err != nil {
		return resp, errdefs.FromStatusCode(err, resp.statusCode)
	}
	err = cli.checkResponseErr(resp)
	return resp, errdefs.FromStatusCode(err, resp.statusCode)
}

func (cli *Client) buildRequest(method, path string, body io.Reader, headers headers) (*http.Request, error) {
	expectedPayload := method == "POST" || method == "PUT"
	if expectedPayload && body == nil {
		body = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(cli.basicAuth.Username, cli.basicAuth.Password)
	cli.addHeaders(req, headers)

	req.URL.Host = cli.host
	req.URL.Scheme = cli.scheme

	if expectedPayload && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (cli *Client) addHeaders(req *http.Request, headers headers) {
	if headers != nil {
		for _, header := range headers {
			for k, v := range header {
				req.Header.Set(k, v)
			}
		}
	}
}

func (cli *Client) doRequest(ctx context.Context, req *http.Request) (ServerResponse, error) {
	serverResp := ServerResponse{statusCode: -1, reqURL: req.URL}

	req = req.WithContext(ctx)
	resp, err := cli.client.Do(req)
	if err != nil {
		if cli.scheme != "https" && strings.Contains(err.Error(), "malformed HTTP response") {
			return serverResp, fmt.Errorf("%v.\n* Are you trying to connect to a TLS-enabled daemon without TLS?", err)
		}

		if cli.scheme == "https" && strings.Contains(err.Error(), "bad certificate") {
			return serverResp, errors.Wrap(err, "The server probably has client authentication (--tlsverify) enabled. Please check your TLS client certification settings")
		}
		switch err {
		case context.Canceled, context.DeadlineExceeded:
			return serverResp, err
		}
		return serverResp, errors.Wrap(err, "error during connect")
	}

	if resp != nil {
		serverResp.statusCode = resp.StatusCode
		serverResp.body = resp.Body
		serverResp.header = resp.Header
	}
	return serverResp, nil
}

func (cli *Client) checkResponseErr(serverResp ServerResponse) error {
	if serverResp.statusCode >= 200 && serverResp.statusCode < 400 {
		return nil
	}

	var body []byte
	var err error
	if serverResp.body != nil {
		bodyMax := 1 * 1024 * 1024 // 1 MiB
		bodyR := &io.LimitedReader{
			R: serverResp.body,
			N: int64(bodyMax),
		}
		body, err = ioutil.ReadAll(bodyR)
		if err != nil {
			return err
		}
		if bodyR.N == 0 {
			return fmt.Errorf("request returned %s with a message (> %d bytes) for API route and version %s, check if the server supports the requested API version", http.StatusText(serverResp.statusCode), bodyMax, serverResp.reqURL)
		}
	}
	if len(body) == 0 {
		return fmt.Errorf("request returned %s for API route and version %s, check if the server supports the requested API version", http.StatusText(serverResp.statusCode), serverResp.reqURL)
	}

	errorMessage := strings.TrimSpace(string(body))

	return errors.Wrap(errors.New(errorMessage), "Error response from daemon")
}

func ensureReaderClosed(response ServerResponse) {
	if response.body != nil {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, response.body, 512)
		response.body.Close()
	}
}

func encodeBody(obj interface{}, headers headers) (io.Reader, headers, error) {
	if obj == nil {
		return nil, headers, nil
	}

	body, err := encodeData(obj)
	if err != nil {
		return nil, headers, err
	}
	if headers == nil {
		headers = make([]map[string]string, 5)
	}
	headers = append(headers, map[string]string{"Content-Type": "application/json"})
	return body, headers, nil
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	params := bytes.NewBuffer(nil)
	if data != nil {
		if err := json.NewEncoder(params).Encode(data); err != nil {
			return nil, err
		}
	}
	return params, nil
}

func wrapResponseError(err error, resp ServerResponse, object, name string) error {
	switch {
	case err == nil:
		return nil
	case resp.statusCode == http.StatusNotFound:
		return objectNotFoundError{object: object, name: name}
	case resp.statusCode == http.StatusInternalServerError:
		return errdefs.System(err)
	default:
		return err
	}
}

type objectNotFoundError struct {
	object string
	name   string
}

func (e objectNotFoundError) NotFound() {}

func (e objectNotFoundError) Error() string {
	return fmt.Sprintf("Error: No such %s: %s", e.object, e.name)
}
