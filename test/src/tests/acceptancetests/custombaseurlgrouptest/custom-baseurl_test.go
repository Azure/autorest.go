package custombaseurlgrouptest

import (
	"context"
	"testing"
	. "tests/generated/custombaseurlgroup"

	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type CustomBaseURLGroupSuite struct{}

var _ = chk.Suite(&CustomBaseURLGroupSuite{})

var custombaseuriClient = getCustomBaseURIClient()

func getCustomBaseURIClient() PathsClient {
	c := NewWithoutDefaults("host:3000")
	c.RetryDuration = 1
	return PathsClient{BaseClient: c}
}

func (s *CustomBaseURLGroupSuite) TestCustomBaseUriPositive(c *chk.C) {
	_, err := custombaseuriClient.GetEmpty(context.Background(), "local")
	c.Assert(err, chk.IsNil)
}

func (s *CustomBaseURLGroupSuite) TestCustomBaseUriNegative(c *chk.C) {
	_, err := custombaseuriClient.GetEmpty(context.Background(), "badhost:3000")
	c.Assert(err, chk.NotNil)

	custombaseuriClient.RetryAttempts = 0
	_, err = custombaseuriClient.GetEmpty(context.Background(), "bad")
	c.Assert(err, chk.NotNil)
}
