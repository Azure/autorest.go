// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azblob

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestContainerPropertiesMarshalRoundTrip(t *testing.T) {
	etag := azcore.ETag("0x8D9B2F4B4E2A7C0")
	lastModified := time.Date(2025, 3, 15, 10, 30, 0, 0, time.UTC)
	leaseDuration := LeaseDurationFixed
	leaseState := LeaseStateAvailable
	leaseStatus := LeaseStatusUnlocked

	original := ContainerProperties{
		ETag:                           &etag,
		LastModified:                   &lastModified,
		DefaultEncryptionScope:         to.Ptr("testscope"),
		HasImmutabilityPolicy:          to.Ptr(true),
		HasLegalHold:                   to.Ptr(false),
		LeaseDuration:                  &leaseDuration,
		LeaseState:                     &leaseState,
		LeaseStatus:                    &leaseStatus,
		PreventEncryptionScopeOverride: to.Ptr(false),
		RemainingRetentionDays:         to.Ptr[int32](7),
	}

	data, err := xml.Marshal(original)
	require.NoError(t, err)

	var roundTripped ContainerProperties
	require.NoError(t, xml.Unmarshal(data, &roundTripped))

	// Verify ETag round-trips correctly
	require.NotNil(t, roundTripped.ETag)
	require.Equal(t, etag, *roundTripped.ETag)

	// Verify LastModified round-trips correctly (compare to second precision since RFC1123 drops sub-seconds)
	require.NotNil(t, roundTripped.LastModified)
	require.True(t, roundTripped.LastModified.Equal(lastModified))

	// Verify other fields
	require.NotNil(t, roundTripped.DefaultEncryptionScope)
	require.Equal(t, "testscope", *roundTripped.DefaultEncryptionScope)
	require.NotNil(t, roundTripped.HasImmutabilityPolicy)
	require.Equal(t, true, *roundTripped.HasImmutabilityPolicy)
	require.NotNil(t, roundTripped.LeaseDuration)
	require.Equal(t, LeaseDurationFixed, *roundTripped.LeaseDuration)
	require.NotNil(t, roundTripped.LeaseState)
	require.Equal(t, LeaseStateAvailable, *roundTripped.LeaseState)
	require.NotNil(t, roundTripped.LeaseStatus)
	require.Equal(t, LeaseStatusUnlocked, *roundTripped.LeaseStatus)
	require.NotNil(t, roundTripped.RemainingRetentionDays)
	require.Equal(t, int32(7), *roundTripped.RemainingRetentionDays)
}

func TestContainerPropertiesUnmarshalFromXML(t *testing.T) {
	// Simulate XML as it would be returned by the Azure Blob Storage service.
	// The ETag value is wrapped in quotes as the service returns it.
	xmlData := []byte(`<ContainerProperties>
	<Etag>"0x8D9B2F4B4E2A7C0"</Etag>
	<Last-Modified>Sat, 15 Mar 2025 10:30:00 GMT</Last-Modified>
	<DefaultEncryptionScope>$account-encryption-key</DefaultEncryptionScope>
	<DenyEncryptionScopeOverride>false</DenyEncryptionScopeOverride>
	<HasImmutabilityPolicy>false</HasImmutabilityPolicy>
	<HasLegalHold>false</HasLegalHold>
	<LeaseState>available</LeaseState>
	<LeaseStatus>unlocked</LeaseStatus>
</ContainerProperties>`)

	var props ContainerProperties
	require.NoError(t, xml.Unmarshal(xmlData, &props))

	// The ETag must preserve the surrounding quotes from the service response
	expectedETag := azcore.ETag(`"0x8D9B2F4B4E2A7C0"`)
	require.NotNil(t, props.ETag)
	require.Equal(t, expectedETag, *props.ETag)

	expectedTime := time.Date(2025, 3, 15, 10, 30, 0, 0, time.UTC)
	require.NotNil(t, props.LastModified)
	require.True(t, props.LastModified.Equal(expectedTime))

	require.NotNil(t, props.DefaultEncryptionScope)
	require.Equal(t, "$account-encryption-key", *props.DefaultEncryptionScope)
	require.NotNil(t, props.PreventEncryptionScopeOverride)
	require.Equal(t, false, *props.PreventEncryptionScopeOverride)
	require.NotNil(t, props.LeaseState)
	require.Equal(t, LeaseStateAvailable, *props.LeaseState)
	require.NotNil(t, props.LeaseStatus)
	require.Equal(t, LeaseStatusUnlocked, *props.LeaseStatus)
}

func TestContainerPropertiesETagWithoutQuotes(t *testing.T) {
	// Test ETag value without surrounding quotes
	xmlData := []byte(`<ContainerProperties>
	<Etag>0x8D9B2F4B4E2A7C0</Etag>
	<Last-Modified>Sat, 15 Mar 2025 10:30:00 GMT</Last-Modified>
</ContainerProperties>`)

	var props ContainerProperties
	require.NoError(t, xml.Unmarshal(xmlData, &props))

	expectedETag := azcore.ETag("0x8D9B2F4B4E2A7C0")
	require.NotNil(t, props.ETag)
	require.Equal(t, expectedETag, *props.ETag)
}
