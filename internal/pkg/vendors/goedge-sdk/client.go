package goedgesdk

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	apiHost     string
	userType    string
	userId      int64
	accessKeyId string
	accessKey   string

	accessToken          string
	accessTokenExpiresAt time.Time

	client *resty.Client
}

func NewClient(apiHost string, userType string, userId int64, accessKeyId string, accessKey string) *Client {
	client := resty.New()

	return &Client{
		apiHost:     strings.TrimRight(apiHost, "/"),
		userType:    userType,
		userId:      userId,
		accessKeyId: accessKeyId,
		accessKey:   accessKey,
		client:      client,
	}
}

func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.client.SetTimeout(timeout)
	return c
}

func (c *Client) sendRequest(path string, params interface{}) (*resty.Response, error) {
	url := c.apiHost + path

	// TODO: access token
	req := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params)
	resp, err := req.Post(url)
	if err != nil {
		return nil, fmt.Errorf("goedge api error: failed to send request: %w", err)
	} else if resp.IsError() {
		return nil, fmt.Errorf("goedge api error: unexpected status code: %d, %s", resp.StatusCode(), resp.Body())
	}

	return resp, nil
}

func (c *Client) sendRequestWithResult(path string, params interface{}, result BaseResponse) error {
	resp, err := c.sendRequest(path, params)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return fmt.Errorf("goedge api error: failed to parse response: %w", err)
	}

	return nil
}
