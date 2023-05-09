module generatortests

go 1.18

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.6.1-0.20230509175118-ce926c4cadfc
	github.com/Azure/azure-sdk-for-go/sdk/tracing/azotel v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.9
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/otel v1.11.2
	go.opentelemetry.io/otel/exporters/jaeger v1.11.0
	go.opentelemetry.io/otel/sdk v1.11.2
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/Azure/azure-sdk-for-go/sdk/tracing/azotel => ../../../../../azure-sdk-for-go/sdk/tracing/azotel
