// Code generated by goa v3.2.5, DO NOT EDIT.
//
// failures client
//
// Command:
// $ goa gen simpleapi/design

package failures

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "failures" service client.
type Client struct {
	GetTimeoutEndpoint goa.Endpoint
	SetTimeoutEndpoint goa.Endpoint
}

// NewClient initializes a "failures" service client given the endpoints.
func NewClient(getTimeout, setTimeout goa.Endpoint) *Client {
	return &Client{
		GetTimeoutEndpoint: getTimeout,
		SetTimeoutEndpoint: setTimeout,
	}
}

// GetTimeout calls the "GetTimeout" endpoint of the "failures" service.
func (c *Client) GetTimeout(ctx context.Context) (res *TimeOutRate, err error) {
	var ires interface{}
	ires, err = c.GetTimeoutEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*TimeOutRate), nil
}

// SetTimeout calls the "SetTimeout" endpoint of the "failures" service.
func (c *Client) SetTimeout(ctx context.Context, p *SetTimeoutPayload) (err error) {
	_, err = c.SetTimeoutEndpoint(ctx, p)
	return
}
