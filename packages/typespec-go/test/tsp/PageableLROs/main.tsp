import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";

using Azure.ResourceManager;

@armProviderNamespace
namespace PageableLROs;

@list
op listPrivateEndPoints is ArmResourceActionAsyncBase<
  SomeResource,
  void,
  ArmResponse<SomeResourceListResult> | ArmAcceptedLroResponse<
    "Accepted. The header contains 'Azure-AsyncOperation' header which can be used to monitor the progress of the operation.",
    ArmLroLocationHeader<FinalResult = SomeResourceListResult>
  >,
  BaseParameters = Azure.ResourceManager.Foundations.DefaultBaseParameters<SomeResource>
>;

model SomeResource
  is Azure.ResourceManager.TrackedResource<SomeResourceProperties> {
  ...ResourceNameParameter<
    Resource = SomeResource,
    KeyName = "resourceName",
    SegmentName = "resourceSegment",
    NamePattern = "^[a-zA-Z0-9][a-zA-Z0-9_-]*$"
  >;
}

model SomeResourceProperties {
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;

  @visibility(Lifecycle.Read)
  resourceIds?: string;
}

model SomeResourceListResult
  is Azure.Core.Page<SomeResourceInfo>;

model SomeResourceInfo {
  @identifiers(#["id"])
  resources?: ReferencedResource[];
}

model ReferencedResource {
  id?: string;
}

union ProvisioningState {
  string,
  Unknown: "Unknown",
  Succeeded: "Succeeded",
  Accepted: "Accepted",
  Failed: "Failed",
  Canceled: "Canceled",
}
