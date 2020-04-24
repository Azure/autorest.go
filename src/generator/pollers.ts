/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { pascalCase } from '@azure-tools/codegen';
import { CodeModel } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { PollerInfo } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';

// Creates the content in pollers.go
export async function generatePollers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('net/http');
  imports.add('time');
  text += imports.text();

  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  for (const poller of values(pollers)) {
    const pollerInterface = pascalCase(poller.name);
    let responseType = '';
    if (poller.schema === undefined) {
      responseType = 'http.Response';
    } else {
      responseType = poller.schema.language.go!.responseType.name;
    }
    text += `// ${pollerInterface} provides polling facilities until the operation completes
type ${pollerInterface} interface {
	Done() bool
	ID() string
	Poll(context.Context) (*${responseType}, error)
	Wait(ctx context.Context, pollingInterval time.Duration) (*${responseType}, error)
}

type ${poller.name} struct {
	// the client for making the request
	client *${poller.client}
	// polling tracker
	pt pollingTracker
}

// Done returns true if the polling operation has terminated either in a success case or failure case,
// otherwise it will return false
func (p *${poller.name}) Done() bool {
	return p.pt.hasTerminated()
}

// ID ... 
func (p *${poller.name}) ID() string {
	return ""
}

func (p *${poller.name}) Poll(ctx context.Context) (*${responseType}, error) {
	if done, err := p.done(ctx); !done || err != nil {
		return nil, err
	}
	resp := p.response()
	result, err := p.client.${poller.operationName}HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *${poller.name}) Wait(ctx context.Context, pollingInterval time.Duration) (*${responseType}, error) {
	return nil, nil
}

// Response returns the last HTTP response.
func (p *${poller.name}) response() *azcore.Response {
	if p.pt == nil {
		return nil
	}
	return p.pt.latestResponse()
}

// done queries the service to see if the operation has completed.
func (p *${poller.name}) done(ctx context.Context) (done bool, err error) {
	if p.pt.hasTerminated() {
		return true, p.pt.pollingError()
	}
	if err := p.pt.pollForStatus(ctx, p.client.p); err != nil {
		return false, err
	}
	if err := p.pt.checkForErrors(); err != nil {
		return p.pt.hasTerminated(), err
	}
	if err := p.pt.updatePollingState(p.pt.provisioningStateApplicable()); err != nil {
		return false, err
	}
	if err := p.pt.initPollingMethod(); err != nil {
		return false, err
	}
	if err := p.pt.updatePollingMethod(); err != nil {
		return false, err
	}
	return p.pt.hasTerminated(), p.pt.pollingError()
}

`;
  }
  return text;
}

// Creates the content in pollers_helper.go
export async function generatePollersHelper(session: Session<CodeModel>): Promise<string> {
	if (session.model.language.go!.pollerTypes === undefined) {
	  return '';
	}
	let text = await contentPreamble(session);
	const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  	pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  
	// add standard imports
	const imports = new ImportManager();
	imports.add('context');
	imports.add('encoding/json');
	imports.add('errors');
	imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
	imports.add('io/ioutil');
	imports.add('net/http');
	imports.add('net/url');
	imports.add('strings');
	imports.add('fmt');
	imports.add('errors');
	text += imports.text();
	let errorCtor = pollers[0].pollingError.language.go!.constructorName;
	// TODO separate this into manageable chunks of text by section of functionality
	  text += `
	  const (
		headerAsyncOperation = "Azure-AsyncOperation"
		headerLocation       = "Location"
	)
	
	const (
		operationInProgress string = "InProgress"
		operationCanceled   string = "Canceled"
		operationFailed     string = "Failed"
		operationSucceeded  string = "Succeeded"
	)
	
	var pollingCodes = [...]int{http.StatusNoContent, http.StatusAccepted, http.StatusCreated, http.StatusOK}
	
	type pollingTracker interface {
		// these methods can differ per tracker
	
		// checks the response headers and status code to determine the polling mechanism
		updatePollingMethod() error
	
		// checks the response for tracker-specific error conditions
		checkForErrors() error
	
		// returns true if provisioning state should be checked
		provisioningStateApplicable() bool
	
		// methods common to all trackers
	
		// initializes a tracker's polling URL and method, called for each iteration.
		// these values can be overridden by each polling tracker as required.
		initPollingMethod() error
	
		// initializes the tracker's internal state, call this when the tracker is created
		initializeState() error
	
		// makes an HTTP request to check the status of the LRO
		pollForStatus(ctx context.Context, client azcore.Pipeline) error
	
		// updates internal tracker state, call this after each call to pollForStatus
		updatePollingState(provStateApl bool) error
	
		// returns the error response from the service, can be nil
		pollingError() error
	
		// returns the polling method being used
		pollingMethod() pollingMethodType
	
		// returns the state of the LRO as returned from the service
		pollingStatus() string
	
		// returns the URL used for polling status
		pollingURL() string
	
		// returns the URL used for the final GET to retrieve the resource
		finalGetURL() string
	
		// returns true if the LRO is in a terminal state
		hasTerminated() bool
	
		// returns true if the LRO is in a failed terminal state
		hasFailed() bool
	
		// returns true if the LRO is in a successful terminal state
		hasSucceeded() bool
	
		// returns the cached HTTP response after a call to pollForStatus(), can be nil
		latestResponse() *azcore.Response
	}
	
	type pollingTrackerBase struct {
		// resp is the last response, either from the submission of the LRO or from polling
		resp *azcore.Response
	
		// method is the HTTP verb, this is needed for deserialization
		Method string \`json:"method"\`
	
		// rawBody is the raw JSON response body
		rawBody map[string]interface{}
	
		// denotes if polling is using async-operation or location header
		Pm pollingMethodType \`json:"pollingMethod"\`
	
		// the URL to poll for status
		URI string \`json:"pollingURI"\`
	
		// the state of the LRO as returned from the service
		State string \`json:"lroState"\`
	
		// the URL to GET for the final result
		FinalGetURI string \`json:"resultURI"\`
	
		// used to hold an error object returned from the service
		Err error \`json:"error,omitempty"\`
	}
	
	func (pt *pollingTrackerBase) initializeState() error {
		// determine the initial polling state based on response body and/or HTTP status
		// code.  this is applicable to the initial LRO response, not polling responses!
		pt.Method = pt.resp.Request.Method
		if err := pt.updateRawBody(); err != nil {
			return err
		}
		switch pt.resp.StatusCode {
		case http.StatusOK:
			if ps := pt.getProvisioningState(); ps != nil {
				pt.State = *ps
				if pt.hasFailed() {
					pt.updateErrorFromResponse()
					return pt.pollingError()
				}
			} else {
				pt.State = operationSucceeded
			}
		case http.StatusCreated:
			if ps := pt.getProvisioningState(); ps != nil {
				pt.State = *ps
			} else {
				pt.State = operationInProgress
			}
		case http.StatusAccepted:
			pt.State = operationInProgress
		case http.StatusNoContent:
			pt.State = operationSucceeded
		default:
			pt.State = operationFailed
			pt.updateErrorFromResponse()
			return pt.pollingError()
		}
		return pt.initPollingMethod()
	}
	
	func (pt pollingTrackerBase) getProvisioningState() *string {
		if pt.rawBody != nil && pt.rawBody["properties"] != nil {
			p := pt.rawBody["properties"].(map[string]interface{})
			if ps := p["provisioningState"]; ps != nil {
				s := ps.(string)
				return &s
			}
		}
		return nil
	}
	
	func (pt *pollingTrackerBase) updateRawBody() error {
		pt.rawBody = map[string]interface{}{}
		if pt.resp.ContentLength != 0 {
			defer pt.resp.Body.Close()
			b, err := ioutil.ReadAll(pt.resp.Body)
			if err != nil {
				return errors.New("failed to read response body")
			}
			// observed in 204 responses over HTTP/2.0; the content length is -1 but body is empty
			if len(b) == 0 {
				return nil
			}
			if err = json.Unmarshal(b, &pt.rawBody); err != nil {
				return errors.New("failed to unmarshal response body")
			}
		}
		return nil
	}
	
	func (pt *pollingTrackerBase) pollForStatus(ctx context.Context, client azcore.Pipeline) error {
		u, err := url.Parse(pt.URI)
		if err != nil {
			return err
		}
		req := azcore.NewRequest(http.MethodGet, *u)
		resp, err := client.Do(ctx, req)
		pt.resp = resp
		if err != nil {
			return errors.New("failed to send HTTP request")
		}
		if pt.resp.HasStatusCode(pollingCodes[:]...) {
			// reset the service error on success case
			pt.Err = nil
			err = pt.updateRawBody()
		} else {
			// check response body for error content
			pt.updateErrorFromResponse()
			err = pt.pollingError()
		}
		return err
	}
	
	// attempts to unmarshal a ServiceError type from the response body.
	// if that fails then make a best attempt at creating something meaningful.
	// NOTE: this assumes that the async operation has failed.
	func (pt *pollingTrackerBase) updateErrorFromResponse() {
		pt.Err = ${errorCtor}(pt.resp)
	}
	
	func (pt *pollingTrackerBase) updatePollingState(provStateApl bool) error {
		if pt.Pm == pollingAsyncOperation && pt.rawBody["status"] != nil {
			pt.State = pt.rawBody["status"].(string)
		} else {
			if pt.resp.StatusCode == http.StatusAccepted {
				pt.State = operationInProgress
			} else if provStateApl {
				if ps := pt.getProvisioningState(); ps != nil {
					pt.State = *ps
				} else {
					pt.State = operationSucceeded
				}
			} else {
				return errors.New("the response from the async operation has an invalid status code")
			}
		}
		// if the operation has failed update the error state
		if pt.hasFailed() {
			pt.updateErrorFromResponse()
		}
		return nil
	}
	
	func (pt pollingTrackerBase) pollingError() error {
		return pt.Err
	}
	
	func (pt pollingTrackerBase) pollingMethod() pollingMethodType {
		return pt.Pm
	}
	
	func (pt pollingTrackerBase) pollingStatus() string {
		return pt.State
	}
	
	func (pt pollingTrackerBase) pollingURL() string {
		return pt.URI
	}
	
	func (pt pollingTrackerBase) finalGetURL() string {
		return pt.FinalGetURI
	}
	
	func (pt pollingTrackerBase) hasTerminated() bool {
		return strings.EqualFold(pt.State, operationCanceled) || strings.EqualFold(pt.State, operationFailed) || strings.EqualFold(pt.State, operationSucceeded)
	}
	
	func (pt pollingTrackerBase) hasFailed() bool {
		return strings.EqualFold(pt.State, operationCanceled) || strings.EqualFold(pt.State, operationFailed)
	}
	
	func (pt pollingTrackerBase) hasSucceeded() bool {
		return strings.EqualFold(pt.State, operationSucceeded)
	}
	
	func (pt pollingTrackerBase) latestResponse() *azcore.Response {
		return pt.resp
	}
	
	// error checking common to all trackers
	func (pt pollingTrackerBase) baseCheckForErrors() error {
		// for Azure-AsyncOperations the response body cannot be nil or empty
		if pt.Pm == pollingAsyncOperation {
			if pt.resp.Body == nil || pt.resp.ContentLength == 0 {
				return errors.New("for Azure-AsyncOperation response body cannot be nil")
			}
			if pt.rawBody["status"] == nil {
				return errors.New("missing status property in Azure-AsyncOperation response body")
			}
		}
		return nil
	}
	
	// default initialization of polling URL/method.  each verb tracker will update this as required.
	func (pt *pollingTrackerBase) initPollingMethod() error {
		if ao, err := getURLFromAsyncOpHeader(pt.resp); err != nil {
			return err
		} else if ao != "" {
			pt.URI = ao
			pt.Pm = pollingAsyncOperation
			return nil
		}
		if lh, err := getURLFromLocationHeader(pt.resp); err != nil {
			return err
		} else if lh != "" {
			pt.URI = lh
			pt.Pm = pollingLocation
			return nil
		}
		// it's ok if we didn't find a polling header, this will be handled elsewhere
		return nil
	}
	
	// DELETE
	
	type pollingTrackerDelete struct {
		pollingTrackerBase
	}
	
	func (pt *pollingTrackerDelete) updatePollingMethod() error {
		// for 201 the Location header is required
		if pt.resp.StatusCode == http.StatusCreated {
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil {
				return err
			} else if lh == "" {
				return errors.New("missing Location header in 201 response")
			} else {
				pt.URI = lh
			}
			pt.Pm = pollingLocation
			pt.FinalGetURI = pt.URI
		}
		// for 202 prefer the Azure-AsyncOperation header but fall back to Location if necessary
		if pt.resp.StatusCode == http.StatusAccepted {
			ao, err := getURLFromAsyncOpHeader(pt.resp)
			if err != nil {
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
			// if the Location header is invalid and we already have a polling URL
			// then we don't care if the Location header URL is malformed.
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil && pt.URI == "" {
				return err
			} else if lh != "" {
				if ao == "" {
					pt.URI = lh
					pt.Pm = pollingLocation
				}
				// when both headers are returned we use the value in the Location header for the final GET
				pt.FinalGetURI = lh
			}
			// make sure a polling URL was found
			if pt.URI == "" {
				return errors.New("didn't get any suitable polling URLs in 202 response")
			}
		}
		return nil
	}
	
	func (pt pollingTrackerDelete) checkForErrors() error {
		return pt.baseCheckForErrors()
	}
	
	func (pt pollingTrackerDelete) provisioningStateApplicable() bool {
		return pt.resp.StatusCode == http.StatusOK || pt.resp.StatusCode == http.StatusNoContent
	}
	
	// PATCH
	
	type pollingTrackerPatch struct {
		pollingTrackerBase
	}
	
	func (pt *pollingTrackerPatch) updatePollingMethod() error {
		// by default we can use the original URL for polling and final GET
		if pt.URI == "" {
			pt.URI = pt.resp.Request.URL.String()
		}
		if pt.FinalGetURI == "" {
			pt.FinalGetURI = pt.resp.Request.URL.String()
		}
		if pt.Pm == pollingUnknown {
			pt.Pm = pollingRequestURI
		}
		// for 201 it's permissible for no headers to be returned
		if pt.resp.StatusCode == http.StatusCreated {
			if ao, err := getURLFromAsyncOpHeader(pt.resp); err != nil {
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
		}
		// for 202 prefer the Azure-AsyncOperation header but fall back to Location if necessary
		// note the absence of the "final GET" mechanism for PATCH
		if pt.resp.StatusCode == http.StatusAccepted {
			ao, err := getURLFromAsyncOpHeader(pt.resp)
			if err != nil {
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
			if ao == "" {
				if lh, err := getURLFromLocationHeader(pt.resp); err != nil {
					return err
				} else if lh == "" {
					return errors.New("didn't get any suitable polling URLs in 202 response")
				} else {
					pt.URI = lh
					pt.Pm = pollingLocation
				}
			}
		}
		return nil
	}
	
	func (pt pollingTrackerPatch) checkForErrors() error {
		return pt.baseCheckForErrors()
	}
	
	func (pt pollingTrackerPatch) provisioningStateApplicable() bool {
		return pt.resp.StatusCode == http.StatusOK || pt.resp.StatusCode == http.StatusCreated
	}
	
	// POST
	
	type pollingTrackerPost struct {
		pollingTrackerBase
	}
	
	func (pt *pollingTrackerPost) updatePollingMethod() error {
		// 201 requires Location header
		if pt.resp.StatusCode == http.StatusCreated {
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil {
				return err
			} else if lh == "" {
				return errors.New("missing Location header in 201 response")
			} else {
				pt.URI = lh
				pt.FinalGetURI = lh
				pt.Pm = pollingLocation
			}
		}
		// for 202 prefer the Azure-AsyncOperation header but fall back to Location if necessary
		if pt.resp.StatusCode == http.StatusAccepted {
			ao, err := getURLFromAsyncOpHeader(pt.resp)
			if err != nil {
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
			// if the Location header is invalid and we already have a polling URL
			// then we don't care if the Location header URL is malformed.
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil && pt.URI == "" {
				return err
			} else if lh != "" {
				if ao == "" {
					pt.URI = lh
					pt.Pm = pollingLocation
				}
				// when both headers are returned we use the value in the Location header for the final GET
				pt.FinalGetURI = lh
			}
			// make sure a polling URL was found
			if pt.URI == "" {
				return errors.New("didn't get any suitable polling URLs in 202 response")
			}
		}
		return nil
	}
	
	func (pt pollingTrackerPost) checkForErrors() error {
		return pt.baseCheckForErrors()
	}
	
	func (pt pollingTrackerPost) provisioningStateApplicable() bool {
		return pt.resp.StatusCode == http.StatusOK || pt.resp.StatusCode == http.StatusNoContent
	}
	
	// PUT
	
	type pollingTrackerPut struct {
		pollingTrackerBase
	}
	
	func (pt *pollingTrackerPut) updatePollingMethod() error {
		// by default we can use the original URL for polling and final GET
		if pt.URI == "" {
			pt.URI = pt.resp.Request.URL.String()
		}
		if pt.FinalGetURI == "" {
			pt.FinalGetURI = pt.resp.Request.URL.String()
		}
		if pt.Pm == pollingUnknown {
			pt.Pm = pollingRequestURI
		}
		// for 201 it's permissible for no headers to be returned
		if pt.resp.StatusCode == http.StatusCreated {
			if ao, err := getURLFromAsyncOpHeader(pt.resp); err != nil {
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
		}
		// for 202 prefer the Azure-AsyncOperation header but fall back to Location if necessary
		if pt.resp.StatusCode == http.StatusAccepted {
			ao, err := getURLFromAsyncOpHeader(pt.resp)
			if err != nil {
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
			// if the Location header is invalid and we already have a polling URL
			// then we don't care if the Location header URL is malformed.
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil && pt.URI == "" {
				return err
			} else if lh != "" {
				if ao == "" {
					pt.URI = lh
					pt.Pm = pollingLocation
				}
			}
			// make sure a polling URL was found
			if pt.URI == "" {
				return errors.New("didn't get any suitable polling URLs in 202 response")
			}
		}
		return nil
	}
	
	func (pt pollingTrackerPut) checkForErrors() error {
		err := pt.baseCheckForErrors()
		if err != nil {
			return err
		}
		// if there are no LRO headers then the body cannot be empty
		ao, err := getURLFromAsyncOpHeader(pt.resp)
		if err != nil {
			return err
		}
		lh, err := getURLFromLocationHeader(pt.resp)
		if err != nil {
			return err
		}
		if ao == "" && lh == "" && len(pt.rawBody) == 0 {
			return errors.New("the response did not contain a body")
		}
		return nil
	}
	
	func (pt pollingTrackerPut) provisioningStateApplicable() bool {
		return pt.resp.StatusCode == http.StatusOK || pt.resp.StatusCode == http.StatusCreated
	}
	
	// creates a polling tracker based on the verb of the original request
	func createPollingTracker(resp *azcore.Response) (pollingTracker, error) {
		var pt pollingTracker
		switch strings.ToUpper(resp.Request.Method) {
		case http.MethodDelete:
			pt = &pollingTrackerDelete{pollingTrackerBase: pollingTrackerBase{resp: resp}}
		case http.MethodPatch:
			pt = &pollingTrackerPatch{pollingTrackerBase: pollingTrackerBase{resp: resp}}
		case http.MethodPost:
			pt = &pollingTrackerPost{pollingTrackerBase: pollingTrackerBase{resp: resp}}
		case http.MethodPut:
			pt = &pollingTrackerPut{pollingTrackerBase: pollingTrackerBase{resp: resp}}
		default:
			return nil, fmt.Errorf("unsupported HTTP method %s", resp.Request.Method)
		}
		if err := pt.initializeState(); err != nil {
			return pt, err
		}
		// this initializes the polling header values, we do this during creation in case the
		// initial response send us invalid values; this way the API call will return a non-nil
		// error (not doing this means the error shows up in Future.Done)
		return pt, pt.updatePollingMethod()
	}
	
	// gets the polling URL from the Azure-AsyncOperation header.
	// ensures the URL is well-formed and absolute.
	func getURLFromAsyncOpHeader(resp *azcore.Response) (string, error) {
		s := resp.Header.Get(http.CanonicalHeaderKey(headerAsyncOperation))
		if s == "" {
			return "", nil
		}
		if !isValidURL(s) {
			return "", fmt.Errorf("invalid polling URL '%s'", s)
		}
		return s, nil
	}
	
	// gets the polling URL from the Location header.
	// ensures the URL is well-formed and absolute.
	func getURLFromLocationHeader(resp *azcore.Response) (string, error) {
		s := resp.Header.Get(http.CanonicalHeaderKey(headerLocation))
		if s == "" {
			return "", nil
		}
		if !isValidURL(s) {
			return "", fmt.Errorf("invalid polling URL '%s'", s)
		}
		return s, nil
	}
	
	// verify that the URL is valid and absolute
	func isValidURL(s string) bool {
		u, err := url.Parse(s)
		return err == nil && u.IsAbs()
	}
	
	// pollingMethodType defines a type used for enumerating polling mechanisms.
	type pollingMethodType string
	
	const (
		// pollingAsyncOperation indicates the polling method uses the Azure-AsyncOperation header.
		pollingAsyncOperation pollingMethodType = "AsyncOperation"
	
		// pollingLocation indicates the polling method uses the Location header.
		pollingLocation pollingMethodType = "Location"
	
		// pollingRequestURI indicates the polling method uses the original request URI.
		pollingRequestURI pollingMethodType = "RequestURI"
	
		// pollingUnknown indicates an unknown polling method and is the default value.
		pollingUnknown pollingMethodType = ""
	)
  `;
	return text;
  }
  
