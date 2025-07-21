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

// Test case 1: Normal internal model without clientName (should get uncapitalize logic)
@global.Azure.ClientGenerator.Core.access(global.Azure.ClientGenerator.Core.Access.internal)
model NormalInternalModel {
  code: string;
}

// Test case 2: Internal model with clientName starting with lowercase (should preserve case)
@global.Azure.ClientGenerator.Core.access(global.Azure.ClientGenerator.Core.Access.internal)
model AcsMessageChannelEventError {
  code: string;
}

@@clientName(AcsMessageChannelEventError,
  "internalACSMessageChannelEventError",
  "go"
);

// Test case 3: Internal model with clientName starting with uppercase (should use normal uncapitalize logic)
@global.Azure.ClientGenerator.Core.access(global.Azure.ClientGenerator.Core.Access.internal)  
model UppercaseInternalModel {
  code: string;
}

@@clientName(UppercaseInternalModel,
  "UppercaseInternalModel",
  "go"
);

@route("/test")
interface TestOperations {
  @route("/normal")
  @get
  @global.Azure.ClientGenerator.Core.access(global.Azure.ClientGenerator.Core.Access.internal)
  getNormal(): NormalInternalModel;

  @route("/lowercase")
  @get
  @global.Azure.ClientGenerator.Core.access(global.Azure.ClientGenerator.Core.Access.internal)
  getLowercase(): AcsMessageChannelEventError;

  @route("/uppercase")
  @get
  @global.Azure.ClientGenerator.Core.access(global.Azure.ClientGenerator.Core.Access.internal)
  getUppercase(): UppercaseInternalModel;
}