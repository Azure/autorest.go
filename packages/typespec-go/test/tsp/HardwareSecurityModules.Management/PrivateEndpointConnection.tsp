import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/openapi";
import "@typespec/rest";
import "./models.tsp";
import "./CloudHsmCluster.tsp";

using TypeSpec.Rest;
using Azure.ResourceManager;
using TypeSpec.Http;
using TypeSpec.OpenAPI;
using Azure.Core;

namespace Microsoft.HardwareSecurityModules;
/**
 * The private endpoint connection resource.
 */
@parentResource(CloudHsmCluster)
model PrivateEndpointConnection
  is Azure.ResourceManager.ProxyResource<PrivateEndpointConnectionProperties> {
  ...ResourceNameParameter<
    Resource = PrivateEndpointConnection,
    KeyName = "peConnectionName",
    SegmentName = "privateEndpointConnections"
  >;

  /**
   * Modified whenever there is a change in the state of private endpoint connection.
   */
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-invalid-envelope-property" "For backward compatibility"
  ...EtagProperty;
}

@armResourceOperations
interface PrivateEndpointConnections {
  /**
   * Gets the private endpoint connection for the Cloud Hsm Cluster.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("CloudHsmClusterPrivateEndpointConnections_Get")
  get is ArmResourceRead<PrivateEndpointConnection>;

  /**
   * Creates or updates the private endpoint connection for the Cloud Hsm Cluster.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-put-operation-response-codes" "For backward compatibility"
  @operationId("CloudHsmClusterPrivateEndpointConnections_Create")
  create is ArmResourceCreateOrReplaceSync<
    PrivateEndpointConnection,
    Response = ArmResourceUpdatedResponse<PrivateEndpointConnection>
  >;

  /**
   * Deletes the private endpoint connection for the Cloud Hsm Cluster.
   */
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "non-standard operations"
  @operationId("CloudHsmClusterPrivateEndpointConnections_Delete")
  delete is ArmResourceDeleteWithoutOkAsync<PrivateEndpointConnection>;

  /**
   * The List operation gets information about the private endpoint connections associated with the Cloud HSM Cluster
   */
  listByCloudHsmCluster is ArmResourceListByParent<PrivateEndpointConnection>;
}

@@maxLength(PrivateEndpointConnection.name, 24);
@@minLength(PrivateEndpointConnection.name, 3);
@@doc(PrivateEndpointConnection.name,
  "Name of the private endpoint connection associated with the Cloud HSM Cluster."
);
@@doc(PrivateEndpointConnection.properties, "Resource properties.");
@@doc(PrivateEndpointConnections.create::parameters.resource,
  "Parameters of the PrivateEndpointConnection"
);
