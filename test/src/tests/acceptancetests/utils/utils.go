package utils

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/go-autorest/autorest/date"
)

func ToDateTimeRFC1123(s string) date.TimeRFC1123 {
	t, _ := time.Parse(time.RFC1123, strings.ToUpper(s))
	return date.TimeRFC1123{t}
}

func ToDateTime(s string) date.Time {
	t, _ := time.Parse(time.RFC3339, strings.ToUpper(s))
	return date.Time{t}
}

func NewPipeline() pipeline.Pipeline {
	f := []pipeline.Factory{
		pipeline.MethodFactoryMarker(),
	}
	return pipeline.NewPipeline(f, pipeline.Options{})
}

func NewPipelineWithRetry() pipeline.Pipeline {
	f := []pipeline.Factory{
		pipeline.MethodFactoryMarker(),
		&SimpleRetryPolicyFactory{},
	}
	return pipeline.NewPipeline(f, pipeline.Options{HTTPSender: &HTTPSenderWithCookiesFactory{}})
}

type HTTPSenderWithCookiesFactory struct{}

func (swc HTTPSenderWithCookiesFactory) New(node pipeline.Policy, config *pipeline.Configuration) pipeline.Policy {
	j, _ := cookiejar.New(nil)
	return &HTTPSenderWithCookiesPolicy{
		sender: &http.Client{
			Jar: j,
		},
	}
}

type HTTPSenderWithCookiesPolicy struct {
	sender *http.Client
}

func (swc HTTPSenderWithCookiesPolicy) Do(ctx context.Context, request pipeline.Request) (pipeline.Response, error) {
	resp, err := swc.sender.Do(request.Request)
	return pipeline.NewHTTPResponse(resp), err
}

type SimpleRetryPolicyFactory struct{}

func (srpf SimpleRetryPolicyFactory) New(node pipeline.Policy, config *pipeline.Configuration) pipeline.Policy {
	return &SimpleRetryPolicy{
		node: node,
		statusCodesForRetry: []int{
			http.StatusRequestTimeout,      // 408
			http.StatusTooManyRequests,     // 429
			http.StatusInternalServerError, // 500
			http.StatusBadGateway,          // 502
			http.StatusServiceUnavailable,  // 503
			http.StatusGatewayTimeout,      // 504
		},
		attempts:             3,
		delayBetweenAttempts: 1 * time.Second,
	}
}

type SimpleRetryPolicy struct {
	node                 pipeline.Policy
	statusCodesForRetry  []int
	attempts             int
	delayBetweenAttempts time.Duration
}

func (srp SimpleRetryPolicy) Do(ctx context.Context, request pipeline.Request) (resp pipeline.Response, err error) {
	for try := 0; try < srp.attempts; try++ {
		reqCopy := request.Copy()
		if try > 0 {
			err = reqCopy.RewindBody()
			if err != nil {
				panic(err)
			}
		}
		resp, err = srp.node.Do(ctx, reqCopy)
		fmt.Printf("retry attempt %v, resp: %v, err: %v\n", try+1, resp, err)
		if err == nil && !srp.shouldRetry(resp.Response().StatusCode) {
			return
		}
		time.Sleep(srp.delayBetweenAttempts)
	}
	return
}

// returns true if statusCode is in the slice of statusCodesForRetry
func (srp SimpleRetryPolicy) shouldRetry(statusCode int) bool {
	for _, v := range srp.statusCodesForRetry {
		if statusCode == v {
			return true
		}
	}
	return false
}
