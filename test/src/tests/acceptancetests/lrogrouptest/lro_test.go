package lrogrouptest

import (
	"testing"
	"time"

	chk "gopkg.in/check.v1"

	"tests/acceptancetests/utils"
	. "tests/generated/lro"
)

func Test(t *testing.T) { chk.TestingT(t) }

type LROSuite struct{}

var _ = chk.Suite(&LROSuite{})

var lroRetryClient = getLRORetrysClient()
var lrosClient = getLROsClient()
var lroSADSClient = getLROSADsClient()
var lroCustomHeaderClient = getLROsCustomHeaderClient()

func getLRORetrysClient() LRORetrysClient {
	c := NewLRORetrysClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getLROsClient() LROsClient {
	c := NewLROsClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getLROSADsClient() LROSADsClient {
	c := NewLROSADsClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getLROsCustomHeaderClient() LROsCustomHeaderClient {
	c := NewLROsCustomHeaderClient()
	c.RetryDuration = 1
	c.BaseURI = utils.GetBaseURI()
	return c
}

func (s *LROSuite) TestDelete202NoRetry204(c *chk.C) {
	future, err := lrosClient.Delete202NoRetry204()
	c.Assert(err, chk.IsNil)

	for done, err := future.Done(lrosClient); !done; done, err = future.Done(lrosClient) {
		c.Assert(err, chk.IsNil)
		dur, ok := future.GetPollingDelay()
		c.Assert(ok, chk.Equals, true)
		time.Sleep(dur)
	}
	c.Assert(future.Response().StatusCode, chk.Equals, 204)
}

func (s *LROSuite) TestDelete202Retry200(c *chk.C) {
	future, err := lroRetryClient.Delete202Retry200()
	c.Assert(err, chk.NotNil)

	future, err = lroRetryClient.Delete202Retry200()
	c.Assert(err, chk.IsNil)

	for done, err := future.Done(lroRetryClient); !done; done, err = future.Done(lroRetryClient) {
		c.Assert(err, chk.IsNil)
		dur, ok := future.GetPollingDelay()
		c.Assert(ok, chk.Equals, true)
		time.Sleep(dur)
	}
	c.Assert(future.Response().StatusCode, chk.Equals, 500)
}

func (s *LROSuite) TestDelete202NonRetry400(c *chk.C) {
	future, err := lroSADSClient.Delete202NonRetry400()
	c.Assert(err, chk.IsNil)

	for done, err := future.Done(lroSADSClient); !done; done, err = future.Done(lroSADSClient) {
		c.Assert(err, chk.IsNil)
		dur, ok := future.GetPollingDelay()
		c.Assert(ok, chk.Equals, true)
		time.Sleep(dur)
	}
	c.Assert(future.Response().StatusCode, chk.Equals, 400)
}
