import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "@typespec/openapi";
import "@typespec/rest";
import "./models.tsp";

using TypeSpec.Rest;
using Azure.ResourceManager;
using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace Microsoft.Web;
/**
 * SSL certificate for an app.
 */
model Certificate is Azure.ResourceManager.TrackedResource<CertificateProperties> {
  ...ResourceNameParameter<Resource = Certificate, KeyName = "name", SegmentName = "certificates", NamePattern = "">;

  /** Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind. */
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-resource-invalid-envelope-property" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
  kind?: string;
}

@armResourceOperations
interface CertificateOps
  extends Azure.ResourceManager.Legacy.LegacyOperations<
      {
        ...ApiVersionParameter,
        ...SubscriptionIdParameter,
        ...ResourceGroupParameter,
        ...Azure.ResourceManager.Legacy.Provider,
      },
      {
        /** Name of the certificate. */
        @path
        @segment("certificates")
        @key
        name: string,
      }
    > {}

@armResourceOperations
interface Certificates {
  /**
   * Description for Get a certificate.
   */
  @summary("Get a certificate.")
  get is CertificateOps.Read<Certificate, OverrideErrorType = DefaultErrorResponse>;

  /**
   * Description for Create or update a certificate.
   */
  #suppress "@azure-tools/typespec-azure-resource-manager/arm-put-operation-response-codes" "FIXME: Update justification, follow aka.ms/tsp/conversion-fix for details"
  @summary("Create or update a certificate.")
  createOrUpdate is CertificateOps.CreateOrUpdateSync<Certificate, Response = ArmResourceUpdatedResponse<Certificate>, OverrideErrorType = DefaultErrorResponse>;

  /**
   * Description for Create or update a certificate.
   */
  @summary("Create or update a certificate.")
  @patch(#{ implicitOptionality: false })
  update is CertificateOps.CustomPatchSync<Certificate, PatchModel = CertificatePatchResource, OverrideErrorType = DefaultErrorResponse>;

  /**
   * Description for Delete a certificate.
   */
  @summary("Delete a certificate.")
  delete is CertificateOps.DeleteSync<Certificate, OverrideErrorType = DefaultErrorResponse>;

  /**
   * Description for Get all certificates in a resource group.
   */
  @summary("Get all certificates in a resource group.")
  listByResourceGroup is CertificateOps.List<Certificate, Response = ArmResponse<CertificateCollection>, OverrideErrorType = DefaultErrorResponse>;

  /**
   * Description for Get all certificates for a subscription.
   */
  @summary("Get all certificates for a subscription.")
  list is Azure.ResourceManager.Extension.ListByTarget<
    Extension.Subscription,
    Certificate,
    Parameters = {
      /**
       * Return only information specified in the filter (using OData syntax). For example: $filter=KeyVaultId eq 'KeyVaultId'
       */
      @query("$filter")
      $filter?: string;
    },
    Response = ArmResponse<CertificateCollection>,
    Error = DefaultErrorResponse
  >;
}
