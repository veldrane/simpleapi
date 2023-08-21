// Code generated by goa v3.2.5, DO NOT EDIT.
//
// root service
//
// Command:
// $ goa gen simpleapi/design

package root

import (
	"context"
)

// Service provide redirect to swagger-ui
type Service interface {
	// Return default redirect
	Default(context.Context) (res *DefaultResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "root"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"default"}

// DefaultResult is the result type of the root service default method.
type DefaultResult struct {
	// Default location string
	Location *string
}
