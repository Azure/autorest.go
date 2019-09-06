package urlmultigrouptest

import (
	"context"
	"net/url"
	"testing"
	"tests/acceptancetests/utils"
	. "tests/generated/urlmultigroup"

	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type URLMultiSuite struct{}

var _ = chk.Suite(&URLMultiSuite{})

var queryClient = getQueryClient()

func getQueryClient() QueriesClient {
	c := NewQueriesClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *URLMultiSuite) TestArrayStringMultiEmpty(c *chk.C) {
	_, err := queryClient.ArrayStringMultiEmpty(context.Background(), []string{""})
	c.Assert(err, chk.IsNil)
}

func (s *URLMultiSuite) TestArrayStringMultiNull(c *chk.C) {
	_, err := queryClient.ArrayStringMultiNull(context.Background(), nil)
	c.Assert(err, chk.IsNil)
}

func (s *URLMultiSuite) TestArrayStringMultiValid(c *chk.C) {
	_, err := queryClient.ArrayStringMultiValid(context.Background(), []string{
		"ArrayQuery1",
		url.QueryEscape("begin!*'();:@ &=+$,/?#[]end"),
		"",
		"",
	})
	c.Assert(err, chk.IsNil)
}
