import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/openapi";
import "@typespec/rest";

using TypeSpec.Rest;
using Azure.ResourceManager;
using TypeSpec.Http;
using TypeSpec.OpenAPI;

@armProviderNamespace
namespace Microsoft.Test;

@doc("A tenant-level resource.")
@tenantResource
model TenantItem is ProxyResource<TenantItemProperties> {
  @doc("Tenant item name.")
  @path
  @key("tenantItemName")
  @segment("tenantItems")
  name: string;
}

@doc("Properties of a tenant item.")
model TenantItemProperties {
  @doc("The description.")
  description?: string;

  @doc("The status.")
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;
}

@armResourceOperations
interface TenantItems {
  @doc("Get a tenant item.")
  get is ArmResourceRead<
    TenantItem,
    BaseParameters = Azure.ResourceManager.Foundations.TenantBaseParameters
  >;

  @doc("List tenant items.")
  list is ArmResourceListByParent<
    TenantItem,
    BaseParameters = Azure.ResourceManager.Foundations.TenantBaseParameters
  >;
}

@doc("A mixed scope resource that has operations at both subscription and tenant level.")
model MixedScopeWidget is TrackedResource<MixedScopeWidgetProperties> {
  @doc("Widget name.")
  @path
  @key("widgetName")
  @segment("mixedScopeWidgets")
  name: string;
}

@doc("Properties of a mixed scope widget.")
model MixedScopeWidgetProperties {
  @doc("The description.")
  description?: string;

  @doc("The status.")
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;
}

@armResourceOperations
interface MixedScopeWidgets {
  @doc("Get a widget at resource group level (uses subscriptionId).")
  get is ArmResourceRead<MixedScopeWidget>;

  @doc("List widgets by subscription (uses subscriptionId).")
  listBySubscription is ArmListBySubscription<MixedScopeWidget>;

  @doc("List all widgets at tenant level (no subscriptionId).")
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-operation" "tenant-level custom operation"
  @operationId("MixedScopeWidgets_ListAll")
  @autoRoute
  @get
  @list
  @action("mixedScopeWidgets")
  listAll is ArmProviderActionSync<
    Response = ResourceListResult<MixedScopeWidget>
  >;
}

