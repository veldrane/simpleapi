// Code generated by goa v3.2.5, DO NOT EDIT.
//
// failures HTTP client encoders and decoders
//
// Command:
// $ goa gen simpleapi/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	failures "simpleapi/gen/failures"
	failuresviews "simpleapi/gen/failures/views"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetTimeoutRequest instantiates a HTTP request object with method and
// path set to call the "failures" service "GetTimeout" endpoint
func (c *Client) BuildGetTimeoutRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetTimeoutFailuresPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("failures", "GetTimeout", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetTimeoutResponse returns a decoder for responses returned by the
// failures GetTimeout endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeGetTimeoutResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetTimeoutResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("failures", "GetTimeout", err)
			}
			p := NewGetTimeoutTimeOutRateOK(&body)
			view := "default"
			vres := &failuresviews.TimeOutRate{Projected: p, View: view}
			if err = failuresviews.ValidateTimeOutRate(vres); err != nil {
				return nil, goahttp.ErrValidationError("failures", "GetTimeout", err)
			}
			res := failures.NewTimeOutRate(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("failures", "GetTimeout", resp.StatusCode, string(body))
		}
	}
}

// BuildSetTimeoutRequest instantiates a HTTP request object with method and
// path set to call the "failures" service "SetTimeout" endpoint
func (c *Client) BuildSetTimeoutRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SetTimeoutFailuresPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("failures", "SetTimeout", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSetTimeoutRequest returns an encoder for requests sent to the failures
// SetTimeout server.
func EncodeSetTimeoutRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*failures.SetTimeoutPayload)
		if !ok {
			return goahttp.ErrInvalidType("failures", "SetTimeout", "*failures.SetTimeoutPayload", v)
		}
		body := NewSetTimeoutRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("failures", "SetTimeout", err)
		}
		return nil
	}
}

// DecodeSetTimeoutResponse returns a decoder for responses returned by the
// failures SetTimeout endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeSetTimeoutResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("failures", "SetTimeout", resp.StatusCode, string(body))
		}
	}
}
