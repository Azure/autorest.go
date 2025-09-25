
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

@armProviderNamespace
@service(#{
  title: "Random",
})
@versioned(Versions)
@doc("Random stuff to test various things.")
namespace Microsoft.Random;

@doc("The available API versions.")
enum Versions {
  /** The initial service version */
  v2024_03_01: "2024-03-01",
}

#suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-operation-missing-decorator" "Resource POST operation must be decorated with @armResourceAction or @armResourceCollectionAction."
interface SomeService {
  @doc("Misc test APIs")
  @post
  @list
  @armResourceCollectionAction
  @route("/providers/Microsoft.Random/listThings")
  listThings(
    ...ApiVersionParameter,
  ): ThingsListResult | ErrorResponse;

  @doc("Return trial status for subscription by region")
  @armResourceAction(Location)
  checkTrialAvailability(
    @doc("Optionally, check for a specific SKU")
    @bodyRoot
    sku?: Azure.ResourceManager.Foundations.Sku,
  ): ArmResponse<Trial> | ErrorResponse;
}

interface Random {
  @doc("Get details of a certificate profile.")
  get is ArmResourceRead<FakeData>;
}

@doc("List of made up things")
model ThingsListResult {
  @doc("The fake items on this page")
  @pageItems
  value: FakeData[];

  @doc("The link to the next page of items")
  @nextLink
  nextLink?: string;

  @doc("More fake things")
  moreStuff?: FakeInfo[];
}

@doc("Response for fake things.")
model FakeData {
  @doc("Some string value")
  someString: string;

  @doc("Some floating point value")
  someFloat: float64;
}

@doc("Fake info that's referenced only via paged result.")
model FakeInfo {
  @doc("Some string value")
  someString: string;

  @doc("Some integer value")
  someNumber: int32;
}

@doc("A location resource")
@subscriptionResource
model Location is ProxyResource<LocationProperties> {
  #suppress "@azure-tools/typespec-azure-core/no-private-usage" "Legacy DO NOT EVER COPY THIS"
  @doc("A location in a subscription")
  @pattern("^[-\\w\\._]+$")
  @key("location")
  @path
  @segment("locations")
  @Azure.ResourceManager.CommonTypes.Private.armCommonParameter(
    "LocationParameter",
    Azure.ResourceManager.CommonTypes.Versions.v5
  )
  name: string;
}

@doc("The properties of a location resource")
model LocationProperties {
  @doc("The provisioning state of the resource.")
  @visibility(Lifecycle.Read)
  provisioningState?: LocationProvisioningState;
}

@doc("location provisioning state")
union LocationProvisioningState {
  string,
  ResourceProvisioningState,
}

@doc("Subscription trial availability")
model Trial {
  @doc("Trial status")
  @visibility(Lifecycle.Read)
  status?: TrialStatus;

  @doc("Number of trial hosts available")
  @visibility(Lifecycle.Read)
  availableHosts?: int32;
}

@doc("trial status")
union TrialStatus {
  string,

  @doc("is available")
  TrialAvailable: "TrialAvailable",

  @doc("is used")
  TrialUsed: "TrialUsed",

  @doc("is disabled")
  TrialDisabled: "TrialDisabled",
}
