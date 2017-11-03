package booleangrouptest

import (
	"context"
	"testing"

	chk "gopkg.in/check.v1"

	"tests/acceptancetests/utils"
	. "tests/generated/body-boolean"
)

func Test(t *testing.T) { chk.TestingT(t) }

type BoolGroupSuite struct{}

var _ = chk.Suite(&BoolGroupSuite{})

var boolClient = getBooleanClient()

func getBooleanClient() BoolClient {
	c := NewBoolClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *BoolGroupSuite) TestGetTrue(c *chk.C) {
	res, err := boolClient.GetTrue(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, true)
}

func (s *BoolGroupSuite) TestPutTrue(c *chk.C) {
	_, err := boolClient.PutTrue(context.Background(), true)
	c.Assert(err, chk.IsNil)
}

func (s *BoolGroupSuite) TestGetFalse(c *chk.C) {
	res, err := boolClient.GetFalse(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.Equals, false)
}

func (s *BoolGroupSuite) TestPutFalse(c *chk.C) {
	_, err := boolClient.PutFalse(context.Background(), false)
	c.Assert(err, chk.IsNil)
}

func (s *BoolGroupSuite) TestGetInvalidBool(c *chk.C) {
	_, err := boolClient.GetInvalid(context.Background())
	c.Assert(err, chk.NotNil)
}

func (s *BoolGroupSuite) TestGetNullBool(c *chk.C) {
	res, err := boolClient.GetNull(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.Value, chk.IsNil)
}
