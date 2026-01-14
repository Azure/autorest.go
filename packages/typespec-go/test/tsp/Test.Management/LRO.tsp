import "@typespec/http";
import "@typespec/rest";
import "@typespec/versioning";
import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.Versioning;
using Azure.Core;
using Azure.ResourceManager;

namespace Microsoft.Test;

@doc("LRO model resource.")
model TestLROModel is TrackedResource<BodyRootProperties> {
  @doc("LRO model resource name.")
  @pattern("^(?=.{3,24}$)[^0-9][A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$")
  @path
  @key("LROModelName")
  @segment("LROModels")
  name: string;
}

@armResourceOperations
interface LRO {
    @post
    @doc("A sample LRO with Azure-AsyncOperation header in 200 response")
    okResponseWithAsyncHeader is  ArmResourceCreateOrReplaceAsync<
    TestLROModel,
    Response = (ArmResponse<TestLROModel> &
      ArmCombinedLroHeaders<FinalResult = TestLROModel>) | (ArmCreatedResponse<TestLROModel> &
      ArmCombinedLroHeaders<FinalResult = TestLROModel>)
  >;
}
