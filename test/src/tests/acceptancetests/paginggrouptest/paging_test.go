package paginggrouptest

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	chk "gopkg.in/check.v1"

	"tests/acceptancetests/utils"
	"tests/generated/paging"
)

func Test(t *testing.T) { chk.TestingT(t) }

type PagingGroupSuite struct{}

var _ = chk.Suite(&PagingGroupSuite{})

var pagingClient = getPagingClient()
var clientID = "client-id"

func getPagingClient() paginggroup.PagingClient {
	c := paginggroup.NewPagingClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *PagingGroupSuite) TestGetMultiplePages(c *chk.C) {
	// Get pages one by one...
	count := 0
	for page, err := pagingClient.GetMultiplePages(context.Background(), clientID, nil, nil); page.NotDone(); err = page.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(page.Values(), chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)

	// Get all!
	count = 0
	for iter, err := pagingClient.GetMultiplePagesComplete(context.Background(), clientID, nil, nil); iter.NotDone(); err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value().Properties, chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)
}

func (s *PagingGroupSuite) TestGetSinglePages(c *chk.C) {
	page, err := pagingClient.GetSinglePages(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(page.NotDone(), chk.Equals, true)
	err = page.Next()
	c.Assert(err, chk.IsNil)
	c.Assert(page.NotDone(), chk.Equals, false)
	err = page.Next()
	c.Assert(err, chk.IsNil)
	c.Assert(page.NotDone(), chk.Equals, false)
}

func (s *PagingGroupSuite) TestGetOdataMultiplePages(c *chk.C) {
	count := 0
	for page, err := pagingClient.GetOdataMultiplePages(context.Background(), clientID, nil, nil); page.NotDone(); err = page.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(page.Values(), chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)

	count = 0
	for iter, err := pagingClient.GetOdataMultiplePagesComplete(context.Background(), clientID, nil, nil); iter.NotDone(); err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value().Properties, chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)
}

func (s *PagingGroupSuite) TestGetMultiplePagesWithOffset(c *chk.C) {
	count := 0
	var id int32
	for page, err := pagingClient.GetMultiplePagesWithOffset(context.Background(), 100, clientID, nil, nil); page.NotDone(); err = page.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(page.Values(), chk.NotNil)
		count++
		id = *page.Values()[0].Properties.ID
	}
	c.Assert(count, chk.Equals, 10)
	c.Assert(id, chk.Equals, int32(110))

	count = 0
	for iter, err := pagingClient.GetMultiplePagesWithOffsetComplete(context.Background(), 100, clientID, nil, nil); iter.NotDone(); err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value().Properties, chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)
}

func (s *PagingGroupSuite) TestGetMultiplePagesRetryFirst(c *chk.C) {
	count := 0
	for page, err := pagingClient.GetMultiplePagesRetryFirst(context.Background()); page.NotDone(); err = page.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(page.Values(), chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)

	count = 0
	for iter, err := pagingClient.GetMultiplePagesRetryFirstComplete(context.Background()); iter.NotDone(); err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value().Properties, chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)
}

func (s *PagingGroupSuite) TestGetMultiplePagesRetrySecond(c *chk.C) {
	count := 0
	for page, err := pagingClient.GetMultiplePagesRetrySecond(context.Background()); page.NotDone(); err = page.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(page.Values(), chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)

	count = 0
	for iter, err := pagingClient.GetMultiplePagesRetrySecondComplete(context.Background()); iter.NotDone(); err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value().Properties, chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)
}

func (s *PagingGroupSuite) TestGetSinglePagesFailure(c *chk.C) {
	page, err := pagingClient.GetSinglePagesFailure(context.Background())
	c.Assert(err, chk.NotNil)
	c.Assert(page.Response().StatusCode, chk.Equals, http.StatusBadRequest)

	count := 0
	for iter, err := pagingClient.GetSinglePagesFailureComplete(context.Background()); iter.NotDone(); err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value().Properties, chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 0)
}

func (s *PagingGroupSuite) TestGetMultiplePagesFailure(c *chk.C) {
	page, err := pagingClient.GetMultiplePagesFailure(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(page.NotDone(), chk.Equals, true)
	err = page.Next()
	c.Assert(err, chk.NotNil)
	detErr, ok := err.(autorest.DetailedError)
	c.Assert(ok, chk.Equals, true)
	c.Assert(detErr.StatusCode, chk.Equals, http.StatusBadRequest)

	count := 0
	for iter, err := pagingClient.GetMultiplePagesFailureComplete(context.Background()); err == nil; err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value().Properties, chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 1)
}

func (s *PagingGroupSuite) TestGetMultiplePagesFailureURI(c *chk.C) {
	page, err := pagingClient.GetMultiplePagesFailureURI(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*page.Response().NextLink, chk.Equals, "*&*#&$")
	err = page.Next()
	c.Assert(err, chk.NotNil)
	c.Assert(err, chk.ErrorMatches, ".*No scheme detected in URL.*")

	count := 0
	for iter, err := pagingClient.GetMultiplePagesFailureURIComplete(context.Background()); err == nil; err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value(), chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 1)
	c.Assert(err, chk.NotNil)
	c.Assert(err, chk.ErrorMatches, ".*No scheme detected in URL.*")
}

func (s *PagingGroupSuite) TestGetMultiplePagesFragmentNextLink(c *chk.C) {
	count := 0
	for page, err := pagingClient.GetMultiplePagesFragmentNextLink(context.Background(), "1.6", "test_user"); page.NotDone(); err = page.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(page.Values(), chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)

	count = 0
	for iter, err := pagingClient.GetMultiplePagesFragmentNextLinkComplete(context.Background(), "1.6", "test_user"); iter.NotDone(); err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value(), chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)
}

func (s *PagingGroupSuite) TestGetMultiplePagesLRO(c *chk.C) {
	future, err := pagingClient.GetMultiplePagesLRO(context.Background(), clientID, nil, nil)
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), pagingClient.Client)
	c.Assert(err, chk.IsNil)
	count := 0
	for page, err := future.Result(pagingClient); page.NotDone(); err = page.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(page.Values(), chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)
	// iterator
	futureIter, err := pagingClient.GetMultiplePagesLROComplete(context.Background(), clientID, nil, nil)
	c.Assert(err, chk.IsNil)
	count = 0
	for iter, err := futureIter.Result(pagingClient); iter.NotDone(); err = iter.Next() {
		c.Assert(err, chk.IsNil)
		c.Assert(iter.Value().Properties, chk.NotNil)
		count++
	}
	c.Assert(count, chk.Equals, 10)
}
