//go:build go1.18
// +build go1.18

package apiversionpath_test

import (
	"fmt"
	"net/url"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test to demonstrate the API version path parameter issue
func TestAPIVersionPathParameterIssue(t *testing.T) {
	// This test demonstrates the issue with API version path parameters
	// The problem occurs when there are multiple path parameters that could be 
	// confused, especially when API version parameters are mixed with other parameters
	
	// Scenario 1: Same operation with confusing parameter names
	// Path: /api/{api-version}/versions/{version}/resources/{resourceId}
	// This could confuse the api-version parameter with the version parameter
	regexStr1 := `/api/(?P<api_version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resources/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex1 := regexp.MustCompile(regexStr1)
	
	path1 := "/api/2023-01-01/versions/v1.5/resources/test-resource"
	matches1 := regex1.FindStringSubmatch(path1)
	require.GreaterOrEqual(t, len(matches1), 4, "Should have at least 4 matches")
	
	// This should work fine - different capture groups
	apiVersionParam, err := url.PathUnescape(matches1[regex1.SubexpIndex("api_version")])
	require.NoError(t, err)
	require.Equal(t, "2023-01-01", apiVersionParam)
	
	versionParam, err := url.PathUnescape(matches1[regex1.SubexpIndex("version")])
	require.NoError(t, err)
	require.Equal(t, "v1.5", versionParam)
	
	resourceIdParam, err := url.PathUnescape(matches1[regex1.SubexpIndex("resourceId")])
	require.NoError(t, err)
	require.Equal(t, "test-resource", resourceIdParam)
	
	fmt.Printf("Pattern SubexpNames: %v\n", regex1.SubexpNames())
	
	// Scenario 2: The real issue might be with parameter ordering and extraction
	// Let's test a more problematic scenario where parameter extraction could fail
	testProblematicScenario(t)
}

func testProblematicScenario(t *testing.T) {
	// Simulate what happens in the actual fake server generation
	// When there are multiple operations with API version parameters
	
	// Pattern similar to what's generated for operations with api-version path param
	regexStr := `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/api-version/(?P<api_version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resources/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	
	// Test with a problematic path that might cause issues
	testPath := "/subscriptions/test-sub/api-version/2023-01-01/resources/test-resource"
	matches := regex.FindStringSubmatch(testPath)
	require.GreaterOrEqual(t, len(matches), 4, "Should have at least 4 matches")
	
	// The issue might be that the parameter extraction order is not consistent
	// or that API version parameters are handled differently than expected
	
	subscriptionId, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	require.NoError(t, err)
	require.Equal(t, "test-sub", subscriptionId)
	
	apiVersion, err := url.PathUnescape(matches[regex.SubexpIndex("api_version")])
	require.NoError(t, err)
	require.Equal(t, "2023-01-01", apiVersion)
	
	resourceName, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	require.NoError(t, err)
	require.Equal(t, "test-resource", resourceName)
}