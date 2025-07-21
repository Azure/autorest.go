import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using Azure.ClientGenerator.Core;
using Azure.Core;
using TypeSpec.Http;

@service(#{
  title: "Test Internal Naming",
})
namespace Microsoft.Test.Internal;

model AcsMessageChannelEventError {
  code: string;
}

@@clientName(AcsMessageChannelEventError,
  "internalACSMessageChannelEventError",
  "go"
);
@@access(AcsMessageChannelEventError, Access.internal, "go");

@route("/test")
interface TestOperations {
  @get
  getError(): AcsMessageChannelEventError;
}