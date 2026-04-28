package azblob

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func timeRFC1123(s string) *time.Time {
	t, err := time.Parse(time.RFC1123, s)
	if err != nil {
		panic(err)
	}
	return &t
}

func timeRFC3339(s string) *time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return &t
}

// TestStorageServicePropertiesRoundTrip tests marshal/unmarshal of StorageServiceProperties
// using the example response body from Get Blob Service Properties.
// https://learn.microsoft.com/en-us/rest/api/storageservices/get-blob-service-properties
func TestStorageServicePropertiesRoundTrip(t *testing.T) {
	// From the sample response in the docs
	xmlData := `<StorageServiceProperties>
    <Logging>
        <Version>1.0</Version>
        <Delete>true</Delete>
        <Read>false</Read>
        <Write>true</Write>
        <RetentionPolicy>
            <Enabled>true</Enabled>
            <Days>7</Days>
        </RetentionPolicy>
    </Logging>
    <HourMetrics>
        <Version>1.0</Version>
        <Enabled>true</Enabled>
        <IncludeAPIs>false</IncludeAPIs>
        <RetentionPolicy>
            <Enabled>true</Enabled>
            <Days>7</Days>
        </RetentionPolicy>
    </HourMetrics>
    <MinuteMetrics>
        <Version>1.0</Version>
        <Enabled>true</Enabled>
        <IncludeAPIs>true</IncludeAPIs>
        <RetentionPolicy>
            <Enabled>true</Enabled>
            <Days>7</Days>
        </RetentionPolicy>
    </MinuteMetrics>
    <Cors>
        <CorsRule>
            <AllowedOrigins>http://www.fabrikam.com,http://www.contoso.com</AllowedOrigins>
            <AllowedMethods>GET,PUT</AllowedMethods>
            <MaxAgeInSeconds>500</MaxAgeInSeconds>
            <ExposedHeaders>x-ms-meta-data*,x-ms-meta-customheader</ExposedHeaders>
            <AllowedHeaders>x-ms-meta-target*,x-ms-meta-customheader</AllowedHeaders>
        </CorsRule>
    </Cors>
    <DefaultServiceVersion>2017-07-29</DefaultServiceVersion>
    <DeleteRetentionPolicy>
        <Enabled>true</Enabled>
        <Days>5</Days>
    </DeleteRetentionPolicy>
    <StaticWebsite>
        <Enabled>true</Enabled>
        <IndexDocument>index.html</IndexDocument>
        <ErrorDocument404Path>error/404.html</ErrorDocument404Path>
    </StaticWebsite>
</StorageServiceProperties>`

	var props StorageServiceProperties
	if err := xml.Unmarshal([]byte(xmlData), &props); err != nil {
		t.Fatal(err)
	}

	// Verify Logging
	if props.Logging == nil {
		t.Fatal("Logging is nil")
	}
	assertEqual(t, "Logging.Version", *props.Logging.Version, "1.0")
	assertEqual(t, "Logging.Delete", *props.Logging.Delete, true)
	assertEqual(t, "Logging.Read", *props.Logging.Read, false)
	assertEqual(t, "Logging.Write", *props.Logging.Write, true)
	assertEqual(t, "Logging.RetentionPolicy.Enabled", *props.Logging.RetentionPolicy.Enabled, true)
	assertEqual(t, "Logging.RetentionPolicy.Days", *props.Logging.RetentionPolicy.Days, int32(7))

	// Verify HourMetrics
	if props.HourMetrics == nil {
		t.Fatal("HourMetrics is nil")
	}
	assertEqual(t, "HourMetrics.IncludeAPIs", *props.HourMetrics.IncludeAPIs, false)

	// Verify MinuteMetrics
	if props.MinuteMetrics == nil {
		t.Fatal("MinuteMetrics is nil")
	}
	assertEqual(t, "MinuteMetrics.IncludeAPIs", *props.MinuteMetrics.IncludeAPIs, true)

	// Verify CORS
	if len(props.CORS) != 1 {
		t.Fatalf("expected 1 CorsRule, got %d", len(props.CORS))
	}
	assertEqual(t, "CorsRule.AllowedOrigins", *props.CORS[0].AllowedOrigins, "http://www.fabrikam.com,http://www.contoso.com")
	assertEqual(t, "CorsRule.AllowedMethods", *props.CORS[0].AllowedMethods, "GET,PUT")
	assertEqual(t, "CorsRule.MaxAgeInSeconds", *props.CORS[0].MaxAgeInSeconds, int32(500))

	// Verify DefaultServiceVersion
	assertEqual(t, "DefaultServiceVersion", *props.DefaultServiceVersion, "2017-07-29")

	// Verify DeleteRetentionPolicy
	assertEqual(t, "DeleteRetentionPolicy.Enabled", *props.DeleteRetentionPolicy.Enabled, true)
	assertEqual(t, "DeleteRetentionPolicy.Days", *props.DeleteRetentionPolicy.Days, int32(5))

	// Verify StaticWebsite
	assertEqual(t, "StaticWebsite.Enabled", *props.StaticWebsite.Enabled, true)
	assertEqual(t, "StaticWebsite.IndexDocument", *props.StaticWebsite.IndexDocument, "index.html")
	assertEqual(t, "StaticWebsite.ErrorDocument404Path", *props.StaticWebsite.ErrorDocument404Path, "error/404.html")

	// Re-marshal and unmarshal to verify round-trip
	out, err := xml.Marshal(props)
	if err != nil {
		t.Fatal(err)
	}
	var props2 StorageServiceProperties
	if err := xml.Unmarshal(out, &props2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip Logging.Version", *props2.Logging.Version, "1.0")
	assertEqual(t, "round-trip CorsRule.AllowedOrigins", *props2.CORS[0].AllowedOrigins, "http://www.fabrikam.com,http://www.contoso.com")
	assertEqual(t, "round-trip StaticWebsite.Enabled", *props2.StaticWebsite.Enabled, true)
}

// TestListContainersSegmentResponseRoundTrip tests marshal/unmarshal of ListContainersSegmentResponse
// using the example response body from List Containers.
// https://learn.microsoft.com/en-us/rest/api/storageservices/list-containers2
func TestListContainersSegmentResponseRoundTrip(t *testing.T) {
	xmlData := `<EnumerationResults ServiceEndpoint="https://myaccount.blob.core.windows.net/">
  <MaxResults>3</MaxResults>
  <Containers>
    <Container>
      <Name>audio</Name>
      <Properties>
        <Last-Modified>Wed, 26 Oct 2016 20:39:39 GMT</Last-Modified>
        <Etag>0x8CACB9BD7C6B1B2</Etag>
        <PublicAccess>container</PublicAccess>
      </Properties>
    </Container>
    <Container>
      <Name>images</Name>
      <Properties>
        <Last-Modified>Wed, 26 Oct 2016 20:39:39 GMT</Last-Modified>
        <Etag>0x8CACB9BD7C1EEEC</Etag>
      </Properties>
    </Container>
    <Container>
      <Name>textfiles</Name>
      <Properties>
        <Last-Modified>Wed, 26 Oct 2016 20:39:39 GMT</Last-Modified>
        <Etag>0x8CACB9BD7BACAC3</Etag>
      </Properties>
    </Container>
  </Containers>
  <NextMarker>video</NextMarker>
</EnumerationResults>`

	var resp ListContainersSegmentResponse
	if err := xml.Unmarshal([]byte(xmlData), &resp); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "ServiceEndpoint", *resp.ServiceEndpoint, "https://myaccount.blob.core.windows.net/")
	assertEqual(t, "MaxResults", *resp.MaxResults, int32(3))
	assertEqual(t, "NextMarker", *resp.NextMarker, "video")

	if len(resp.ContainerItems) != 3 {
		t.Fatalf("expected 3 containers, got %d", len(resp.ContainerItems))
	}

	assertEqual(t, "Container[0].Name", *resp.ContainerItems[0].Name, "audio")
	assertEqual(t, "Container[0].Etag", string(*resp.ContainerItems[0].Properties.ETag), "0x8CACB9BD7C6B1B2")
	assertEqual(t, "Container[0].PublicAccess", string(*resp.ContainerItems[0].Properties.PublicAccess), "container")

	assertEqual(t, "Container[1].Name", *resp.ContainerItems[1].Name, "images")
	assertEqual(t, "Container[2].Name", *resp.ContainerItems[2].Name, "textfiles")

	// Verify time parsing
	expectedTime := timeRFC1123("Wed, 26 Oct 2016 20:39:39 GMT")
	if !resp.ContainerItems[0].Properties.LastModified.Equal(*expectedTime) {
		t.Fatalf("expected Last-Modified %v, got %v", expectedTime, resp.ContainerItems[0].Properties.LastModified)
	}

	// Round-trip
	out, err := xml.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var resp2 ListContainersSegmentResponse
	if err := xml.Unmarshal(out, &resp2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip ServiceEndpoint", *resp2.ServiceEndpoint, "https://myaccount.blob.core.windows.net/")
	if len(resp2.ContainerItems) != 3 {
		t.Fatalf("round-trip expected 3 containers, got %d", len(resp2.ContainerItems))
	}
	assertEqual(t, "round-trip Container[0].Name", *resp2.ContainerItems[0].Name, "audio")
}

// TestContainerItemWithMetadataRoundTrip tests marshal/unmarshal of ContainerItem with Metadata.
// https://learn.microsoft.com/en-us/rest/api/storageservices/list-containers2
func TestContainerItemWithMetadataRoundTrip(t *testing.T) {
	xmlData := `<Container>
  <Name>mycontainer</Name>
  <Properties>
    <Last-Modified>Wed, 26 Oct 2016 20:39:39 GMT</Last-Modified>
    <Etag>0x8CACB9BD7C6B1B2</Etag>
    <LeaseStatus>unlocked</LeaseStatus>
    <LeaseState>available</LeaseState>
    <HasImmutabilityPolicy>false</HasImmutabilityPolicy>
    <HasLegalHold>false</HasLegalHold>
  </Properties>
  <Metadata>
    <MyMetadata1>first value</MyMetadata1>
    <MyMetadata2>second value</MyMetadata2>
  </Metadata>
</Container>`

	var item ContainerItem
	if err := xml.Unmarshal([]byte(xmlData), &item); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Name", *item.Name, "mycontainer")
	assertEqual(t, "Etag", string(*item.Properties.ETag), "0x8CACB9BD7C6B1B2")
	assertEqual(t, "LeaseStatus", string(*item.Properties.LeaseStatus), "unlocked")
	assertEqual(t, "LeaseState", string(*item.Properties.LeaseState), "available")
	assertEqual(t, "HasImmutabilityPolicy", *item.Properties.HasImmutabilityPolicy, false)
	assertEqual(t, "HasLegalHold", *item.Properties.HasLegalHold, false)

	if len(item.Metadata) != 2 {
		t.Fatalf("expected 2 metadata entries, got %d", len(item.Metadata))
	}

	// Round-trip
	out, err := xml.Marshal(item)
	if err != nil {
		t.Fatal(err)
	}
	var item2 ContainerItem
	if err := xml.Unmarshal(out, &item2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip Name", *item2.Name, "mycontainer")
	if len(item2.Metadata) != 2 {
		t.Fatalf("round-trip expected 2 metadata entries, got %d", len(item2.Metadata))
	}
}

// TestBlockListRoundTrip tests marshal/unmarshal of BlockList
// using the example response body from Get Block List.
// https://learn.microsoft.com/en-us/rest/api/storageservices/get-block-list
func TestBlockListRoundTrip(t *testing.T) {
	xmlData := `<BlockList>
  <CommittedBlocks>
    <Block>
      <Name>QmxvY2tJZDAwMQ==</Name>
      <Size>4194304</Size>
    </Block>
    <Block>
      <Name>QmxvY2tJZDAwMg==</Name>
      <Size>4194304</Size>
    </Block>
  </CommittedBlocks>
  <UncommittedBlocks>
    <Block>
      <Name>QmxvY2tJZDAwMw==</Name>
      <Size>4194304</Size>
    </Block>
    <Block>
      <Name>QmxvY2tJZDAwNA==</Name>
      <Size>1024000</Size>
    </Block>
  </UncommittedBlocks>
</BlockList>`

	var blockList BlockList
	if err := xml.Unmarshal([]byte(xmlData), &blockList); err != nil {
		t.Fatal(err)
	}

	if len(blockList.CommittedBlocks) != 2 {
		t.Fatalf("expected 2 committed blocks, got %d", len(blockList.CommittedBlocks))
	}
	if len(blockList.UncommittedBlocks) != 2 {
		t.Fatalf("expected 2 uncommitted blocks, got %d", len(blockList.UncommittedBlocks))
	}
	assertEqual(t, "CommittedBlocks[0].Name", *blockList.CommittedBlocks[0].Name, "QmxvY2tJZDAwMQ==")
	assertEqual(t, "CommittedBlocks[0].Size", *blockList.CommittedBlocks[0].Size, int64(4194304))
	assertEqual(t, "CommittedBlocks[1].Name", *blockList.CommittedBlocks[1].Name, "QmxvY2tJZDAwMg==")
	assertEqual(t, "UncommittedBlocks[0].Name", *blockList.UncommittedBlocks[0].Name, "QmxvY2tJZDAwMw==")
	assertEqual(t, "UncommittedBlocks[1].Name", *blockList.UncommittedBlocks[1].Name, "QmxvY2tJZDAwNA==")
	assertEqual(t, "UncommittedBlocks[1].Size", *blockList.UncommittedBlocks[1].Size, int64(1024000))

	// Round-trip
	out, err := xml.Marshal(blockList)
	if err != nil {
		t.Fatal(err)
	}
	var blockList2 BlockList
	if err := xml.Unmarshal(out, &blockList2); err != nil {
		t.Fatal(err)
	}
	if len(blockList2.CommittedBlocks) != 2 {
		t.Fatalf("round-trip expected 2 committed blocks, got %d", len(blockList2.CommittedBlocks))
	}
	if len(blockList2.UncommittedBlocks) != 2 {
		t.Fatalf("round-trip expected 2 uncommitted blocks, got %d", len(blockList2.UncommittedBlocks))
	}
}

// TestBlockLookupListRoundTrip tests marshal/unmarshal of BlockLookupList
// using the example request body from Put Block List.
// https://learn.microsoft.com/en-us/rest/api/storageservices/put-block-list
func TestBlockLookupListRoundTrip(t *testing.T) {
	xmlData := `<BlockList>
  <Latest>AAAAAA==</Latest>
  <Latest>AQAAAA==</Latest>
  <Latest>AZAAAA==</Latest>
</BlockList>`

	var bll BlockLookupList
	if err := xml.Unmarshal([]byte(xmlData), &bll); err != nil {
		t.Fatal(err)
	}

	if len(bll.Latest) != 3 {
		t.Fatalf("expected 3 latest blocks, got %d", len(bll.Latest))
	}
	assertEqual(t, "Latest[0]", *bll.Latest[0], "AAAAAA==")
	assertEqual(t, "Latest[1]", *bll.Latest[1], "AQAAAA==")
	assertEqual(t, "Latest[2]", *bll.Latest[2], "AZAAAA==")

	// Round-trip
	out, err := xml.Marshal(bll)
	if err != nil {
		t.Fatal(err)
	}
	var bll2 BlockLookupList
	if err := xml.Unmarshal(out, &bll2); err != nil {
		t.Fatal(err)
	}
	if len(bll2.Latest) != 3 {
		t.Fatalf("round-trip expected 3 latest blocks, got %d", len(bll2.Latest))
	}
}

// TestBlockLookupListMixedRoundTrip tests mixed committed/uncommitted block lists.
// https://learn.microsoft.com/en-us/rest/api/storageservices/put-block-list
func TestBlockLookupListMixedRoundTrip(t *testing.T) {
	xmlData := `<BlockList>
  <Uncommitted>ANAAAA==</Uncommitted>
  <Committed>AQAAAA==</Committed>
  <Uncommitted>AZAAAA==</Uncommitted>
</BlockList>`

	var bll BlockLookupList
	if err := xml.Unmarshal([]byte(xmlData), &bll); err != nil {
		t.Fatal(err)
	}

	if len(bll.Uncommitted) != 2 {
		t.Fatalf("expected 2 uncommitted blocks, got %d", len(bll.Uncommitted))
	}
	if len(bll.Committed) != 1 {
		t.Fatalf("expected 1 committed block, got %d", len(bll.Committed))
	}
	assertEqual(t, "Uncommitted[0]", *bll.Uncommitted[0], "ANAAAA==")
	assertEqual(t, "Committed[0]", *bll.Committed[0], "AQAAAA==")

	// Round-trip
	out, err := xml.Marshal(bll)
	if err != nil {
		t.Fatal(err)
	}
	var bll2 BlockLookupList
	if err := xml.Unmarshal(out, &bll2); err != nil {
		t.Fatal(err)
	}
	if len(bll2.Uncommitted) != 2 {
		t.Fatalf("round-trip expected 2 uncommitted blocks, got %d", len(bll2.Uncommitted))
	}
	if len(bll2.Committed) != 1 {
		t.Fatalf("round-trip expected 1 committed block, got %d", len(bll2.Committed))
	}
}

// TestPageListRoundTrip tests marshal/unmarshal of PageList
// using the example response body from Get Page Ranges.
// https://learn.microsoft.com/en-us/rest/api/storageservices/get-page-ranges
func TestPageListRoundTrip(t *testing.T) {
	xmlData := `<PageList>
   <PageRange>
      <Start>0</Start>
      <End>511</End>
   </PageRange>
   <PageRange>
      <Start>1024</Start>
      <End>1535</End>
   </PageRange>
</PageList>`

	var pageList PageList
	if err := xml.Unmarshal([]byte(xmlData), &pageList); err != nil {
		t.Fatal(err)
	}

	if len(pageList.PageRange) != 2 {
		t.Fatalf("expected 2 page ranges, got %d", len(pageList.PageRange))
	}
	assertEqual(t, "PageRange[0].Start", *pageList.PageRange[0].Start, int64(0))
	assertEqual(t, "PageRange[0].End", *pageList.PageRange[0].End, int64(511))
	assertEqual(t, "PageRange[1].Start", *pageList.PageRange[1].Start, int64(1024))
	assertEqual(t, "PageRange[1].End", *pageList.PageRange[1].End, int64(1535))

	// Round-trip
	out, err := xml.Marshal(pageList)
	if err != nil {
		t.Fatal(err)
	}
	var pageList2 PageList
	if err := xml.Unmarshal(out, &pageList2); err != nil {
		t.Fatal(err)
	}
	if len(pageList2.PageRange) != 2 {
		t.Fatalf("round-trip expected 2 page ranges, got %d", len(pageList2.PageRange))
	}
}

// TestPageListWithClearRangesRoundTrip tests diff page ranges with clear ranges.
// https://learn.microsoft.com/en-us/rest/api/storageservices/get-page-ranges
func TestPageListWithClearRangesRoundTrip(t *testing.T) {
	xmlData := `<PageList>
   <PageRange>
      <Start>0</Start>
      <End>511</End>
   </PageRange>
   <ClearRange>
      <Start>512</Start>
      <End>1023</End>
   </ClearRange>
   <PageRange>
      <Start>1024</Start>
      <End>1535</End>
   </PageRange>
</PageList>`

	var pageList PageList
	if err := xml.Unmarshal([]byte(xmlData), &pageList); err != nil {
		t.Fatal(err)
	}

	if len(pageList.PageRange) != 2 {
		t.Fatalf("expected 2 page ranges, got %d", len(pageList.PageRange))
	}
	if len(pageList.ClearRange) != 1 {
		t.Fatalf("expected 1 clear range, got %d", len(pageList.ClearRange))
	}
	assertEqual(t, "ClearRange[0].Start", *pageList.ClearRange[0].Start, int64(512))
	assertEqual(t, "ClearRange[0].End", *pageList.ClearRange[0].End, int64(1023))

	// Round-trip
	out, err := xml.Marshal(pageList)
	if err != nil {
		t.Fatal(err)
	}
	var pageList2 PageList
	if err := xml.Unmarshal(out, &pageList2); err != nil {
		t.Fatal(err)
	}
	if len(pageList2.PageRange) != 2 {
		t.Fatalf("round-trip expected 2 page ranges, got %d", len(pageList2.PageRange))
	}
	if len(pageList2.ClearRange) != 1 {
		t.Fatalf("round-trip expected 1 clear range, got %d", len(pageList2.ClearRange))
	}
}

// TestFilterBlobSegmentRoundTrip tests marshal/unmarshal of FilterBlobSegment
// using the example response body from Find Blobs by Tags.
// https://learn.microsoft.com/en-us/rest/api/storageservices/find-blobs-by-tags
func TestFilterBlobSegmentRoundTrip(t *testing.T) {
	xmlData := `<EnumerationResults ServiceEndpoint="http://myaccount.blob.core.windows.net/">
  <Where>Status = 'In Progress'</Where>
  <Blobs>
    <Blob>
      <Name>my-blob</Name>
      <ContainerName>my-container</ContainerName>
      <Tags>
        <TagSet>
          <Tag>
            <Key>Status</Key>
            <Value>In Progress</Value>
          </Tag>
        </TagSet>
      </Tags>
    </Blob>
  </Blobs>
  <NextMarker />
</EnumerationResults>`

	var resp FilterBlobSegment
	if err := xml.Unmarshal([]byte(xmlData), &resp); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "ServiceEndpoint", *resp.ServiceEndpoint, "http://myaccount.blob.core.windows.net/")
	assertEqual(t, "Where", *resp.Where, "Status = 'In Progress'")
	if len(resp.Blobs) != 1 {
		t.Fatalf("expected 1 blob, got %d", len(resp.Blobs))
	}
	assertEqual(t, "Blobs[0].Name", *resp.Blobs[0].Name, "my-blob")
	assertEqual(t, "Blobs[0].ContainerName", *resp.Blobs[0].ContainerName, "my-container")
	if len(resp.Blobs[0].Tags.BlobTagSet) != 1 {
		t.Fatalf("expected 1 tag, got %d", len(resp.Blobs[0].Tags.BlobTagSet))
	}
	assertEqual(t, "Tag.Key", *resp.Blobs[0].Tags.BlobTagSet[0].Key, "Status")
	assertEqual(t, "Tag.Value", *resp.Blobs[0].Tags.BlobTagSet[0].Value, "In Progress")

	// Round-trip
	out, err := xml.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var resp2 FilterBlobSegment
	if err := xml.Unmarshal(out, &resp2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip Where", *resp2.Where, "Status = 'In Progress'")
	if len(resp2.Blobs) != 1 {
		t.Fatalf("round-trip expected 1 blob, got %d", len(resp2.Blobs))
	}
}

// TestListBlobsFlatSegmentResponseRoundTrip tests marshal/unmarshal of ListBlobsFlatSegmentResponse
// using the example response body from List Blobs.
// https://learn.microsoft.com/en-us/rest/api/storageservices/list-blobs
func TestListBlobsFlatSegmentResponseRoundTrip(t *testing.T) {
	xmlData := `<EnumerationResults ServiceEndpoint="http://myaccount.blob.core.windows.net/" ContainerName="mycontainer">
  <Prefix>my</Prefix>
  <MaxResults>10</MaxResults>
  <Blobs>
    <Blob>
      <Name>my-blob.txt</Name>
      <Snapshot></Snapshot>
      <Deleted>false</Deleted>
      <Properties>
        <Creation-Time>Wed, 26 Oct 2016 20:39:39 GMT</Creation-Time>
        <Last-Modified>Wed, 26 Oct 2016 20:39:39 GMT</Last-Modified>
        <Etag>0x8CACB9BD7C6B1B2</Etag>
        <Content-Length>1024</Content-Length>
        <Content-Type>application/octet-stream</Content-Type>
        <BlobType>BlockBlob</BlobType>
        <AccessTier>Hot</AccessTier>
        <LeaseStatus>unlocked</LeaseStatus>
        <LeaseState>available</LeaseState>
        <ServerEncrypted>true</ServerEncrypted>
      </Properties>
      <Tags>
        <TagSet>
          <Tag>
            <Key>TagName</Key>
            <Value>TagValue</Value>
          </Tag>
        </TagSet>
      </Tags>
    </Blob>
  </Blobs>
  <NextMarker />
</EnumerationResults>`

	var resp ListBlobsFlatSegmentResponse
	if err := xml.Unmarshal([]byte(xmlData), &resp); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "ServiceEndpoint", *resp.ServiceEndpoint, "http://myaccount.blob.core.windows.net/")
	assertEqual(t, "ContainerName", *resp.ContainerName, "mycontainer")
	assertEqual(t, "Prefix", *resp.Prefix, "my")
	assertEqual(t, "MaxResults", *resp.MaxResults, int32(10))

	if resp.Segment == nil || len(resp.Segment.BlobItems) != 1 {
		t.Fatal("expected 1 blob item")
	}
	blob := resp.Segment.BlobItems[0]
	assertEqual(t, "Blob.Name", *blob.Name, "my-blob.txt")
	assertEqual(t, "Blob.Deleted", *blob.Deleted, false)
	assertEqual(t, "Blob.Properties.ETag", string(*blob.Properties.ETag), "0x8CACB9BD7C6B1B2")
	assertEqual(t, "Blob.Properties.ContentLength", *blob.Properties.ContentLength, int64(1024))
	assertEqual(t, "Blob.Properties.ContentType", *blob.Properties.ContentType, "application/octet-stream")
	assertEqual(t, "Blob.Properties.BlobType", string(*blob.Properties.BlobType), "BlockBlob")
	assertEqual(t, "Blob.Properties.AccessTier", string(*blob.Properties.AccessTier), "Hot")
	assertEqual(t, "Blob.Properties.ServerEncrypted", *blob.Properties.ServerEncrypted, true)

	// Tags
	if blob.BlobTags == nil || len(blob.BlobTags.BlobTagSet) != 1 {
		t.Fatal("expected 1 blob tag")
	}
	assertEqual(t, "Tag.Key", *blob.BlobTags.BlobTagSet[0].Key, "TagName")
	assertEqual(t, "Tag.Value", *blob.BlobTags.BlobTagSet[0].Value, "TagValue")

}

// TestListBlobsHierarchySegmentResponseRoundTrip tests marshal/unmarshal of ListBlobsHierarchySegmentResponse.
// https://learn.microsoft.com/en-us/rest/api/storageservices/list-blobs
func TestListBlobsHierarchySegmentResponseRoundTrip(t *testing.T) {
	xmlData := `<EnumerationResults ServiceEndpoint="http://myaccount.blob.core.windows.net/" ContainerName="mycontainer">
  <Delimiter>/</Delimiter>
  <Blobs>
    <Blob>
      <Name>file1.txt</Name>
      <Snapshot></Snapshot>
      <Deleted>false</Deleted>
      <Properties>
        <Last-Modified>Wed, 26 Oct 2016 20:39:39 GMT</Last-Modified>
        <Etag>0x8CACB9BD7C6B1B2</Etag>
        <Content-Length>512</Content-Length>
        <Content-Type>text/plain</Content-Type>
        <BlobType>BlockBlob</BlobType>
        <LeaseStatus>unlocked</LeaseStatus>
        <LeaseState>available</LeaseState>
      </Properties>
    </Blob>
    <BlobPrefix>
      <Name>subdir/</Name>
    </BlobPrefix>
  </Blobs>
  <NextMarker />
</EnumerationResults>`

	var resp ListBlobsHierarchySegmentResponse
	if err := xml.Unmarshal([]byte(xmlData), &resp); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Delimiter", *resp.Delimiter, "/")
	if resp.Segment == nil {
		t.Fatal("Segment is nil")
	}
	if len(resp.Segment.BlobItems) != 1 {
		t.Fatalf("expected 1 blob item, got %d", len(resp.Segment.BlobItems))
	}
	if len(resp.Segment.BlobPrefixes) != 1 {
		t.Fatalf("expected 1 blob prefix, got %d", len(resp.Segment.BlobPrefixes))
	}
	assertEqual(t, "BlobPrefix[0].Name", *resp.Segment.BlobPrefixes[0].Name, "subdir/")
}

// TestGeoReplicationRoundTrip tests marshal/unmarshal of GeoReplication with RFC1123 time.
func TestGeoReplicationRoundTrip(t *testing.T) {
	xmlData := `<GeoReplication>
  <Status>live</Status>
  <LastSyncTime>Wed, 26 Oct 2016 20:39:39 GMT</LastSyncTime>
</GeoReplication>`

	var geo GeoReplication
	if err := xml.Unmarshal([]byte(xmlData), &geo); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Status", string(*geo.Status), "live")
	expectedTime := timeRFC1123("Wed, 26 Oct 2016 20:39:39 GMT")
	if !geo.LastSyncTime.Equal(*expectedTime) {
		t.Fatalf("expected LastSyncTime %v, got %v", expectedTime, geo.LastSyncTime)
	}

	// Round-trip
	out, err := xml.Marshal(geo)
	if err != nil {
		t.Fatal(err)
	}
	var geo2 GeoReplication
	if err := xml.Unmarshal(out, &geo2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip Status", string(*geo2.Status), "live")
	if !geo2.LastSyncTime.Equal(*expectedTime) {
		t.Fatalf("round-trip expected LastSyncTime %v, got %v", expectedTime, geo2.LastSyncTime)
	}
}

// TestUserDelegationKeyRoundTrip tests marshal/unmarshal of UserDelegationKey with RFC3339 time.
func TestUserDelegationKeyRoundTrip(t *testing.T) {
	xmlData := `<UserDelegationKey>
  <SignedOid>00000000-0000-0000-0000-000000000000</SignedOid>
  <SignedTid>00000000-0000-0000-0000-000000000001</SignedTid>
  <SignedStart>2020-01-01T00:00:00Z</SignedStart>
  <SignedExpiry>2020-01-02T00:00:00Z</SignedExpiry>
  <SignedService>b</SignedService>
  <SignedVersion>2019-12-12</SignedVersion>
  <Value>dGVzdGtleQ==</Value>
</UserDelegationKey>`

	var key UserDelegationKey
	if err := xml.Unmarshal([]byte(xmlData), &key); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "SignedOid", *key.SignedOID, "00000000-0000-0000-0000-000000000000")
	assertEqual(t, "SignedTid", *key.SignedTID, "00000000-0000-0000-0000-000000000001")
	assertEqual(t, "SignedService", *key.SignedService, "b")
	assertEqual(t, "SignedVersion", *key.SignedVersion, "2019-12-12")
	assertEqual(t, "Value", *key.Value, "dGVzdGtleQ==")

	expectedStart := timeRFC3339("2020-01-01T00:00:00Z")
	if !key.SignedStart.Equal(*expectedStart) {
		t.Fatalf("expected SignedStart %v, got %v", expectedStart, key.SignedStart)
	}
	expectedExpiry := timeRFC3339("2020-01-02T00:00:00Z")
	if !key.SignedExpiry.Equal(*expectedExpiry) {
		t.Fatalf("expected SignedExpiry %v, got %v", expectedExpiry, key.SignedExpiry)
	}

	// Round-trip
	out, err := xml.Marshal(key)
	if err != nil {
		t.Fatal(err)
	}
	var key2 UserDelegationKey
	if err := xml.Unmarshal(out, &key2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip SignedOid", *key2.SignedOID, "00000000-0000-0000-0000-000000000000")
	if !key2.SignedStart.Equal(*expectedStart) {
		t.Fatalf("round-trip expected SignedStart %v, got %v", expectedStart, key2.SignedStart)
	}
}

// TestTagsRoundTrip tests marshal/unmarshal of Tags.
func TestTagsRoundTrip(t *testing.T) {
	xmlData := `<Tags>
  <TagSet>
    <Tag>
      <Key>Project</Key>
      <Value>Contoso</Value>
    </Tag>
    <Tag>
      <Key>Status</Key>
      <Value>Active</Value>
    </Tag>
  </TagSet>
</Tags>`

	var tags Tags
	if err := xml.Unmarshal([]byte(xmlData), &tags); err != nil {
		t.Fatal(err)
	}

	if len(tags.BlobTagSet) != 2 {
		t.Fatalf("expected 2 tags, got %d", len(tags.BlobTagSet))
	}
	assertEqual(t, "Tag[0].Key", *tags.BlobTagSet[0].Key, "Project")
	assertEqual(t, "Tag[0].Value", *tags.BlobTagSet[0].Value, "Contoso")
	assertEqual(t, "Tag[1].Key", *tags.BlobTagSet[1].Key, "Status")
	assertEqual(t, "Tag[1].Value", *tags.BlobTagSet[1].Value, "Active")

	// Round-trip
	out, err := xml.Marshal(tags)
	if err != nil {
		t.Fatal(err)
	}
	var tags2 Tags
	if err := xml.Unmarshal(out, &tags2); err != nil {
		t.Fatal(err)
	}
	if len(tags2.BlobTagSet) != 2 {
		t.Fatalf("round-trip expected 2 tags, got %d", len(tags2.BlobTagSet))
	}
}

// TestSignedIdentifierRoundTrip tests marshal/unmarshal of SignedIdentifier with time values.
func TestSignedIdentifierRoundTrip(t *testing.T) {
	xmlData := `<SignedIdentifier>
  <Id>MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=</Id>
  <AccessPolicy>
    <Start>2009-09-28T08:49:37.0000000Z</Start>
    <Expiry>2009-09-29T08:49:37.0000000Z</Expiry>
    <Permission>rwd</Permission>
  </AccessPolicy>
</SignedIdentifier>`

	var si SignedIdentifier
	if err := xml.Unmarshal([]byte(xmlData), &si); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Id", *si.ID, "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=")
	assertEqual(t, "Permission", *si.AccessPolicy.Permission, "rwd")

	// Round-trip
	out, err := xml.Marshal(si)
	if err != nil {
		t.Fatal(err)
	}
	var si2 SignedIdentifier
	if err := xml.Unmarshal(out, &si2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip Id", *si2.ID, "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=")
	assertEqual(t, "round-trip Permission", *si2.AccessPolicy.Permission, "rwd")
}

// TestStorageServicePropertiesMarshal tests that constructing a StorageServiceProperties and marshaling
// produces valid XML matching the Set Blob Service Properties request body.
// https://learn.microsoft.com/en-us/rest/api/storageservices/set-blob-service-properties
func TestStorageServicePropertiesMarshal(t *testing.T) {
	props := StorageServiceProperties{
		Logging: &Logging{
			Version: to.Ptr("1.0"),
			Delete:  to.Ptr(true),
			Read:    to.Ptr(false),
			Write:   to.Ptr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(true),
				Days:    to.Ptr(int32(7)),
			},
		},
		HourMetrics: &Metrics{
			Version:     to.Ptr("1.0"),
			Enabled:     to.Ptr(true),
			IncludeAPIs: to.Ptr(false),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(true),
				Days:    to.Ptr(int32(7)),
			},
		},
		CORS: []*CORSRule{
			{
				AllowedOrigins:  to.Ptr("http://www.fabrikam.com,http://www.contoso.com"),
				AllowedMethods:  to.Ptr("GET,PUT"),
				MaxAgeInSeconds: to.Ptr(int32(500)),
				ExposedHeaders:  to.Ptr("x-ms-meta-data*,x-ms-meta-customheader"),
				AllowedHeaders:  to.Ptr("x-ms-meta-target*,x-ms-meta-customheader"),
			},
		},
		DefaultServiceVersion: to.Ptr("2018-03-28"),
		DeleteRetentionPolicy: &RetentionPolicy{
			Enabled: to.Ptr(true),
			Days:    to.Ptr(int32(5)),
		},
		StaticWebsite: &StaticWebsite{
			Enabled:              to.Ptr(true),
			IndexDocument:        to.Ptr("index.html"),
			ErrorDocument404Path: to.Ptr("error/404.html"),
		},
	}

	data, err := xml.Marshal(props)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal back and verify
	var props2 StorageServiceProperties
	if err := xml.Unmarshal(data, &props2); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Logging.Version", *props2.Logging.Version, "1.0")
	assertEqual(t, "Logging.Delete", *props2.Logging.Delete, true)
	if len(props2.CORS) != 1 {
		t.Fatalf("expected 1 CorsRule, got %d", len(props2.CORS))
	}
	assertEqual(t, "CorsRule.MaxAgeInSeconds", *props2.CORS[0].MaxAgeInSeconds, int32(500))
	assertEqual(t, "DefaultServiceVersion", *props2.DefaultServiceVersion, "2018-03-28")
	if props2.DeleteRetentionPolicy == nil {
		t.Fatal("DeleteRetentionPolicy is nil after round-trip")
	}
	assertEqual(t, "DeleteRetentionPolicy.Enabled", *props2.DeleteRetentionPolicy.Enabled, true)
	if props2.StaticWebsite == nil {
		t.Fatal("StaticWebsite is nil after round-trip")
	}
	assertEqual(t, "StaticWebsite.IndexDocument", *props2.StaticWebsite.IndexDocument, "index.html")
}

// TestPropertiesInternalTimeFieldsRoundTrip tests that time fields on PropertiesInternal
// round-trip correctly through XML.
func TestPropertiesInternalTimeFieldsRoundTrip(t *testing.T) {
	xmlData := `<Properties>
  <Creation-Time>Wed, 26 Oct 2016 20:39:39 GMT</Creation-Time>
  <Last-Modified>Thu, 27 Oct 2016 10:00:00 GMT</Last-Modified>
  <Etag>0x8CACB9BD7C6B1B2</Etag>
  <Content-Length>2048</Content-Length>
  <Content-Type>text/plain</Content-Type>
  <BlobType>BlockBlob</BlobType>
  <AccessTier>Hot</AccessTier>
  <AccessTierInferred>true</AccessTierInferred>
  <ServerEncrypted>true</ServerEncrypted>
  <LeaseStatus>unlocked</LeaseStatus>
  <LeaseState>available</LeaseState>
</Properties>`

	var props Properties
	if err := xml.Unmarshal([]byte(xmlData), &props); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Etag", string(*props.ETag), "0x8CACB9BD7C6B1B2")
	assertEqual(t, "ContentLength", *props.ContentLength, int64(2048))
	assertEqual(t, "ContentType", *props.ContentType, "text/plain")
	assertEqual(t, "BlobType", string(*props.BlobType), "BlockBlob")
	assertEqual(t, "AccessTier", string(*props.AccessTier), "Hot")
	assertEqual(t, "AccessTierInferred", *props.AccessTierInferred, true)
	assertEqual(t, "ServerEncrypted", *props.ServerEncrypted, true)

	expectedCreation := timeRFC1123("Wed, 26 Oct 2016 20:39:39 GMT")
	if !props.CreationTime.Equal(*expectedCreation) {
		t.Fatalf("expected CreationTime %v, got %v", expectedCreation, props.CreationTime)
	}
	expectedModified := timeRFC1123("Thu, 27 Oct 2016 10:00:00 GMT")
	if !props.LastModified.Equal(*expectedModified) {
		t.Fatalf("expected LastModified %v, got %v", expectedModified, props.LastModified)
	}

	// Round-trip
	out, err := xml.Marshal(props)
	if err != nil {
		t.Fatal(err)
	}
	var props2 Properties
	if err := xml.Unmarshal(out, &props2); err != nil {
		t.Fatal(err)
	}
	if !props2.CreationTime.Equal(*expectedCreation) {
		t.Fatalf("round-trip expected CreationTime %v, got %v", expectedCreation, props2.CreationTime)
	}
	if !props2.LastModified.Equal(*expectedModified) {
		t.Fatalf("round-trip expected LastModified %v, got %v", expectedModified, props2.LastModified)
	}
}

// TestArrowConfigurationRoundTrip tests marshal/unmarshal of ArrowConfiguration.
func TestArrowConfigurationRoundTrip(t *testing.T) {
	xmlData := `<ArrowConfiguration>
  <Schema>
    <Field>
      <Type>INT64</Type>
      <Name>id</Name>
    </Field>
    <Field>
      <Type>DOUBLE</Type>
      <Name>value</Name>
      <Precision>10</Precision>
      <Scale>2</Scale>
    </Field>
  </Schema>
</ArrowConfiguration>`

	var ac ArrowConfiguration
	if err := xml.Unmarshal([]byte(xmlData), &ac); err != nil {
		t.Fatal(err)
	}

	if len(ac.Schema) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(ac.Schema))
	}
	assertEqual(t, "Field[0].Type", *ac.Schema[0].Type, "INT64")
	assertEqual(t, "Field[0].Name", *ac.Schema[0].Name, "id")
	assertEqual(t, "Field[1].Type", *ac.Schema[1].Type, "DOUBLE")
	assertEqual(t, "Field[1].Precision", *ac.Schema[1].Precision, int32(10))
	assertEqual(t, "Field[1].Scale", *ac.Schema[1].Scale, int32(2))

	// Round-trip
	out, err := xml.Marshal(ac)
	if err != nil {
		t.Fatal(err)
	}
	var ac2 ArrowConfiguration
	if err := xml.Unmarshal(out, &ac2); err != nil {
		t.Fatal(err)
	}
	if len(ac2.Schema) != 2 {
		t.Fatalf("round-trip expected 2 fields, got %d", len(ac2.Schema))
	}
	assertEqual(t, "round-trip Field[0].Type", *ac2.Schema[0].Type, "INT64")
}

// TestStorageServiceStatsRoundTrip tests StorageServiceStats with GeoReplication.
func TestStorageServiceStatsRoundTrip(t *testing.T) {
	xmlData := `<StorageServiceStats>
  <GeoReplication>
    <Status>live</Status>
    <LastSyncTime>Wed, 26 Oct 2016 20:39:39 GMT</LastSyncTime>
  </GeoReplication>
</StorageServiceStats>`

	var stats StorageServiceStats
	if err := xml.Unmarshal([]byte(xmlData), &stats); err != nil {
		t.Fatal(err)
	}

	if stats.GeoReplication == nil {
		t.Fatal("GeoReplication is nil")
	}
	assertEqual(t, "Status", string(*stats.GeoReplication.Status), "live")

	// Round-trip
	out, err := xml.Marshal(stats)
	if err != nil {
		t.Fatal(err)
	}
	var stats2 StorageServiceStats
	if err := xml.Unmarshal(out, &stats2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip Status", string(*stats2.GeoReplication.Status), "live")
}

// TestBlockListCommittedOnlyUnmarshal tests that committed-only block list responses parse correctly.
// https://learn.microsoft.com/en-us/rest/api/storageservices/get-block-list
func TestBlockListCommittedOnlyUnmarshal(t *testing.T) {
	xmlData := `<BlockList>
  <CommittedBlocks>
    <Block>
      <Name>QmxvY2tJZDAwMQ==</Name>
      <Size>4194304</Size>
    </Block>
    <Block>
      <Name>QmxvY2tJZDAwMg==</Name>
      <Size>4194304</Size>
    </Block>
  </CommittedBlocks>
</BlockList>`

	var blockList BlockList
	if err := xml.Unmarshal([]byte(xmlData), &blockList); err != nil {
		t.Fatal(err)
	}

	if len(blockList.CommittedBlocks) != 2 {
		t.Fatalf("expected 2 committed blocks, got %d", len(blockList.CommittedBlocks))
	}
	if blockList.UncommittedBlocks != nil {
		t.Fatalf("expected nil uncommitted blocks, got %d", len(blockList.UncommittedBlocks))
	}
}

// TestEmptySlicesRoundTrip tests that empty/nil slices are handled gracefully.
func TestEmptySlicesRoundTrip(t *testing.T) {
	props := StorageServiceProperties{
		Logging: &Logging{
			Version: to.Ptr("1.0"),
			Delete:  to.Ptr(false),
			Read:    to.Ptr(false),
			Write:   to.Ptr(false),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(false),
			},
		},
	}

	data, err := xml.Marshal(props)
	if err != nil {
		t.Fatal(err)
	}

	xmlStr := string(data)
	assertNotContains(t, "StorageServiceProperties with nil CORS", xmlStr, "<Cors>")

	var props2 StorageServiceProperties
	if err := xml.Unmarshal(data, &props2); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Logging.Version", *props2.Logging.Version, "1.0")
	if props2.CORS != nil {
		t.Fatalf("expected nil Cors, got %d rules", len(props2.CORS))
	}
}

// TestNilSlicesOmitXMLTags tests that nil slices omit their XML wrapper tags during marshaling.
func TestNilSlicesOmitXMLTags(t *testing.T) {
	// ArrowConfiguration with nil Schema
	ac := ArrowConfiguration{}
	data, err := xml.Marshal(ac)
	if err != nil {
		t.Fatal(err)
	}
	assertNotContains(t, "ArrowConfiguration with nil Schema", string(data), "<Schema>")

	// BlockList with nil CommittedBlocks and UncommittedBlocks
	bl := BlockList{}
	data, err = xml.Marshal(bl)
	if err != nil {
		t.Fatal(err)
	}
	xmlStr := string(data)
	assertNotContains(t, "BlockList with nil CommittedBlocks", xmlStr, "<CommittedBlocks>")
	assertNotContains(t, "BlockList with nil UncommittedBlocks", xmlStr, "<UncommittedBlocks>")

	// BlockLookupList with nil slices
	bll := BlockLookupList{}
	data, err = xml.Marshal(bll)
	if err != nil {
		t.Fatal(err)
	}
	xmlStr = string(data)
	assertNotContains(t, "BlockLookupList with nil Committed", xmlStr, "<Committed>")
	assertNotContains(t, "BlockLookupList with nil Latest", xmlStr, "<Latest>")
	assertNotContains(t, "BlockLookupList with nil Uncommitted", xmlStr, "<Uncommitted>")

	// FilterBlobSegment with nil Blobs
	fbs := FilterBlobSegment{
		ServiceEndpoint: to.Ptr("https://example.com/"),
		Where:           to.Ptr("tag = 'value'"),
	}
	data, err = xml.Marshal(fbs)
	if err != nil {
		t.Fatal(err)
	}
	assertNotContains(t, "FilterBlobSegment with nil Blobs", string(data), "<Blobs>")

	// PageList with nil PageRange and ClearRange
	pl := PageList{}
	data, err = xml.Marshal(pl)
	if err != nil {
		t.Fatal(err)
	}
	xmlStr = string(data)
	assertNotContains(t, "PageList with nil PageRange", xmlStr, "<PageRange>")
	assertNotContains(t, "PageList with nil ClearRange", xmlStr, "<ClearRange>")

	// Tags with nil BlobTagSet
	tags := Tags{}
	data, err = xml.Marshal(tags)
	if err != nil {
		t.Fatal(err)
	}
	assertNotContains(t, "Tags with nil BlobTagSet", string(data), "<TagSet>")

	// ListContainersSegmentResponse with nil ContainerItems
	lcsr := ListContainersSegmentResponse{
		ServiceEndpoint: to.Ptr("https://example.com/"),
	}
	data, err = xml.Marshal(lcsr)
	if err != nil {
		t.Fatal(err)
	}
	assertNotContains(t, "ListContainersSegmentResponse with nil ContainerItems", string(data), "<Containers>")

	// SignedIdentifiers with nil Items
	type wrapper struct {
		XMLName      xml.Name             `xml:"SignedIdentifiers"`
		ContainerACL *[]*SignedIdentifier `xml:"SignedIdentifier"`
	}
	data, err = xml.Marshal(wrapper{})
	if err != nil {
		t.Fatal(err)
	}
	assertNotContains(t, "SignedIdentifiers with nil Items", string(data), "<SignedIdentifier>")
}

// TestKeyInfoRoundTrip tests marshal/unmarshal of KeyInfo.
func TestKeyInfoRoundTrip(t *testing.T) {
	xmlData := `<KeyInfo>
  <Start>2020-01-01T00:00:00Z</Start>
  <Expiry>2020-01-02T00:00:00Z</Expiry>
</KeyInfo>`

	var ki KeyInfo
	if err := xml.Unmarshal([]byte(xmlData), &ki); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "Start", *ki.Start, "2020-01-01T00:00:00Z")
	assertEqual(t, "Expiry", *ki.Expiry, "2020-01-02T00:00:00Z")

	// Round-trip
	out, err := xml.Marshal(ki)
	if err != nil {
		t.Fatal(err)
	}
	var ki2 KeyInfo
	if err := xml.Unmarshal(out, &ki2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "round-trip Start", *ki2.Start, "2020-01-01T00:00:00Z")
	assertEqual(t, "round-trip Expiry", *ki2.Expiry, "2020-01-02T00:00:00Z")
}

// TestQueryRequestRoundTrip tests marshal/unmarshal of QueryRequest.
func TestQueryRequestRoundTrip(t *testing.T) {
	queryType := QueryRequestType("SQL")
	qr := QueryRequest{
		QueryType:  &queryType,
		Expression: to.Ptr("SELECT * FROM BlobStorage"),
	}

	data, err := xml.Marshal(qr) //nolint:staticcheck // we use custom helper for map[string]any
	if err != nil {
		t.Fatal(err)
	}

	var qr2 QueryRequest
	if err := xml.Unmarshal(data, &qr2); err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "QueryType", string(*qr2.QueryType), "SQL")
	assertEqual(t, "Expression", *qr2.Expression, "SELECT * FROM BlobStorage")
}

// TestJSONTextConfigurationMarshalElementName tests that JSONTextConfiguration.MarshalXML
// produces the correct element name "JsonTextConfiguration".
func TestJSONTextConfigurationMarshalElementName(t *testing.T) {
	jtc := JSONTextConfiguration{
		RecordSeparator: to.Ptr("\n"),
	}

	data, err := xml.Marshal(jtc)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the element name is "JsonTextConfiguration", not "JSONTextConfiguration"
	xmlStr := string(data)
	if !strings.Contains(xmlStr, "<JsonTextConfiguration>") {
		t.Fatalf("expected element name JsonTextConfiguration, got: %s", xmlStr)
	}

	// Round-trip via QueryFormat which embeds JSONTextConfiguration
	formatType := QueryFormatType("json")
	qf := QueryFormat{
		Type:                  &formatType,
		JSONTextConfiguration: &jtc,
	}

	data, err = xml.Marshal(qf) //nolint:staticcheck // we use custom helper for map[string]any
	if err != nil {
		t.Fatal(err)
	}

	var qf2 QueryFormat
	if err := xml.Unmarshal(data, &qf2); err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "QueryFormat.Type", string(*qf2.Type), "json")
	if qf2.JSONTextConfiguration == nil {
		t.Fatal("JSONTextConfiguration is nil after round-trip")
	}
	assertEqual(t, "RecordSeparator", *qf2.JSONTextConfiguration.RecordSeparator, "\n")
}

func assertEqual[T comparable](t *testing.T, name string, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("%s: got %v, want %v", name, got, want)
	}
}

func assertNotContains(t *testing.T, context, xmlStr, forbidden string) {
	t.Helper()
	if strings.Contains(xmlStr, forbidden) {
		t.Fatalf("%s: XML should not contain %q, got: %s", context, forbidden, xmlStr)
	}
}
