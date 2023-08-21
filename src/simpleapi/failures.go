package simpleapi

import (
	"context"
	"log"
	"fmt"
	failures "simpleapi/gen/failures"
)

var TimeOutRateContent failures.TimeOutRate

// failures service example implementation.
// The example methods log the requests and return zero values.
type failuressrvc struct {
	logger *log.Logger
}

// NewFailures returns the failures service implementation.
func NewFailures(logger *log.Logger) failures.Service {
	rate := 0
	TimeOutRateContent = failures.TimeOutRate{
		Timeoutrate: &rate,
	}
	return &failuressrvc{logger}
}

// List all stored articles
func (s *failuressrvc) GetTimeout(ctx context.Context) (res *failures.TimeOutRate, err error) {
	res = &TimeOutRateContent
	s.logger.Print("failures.GetTimeout")
	return
}

// Manage timeouts on the /article/{id} endpoint
func (s *failuressrvc) SetTimeout(ctx context.Context, p *failures.SetTimeoutPayload) (err error) {

	if ((p.TimeOutRate > 100) || (p.TimeOutRate < 0)) {
		return failures.MakeBadRequest(fmt.Errorf("Timeout has to be between 0 and 100"))
	}
	rate := failures.TimeOutRate {
		Timeoutrate: &p.TimeOutRate,
	}
	TimeOutRateContent = rate
	s.logger.Print("failures.SetTimeout")
	return
}
