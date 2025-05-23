import "@typespec/rest";
import "@typespec/http";
import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";

using TypeSpec.Rest;
using TypeSpec.Http;
using Azure.ResourceManager;
using Azure.ResourceManager.Foundations;

namespace Microsoft.HardwareSecurityModules;

@@nextLink(CommonTypes.PrivateLinkResourceListResult.nextLink);
@@Azure.Core.items(CommonTypes.PrivateLinkResourceListResult.value);

/**
 * State of security domain activation
 */
union ActivationState {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  NotDefined: "NotDefined",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  NotActivated: "NotActivated",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Active: "Active",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Failed: "Failed",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Unknown: "Unknown",
}

/**
 * The Cloud HSM Cluster's auto-generated Domain Name Label Scope
 */
union AutoGeneratedDomainNameLabelScope {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  TenantReuse: "TenantReuse",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  SubscriptionReuse: "SubscriptionReuse",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  ResourceGroupReuse: "ResourceGroupReuse",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  NoReuse: "NoReuse",
}

/**
 * The private endpoint connection status.
 */
union PrivateEndpointServiceConnectionStatus {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Pending: "Pending",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Approved: "Approved",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Rejected: "Rejected",
}

/**
 * The current provisioning state.
 */
union PrivateEndpointConnectionProvisioningState {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Succeeded: "Succeeded",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Creating: "Creating",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Deleting: "Deleting",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Failed: "Failed",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Updating: "Updating",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  InternalError: "InternalError",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Canceled: "Canceled",
}

/**
 * The Cloud HSM Cluster's provisioningState
 */
union ProvisioningState {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Provisioning: "Provisioning",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Succeeded: "Succeeded",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Failed: "Failed",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Deleting: "Deleting",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Canceled: "Canceled",
}

/**
 * The Cloud HSM Cluster public network access
 */
union PublicNetworkAccess {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Disabled: "Disabled",
}

/**
 * Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
 */
union ManagedServiceIdentityType {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  None: "None",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  SystemAssigned: "SystemAssigned",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  UserAssigned: "UserAssigned",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  `SystemAssigned,UserAssigned`: "SystemAssigned,UserAssigned",
}

/**
 * Sku family of the Cloud HSM Cluster
 */
union CloudHsmClusterSkuFamily {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  B: "B",
}

/**
 * Status of the backup/restore operation
 */
union BackupRestoreOperationStatus {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  InProgress: "InProgress",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Succeeded: "Succeeded",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Failed: "Failed",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Cancelled: "Cancelled",
}

/**
 * SKU of the dedicated HSM
 */
union SkuName {
  string,

  /**
   * The dedicated HSM is a Safenet Luna Network HSM A790 device.
   */
  `SafeNet Luna Network HSM A790`: "SafeNet Luna Network HSM A790",

  /**
   * The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 1 local master key which supports up to 60 calls per second.
   */
  payShield10K_LMK1_CPS60: "payShield10K_LMK1_CPS60",

  /**
   * The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 1 local master key which supports up to 250 calls per second.
   */
  payShield10K_LMK1_CPS250: "payShield10K_LMK1_CPS250",

  /**
   * The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 1 local master key which supports up to 2500 calls per second.
   */
  payShield10K_LMK1_CPS2500: "payShield10K_LMK1_CPS2500",

  /**
   * The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 2 local master keys which supports up to 60 calls per second.
   */
  payShield10K_LMK2_CPS60: "payShield10K_LMK2_CPS60",

  /**
   * The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 2 local master keys which supports up to 250 calls per second.
   */
  payShield10K_LMK2_CPS250: "payShield10K_LMK2_CPS250",

  /**
   * The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 2 local master keys which supports up to 2500 calls per second.
   */
  payShield10K_LMK2_CPS2500: "payShield10K_LMK2_CPS2500",
}

/**
 * Provisioning state.
 */
#suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-provisioning-state" "For backward compatibility"
union JsonWebKeyType {
  string,

  /**
   * The dedicated HSM has been fully provisioned.
   */
  Succeeded: "Succeeded",

  /**
   * The dedicated HSM is currently being provisioned.
   */
  Provisioning: "Provisioning",

  /**
   * A device is currently being allocated for the dedicated HSM resource.
   */
  Allocating: "Allocating",

  /**
   * The dedicated HSM is being connected to the virtual network.
   */
  Connecting: "Connecting",

  /**
   * Provisioning of the dedicated HSM has failed.
   */
  Failed: "Failed",

  /**
   * Validating the subscription has sufficient quota to allocate a dedicated HSM device.
   */
  CheckingQuota: "CheckingQuota",

  /**
   * The dedicated HSM is currently being deleted.
   */
  Deleting: "Deleting",
}

/**
 * The type of identity.
 */
union IdentityType {
  string,
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  User: "User",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Application: "Application",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  ManagedIdentity: "ManagedIdentity",
  #suppress "@azure-tools/typespec-azure-core/documentation-required" "For backward compatibility"
  Key: "Key",
}

/**
 * Sku name of the Cloud HSM Cluster
 */
#suppress "@azure-tools/typespec-azure-core/no-enum" "For backward compatibility"
enum CloudHsmClusterSkuName {
  Standard_B1,
  `Standard B10`,
}

/**
 * Properties of a Cloud HSM Cluster.
 */
model CloudHsmClusterProperties {
  /**
   * State of security domain activation
   */
  @visibility(Lifecycle.Read)
  activationState?: ActivationState;

  /**
   * The Cloud HSM Cluster's auto-generated Domain Name Label Scope
   */
  autoGeneratedDomainNameLabelScope?: AutoGeneratedDomainNameLabelScope;

  /**
   * An array of Cloud HSM Cluster's HSMs
   */
  @visibility(Lifecycle.Read)
  @OpenAPI.extension("x-ms-identifiers", #["fqdn"])
  hsms?: CloudHsmProperties[];

  /**
   * List of private endpoint connection resources
   */
  @visibility(Lifecycle.Read)
  @OpenAPI.extension("x-ms-identifiers", #["name"])
  privateEndpointConnections?: PrivateEndpointConnection[];

  /**
   * The Cloud HSM Cluster's provisioningState
   */
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;

  /**
   * The Cloud HSM Cluster public network access
   */
  publicNetworkAccess?: PublicNetworkAccess;

  /**
   * Cloud HSM Cluster status message
   */
  @visibility(Lifecycle.Read)
  statusMessage?: string;
}

/**
 * The Cloud HSM Properties
 */
model CloudHsmProperties {
  /**
   * FQDN of the Cloud HSM
   */
  fqdn?: string;

  /**
   * The Cloud HSM State. Values are: Deploying, ConfiguringSlb, Starting, Starting, Failed, Failed, Deleting, DeletingSlbEntry, InitialProvisioning, Updating
   */
  state?: string;

  /**
   * The Cloud HSM State message
   */
  stateMessage?: string;
}

/**
 * Properties of the private endpoint connection.
 */
model PrivateEndpointConnectionProperties {
  /**
   * The private endpoint resource.
   */
  privateEndpoint?: CommonTypes.PrivateEndpoint;

  /**
   * A collection of information about the state of the connection between service consumer and provider.
   */
  privateLinkServiceConnectionState: CommonTypes.PrivateLinkServiceConnectionState;

  /**
   * The provisioning state of the private endpoint connection resource.
   */
  @visibility(Lifecycle.Read)
  provisioningState?: PrivateEndpointConnectionProvisioningState;

  /**
   * The group ids for the private endpoint resource.
   */
  @visibility(Lifecycle.Read)
  groupIds?: string[];
}

/**
 * The private endpoint resource.
 */
model PrivateEndpoint {
  /**
   * The ARM identifier for private endpoint.
   */
  @visibility(Lifecycle.Read)
  id?: string;
}

/**
 * A collection of information about the state of the connection between service consumer and provider.
 */
model PrivateLinkServiceConnectionState {
  /**
   * Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
   */
  status?: PrivateEndpointServiceConnectionStatus;

  /**
   * The reason for approval/rejection of the connection.
   */
  description?: string;

  /**
   * A message indicating if changes on the service provider require any updates on the consumer.
   */
  actionsRequired?: string;
}

/**
 * Common fields that are returned in the response for all Azure Resource Manager resources
 */
model Resource {
  /**
   * Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
   */
  @visibility(Lifecycle.Read)
  id?: Azure.Core.armResourceIdentifier;

  /**
   * The name of the resource
   */
  @visibility(Lifecycle.Read)
  name?: string;

  /**
   * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
   */
  @visibility(Lifecycle.Read)
  type?: string;

  /**
   * Azure Resource Manager metadata containing createdBy and modifiedBy information.
   */
  @visibility(Lifecycle.Read)
  systemData?: SystemData;
}

/**
 * Cloud Hsm Cluster SKU information
 */
model CloudHsmClusterSku {
  /**
   * Sku family of the Cloud HSM Cluster
   */
  family: CloudHsmClusterSkuFamily;

  /**
   * Sku name of the Cloud HSM Cluster
   */
  name: CloudHsmClusterSkuName;

  /**
   * Sku capacity
   */
  capacity?: int32;
}

/**
 * Patchable properties of the Cloud HSM Cluster
 */
#suppress "@azure-tools/typespec-azure-resource-manager/patch-envelope"
model CloudHsmClusterPatchParameters {
  /**
   * The Cloud HSM Cluster's tags
   */
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-no-record" "For backward compatibility"
  tags?: Record<string>;

  /**
   * Managed service identity (system assigned and/or user assigned identities)
   */
  identity?: CommonTypes.ManagedServiceIdentity;
}

/**
 * A private link resource.
 */
#suppress "@azure-tools/typespec-azure-core/composition-over-inheritance" "For backward compatibility"
model PrivateLinkResource extends Resource {
  /**
   * Resource properties.
   */
  #suppress "@azure-tools/typespec-azure-core/no-private-usage" "For backward compatibility"
  @Azure.ResourceManager.Private.conditionalClientFlatten
  properties?: PrivateLinkResourceProperties;
}

/**
 * Properties of a private link resource.
 */
model PrivateLinkResourceProperties {
  /**
   * The private link resource group id.
   */
  @visibility(Lifecycle.Read)
  groupId?: string;

  /**
   * The private link resource required member names.
   */
  @visibility(Lifecycle.Read)
  requiredMembers?: string[];

  /**
   * The private link resource private link DNS zone name.
   */
  requiredZoneNames?: string[];
}

/**
 * Backup properties
 */
#suppress "@azure-tools/typespec-azure-core/composition-over-inheritance" "For backward compatibility"
#suppress "@azure-tools/typespec-azure-resource-manager/no-empty-model"
model BackupRequestProperties extends BackupRestoreRequestBaseProperties {}

/**
 * Backup and Restore Common properties
 */
model BackupRestoreRequestBaseProperties {
  /**
   * The Azure blob storage container Uri which contains the backup
   */
  azureStorageBlobContainerUri: url;

  /**
   * The SAS token pointing to an Azure blob storage container. This property is reserved for Azure Backup Service.
   */
  @secret
  token?: string;
}

/**
 * Backup operation Result
 */
model BackupResult {
  /**
   * Properties of the Cloud HSM Cluster
   */
  #suppress "@azure-tools/typespec-azure-core/no-private-usage" "For backward compatibility"
  @Azure.ResourceManager.Private.conditionalClientFlatten
  properties?: BackupResultProperties;
}

/**
 * Properties of the Cloud HSM Cluster
 */
#suppress "@azure-tools/typespec-azure-core/composition-over-inheritance" "For backward compatibility"
model BackupResultProperties extends BackupRestoreBaseResultProperties {
  /**
   * The Azure blob storage container Uri which contains the backup
   */
  azureStorageBlobContainerUri?: url;

  /**
   * The ID of the backup.
   */
  backupId?: string;
}

/**
 * Backup and Restore operation common properties
 */
model BackupRestoreBaseResultProperties {
  /**
   * Status of the backup/restore operation
   */
  @visibility(Lifecycle.Read)
  status?: BackupRestoreOperationStatus;

  /**
   * The status details of backup/restore operation
   */
  statusDetails?: string;

  /**
   * Error encountered, if any, during the backup/restore operation.
   */
  error?: ErrorDetail;

  /**
   * The start time of the backup/restore operation in UTC
   */
  @visibility(Lifecycle.Read)
  // FIXME: (utcDateTime) Please double check that this is the correct type for your scenario.
  startTime?: utcDateTime;

  /**
   * The end time of the backup/restore operation in UTC
   */
  #suppress "@azure-tools/typespec-azure-core/no-nullable"
  @visibility(Lifecycle.Read)
  // FIXME: (utcDateTime) Please double check that this is the correct type for your scenario.
  endTime?: utcDateTime | null;

  /**
   * Identifier for the backup/restore operation.
   */
  jobId?: string;
}

/**
 * Cloud Hsm Cluster restore information
 */
#suppress "@azure-tools/typespec-azure-core/composition-over-inheritance" "For backward compatibility"
model RestoreRequestProperties extends BackupRestoreRequestBaseProperties {
  /**
   * An autogenerated unique string ID for labeling the backup. It contains both a UUID and a date timestamp.
   */
  backupId: string;
}

/**
 * Restore operation properties
 */
model RestoreResult {
  /**
   * Backup and Restore operation common properties
   */
  properties?: BackupRestoreBaseResultProperties;
}

/**
 * SKU of the dedicated HSM
 */
model Sku {
  /**
   * SKU of the dedicated HSM
   */
  name?: SkuName;
}

/**
 * Properties of the dedicated hsm
 */
model DedicatedHsmProperties {
  /**
   * Specifies the network interfaces of the dedicated hsm.
   */
  networkProfile?: NetworkProfile;

  /**
   * Specifies the management network interfaces of the dedicated hsm.
   */
  managementNetworkProfile?: NetworkProfile;

  /**
   * This field will be used when RP does not support Availability zones.
   */
  stampId?: string;

  /**
   * Resource Status Message.
   */
  @visibility(Lifecycle.Read)
  statusMessage?: string;

  /**
   * Provisioning state.
   */
  @visibility(Lifecycle.Read)
  provisioningState?: JsonWebKeyType;
}

/**
 * The network profile definition.
 */
model NetworkProfile {
  /**
   * Specifies the identifier of the subnet.
   */
  subnet?: ApiEntityReference;

  /**
   * Specifies the list of resource Ids for the network interfaces associated with the dedicated HSM.
   */
  @OpenAPI.extension("x-ms-identifiers", #[])
  networkInterfaces?: NetworkInterface[];
}

/**
 * The API entity reference.
 */
model ApiEntityReference {
  /**
   * The Azure resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
   */
  resourceId?: Azure.Core.armResourceIdentifier;
}

/**
 * The network interface definition.
 */
model NetworkInterface {
  /**
   * The Azure resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
   */
  @visibility(Lifecycle.Read)
  resourceId?: Azure.Core.armResourceIdentifier;

  /**
   * Private Ip address of the interface
   */
  privateIpAddress?: string;
}

/**
 * The error exception.
 */
@error
model DedicatedHsmError {
  /**
   * The error detail of the operation if any.
   */
  @visibility(Lifecycle.Read)
  error?: Error;
}

/**
 * The key vault server error.
 */
model Error {
  /**
   * The error code.
   */
  @visibility(Lifecycle.Read)
  code?: string;

  /**
   * The error message.
   */
  @visibility(Lifecycle.Read)
  message?: string;

  /**
   * Contains more specific error that narrows down the cause. May be null.
   */
  @visibility(Lifecycle.Read)
  innererror?: Error;
}

/**
 * Patchable properties of the dedicated HSM
 */
#suppress "@azure-tools/typespec-azure-resource-manager/patch-envelope"
model DedicatedHsmPatchParameters {
  /**
   * Resource tags
   */
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-no-record" "For backward compatibility"
  tags?: Record<string>;
}

/**
 * Collection of OutboundEnvironmentEndpoint
 */
model OutboundEnvironmentEndpointCollection
  is Azure.Core.Page<OutboundEnvironmentEndpoint>;

/**
 * Egress endpoints which dedicated hsm service connects to for common purpose.
 */
model OutboundEnvironmentEndpoint {
  /**
   * The category of endpoints accessed by the dedicated hsm service, e.g. azure-resource-management, apiserver, etc.
   */
  category?: string;

  /**
   * The endpoints that dedicated hsm service connects to
   */
  @OpenAPI.extension("x-ms-identifiers", #[])
  endpoints?: EndpointDependency[];
}

/**
 * A domain name that dedicated hsm services are reaching at.
 */
model EndpointDependency {
  /**
   * The domain name of the dependency.
   */
  domainName?: string;

  /**
   * The Ports and Protocols used when connecting to domainName.
   */
  @OpenAPI.extension("x-ms-identifiers", #[])
  endpointDetails?: EndpointDetail[];
}

/**
 * Connect information from the dedicated hsm service to a single endpoint.
 */
model EndpointDetail {
  /**
   * An IP Address that Domain Name currently resolves to.
   */
  ipAddress?: string;

  /**
   * The port an endpoint is connected to.
   */
  port?: int32;

  /**
   * The protocol used for connection
   */
  protocol?: string;

  /**
   * Description of the detail
   */
  description?: string;
}
