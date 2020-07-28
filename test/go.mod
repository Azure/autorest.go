module generatortests

go 1.13

require (
	armnetwork v0.0.0-00010101000000-000000000000
	azblob v0.0.0-00010101000000-000000000000
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.9.2
	github.com/Azure/azure-sdk-for-go/sdk/to v0.1.0
)

replace (
	armnetwork => ./network/2020-03-01/armnetwork
	azblob => ./storage/2019-07-07/azblob
)
