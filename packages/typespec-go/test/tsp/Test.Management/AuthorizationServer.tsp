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

@doc("A tenant-level authorization resource.")
@tenantResource
model AuthorizationResource is ProxyResource<AuthorizationProperties> {
  @doc("Authorization name.")
  @path
  @key("authorizationName")
  @segment("authorizations")
  name: string;
}

@doc("Properties of an authorization.")
model AuthorizationProperties {
  @doc("The description.")
  description?: string;

  @doc("The provisioning state.")
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;
}

@armResourceOperations
interface Authorization {
  @doc("Get an authorization.")
  get is ArmResourceRead<
    AuthorizationResource,
    BaseParameters = Azure.ResourceManager.Foundations.TenantBaseParameters
  >;
}

@doc("A tenant-level authorization server resource.")
@tenantResource
model AuthorizationServerResource
  is ProxyResource<AuthorizationServerProperties> {
  @doc("Authorization server name.")
  @path
  @key("authorizationServerName")
  @segment("authorizationServers")
  name: string;
}

@doc("Properties of an authorization server.")
model AuthorizationServerProperties {
  @doc("The description.")
  description?: string;

  @doc("The provisioning state.")
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;
}

@armResourceOperations
interface AuthorizationServer {
  @doc("Get an authorization server.")
  get is ArmResourceRead<
    AuthorizationServerResource,
    BaseParameters = Azure.ResourceManager.Foundations.TenantBaseParameters
  >;
}
