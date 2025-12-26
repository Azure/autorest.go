import "@typespec/rest";
import "@typespec/versioning";
import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";

using TypeSpec.Rest;
using TypeSpec.Http;
using Azure.ResourceManager.Foundations;
using Azure.Core;
using Azure.ResourceManager;
using TypeSpec.Versioning;

@doc("Additional property test resource provider.")
@armProviderNamespace
@service(#{
  title: "AdditionalPropertyManagementClient",
})
@versioned(Versions)
@armCommonTypesVersion(Azure.ResourceManager.CommonTypes.Versions.v5)
namespace Microsoft.AdditionalProperty;

@doc("The available API versions.")
enum Versions {
  @doc("The 2025-01-01 API version.")
  // @useDependency(Azure.ResourceManager.Versions.v1_0_Preview_1)
  // @useDependency(Azure.Core.Versions.v1_0_Preview_2)
  v2025_01_01: "2025-01-01",
}

@armResourceOperations
interface TestResources {
  op createUser(user: CreateUserRequest): CreateUserResponse;
  @route("/test/anonymous1")
  op createOrder(order: Order1): CreateOrderResponseDesigned;
  @route("/test/anonymous2")
  op createNamedOrder(order: Order): CreateOrderResponseDesigned;

  @route("/test/anonymous")
  op createUserTest(
  user: {
    username: string;
    email: string;
  }
): {
  id: string;
  createdAt: utcDateTime;
};
}

model CreateUserRequest {
  username: string;
  email: string;
}

model CreateUserResponse {
  id: string;
  createdAt: utcDateTime;
}

model CreateOrderRequestDesigned {
  username: string;
  email: string;
}

model CreateOrderResponseDesigned {
  id: string;
  createdAt: utcDateTime;
}

model Order1 {
  item: string;
  shippingAddress: {
    street: string;
    city: string;
    zipCode: string;
  };
}

model Address {
  street: string;
  city: string;
  zipCode: string;
}

model Order {
  item: string;
  shippingAddress: Address;
}
