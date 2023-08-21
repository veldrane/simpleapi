package simpleapi

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	stuff "simpleapi/gen/stuff"
	"time"
)

const timeout = 1000

var articles stuff.StoredArticleCollection
var admin string
var count int

// articles service example implementation.
// The example methods log the requests and return zero values.
type stuffsrvc struct {
	logger *log.Logger
}

// NewStuff returns the Articles service implementation.
func NewStuff(logger *log.Logger) stuff.Service {
	count = 0
	rand.Seed(time.Now().UnixNano())
	return &stuffsrvc{logger}
}

// List all stored articles
func (s *stuffsrvc) List(ctx context.Context) (res stuff.StoredArticleCollection, err error) {

	res = articles
	return res, nil
}

// Add article to the articles
func (s *stuffsrvc) Add(ctx context.Context, p *stuff.AddPayload) (err error) {

	var article stuff.StoredArticle
	id := (count + 1)

	s.logger.Print("reference has been added ", p.Admin)

	if p.Admin == nil {
		article = stuff.StoredArticle{
			ID:      &id,
			Title:   &p.Title,
			Author:  &p.Author,
			Desc:    p.Desc,
			Content: p.Content,
		}
	} else {
		article = stuff.StoredArticle{
			ID:      &id,
			Title:   &p.Title,
			Author:  &p.Author,
			Desc:    p.Desc,
			Content: p.Content,
			Admin:   p.Admin,
		}
	}

	articles = append(articles, &article)
	count++
	s.logger.Print("New article ", p.Title, " from author ", p.Author, " has been added")
	return
}

func getarticle(id int, timeout int) (res *stuff.StoredArticle) {

	time.Sleep(time.Duration(timeout) * time.Millisecond)

	for _, k := range articles {
		if *k.ID == id {
			res = k
			break
		}
	}

	return res
}

func (s *stuffsrvc) Show(ctx context.Context, p *stuff.ShowPayload) (res *stuff.StoredArticle, err error) {

	random := 1
	if *TimeOutRateContent.Timeoutrate == 100 {
		return nil, stuff.MakeInternalError(fmt.Errorf("Internal Error"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()

	if *TimeOutRateContent.Timeoutrate > 0 {
		probability := (1 / (float64(100.0-*TimeOutRateContent.Timeoutrate) / 100.0))
		s.logger.Print("Probability has been computed to value: ", probability)
		random = rand.Intn(int(probability * timeout))
		s.logger.Print("Random timeout has been generated to value: ", random)
	}
	// Use a channel to wait for either timeout or completion of division
	c := make(chan *stuff.StoredArticle, 1)

	// Call division operation in separate Go routine
	go func() { c <- getarticle(p.ID, random) }()

	// Wait for results or timeout
	select {
	case <-ctx.Done():
		// Timeout triggered, return timeout error
		return nil, stuff.MakeTimeout(ctx.Err())
	case res := <-c:

		if res == nil {
			return res, stuff.MakeNotFound(fmt.Errorf("Article with id %d is not found", p.ID))
		}

		return res, nil
	}
}
