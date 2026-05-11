import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/openapi";
import "@typespec/rest";
import "./models.tsp";

using TypeSpec.Rest;
using Azure.ResourceManager;
using TypeSpec.Http;
using TypeSpec.OpenAPI;

@armProviderNamespace
namespace Microsoft.Test;
@doc("Body root resource.")
model BodyRoot is TrackedResource<BodyRootProperties> {
  @doc("Body root resource name.")
  @pattern("^(?=.{3,24}$)[^0-9][A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$")
  @path
  @key("bodyRootName")
  @segment("bodyRoots")
  name: string;
}

@armResourceOperations
interface BodyRoots {
  @doc("Create or update a body root resource.")
  put is ArmResourceCreateOrReplaceSync<BodyRoot>;

  @doc("Get details of a certificate profile.")
  get is ArmResourceRead<BodyRoot>;

  @doc("Revoke a certificate under a certificate profile.")
  action is ArmResourceActionNoContentSync<
    BodyRoot,
    {
      @bodyRoot
      action: ActionRequest;
    },
    Parameters = {
      @doc("The constant filter value.")
      @query("filter")
      filter: "constantValue";
    }
  >;
}
