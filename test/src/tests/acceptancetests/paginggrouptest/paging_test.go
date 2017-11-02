package paginggrouptest

import (
	"context"
	"net/http"
	"testing"

	chk "gopkg.in/check.v1"

	"tests/acceptancetests/utils"
	. "tests/generated/paging"
)

func Test(t *testing.T) { chk.TestingT(t) }

type PagingGroupSuite struct{}

var _ = chk.Suite(&PagingGroupSuite{})

var pagingClient = getPagingClient()
var clientID = "client-id"

func getPagingClient() PagingClient {
	c := NewPagingClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *PagingGroupSuite) TestGetMultiplePages(c *chk.C) {
	// Get pages one by one...
	res, err := pagingClient.GetMultiplePages(context.Background(), clientID, nil, nil)
	c.Assert(err, chk.IsNil)
	c.Assert(res.NextLink, chk.NotNil)
	count := 1
	for res.NextLink != nil {
		count++
		resNext, err := pagingClient.GetMultiplePagesNextResults(res)
		c.Assert(err, chk.IsNil)
		res = resNext
	}
	c.Assert(count, chk.Equals, 10)

	// Get all!
	resChan, errChan := pagingClient.GetMultiplePagesComplete(context.Background(), clientID, nil, nil)
	count = 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 10)
	c.Assert(<-errChan, chk.IsNil)

	// Get some and then cancel
	ctx, cancel := context.WithCancel(context.Background())
	resChan, errChan = pagingClient.GetMultiplePagesComplete(ctx, clientID, nil, nil)
	for i := 0; i < 3; i++ {
		_, ok := <-resChan
		c.Assert(ok, chk.Equals, true)
	}
	cancel()
	c.Assert(<-errChan, chk.ErrorMatches, "context canceled")
	_, ok := <-resChan
	c.Assert(ok, chk.Equals, false)
}

func (s *PagingGroupSuite) TestGetSinglePages(c *chk.C) {
	res, err := pagingClient.GetSinglePages(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.NextLink, chk.IsNil)
}

func (s *PagingGroupSuite) TestGetOdataMultiplePages(c *chk.C) {
	res, err := pagingClient.GetOdataMultiplePages(context.Background(), clientID, nil, nil)
	c.Assert(err, chk.IsNil)
	c.Assert(res.OdataNextLink, chk.NotNil)
	count := 1
	for res.OdataNextLink != nil {
		count++
		resNext, err := pagingClient.GetOdataMultiplePagesNextResults(res)
		c.Assert(err, chk.IsNil)
		res = resNext
	}
	c.Assert(count, chk.Equals, 10)

	resChan, errChan := pagingClient.GetOdataMultiplePagesComplete(context.Background(), clientID, nil, nil)
	count = 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 10)
	c.Assert(<-errChan, chk.IsNil)
}

func (s *PagingGroupSuite) TestGetMultiplePagesWithOffset(c *chk.C) {
	res, err := pagingClient.GetMultiplePagesWithOffset(context.Background(), 100, clientID, nil, nil)
	c.Assert(err, chk.IsNil)
	c.Assert(res.NextLink, chk.NotNil)
	count := 1
	for res.NextLink != nil {
		count++
		resNext, err := pagingClient.GetMultiplePagesWithOffsetNextResults(res)
		c.Assert(err, chk.IsNil)
		res = resNext
	}
	c.Assert(count, chk.Equals, 10)
	c.Assert(*(*res.Values)[0].Properties.ID, chk.Equals, int32(110))

	resChan, errChan := pagingClient.GetMultiplePagesWithOffsetComplete(context.Background(), 100, clientID, nil, nil)
	count = 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 10)
	c.Assert(<-errChan, chk.IsNil)
}

func (s *PagingGroupSuite) TestGetMultiplePagesRetryFirst(c *chk.C) {
	res, err := pagingClient.GetMultiplePagesRetryFirst(context.Background())
	c.Assert(err, chk.IsNil)
	count := 1
	for res.NextLink != nil {
		count++
		resNext, err := pagingClient.GetMultiplePagesRetryFirstNextResults(res)
		c.Assert(err, chk.IsNil)
		res = resNext
	}
	c.Assert(count, chk.Equals, 10)

	resChan, errChan := pagingClient.GetMultiplePagesRetryFirstComplete(context.Background())
	count = 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 10)
	c.Assert(<-errChan, chk.IsNil)
}

func (s *PagingGroupSuite) TestGetMultiplePagesRetrySecond(c *chk.C) {
	res, err := pagingClient.GetMultiplePagesRetrySecond(context.Background())
	c.Assert(err, chk.IsNil)
	count := 1
	for res.NextLink != nil {
		count++
		resNext, err := pagingClient.GetMultiplePagesRetrySecondNextResults(res)
		c.Assert(err, chk.IsNil)
		res = resNext
	}
	c.Assert(count, chk.Equals, 10)

	resChan, errChan := pagingClient.GetMultiplePagesRetrySecondComplete(context.Background())
	count = 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 10)
	c.Assert(<-errChan, chk.IsNil)
}

func (s *PagingGroupSuite) TestGetSinglePagesFailure(c *chk.C) {
	res, err := pagingClient.GetSinglePagesFailure(context.Background())
	c.Assert(err, chk.NotNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusBadRequest)

	resChan, errChan := pagingClient.GetSinglePagesFailureComplete(context.Background())
	count := 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 0)
	c.Assert(<-errChan, chk.NotNil)
}

func (s *PagingGroupSuite) TestGetMultiplePagesFailure(c *chk.C) {
	res, err := pagingClient.GetMultiplePagesFailure(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.NextLink, chk.NotNil)
	res, err = pagingClient.GetMultiplePagesFailureNextResults(res)
	c.Assert(err, chk.NotNil)
	c.Assert(res.StatusCode, chk.Equals, http.StatusBadRequest)

	resChan, errChan := pagingClient.GetMultiplePagesFailureComplete(context.Background())
	count := 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 1)
	c.Assert(<-errChan, chk.NotNil)
}

func (s *PagingGroupSuite) TestGetMultiplePagesFailureURI(c *chk.C) {
	res, err := pagingClient.GetMultiplePagesFailureURI(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.NextLink, chk.Equals, "*&*#&$")
	_, err = pagingClient.GetMultiplePagesFailureURINextResults(res)
	c.Assert(err, chk.NotNil)
	c.Assert(err, chk.ErrorMatches, ".*No scheme detected in URL.*")

	resChan, errChan := pagingClient.GetMultiplePagesFailureURIComplete(context.Background())
	count := 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 1)
	err = <-errChan
	c.Assert(err, chk.NotNil)
	c.Assert(err, chk.ErrorMatches, ".*No scheme detected in URL.*")
}

func (s *PagingGroupSuite) TestGetMultiplePagesFragmentNextLink(c *chk.C) {
	res, err := pagingClient.GetMultiplePagesFragmentNextLink(context.Background(), "1.6", "test_user")
	c.Assert(err, chk.IsNil)
	count := 1
	for res.OdataNextLink != nil {
		count++
		resNext, err := pagingClient.NextFragment(context.Background(), "1.6", "test_user", *res.OdataNextLink)
		c.Assert(err, chk.IsNil)
		res = resNext
	}
	c.Assert(count, chk.Equals, 10)

	resChan, errChan := pagingClient.GetMultiplePagesFragmentNextLinkComplete(context.Background(), "1.6", "test_user")
	count = 0
	for item := range resChan {
		count++
		c.Assert(item, chk.NotNil)
	}
	c.Assert(count, chk.Equals, 10)
	c.Assert(<-errChan, chk.IsNil)
}
