import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/openapi";
import "@typespec/rest";
import "./models.tsp";

using TypeSpec.Rest;
using Azure.ResourceManager;
using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace Microsoft.BillingBenefits;
/**
 * Savings plan order
 */
#suppress "@azure-tools/typespec-azure-core/no-private-usage"
#suppress "@azure-tools/typespec-azure-core/composition-over-inheritance"
@Azure.ResourceManager.Private.armResourceInternal(SavingsPlanOrderModelProperties)
@TypeSpec.Http.Private.includeInapplicableMetadataInPayload(false)
@tenantResource
model SavingsPlanOrderModel extends Foundations.ProxyResource {
  ...ResourceNameParameter<
    Resource = SavingsPlanOrderModel,
    KeyName = "savingsPlanOrderId",
    SegmentName = "savingsPlanOrders",
    NamePattern = ""
  >;

  /**
   * Savings plan SKU
   */
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-invalid-envelope-property" "For backward compatibility"
  sku: ResourceSku;

  #suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-invalid-envelope-property" "For backward compatibility"
  #suppress "@azure-tools/typespec-azure-core/no-private-usage"
  @doc("The resource-specific properties for this resource.")
  @Azure.ResourceManager.Private.conditionalClientFlatten
  properties?: SavingsPlanOrderModelProperties;
}

@armResourceOperations
interface SavingsPlanOrderModels {
  /**
   * Get a savings plan order.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("SavingsPlanOrder_Get")
  get is ArmResourceRead<
    SavingsPlanOrderModel,
    BaseParameters = Azure.ResourceManager.Foundations.TenantBaseParameters,
    Parameters = {
      /**
       * May be used to expand the detail information of some properties.
       */
      @query("$expand")
      $expand?: string;
    }
  >;

  /**
   * List all Savings plan orders.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("SavingsPlanOrder_List")
  list is ArmResourceListByParent<
    SavingsPlanOrderModel,
    BaseParameters = Azure.ResourceManager.Foundations.TenantBaseParameters,
    Response = ArmResponse<SavingsPlanOrderModelList>
  >;

  /**
   * Elevate as owner on savings plan order based on billing permissions.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("SavingsPlanOrder_Elevate")
  elevate is ArmResourceActionSync<
    SavingsPlanOrderModel,
    void,
    ArmResponse<RoleAssignmentEntity>,
    BaseParameters = Azure.ResourceManager.Foundations.TenantBaseParameters
  >;
}

@@doc(SavingsPlanOrderModel.name, "Order ID of the savings plan");
@@doc(SavingsPlanOrderModel.properties, "Savings plan order properties");
