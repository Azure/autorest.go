import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using Azure.ClientGenerator.Core;
using Azure.Core;
using TypeSpec.Http;

@service(#{
  title: "Name Collision Test",
})
namespace Search;

// Union type that will become an enum
union IndexerStatus {
  string,
  Running: "running",
  Stopped: "stopped",
  Failed: "failed",
}

// Model type that has package prefix "Search" which will be trimmed
// After trimming "Search", this becomes "IndexerStatus", which collides with the union name
model SearchIndexerStatus {
  status: IndexerStatus;
  message: string;
}

@route("/indexer")
@tag("Indexer")
interface IndexerOperations {
  @get
  getStatus(): SearchIndexerStatus;
}
