import "./main.tsp";
import "@azure-tools/typespec-client-generator-core";

using Azure.ClientGenerator.Core;
using TypeSpec.Versioning;
using TypeSpec.Http;

@versioned(KeyVault.Versions)
namespace Customizations;

@@clientName(KeyVault, "Client", "go");

using KeyVault;

/**
 * The full key identifier, attributes, and tags are provided in the response.
 * This operation requires the keys/list permission.
 */
#suppress "@azure-tools/typespec-azure-core/use-standard-operations" "Foundations.Operation is necessary for Key Vault"
#suppress "@azure-tools/typespec-azure-core/use-standard-names" "Operation name is already established"
@summary("Retrieves a list of individual key versions with the same key name.")
@route("/keys/{key-name}/versions")
@get
op listKeyPropertiesVersions is Azure.Core.Foundations.Operation<
  {
    /**
     * The name of the key.
     */
    @path("key-name")
    keyName: string;
  },
  KeyListResult,
  {},
  KeyVaultError
>;

/**s
 * Retrieves a list of the keys in the Key Vault as JSON Web Key structures that
 * contain the public part of a stored key. The LIST operation is applicable to
 * all key types, however only the base key identifier, attributes, and tags are
 * provided in the response. Individual versions of a key are not listed in the
 * response. This operation requires the keys/list permission.
 */
#suppress "@azure-tools/typespec-azure-core/use-standard-operations" "Foundations.Operation is necessary for Key Vault"
#suppress "@azure-tools/typespec-azure-core/use-standard-names" "Operation name is already established"
@summary("List keys in the specified vault.")
@route("/keys")
@get
op listKeyProperties is Azure.Core.Foundations.Operation<
  {},
  KeyListResult,
  {},
  KeyVaultError
>;

/**
 * Retrieves a list of the keys in the Key Vault as JSON Web Key structures that
 * contain the public part of a deleted key. This operation includes
 * deletion-specific information. The Get Deleted Keys operation is applicable for
 * vaults enabled for soft-delete. While the operation can be invoked on any
 * vault, it will return an error if invoked on a non soft-delete enabled vault.
 * This operation requires the keys/list permission.
 */
#suppress "@azure-tools/typespec-azure-core/use-standard-operations" "Foundations.Operation is necessary for Key Vault"
#suppress "@azure-tools/typespec-azure-core/use-standard-names" "Operation name is already established"
@summary("Lists the deleted keys in the specified vault.")
@route("/deletedkeys")
@get
op listDeletedKeyProperties is Azure.Core.Foundations.Operation<
  {},
  DeletedKeyListResult,
  {},
  KeyVaultError
>;

@@override(getKeyVersions, listKeyPropertiesVersions, "go");
@@override(getKeys, listKeyProperties, "go");
@@override(getDeletedKeys, listDeletedKeyProperties, "go");

@@clientName(KeyCreateParameters, "CreateKeyParameters", "go");
@@clientName(KeyExportParameters, "ExportKeyParameters", "go");
@@clientName(KeyImportParameters, "ImportKeyParameters", "go");
@@clientName(KeyReleaseParameters, "ReleaseParameters", "go");
@@clientName(KeyRestoreParameters, "RestoreKeyParameters", "go");
@@clientName(KeySignParameters, "SignParameters", "go"); 
@@clientName(KeyUpdateParameters, "UpdateKeyParameters", "go");  
@@clientName(KeyVerifyParameters, "VerifyParameters", "go"); 
@@clientName(GetRandomBytesRequest, "GetRandomBytesParameters", "go"); 
@@clientName(getDeletedKeys, "ListDeletedKeyProperties", "go"); 
@@clientName(getKeys, "ListKeyProperties", "go"); 
@@clientName(getKeyVersions, "ListKeyPropertiesVersions", "go"); 
@@clientName(DeletedKeyBundle, "DeletedKey", "go"); 
@@clientName(KeyProperties, "KeyP", "go"); 
@@clientName(KeyItem, "KeyProperties", "go"); 
@@clientName(DeletedKeyItem, "DeletedKeyProperties", "go"); 
@@clientName(DeletedKeyListResult, "DeletedKeyPropertiesListResult", "go"); 
@@clientName(KeyListResult, "KeyPropertiesListResult", "go");
@@clientName(LifetimeActions, "LifetimeAction", "go");  
@@clientName(LifetimeActionsType, "LifetimeActionType", "go");  
@@clientName(LifetimeActionsTrigger, "LifetimeActionTrigger", "go");  
@@clientName(KeyAttributes.hsmPlatform, "HSMPlatform", "go");  
@@clientName(KeyRestoreParameters.keyBundleBackup, "KeyBackup", "go");
@@clientName(KeyOperationsParameters, "KeyOperationParameters", "go");  

@@clientName(KeyImportParameters.hsm, "HSM", "go");  
@@clientName(KeyReleaseParameters.enc, "algorithm", "go");   
@@clientName(KeyOperationsParameters.aad, "AdditionalAuthenticatedData", "go");  
@@clientName(KeyOperationsParameters.tag, "AuthenticationTag", "go");  

@@clientName(JsonWebKeyType, "KeyType", "go");  
@@clientName(JsonWebKeyOperation, "KeyOperation", "go");  
@@clientName(JsonWebKeyEncryptionAlgorithm, "EncryptionAlgorithm", "go");  
@@clientName(JsonWebKeyCurveName, "CurveName", "go");  
@@clientName(JsonWebKeySignatureAlgorithm, "SignatureAlgorithm", "go"); 

@@clientName(JsonWebKey.kid, "KID", "go");  
@@clientName(KeyItem.kid, "KID", "go"); 
@@clientName(KeyOperationResult.kid, "KID", "go"); 
@@clientName(KeyOperationsParameters.iv, "IV", "go");
@@clientName(KeyOperationResult.iv, "IV", "go"); 

@@clientName(JsonWebKey.dp, "DP", "go"); 
@@clientName(JsonWebKey.dq, "DQ", "go"); 
@@clientName(JsonWebKey.qi, "QI", "go"); 
