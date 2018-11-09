package lrogrouptest

import (
	"context"
	"net/http"
	"testing"
	"tests/acceptancetests/utils"
	"tests/generated/lrogroup"
	"time"

	"github.com/Azure/go-autorest/autorest/to"
	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type LROSuite struct{}

var _ = chk.Suite(&LROSuite{})

var lroRetryClient = getLRORetrysClient()
var lrosClient = getLROsClient()
var lroSADSClient = getLROSADsClient()
var lroCustomHeaderClient = getLROsCustomHeaderClient()

func getLRORetrysClient() lrogroup.LRORetrysClient {
	c := lrogroup.NewLRORetrysClient()
	c.RetryDuration = 1
	c.PollingDelay = time.Second
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getLROsClient() lrogroup.LROsClient {
	c := lrogroup.NewLROsClient()
	c.RetryDuration = 1
	c.PollingDelay = time.Second
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getLROSADsClient() lrogroup.LROSADsClient {
	c := lrogroup.NewLROSADsClient()
	c.RetryDuration = 1
	c.PollingDelay = time.Second
	c.BaseURI = utils.GetBaseURI()
	return c
}

func getLROsCustomHeaderClient() lrogroup.LROsCustomHeaderClient {
	c := lrogroup.NewLROsCustomHeaderClient()
	c.RetryDuration = 1
	c.PollingDelay = time.Second
	c.BaseURI = utils.GetBaseURI()
	return c
}

// retry client

func (s *LROSuite) TestRetryDelete202Retry200(c *chk.C) {
	future, err := lroRetryClient.Delete202Retry200(context.Background())
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroRetryClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lroRetryClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
}

func (s *LROSuite) TestRetryDeleteAsyncRelativeRetrySucceeded(c *chk.C) {
	future, err := lroRetryClient.DeleteAsyncRelativeRetrySucceeded(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(future.Response().StatusCode, chk.Equals, http.StatusAccepted)
	err = future.WaitForCompletionRef(context.Background(), lroRetryClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestRetryDeleteProvisioning202Accepted200Succeeded(c *chk.C) {
	future, err := lroRetryClient.DeleteProvisioning202Accepted200Succeeded(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(future.Response().StatusCode, chk.Equals, http.StatusAccepted)
	err = future.WaitForCompletionRef(context.Background(), lroRetryClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestRetryPost202Retry200(c *chk.C) {
	future, err := lroRetryClient.Post202Retry200(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroRetryClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lroRetryClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
}

func (s *LROSuite) TestRetryPostAsyncRelativeRetrySucceeded(c *chk.C) {
	future, err := lroRetryClient.PostAsyncRelativeRetrySucceeded(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroRetryClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lroRetryClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
}

func (s *LROSuite) TestRetryPut201CreatingSucceeded200(c *chk.C) {
	future, err := lroRetryClient.Put201CreatingSucceeded200(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroRetryClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lroRetryClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
	c.Assert(r.Name, chk.NotNil)
}

func (s *LROSuite) TestRetryPutAsyncRelativeRetrySucceeded(c *chk.C) {
	future, err := lroRetryClient.PutAsyncRelativeRetrySucceeded(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroRetryClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lroRetryClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
	c.Assert(r.Name, chk.NotNil)
}

// vanilla client

func (s *LROSuite) TestDelete202NoRetry204(c *chk.C) {
	future, err := lrosClient.Delete202NoRetry204(context.Background())
	c.Assert(err, chk.IsNil)
	for done, err := future.Done(lrosClient); !done; done, err = future.Done(lrosClient) {
		c.Assert(err, chk.IsNil)
	}
	c.Assert(future.Response().StatusCode, chk.Equals, 204)
}

func (s *LROSuite) TestDelete202Retry200(c *chk.C) {
	future, err := lrosClient.Delete202Retry200(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(future.Response().StatusCode, chk.Equals, http.StatusAccepted)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, true)
	p, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(p.Name, chk.IsNil)
}

func (s *LROSuite) TestDelete204Succeeded(c *chk.C) {
	future, err := lrosClient.Delete204Succeeded(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, true)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusNoContent)
}

func (s *LROSuite) TestDeleteAsyncNoHeaderInRetry(c *chk.C) {
	future, err := lrosClient.DeleteAsyncNoHeaderInRetry(context.Background())
	c.Assert(err, chk.IsNil)
	c.Assert(future.Response().StatusCode, chk.Equals, http.StatusAccepted)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, false)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
}

func (s *LROSuite) TestDeleteAsyncNoRetrySucceeded(c *chk.C) {
	future, err := lrosClient.DeleteAsyncNoRetrySucceeded(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, false)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
}

func (s *LROSuite) TestDeleteAsyncRetrycanceled(c *chk.C) {
	future, err := lrosClient.DeleteAsyncRetrycanceled(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, false)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.NotNil)
	c.Assert(r.Response, chk.IsNil)
}

func (s *LROSuite) TestDeleteAsyncRetryFailed(c *chk.C) {
	future, err := lrosClient.DeleteAsyncRetryFailed(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, false)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.NotNil)
	c.Assert(r.Response, chk.IsNil)
}

func (s *LROSuite) TestDeleteAsyncRetrySucceeded(c *chk.C) {
	future, err := lrosClient.DeleteAsyncRetrySucceeded(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, false)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
}

func (s *LROSuite) TestDeleteNoHeaderInRetry(c *chk.C) {
	future, err := lrosClient.DeleteNoHeaderInRetry(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, false)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusNoContent)
}

func (s *LROSuite) TestDeleteProvisioning202Accepted200Succeeded(c *chk.C) {
	future, err := lrosClient.DeleteProvisioning202Accepted200Succeeded(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, true)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
	c.Assert(r.Name, chk.NotNil)
}

func (s *LROSuite) TestDeleteProvisioning202Deletingcanceled200(c *chk.C) {
	future, err := lrosClient.DeleteProvisioning202Deletingcanceled200(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.NotNil)
	c.Assert(done, chk.Equals, true)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.NotNil)
	c.Assert(r.Response, chk.NotNil)
}

func (s *LROSuite) TestDeleteProvisioning202DeletingFailed200(c *chk.C) {
	future, err := lrosClient.DeleteProvisioning202DeletingFailed200(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lrosClient)
	c.Assert(err, chk.NotNil)
	c.Assert(done, chk.Equals, true)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.NotNil)
	c.Assert(r.Response, chk.NotNil)
}

func (s *LROSuite) TestPost200WithPayload(c *chk.C) {
	future, err := lrosClient.Post200WithPayload(context.Background())
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.ID, chk.NotNil)
}

func (s *LROSuite) TestPost202NoRetry204(c *chk.C) {
	future, err := lrosClient.Post202NoRetry204(context.Background(), &lrogroup.Product{
		Location: to.StringPtr("West US"),
	})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.Response.StatusCode, chk.Equals, http.StatusNoContent)
	c.Assert(r.ID, chk.IsNil)
}

func (s *LROSuite) TestPost202Retry200(c *chk.C) {
	future, err := lrosClient.Post202Retry200(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)
}

func (s *LROSuite) TestPostAsyncNoRetrySucceeded(c *chk.C) {
	future, err := lrosClient.PostAsyncNoRetrySucceeded(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lrosClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.ID, chk.NotNil)
}

func (s *LROSuite) TestPostAsyncRetrycanceled(c *chk.C) {
	future, err := lrosClient.PostAsyncRetrycanceled(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestPostAsyncRetryFailed(c *chk.C) {
	future, err := lrosClient.PostAsyncRetryFailed(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestPostAsyncRetrySucceeded(c *chk.C) {
	future, err := lrosClient.PostAsyncRetrySucceeded(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPut200Acceptedcanceled200(c *chk.C) {
	future, err := lrosClient.Put200Acceptedcanceled200(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestPut200Succeeded(c *chk.C) {
	future, err := lrosClient.Put200Succeeded(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPut200SucceededNoState(c *chk.C) {
	future, err := lrosClient.Put200SucceededNoState(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPut200UpdatingSucceeded204(c *chk.C) {
	future, err := lrosClient.Put200UpdatingSucceeded204(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPut201CreatingFailed200(c *chk.C) {
	future, err := lrosClient.Put201CreatingFailed200(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestPut201CreatingSucceeded200(c *chk.C) {
	future, err := lrosClient.Put201CreatingSucceeded200(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPut202Retry200(c *chk.C) {
	future, err := lrosClient.Put202Retry200(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPutAsyncNoHeaderInRetry(c *chk.C) {
	future, err := lrosClient.PutAsyncNoHeaderInRetry(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPutAsyncNonResource(c *chk.C) {
	future, err := lrosClient.PutAsyncNonResource(context.Background(), &lrogroup.Sku{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPutAsyncNoRetrycanceled(c *chk.C) {
	future, err := lrosClient.PutAsyncNoRetrycanceled(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestPutAsyncNoRetrySucceeded(c *chk.C) {
	future, err := lrosClient.PutAsyncNoRetrySucceeded(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPutAsyncRetryFailed(c *chk.C) {
	future, err := lrosClient.PutAsyncRetryFailed(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestPutAsyncRetrySucceeded(c *chk.C) {
	future, err := lrosClient.PutAsyncRetrySucceeded(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPutAsyncSubResource(c *chk.C) {
	future, err := lrosClient.PutAsyncSubResource(context.Background(), &lrogroup.SubProduct{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPutNoHeaderInRetry(c *chk.C) {
	future, err := lrosClient.PutNoHeaderInRetry(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPutNonResource(c *chk.C) {
	future, err := lrosClient.PutNonResource(context.Background(), &lrogroup.Sku{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

func (s *LROSuite) TestPutSubResource(c *chk.C) {
	future, err := lrosClient.PutSubResource(context.Background(), &lrogroup.SubProduct{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lrosClient.Client)
	c.Assert(err, chk.IsNil)
}

// sads client

func (s *LROSuite) TestSADsDelete202NonRetry400(c *chk.C) {
	future, err := lroSADSClient.Delete202NonRetry400(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lroSADSClient)
	c.Assert(done, chk.Equals, false)
	c.Assert(err, chk.NotNil)
	c.Assert(future.Response().StatusCode, chk.Equals, 400)
}

func (s *LROSuite) TestSADsDelete202RetryInvalidHeader(c *chk.C) {
	_, err := lroSADSClient.Delete202RetryInvalidHeader(context.Background())
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsDelete204Succeeded(c *chk.C) {
	future, err := lroSADSClient.Delete204Succeeded(context.Background())
	c.Assert(err, chk.IsNil)
	done, err := future.Done(lroSADSClient)
	c.Assert(err, chk.IsNil)
	c.Assert(done, chk.Equals, true)
}

func (s *LROSuite) TestSADsDeleteAsyncRelativeRetry400(c *chk.C) {
	future, err := lroSADSClient.DeleteAsyncRelativeRetry400(context.Background())
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroSADSClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsDeleteAsyncRelativeRetryInvalidHeader(c *chk.C) {
	_, err := lroSADSClient.DeleteAsyncRelativeRetryInvalidHeader(context.Background())
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsDeleteAsyncRelativeRetryInvalidJSONPolling(c *chk.C) {
	future, err := lroSADSClient.DeleteAsyncRelativeRetryInvalidJSONPolling(context.Background())
	c.Assert(err, chk.IsNil)
	_, err = future.Done(lroSADSClient)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsDeleteAsyncRelativeRetryNoStatus(c *chk.C) {
	future, err := lroSADSClient.DeleteAsyncRelativeRetryNoStatus(context.Background())
	c.Assert(err, chk.IsNil)
	_, err = future.Done(lroSADSClient)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsDeleteNonRetry400(c *chk.C) {
	_, err := lroSADSClient.DeleteNonRetry400(context.Background())
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPost202NoLocation(c *chk.C) {
	_, err := lroSADSClient.Post202NoLocation(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPost202NonRetry400(c *chk.C) {
	future, err := lroSADSClient.Post202NonRetry400(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	b, err := future.Done(lroSADSClient)
	c.Assert(err, chk.NotNil)
	c.Assert(b, chk.Equals, false)
}

func (s *LROSuite) TestSADsPost202RetryInvalidHeader(c *chk.C) {
	_, err := lroSADSClient.Post202RetryInvalidHeader(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPostAsyncRelativeRetry400(c *chk.C) {
	future, err := lroSADSClient.PostAsyncRelativeRetry400(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	b, err := future.Done(lroSADSClient)
	c.Assert(err, chk.NotNil)
	c.Assert(b, chk.Equals, false)
}

func (s *LROSuite) TestSADsPostAsyncRelativeRetryInvalidHeader(c *chk.C) {
	_, err := lroSADSClient.PostAsyncRelativeRetryInvalidHeader(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPostAsyncRelativeRetryInvalidJSONPolling(c *chk.C) {
	future, err := lroSADSClient.PostAsyncRelativeRetryInvalidJSONPolling(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	_, err = future.Done(lroSADSClient)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPostAsyncRelativeRetryNoPayload(c *chk.C) {
	future, err := lroSADSClient.PostAsyncRelativeRetryNoPayload(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	_, err = future.Done(lroSADSClient)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPostNonRetry400(c *chk.C) {
	future, err := lroSADSClient.PostNonRetry400(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.NotNil)
	_, err = future.Done(lroSADSClient)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPut200InvalidJSON(c *chk.C) {
	_, err := lroSADSClient.Put200InvalidJSON(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutAsyncRelativeRetry400(c *chk.C) {
	future, err := lroSADSClient.PutAsyncRelativeRetry400(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroSADSClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutAsyncRelativeRetryInvalidHeader(c *chk.C) {
	_, err := lroSADSClient.PutAsyncRelativeRetryInvalidHeader(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutAsyncRelativeRetryInvalidJSONPolling(c *chk.C) {
	future, err := lroSADSClient.PutAsyncRelativeRetryInvalidJSONPolling(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroSADSClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutAsyncRelativeRetryNoStatus(c *chk.C) {
	future, err := lroSADSClient.PutAsyncRelativeRetryNoStatus(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroSADSClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutAsyncRelativeRetryNoStatusPayload(c *chk.C) {
	future, err := lroSADSClient.PutAsyncRelativeRetryNoStatusPayload(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroSADSClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutError201NoProvisioningStatePayload(c *chk.C) {
	future, err := lroSADSClient.PutError201NoProvisioningStatePayload(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	_, err = future.Result(lroSADSClient)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutNonRetry201Creating400(c *chk.C) {
	future, err := lroSADSClient.PutNonRetry201Creating400(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroSADSClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutNonRetry201Creating400InvalidJSON(c *chk.C) {
	future, err := lroSADSClient.PutNonRetry201Creating400InvalidJSON(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroSADSClient.Client)
	c.Assert(err, chk.NotNil)
}

func (s *LROSuite) TestSADsPutNonRetry400(c *chk.C) {
	_, err := lroSADSClient.PutNonRetry400(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.NotNil)
}

// custom header client

// NOTE: at this time the Go SDK doesn't expose means to include custom headers in LROs.
//       while you can add it to the initial request by calling the preparer and sender
//       explicitly (see below) there is no way to include the custom header data in the
//       LRO polling request.  as a result these tests will all fail so they are empty for now.

func (s *LROSuite) TestCustomHeaderPost202Retry200(c *chk.C) {
	/*req, err := lroCustomHeaderClient.Post202Retry200Preparer(context.Background(), &lrogroup.Product{})
	c.Assert(err, chk.IsNil)
	req.Header.Add("x-ms-client-request-id", "9C4D50EE-2D56-4CD3-8152-34347DC9F2B0")
	future, err := lroCustomHeaderClient.Post202Retry200Sender(req)
	c.Assert(err, chk.IsNil)
	err = future.WaitForCompletionRef(context.Background(), lroCustomHeaderClient.Client)
	c.Assert(err, chk.IsNil)
	r, err := future.Result(lroCustomHeaderClient)
	c.Assert(err, chk.IsNil)
	c.Assert(r.StatusCode, chk.Equals, http.StatusOK)*/
}

func (s *LROSuite) TestCustomHeaderPostAsyncRetrySucceeded(c *chk.C) {
	// TODO
}

func (s *LROSuite) TestCustomHeaderPut201CreatingSucceeded200(c *chk.C) {
	// TODO
}

func (s *LROSuite) TestCustomHeaderPutAsyncRetrySucceeded(c *chk.C) {
	// TODO
}
