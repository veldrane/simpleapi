// Code generated by goa v3.2.5, DO NOT EDIT.
//
// stuff HTTP server encoders and decoders
//
// Command:
// $ goa gen simpleapi/design

package server

import (
	"context"
	"io"
	"net/http"
	stuffviews "simpleapi/gen/stuff/views"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the stuff
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(stuffviews.StoredArticleCollection)
		enc := encoder(ctx, w)
		body := NewStoredArticleResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeAddResponse returns an encoder for responses returned by the stuff Add
// endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeAddRequest returns a decoder for requests sent to the stuff Add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			admin *string
		)
		adminRaw := r.Header.Get("X-Simple-Admin")
		if adminRaw != "" {
			admin = &adminRaw
		}
		payload := NewAddPayload(&body, admin)

		return payload, nil
	}
}

// EncodeShowResponse returns an encoder for responses returned by the stuff
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*stuffviews.StoredArticle)
		enc := encoder(ctx, w)
		body := NewShowResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the stuff show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(id)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show stuff
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "NotFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "Timeout":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowTimeoutResponseBody(res)
			}
			w.Header().Set("goa-error", "Timeout")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "InternalError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalStuffviewsStoredArticleViewToStoredArticleResponse builds a value of
// type *StoredArticleResponse from a value of type
// *stuffviews.StoredArticleView.
func marshalStuffviewsStoredArticleViewToStoredArticleResponse(v *stuffviews.StoredArticleView) *StoredArticleResponse {
	res := &StoredArticleResponse{
		ID:      v.ID,
		Title:   v.Title,
		Author:  v.Author,
		Desc:    v.Desc,
		Content: v.Content,
		Admin:   v.Admin,
	}

	return res
}
