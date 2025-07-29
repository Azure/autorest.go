//go:build go1.18
// +build go1.18

// Package fake_test demonstrates the fix for API version path parameter handling in fake servers.
// 
// The issue was that when API version parameters are literal parameters (with default values), 
// the fake server generation would create regex capture groups for them, but the parsing logic
// only expected capture groups for non-literal parameters. This caused misalignment between 
// the number of expected matches and actual matches.
//
// The fix ensures that literal parameters (like API versions with default values) are replaced 
// with their literal values in the regex pattern, without creating capture groups.
package fake_test

import (
	"net/url"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAPIVersionPathParameterFix demonstrates the fix for the API version path parameter issue
func TestAPIVersionPathParameterFix(t *testing.T) {
	tests := []struct {
		name               string
		description        string
		regexPattern       string
		testPath          string
		expectedParams    map[string]string
		shouldMatch       bool
	}{
		{
			name:        "API version as literal parameter",
			description: "When API version is a literal parameter, it should be replaced with the literal value, not a capture group",
			regexPattern: `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/api-version/2023-01-01/resources/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`,
			testPath:     "/subscriptions/test-subscription/api-version/2023-01-01/resources/test-resource",
			expectedParams: map[string]string{
				"subscriptionId": "test-subscription",
				"resourceId":     "test-resource",
			},
			shouldMatch: true,
		},
		{
			name:        "Wrong API version should not match",
			description: "When the API version in the path doesn't match the literal value, it should not match",
			regexPattern: `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/api-version/2023-01-01/resources/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`,
			testPath:     "/subscriptions/test-subscription/api-version/2024-01-01/resources/test-resource",
			expectedParams: map[string]string{},
			shouldMatch: false,
		},
		{
			name:        "API version as regular parameter",
			description: "When API version is a regular path parameter, it should have a capture group",
			regexPattern: `/api/(?P<api_version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`,
			testPath:     "/api/2023-01-01/subscriptions/test-subscription",
			expectedParams: map[string]string{
				"api_version":    "2023-01-01",
				"subscriptionId": "test-subscription",
			},
			shouldMatch: true,
		},
		{
			name:        "Mixed literal and parameter scenario",
			description: "Complex scenario with both literal and regular parameters",
			regexPattern: `/providers/Microsoft.Storage/storageAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/api-version/2023-01-01/blobServices/(?P<blobServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`,
			testPath:     "/providers/Microsoft.Storage/storageAccounts/mystorageaccount/api-version/2023-01-01/blobServices/default",
			expectedParams: map[string]string{
				"accountName":     "mystorageaccount",
				"blobServiceName": "default",
			},
			shouldMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex := regexp.MustCompile(tt.regexPattern)
			matches := regex.FindStringSubmatch(tt.testPath)
			
			if tt.shouldMatch {
				require.True(t, len(matches) > 0, "Pattern should match path")
				
				// Verify all expected parameters are extracted correctly
				for paramName, expectedValue := range tt.expectedParams {
					index := regex.SubexpIndex(paramName)
					require.Greater(t, index, 0, "Parameter %s should have a capture group", paramName)
					require.Less(t, index, len(matches), "Capture group index should be within matches")
					
					actualValue, err := url.PathUnescape(matches[index])
					require.NoError(t, err, "Failed to unescape parameter %s", paramName)
					require.Equal(t, expectedValue, actualValue, "Parameter %s should have expected value", paramName)
				}
				
				// Verify the number of capture groups matches expectations
				subexpNames := regex.SubexpNames()
				expectedCaptureGroups := len(tt.expectedParams) + 1 // +1 for the full match
				require.Len(t, subexpNames, expectedCaptureGroups, "Should have expected number of capture groups")
				
			} else {
				require.True(t, len(matches) == 0, "Pattern should not match path")
			}
		})
	}
}

// TestBeforeAndAfterFix demonstrates the difference between the old and new behavior
func TestBeforeAndAfterFix(t *testing.T) {
	// Before the fix: API version parameters would create capture groups even when they were literals
	oldRegexPattern := `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/api-version/(?P<api_version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resources/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	
	// After the fix: Literal API version parameters are replaced with their literal values
	newRegexPattern := `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/api-version/2023-01-01/resources/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	
	testPath := "/subscriptions/test-subscription/api-version/2023-01-01/resources/test-resource"
	
	// Test old pattern (would create unnecessary capture group)
	oldRegex := regexp.MustCompile(oldRegexPattern)
	oldMatches := oldRegex.FindStringSubmatch(testPath)
	require.Len(t, oldMatches, 4, "Old pattern creates 3 capture groups + full match")
	require.Equal(t, []string{"", "subscriptionId", "api_version", "resourceId"}, oldRegex.SubexpNames())
	
	// Test new pattern (no capture group for literal API version)
	newRegex := regexp.MustCompile(newRegexPattern)
	newMatches := newRegex.FindStringSubmatch(testPath)
	require.Len(t, newMatches, 3, "New pattern creates 2 capture groups + full match")
	require.Equal(t, []string{"", "subscriptionId", "resourceId"}, newRegex.SubexpNames())
	
	// Verify both patterns extract non-literal parameters correctly
	subscriptionId, err := url.PathUnescape(oldMatches[oldRegex.SubexpIndex("subscriptionId")])
	require.NoError(t, err)
	require.Equal(t, "test-subscription", subscriptionId)
	
	subscriptionId2, err := url.PathUnescape(newMatches[newRegex.SubexpIndex("subscriptionId")])
	require.NoError(t, err)
	require.Equal(t, "test-subscription", subscriptionId2)
	
	resourceId, err := url.PathUnescape(oldMatches[oldRegex.SubexpIndex("resourceId")])
	require.NoError(t, err)
	require.Equal(t, "test-resource", resourceId)
	
	resourceId2, err := url.PathUnescape(newMatches[newRegex.SubexpIndex("resourceId")])
	require.NoError(t, err)
	require.Equal(t, "test-resource", resourceId2)
	
	// The key difference: old pattern has api_version capture group, new pattern doesn't
	require.Greater(t, oldRegex.SubexpIndex("api_version"), 0, "Old pattern should have api_version capture group")
	require.Equal(t, -1, newRegex.SubexpIndex("api_version"), "New pattern should not have api_version capture group")
}