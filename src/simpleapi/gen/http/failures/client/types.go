// Code generated by goa v3.2.5, DO NOT EDIT.
//
// failures HTTP client types
//
// Command:
// $ goa gen simpleapi/design

package client

import (
	failures "simpleapi/gen/failures"
	failuresviews "simpleapi/gen/failures/views"
)

// SetTimeoutRequestBody is the type of the "failures" service "SetTimeout"
// endpoint HTTP request body.
type SetTimeoutRequestBody struct {
	// Rate of the failures on /article/{id}
	TimeOutRate int `form:"TimeOutRate" json:"TimeOutRate" xml:"TimeOutRate"`
}

// GetTimeoutResponseBody is the type of the "failures" service "GetTimeout"
// endpoint HTTP response body.
type GetTimeoutResponseBody struct {
	// rate ot the timeouts in percent
	Timeoutrate *int `form:"timeoutrate,omitempty" json:"timeoutrate,omitempty" xml:"timeoutrate,omitempty"`
}

// NewSetTimeoutRequestBody builds the HTTP request body from the payload of
// the "SetTimeout" endpoint of the "failures" service.
func NewSetTimeoutRequestBody(p *failures.SetTimeoutPayload) *SetTimeoutRequestBody {
	body := &SetTimeoutRequestBody{
		TimeOutRate: p.TimeOutRate,
	}
	return body
}

// NewGetTimeoutTimeOutRateOK builds a "failures" service "GetTimeout" endpoint
// result from a HTTP "OK" response.
func NewGetTimeoutTimeOutRateOK(body *GetTimeoutResponseBody) *failuresviews.TimeOutRateView {
	v := &failuresviews.TimeOutRateView{
		Timeoutrate: body.Timeoutrate,
	}

	return v
}