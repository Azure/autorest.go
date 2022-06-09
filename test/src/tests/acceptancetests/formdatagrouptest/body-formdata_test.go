package formdatagrouptest

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"
	"tests/acceptancetests/utils"
	. "tests/generated/formdatagroup"

	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type FormdataSuite struct{}

var _ = chk.Suite(&FormdataSuite{})
var formdataClient = getFormdataClient()

func getFormdataClient() FormdataClient {
	c := NewFormdataClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *FormdataSuite) TestUploadFileViaBody(c *chk.C) {
	f, err := ioutil.ReadFile("../sample.png")
	c.Assert(err, chk.IsNil)
	res, err := formdataClient.UploadFileViaBody(context.Background(), ioutil.NopCloser(bytes.NewReader(f)))
	c.Assert(err, chk.IsNil)
	buf := new(bytes.Buffer)
	buf.ReadFrom(*res.Value)
	b := buf.Bytes()
	defer (*res.Value).Close()
	c.Assert(len(b), chk.Equals, len(f))
	c.Assert(string(b), chk.Equals, string(f))
}

func (s *FormdataSuite) TestUploadFile(c *chk.C) {
	c.Skip("server returning HTTP 500, needs investigation")
	f, err := ioutil.ReadFile("../sample.png")
	c.Assert(err, chk.IsNil)
	res, err := formdataClient.UploadFile(context.Background(), ioutil.NopCloser(bytes.NewReader(f)), "samplefile")
	c.Assert(err, chk.IsNil)
	buf := new(bytes.Buffer)
	buf.ReadFrom(*res.Value)
	b := buf.Bytes()
	defer (*res.Value).Close()
	c.Assert(len(b), chk.Equals, len(f))
	c.Assert(string(b), chk.Equals, string(f))
}
