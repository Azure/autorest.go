import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/openapi";
import "@typespec/rest";
import "./models.tsp";

using TypeSpec.Rest;
using Azure.ResourceManager;
using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace Microsoft.HardwareSecurityModules;
/**
 * Resource information with extended details.
 */
#suppress "@azure-tools/typespec-azure-core/composition-over-inheritance"
#suppress "@azure-tools/typespec-azure-core/no-private-usage"
@Azure.ResourceManager.Private.armResourceInternal(DedicatedHsmProperties)
@TypeSpec.Http.Private.includeInapplicableMetadataInPayload(false)
model DedicatedHsm extends Foundations.TrackedResource {
  ...ResourceNameParameter<
    Resource = DedicatedHsm,
    KeyName = "name",
    SegmentName = "dedicatedHSMs"
  >;

  /**
   * SKU details
   */
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-invalid-envelope-property" "For backward compatibility"
  sku: Sku;

  ...Azure.ResourceManager.AvailabilityZonesProperty;

  /**
   * Properties of the dedicated HSM
   */
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-invalid-envelope-property"
  @Azure.ResourceManager.Private.conditionalClientFlatten
  @Azure.ResourceManager.Private.armResourcePropertiesOptionality(false)
  properties: DedicatedHsmProperties;
}

@armResourceOperations
interface DedicatedHsms {
  /**
   * Gets the specified Azure dedicated HSM.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("DedicatedHsm_Get")
  get is ArmResourceRead<DedicatedHsm, Error = DedicatedHsmError>;

  /**
   * Create or Update a dedicated HSM in the specified subscription.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("DedicatedHsm_CreateOrUpdate")
  createOrUpdate is ArmResourceCreateOrReplaceAsync<
    DedicatedHsm,
    LroHeaders = ArmLroLocationHeader &
      Azure.Core.Foundations.RetryAfterHeader,
    Error = DedicatedHsmError
  >;

  /**
   * Update a dedicated HSM in the specified subscription.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @patch(#{ implicitOptionality: false })
  @operationId("DedicatedHsm_Update")
  update is ArmCustomPatchAsync<
    DedicatedHsm,
    PatchModel = DedicatedHsmPatchParameters,
    Error = DedicatedHsmError
  >;

  /**
   * Deletes the specified Azure Dedicated HSM.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("DedicatedHsm_Delete")
  delete is ArmResourceDeleteWithoutOkAsync<
    DedicatedHsm,
    Error = DedicatedHsmError
  >;

  /**
   * The List operation gets information about the dedicated hsms associated with the subscription and within the specified resource group.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("DedicatedHsm_ListByResourceGroup")
  listByResourceGroup is ArmResourceListByParent<
    DedicatedHsm,
    Parameters = {
      /**
       * Maximum number of results to return.
       */
      @query("$top")
      $top?: int32;
    },
    Error = DedicatedHsmError
  >;

  /**
   * The List operation gets information about the dedicated HSMs associated with the subscription.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("DedicatedHsm_ListBySubscription")
  listBySubscription is ArmListBySubscription<
    DedicatedHsm,
    Parameters = {
      /**
       * Maximum number of results to return.
       */
      @query("$top")
      $top?: int32;
    },
    Error = DedicatedHsmError
  >;

  /**
   * Gets a list of egress endpoints (network endpoints of all outbound dependencies) in the specified dedicated hsm resource. The operation returns properties of each egress endpoint.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @get
  @action("outboundNetworkDependenciesEndpoints")
  @operationId("DedicatedHsm_ListOutboundNetworkDependenciesEndpoints")
  listOutboundNetworkDependenciesEndpoints is ArmResourceActionSync<
    DedicatedHsm,
    void,
    ArmResponse<OutboundEnvironmentEndpointCollection>,
    Error = DedicatedHsmError
  >;
}

@@doc(DedicatedHsm.name, "Name of the dedicated Hsm");
@@doc(DedicatedHsm.properties, "Properties of the dedicated HSM");
@@doc(DedicatedHsms.createOrUpdate::parameters.resource,
  "Parameters to create or update the dedicated hsm"
);
@@doc(DedicatedHsms.update::parameters.properties,
  "Parameters to patch the dedicated HSM"
);
