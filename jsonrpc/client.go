package jsonrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client is interface of the JSON-RPC client
type Client interface {
	// Call with requests. When given multiple requests
	Call(...*Request) (map[RequestID]*Response, error)
}

// NewClient creates a new JSON-RPC Client
func NewClient(url string) Client {
	return &client{url: url, http: &http.Client{Transport: &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}}}
}

type client struct {
	url  string
	http *http.Client
}

// Call implements Client interface
func (c *client) Call(requests ...*Request) (map[RequestID]*Response, error) {
	switch len(requests) {
	case 0:
		return nil, errors.New("no requests")
	case 1:
		request := requests[0]
		reqBody, err := json.Marshal(request)
		if err != nil {
			return nil, newError(SerializeRequestJsonError, err)
		}
		var resp Response
		if err = c.httpPost(reqBody, &resp); err != nil {
			return nil, err
		}
		return valid(requests, &resp)
	default:
		reqBody, err := json.Marshal(requests)
		if err != nil {
			return nil, newError(SerializeRequestJsonError, err)
		}
		var resps []*Response
		if err = c.httpPost(reqBody, &resps); err != nil {
			return nil, err
		}
		return valid(requests, resps...)
	}
}

func (c *client) httpPost(body []byte, ret interface{}) error {
	resp, err := c.http.Post(c.url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return newError(HttpCallError, err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return newError(ReadHttpResponseBodyError, err)
	}

	if resp.StatusCode != 200 {
		return newError(HttpCallError, fmt.Errorf(
			"Failed https call: %d, %s", resp.StatusCode, string(body)))
	}

	if err = json.Unmarshal(body, ret); err != nil {
		return newError(ParseResponseJsonError, err)
	}
	return nil
}

func valid(requests []*Request, resps ...*Response) (map[RequestID]*Response, error) {
	ret := make(map[RequestID]*Response)
	for _, resp := range resps {
		if err := resp.Validate(); err != nil {
			return nil, err
		}
		if resp.ID != nil {
			ret[*resp.ID] = resp
		}
	}
	var missing []string
	for _, req := range requests {
		if _, ok := ret[req.ID]; !ok {
			missing = append(missing, req.ToString())
		}
	}
	if len(missing) > 0 {
		return ret, newError(InvalidJsonRpcResponseError, fmt.Errorf(
			"missing responses for requests: \n%s", strings.Join(missing, "\n")))
	}
	return ret, nil
}
