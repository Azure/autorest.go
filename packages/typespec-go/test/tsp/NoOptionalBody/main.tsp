import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using Azure.ClientGenerator.Core;
using Azure.Core;
using TypeSpec.Http;

@service(#{
  title: "No optional body params",
})
namespace Microsoft.No.Optional.Body;

model Widget {
  weight: int32;
}

@route("/widgets")
@tag("Widgets")
interface Widgets {
  @put
  put(@body body?: Widget): void
}
