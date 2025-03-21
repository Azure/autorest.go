
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
  @useDependency(Azure.ResourceManager.Versions.v1_0_Preview_1)
  @useDependency(Azure.Core.Versions.v1_0_Preview_1)
  v2024_03_01: "2024-03-01",
}

#suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-operation-missing-decorator" "Resource POST operation must be decorated with @armResourceAction or @armResourceCollectionAction."
interface SomeService {
  @doc("Misc test APIs")
  @post
  @armResourceCollectionAction
  @route("/providers/Microsoft.Random/listThings")
  listThings(
    ...ApiVersionParameter,
  ): ThingsListResult | ErrorResponse;
}

@doc("List of made up things")
@Azure.Core.pagedResult
model ThingsListResult {
  @doc("The fake items on this page")
  @items
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
