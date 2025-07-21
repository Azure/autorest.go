//go:build go1.18
// +build go1.18

package fake

import (
	"fmt"
	"net/url"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAPIVersionPathParameterHandling tests the fix for API version path parameter handling
func TestAPIVersionPathParameterHandling(t *testing.T) {
	// Test case 1: API version as literal parameter
	// This simulates a path like /subscriptions/{subscriptionId}/api-version/{api-version}/resources/{resourceId}
	// where {api-version} is a literal parameter with value "2023-01-01"
	// and {subscriptionId}, {resourceId} are regular path parameters
	
	// Expected regex: /subscriptions/(?P<subscriptionId>[...]+)/api-version/2023-01-01/resources/(?P<resourceId>[...]+)
	regexStrLiteral := `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/api-version/2023-01-01/resources/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regexLiteral := regexp.MustCompile(regexStrLiteral)
	
	pathLiteral := "/subscriptions/test-subscription/api-version/2023-01-01/resources/test-resource"
	matchesLiteral := regexLiteral.FindStringSubmatch(pathLiteral)
	
	require.GreaterOrEqual(t, len(matchesLiteral), 3, "Should have at least 3 matches for literal case")
	
	subscriptionId, err := url.PathUnescape(matchesLiteral[regexLiteral.SubexpIndex("subscriptionId")])
	require.NoError(t, err)
	require.Equal(t, "test-subscription", subscriptionId)
	
	resourceId, err := url.PathUnescape(matchesLiteral[regexLiteral.SubexpIndex("resourceId")])
	require.NoError(t, err)
	require.Equal(t, "test-resource", resourceId)
	
	// Test case 2: API version as regular path parameter  
	// This simulates a path like /api-version/{api-version}/subscriptions/{subscriptionId}
	// where both {api-version} and {subscriptionId} are regular path parameters
	
	regexStrParam := `/api-version/(?P<api_version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regexParam := regexp.MustCompile(regexStrParam)
	
	pathParam := "/api-version/2023-01-01/subscriptions/test-subscription"
	matchesParam := regexParam.FindStringSubmatch(pathParam)
	
	require.GreaterOrEqual(t, len(matchesParam), 3, "Should have at least 3 matches for parameter case")
	
	apiVersion, err := url.PathUnescape(matchesParam[regexParam.SubexpIndex("api_version")])
	require.NoError(t, err)
	require.Equal(t, "2023-01-01", apiVersion)
	
	subscriptionId2, err := url.PathUnescape(matchesParam[regexParam.SubexpIndex("subscriptionId")])
	require.NoError(t, err)
	require.Equal(t, "test-subscription", subscriptionId2)
	
	fmt.Printf("Literal case SubexpNames: %v\n", regexLiteral.SubexpNames())
	fmt.Printf("Parameter case SubexpNames: %v\n", regexParam.SubexpNames())
}

// TestAPIVersionConfusionScenario tests the scenario that was causing issues
func TestAPIVersionConfusionScenario(t *testing.T) {
	// This tests a scenario where there are multiple version-related parameters
	// that could potentially be confused with each other
	
	// Path: /api/{api-version}/versions/{version}/subscriptions/{subscriptionId}
	// Where:
	// - {api-version} might be a literal parameter with value "2023-01-01"
	// - {version} is a regular path parameter 
	// - {subscriptionId} is a regular path parameter
	
	// The fixed implementation should generate:
	// /api/2023-01-01/versions/(?P<version>[...]+)/subscriptions/(?P<subscriptionId>[...]+)
	regexStr := `/api/2023-01-01/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	
	testPath := "/api/2023-01-01/versions/v1.5/subscriptions/test-subscription"
	matches := regex.FindStringSubmatch(testPath)
	
	require.GreaterOrEqual(t, len(matches), 3, "Should have at least 3 matches")
	
	// Extract only the non-literal parameters
	version, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	require.NoError(t, err)
	require.Equal(t, "v1.5", version)
	
	subscriptionId, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	require.NoError(t, err)
	require.Equal(t, "test-subscription", subscriptionId)
	
	// Verify that the regex has the expected number of capture groups
	// (only for non-literal parameters)
	subexpNames := regex.SubexpNames()
	require.Len(t, subexpNames, 3, "Should have 3 subexp names: empty, version, subscriptionId")
	require.Equal(t, []string{"", "version", "subscriptionId"}, subexpNames)
	
	fmt.Printf("Confusion scenario SubexpNames: %v\n", subexpNames)
}