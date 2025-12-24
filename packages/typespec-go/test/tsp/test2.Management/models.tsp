import "@typespec/rest";
import "@typespec/http";
import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";

using TypeSpec.Rest;
using TypeSpec.Http;
using Azure.ResourceManager;
using Azure.ResourceManager.Foundations;
using Azure.Core;

namespace Microsoft.Web;

/**
 * Collection of certificates.
 */
model CertificateCollection is Azure.Core.Page<Certificate>;

/**
 * Certificate resource specific properties
 */
#suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-provisioning-state" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
model CertificateProperties {
  /**
   * Certificate password.
   */
  #suppress "@azure-tools/typespec-azure-resource-manager/secret-prop" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
  password?: string;

  /**
   * Friendly name of the certificate.
   */
  @visibility(Lifecycle.Read)
  friendlyName?: string;

  /**
   * Subject name of the certificate.
   */
  @visibility(Lifecycle.Read)
  subjectName?: string;

  /**
   * Host names the certificate applies to.
   */
  hostNames?: string[];

  /**
   * Pfx blob.
   */
  pfxBlob?: bytes;

  /**
   * App name.
   */
  @visibility(Lifecycle.Read)
  siteName?: string;

  /**
   * Self link.
   */
  @visibility(Lifecycle.Read)
  selfLink?: string;

  /**
   * Certificate issuer.
   */
  @visibility(Lifecycle.Read)
  issuer?: string;

  /**
   * Certificate issue Date.
   */
  @visibility(Lifecycle.Read)
  // FIXME: (utcDateTime) Please double check that this is the correct type for your scenario.
  issueDate?: utcDateTime;

  /**
   * Certificate expiration date.
   */
  @visibility(Lifecycle.Read)
  // FIXME: (utcDateTime) Please double check that this is the correct type for your scenario.
  expirationDate?: utcDateTime;

  /**
   * Certificate thumbprint.
   */
  @visibility(Lifecycle.Read)
  thumbprint?: string;

  /**
   * Is the certificate valid?.
   */
  @visibility(Lifecycle.Read)
  valid?: boolean;

  /**
   * Raw bytes of .cer file
   */
  @visibility(Lifecycle.Read)
  cerBlob?: bytes;

  /**
   * Public key hash.
   */
  @visibility(Lifecycle.Read)
  publicKeyHash?: string;

  /**
   * Specification for the App Service Environment to use for the certificate.
   */
  @visibility(Lifecycle.Read)
  hostingEnvironmentProfile?: HostingEnvironmentProfile;

  /**
   * Azure Key Vault Csm resource Id.
   */
  keyVaultId?: Azure.Core.armResourceIdentifier<[
    {
      type: "Microsoft.KeyVault/vaults";
    }
  ]>;

  /**
   * Azure Key Vault secret name.
   */
  keyVaultSecretName?: string;

  /**
   * Status of the Key Vault secret.
   */
  @visibility(Lifecycle.Read)
  keyVaultSecretStatus?: KeyVaultSecretStatus;

  /**
   * Resource ID of the associated App Service plan.
   */
  serverFarmId?: Azure.Core.armResourceIdentifier<[
    {
      type: "Microsoft.Web/serverfarms";
    }
  ]>;

  /**
   * CNAME of the certificate to be issued via free certificate
   */
  canonicalName?: string;

  /**
   * Method of domain validation for free cert
   */
  domainValidationMethod?: string;
}

/**
 * ARM resource for a certificate.
 */
#suppress "@azure-tools/typespec-azure-resource-manager/patch-envelope" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
#suppress "@azure-tools/typespec-azure-core/composition-over-inheritance" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
model CertificatePatchResource extends ProxyOnlyResource {
  /**
   * CertificatePatchResource resource specific properties
   */
  #suppress "@azure-tools/typespec-azure-core/no-private-usage" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
  properties?: CertificatePatchResourceProperties;
}

/**
 * CertificatePatchResource resource specific properties
 */
model CertificatePatchResourceProperties {
  /**
   * Certificate password.
   */
  @visibility(Lifecycle.Read)
  password?: string;

  /**
   * Friendly name of the certificate.
   */
  @visibility(Lifecycle.Read)
  friendlyName?: string;

  /**
   * Subject name of the certificate.
   */
  @visibility(Lifecycle.Read)
  subjectName?: string;

  /**
   * Host names the certificate applies to.
   */
  hostNames?: string[];

  /**
   * Pfx blob.
   */
  pfxBlob?: bytes;

  /**
   * App name.
   */
  @visibility(Lifecycle.Read)
  siteName?: string;

  /**
   * Self link.
   */
  @visibility(Lifecycle.Read)
  selfLink?: string;

  /**
   * Certificate issuer.
   */
  @visibility(Lifecycle.Read)
  issuer?: string;

  /**
   * Certificate issue Date.
   */
  @visibility(Lifecycle.Read)
  // FIXME: (utcDateTime) Please double check that this is the correct type for your scenario.
  issueDate?: utcDateTime;

  /**
   * Certificate expiration date.
   */
  @visibility(Lifecycle.Read)
  // FIXME: (utcDateTime) Please double check that this is the correct type for your scenario.
  expirationDate?: utcDateTime;

  /**
   * Certificate thumbprint.
   */
  @visibility(Lifecycle.Read)
  thumbprint?: string;

  /**
   * Is the certificate valid?.
   */
  @visibility(Lifecycle.Read)
  valid?: boolean;

  /**
   * Raw bytes of .cer file
   */
  @visibility(Lifecycle.Read)
  cerBlob?: bytes;

  /**
   * Public key hash.
   */
  @visibility(Lifecycle.Read)
  publicKeyHash?: string;

  /**
   * Specification for the App Service Environment to use for the certificate.
   */
  @visibility(Lifecycle.Read)
  hostingEnvironmentProfile?: HostingEnvironmentProfile;

  /**
   * Key Vault Csm resource Id.
   */
  keyVaultId?: string;

  /**
   * Key Vault secret name.
   */
  keyVaultSecretName?: string;

  /**
   * Status of the Key Vault secret.
   */
  @visibility(Lifecycle.Read)
  keyVaultSecretStatus?: KeyVaultSecretStatus;

  /**
   * Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
   */
  serverFarmId?: string;

  /**
   * CNAME of the certificate to be issued via free certificate
   */
  canonicalName?: string;

  /**
   * Method of domain validation for free cert
   */
  domainValidationMethod?: string;
}

/**
 * Status of the Key Vault secret.
 */
#suppress "@azure-tools/typespec-azure-core/no-enum" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
enum KeyVaultSecretStatus {
  Initialized,
  WaitingOnCertificateOrder,
  Succeeded,
  CertificateOrderFailed,
  OperationNotPermittedOnKeyVault,
  AzureServiceUnauthorizedToAccessKeyVault,
  KeyVaultDoesNotExist,
  KeyVaultSecretDoesNotExist,
  UnknownError,
  ExternalPrivateKey,
  Unknown,
}

/**
 * App Service error response.
 */
@error
model DefaultErrorResponse {
  /**
   * Error model.
   */
  @visibility(Lifecycle.Read)
  error?: DefaultErrorResponseError;
}

/**
 * Error model.
 */
@error
model DefaultErrorResponseError {
  /**
   * Standardized string to programmatically identify the error.
   */
  @visibility(Lifecycle.Read)
  code?: string;

  /**
   * Detailed error description and debugging information.
   */
  @visibility(Lifecycle.Read)
  message?: string;

  /**
   * Detailed error description and debugging information.
   */
  @visibility(Lifecycle.Read)
  target?: string;

  #suppress "@azure-tools/typespec-azure-core/documentation-required" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
  @identifiers(#[])
  details?: DefaultErrorResponseErrorDetailsItem[];

  /**
   * More information to debug error.
   */
  @visibility(Lifecycle.Read)
  innererror?: string;
}

/**
 * Detailed errors.
 */
model DefaultErrorResponseErrorDetailsItem {
  /**
   * Standardized string to programmatically identify the error.
   */
  @visibility(Lifecycle.Read)
  code?: string;

  /**
   * Detailed error description and debugging information.
   */
  @visibility(Lifecycle.Read)
  message?: string;

  /**
   * Detailed error description and debugging information.
   */
  @visibility(Lifecycle.Read)
  target?: string;
}

/**
 * Azure proxy only resource. This resource is not tracked by Azure Resource Manager.
 */
model ProxyOnlyResource {
  /**
   * Resource Id.
   */
  @visibility(Lifecycle.Read)
  id?: string;

  /**
   * Resource Name.
   */
  @visibility(Lifecycle.Read)
  name?: string;

  /**
   * Kind of resource.
   */
  kind?: string;

  /**
   * Resource type.
   */
  @visibility(Lifecycle.Read)
  type?: string;
}

/**
 * Specification for an App Service Environment to use for this resource.
 */
model HostingEnvironmentProfile {
  /**
   * Resource ID of the App Service Environment.
   */
  id?: string;

  /**
   * Name of the App Service Environment.
   */
  @visibility(Lifecycle.Read)
  name?: string;

  /**
   * Resource type of the App Service Environment.
   */
  @visibility(Lifecycle.Read)
  type?: string;
}
