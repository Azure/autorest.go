package bytegrouptest

import (
	"bytes"
	"context"
	"testing"

	chk "gopkg.in/check.v1"

	"tests/acceptancetests/utils"
	. "tests/generated/body-byte"
)

func Test(t *testing.T) { chk.TestingT(t) }

type ByteGroupSuite struct{}

var _ = chk.Suite(&ByteGroupSuite{})

var byteClient = getByteClient()

func getByteClient() ByteClient {
	c := NewByteClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *ByteGroupSuite) TestGetNonASCII(c *chk.C) {
	res, err := byteClient.GetNonASCII(context.Background())
	c.Assert(err, chk.IsNil)
	if !bytes.Equal(*res.Value, []byte{255, 254, 253, 252, 251, 250, 249, 248, 247, 246}) {
		c.Errorf("%v\n", *res.Value)
	}
}

func (s *ByteGroupSuite) TestGetEmptyByte(c *chk.C) {
	res, err := byteClient.GetEmpty(context.Background())
	c.Assert(err, chk.IsNil)
	if !bytes.Equal(*res.Value, nil) {
		c.Errorf("%v\n", *res.Value)
	}
}

func (s *ByteGroupSuite) TestGetInvalidByte(c *chk.C) {
	_, err := byteClient.GetInvalid(context.Background())
	c.Assert(err, chk.NotNil)
}

func (s *ByteGroupSuite) TestGetNullByte(c *chk.C) {
	res, err := byteClient.GetNull(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(res.Value, chk.IsNil)
}

func (s *ByteGroupSuite) TestPutNonASCII(c *chk.C) {
	_, err := byteClient.PutNonASCII(context.Background(), []byte{255, 254, 253, 252, 251, 250, 249, 248, 247, 246})
	c.Assert(err, chk.IsNil)
}
