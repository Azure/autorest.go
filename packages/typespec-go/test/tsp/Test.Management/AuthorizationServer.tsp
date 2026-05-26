import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/openapi";
import "@typespec/rest";

using TypeSpec.Rest;
using Azure.ResourceManager;
using TypeSpec.Http;
using TypeSpec.OpenAPI;

// Regression test for https://github.com/Azure/autorest.go/issues/<duplicate-interceptor>
// Two clients whose fake server names differ only by a trailing duplicate word
// ("AuthorizationServer" vs. "AuthorizationServerServer") must not collide on
// the generated `*TransportInterceptor` variable in the fake package.
// This mirrors the pattern observed in armapimanagement.

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
