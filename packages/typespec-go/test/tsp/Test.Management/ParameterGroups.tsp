import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/rest";
import "@typespec/http";
import "./models.tsp";

using Azure.ResourceManager;
using TypeSpec.Rest;
using TypeSpec.Http;

namespace Microsoft.Test;

// Test parameter group model - should preserve "TestManagement" prefix
// and not get stuttered to just "RequestOptions"
@doc("Request options for test operations")
model `test-management-request-options` {
  @doc("Client request ID for tracking")
  @header("x-ms-client-request-id")
  clientRequestId: string;

  @doc("Optional correlation ID")
  @header("x-ms-correlation-id")
  correlationId?: string;
}

@doc("Widget resource for testing parameter groups")
model Widget is TrackedResource<WidgetProperties> {
  @doc("The name of the widget")
  @key("widgetName")
  @segment("widgets")
  @path
  name: string;
}

@armResourceOperations
interface ParameterGroupOperations {
  @doc("Test operation 1 using parameter group")
  @get
  @armResourceRead(Widget)
  testOperation1(
    ...ResourceInstanceParameters<Widget>,
    @doc("Parameter group options")
    params: `test-management-request-options`,
  ): ArmResponse<Widget> | ErrorResponse;

  @doc("Test operation 2 using the same parameter group - should reuse, not duplicate")
  @post
  @armResourceAction(Widget)
  testOperation2(
    ...ResourceInstanceParameters<Widget>,
    @doc("Parameter group options")
    params: `test-management-request-options`,
  ): ArmResponse<Widget> | ErrorResponse;
}
