// Code generated by goa v3.2.5, DO NOT EDIT.
//
// failures endpoints
//
// Command:
// $ goa gen simpleapi/design

package failures

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "failures" service endpoints.
type Endpoints struct {
	GetTimeout goa.Endpoint
	SetTimeout goa.Endpoint
}

// NewEndpoints wraps the methods of the "failures" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetTimeout: NewGetTimeoutEndpoint(s),
		SetTimeout: NewSetTimeoutEndpoint(s),
	}
}

// Use applies the given middleware to all the "failures" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetTimeout = m(e.GetTimeout)
	e.SetTimeout = m(e.SetTimeout)
}

// NewGetTimeoutEndpoint returns an endpoint function that calls the method
// "GetTimeout" of service "failures".
func NewGetTimeoutEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.GetTimeout(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTimeOutRate(res, "default")
		return vres, nil
	}
}

// NewSetTimeoutEndpoint returns an endpoint function that calls the method
// "SetTimeout" of service "failures".
func NewSetTimeoutEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SetTimeoutPayload)
		return nil, s.SetTimeout(ctx, p)
	}
}
