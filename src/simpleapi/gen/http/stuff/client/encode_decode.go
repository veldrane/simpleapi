// Code generated by goa v3.2.5, DO NOT EDIT.
//
// stuff HTTP client encoders and decoders
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
	stuff "simpleapi/gen/stuff"
	stuffviews "simpleapi/gen/stuff/views"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "stuff" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListStuffPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("stuff", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the stuff
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("stuff", "list", err)
			}
			p := NewListStoredArticleCollectionOK(body)
			view := "default"
			vres := stuffviews.StoredArticleCollection{Projected: p, View: view}
			if err = stuffviews.ValidateStoredArticleCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("stuff", "list", err)
			}
			res := stuff.NewStoredArticleCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("stuff", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "stuff" service "Add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddStuffPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("stuff", "Add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the stuff Add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*stuff.AddPayload)
		if !ok {
			return goahttp.ErrInvalidType("stuff", "Add", "*stuff.AddPayload", v)
		}
		if p.Admin != nil {
			head := *p.Admin
			req.Header.Set("X-Simple-Admin", head)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("stuff", "Add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the stuff Add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
		case http.StatusCreated:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("stuff", "Add", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "stuff" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*stuff.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("stuff", "show", "*stuff.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowStuffPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("stuff", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeShowResponse returns a decoder for responses returned by the stuff
// show endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeShowResponse may return the following errors:
//	- "NotFound" (type *goa.ServiceError): http.StatusNotFound
//	- "Timeout" (type *goa.ServiceError): http.StatusInternalServerError
//	- "InternalError" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("stuff", "show", err)
			}
			p := NewShowStoredArticleOK(&body)
			view := "default"
			vres := &stuffviews.StoredArticle{Projected: p, View: view}
			if err = stuffviews.ValidateStoredArticle(vres); err != nil {
				return nil, goahttp.ErrValidationError("stuff", "show", err)
			}
			res := stuff.NewStoredArticle(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body ShowNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("stuff", "show", err)
			}
			err = ValidateShowNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("stuff", "show", err)
			}
			return nil, NewShowNotFound(&body)
		case http.StatusInternalServerError:
			en := resp.Header.Get("goa-error")
			switch en {
			case "Timeout":
				var (
					body ShowTimeoutResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("stuff", "show", err)
				}
				err = ValidateShowTimeoutResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("stuff", "show", err)
				}
				return nil, NewShowTimeout(&body)
			case "InternalError":
				var (
					body ShowInternalErrorResponseBody
					err  error
				)
				err = decoder(resp).Decode(&body)
				if err != nil {
					return nil, goahttp.ErrDecodingError("stuff", "show", err)
				}
				err = ValidateShowInternalErrorResponseBody(&body)
				if err != nil {
					return nil, goahttp.ErrValidationError("stuff", "show", err)
				}
				return nil, NewShowInternalError(&body)
			default:
				body, _ := ioutil.ReadAll(resp.Body)
				return nil, goahttp.ErrInvalidResponse("stuff", "show", resp.StatusCode, string(body))
			}
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("stuff", "show", resp.StatusCode, string(body))
		}
	}
}

// unmarshalStoredArticleResponseToStuffviewsStoredArticleView builds a value
// of type *stuffviews.StoredArticleView from a value of type
// *StoredArticleResponse.
func unmarshalStoredArticleResponseToStuffviewsStoredArticleView(v *StoredArticleResponse) *stuffviews.StoredArticleView {
	res := &stuffviews.StoredArticleView{
		ID:      v.ID,
		Title:   v.Title,
		Author:  v.Author,
		Desc:    v.Desc,
		Content: v.Content,
		Admin:   v.Admin,
	}

	return res
}