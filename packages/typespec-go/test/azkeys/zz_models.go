// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package azkeys

import "time"

// BackupKeyResult - The backup key result, containing the backup blob.
type BackupKeyResult struct {
	// READ-ONLY; The backup blob containing the backed up key.
	Value []byte
}

// DeletedKeyBundle - A DeletedKeyBundle consisting of a WebKey plus its Attributes and deletion info
type DeletedKeyBundle struct {
	// The key management attributes.
	Attributes *KeyAttributes

	// The Json web key.
	Key *JSONWebKey

	// The url of the recovery object, used to identify and recover the deleted key.
	RecoveryID *string

	// The policy rules under which the key can be exported.
	ReleasePolicy *KeyReleasePolicy

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string

	// READ-ONLY; The time when the key was deleted, in UTC
	DeletedDate *time.Time

	// READ-ONLY; True if the key's lifetime is managed by key vault. If this is a key backing a
	// certificate, then managed will be true.
	Managed *bool

	// READ-ONLY; The time when the key is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time
}

// DeletedKeyItem - The deleted key item containing the deleted key metadata and information about
// deletion.
type DeletedKeyItem struct {
	// The key management attributes.
	Attributes *KeyAttributes

	// Key identifier.
	Kid *string

	// The url of the recovery object, used to identify and recover the deleted key.
	RecoveryID *string

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string

	// READ-ONLY; The time when the key was deleted, in UTC
	DeletedDate *time.Time

	// READ-ONLY; True if the key's lifetime is managed by key vault. If this is a key backing a
	// certificate, then managed will be true.
	Managed *bool

	// READ-ONLY; The time when the key is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time
}

// DeletedKeyListResult - A list of keys that have been deleted in this vault.
type DeletedKeyListResult struct {
	// READ-ONLY; The URL to get the next set of deleted keys.
	NextLink *string

	// READ-ONLY; A response message containing a list of deleted keys in the key vault along with a link to the next page of
	// deleted keys.
	Value []*DeletedKeyItem
}

// GetRandomBytesRequest - The get random bytes request object.
type GetRandomBytesRequest struct {
	// REQUIRED; The requested number of random bytes.
	Count *int32
}

// JSONWebKey - As of http://tools.ietf.org/html/draft-ietf-jose-json-web-key-18
type JSONWebKey struct {
	// Elliptic curve name. For valid values, see JsonWebKeyCurveName.
	Crv *JSONWebKeyCurveName

	// RSA private exponent, or the D component of an EC private key.
	D []byte

	// RSA private key parameter.
	Dp []byte

	// RSA private key parameter.
	Dq []byte

	// RSA public exponent.
	E []byte

	// Symmetric key.
	K []byte

	// Json web key operations. For more information on possible key operations, see
	// JsonWebKeyOperation.
	KeyOps []*string

	// Key identifier.
	Kid *string

	// JsonWebKey Key Type (kty), as defined in
	// https://tools.ietf.org/html/draft-ietf-jose-json-web-algorithms-40.
	Kty *JSONWebKeyType

	// RSA modulus.
	N []byte

	// RSA secret prime.
	P []byte

	// RSA secret prime, with p < q.
	Q []byte

	// RSA private key parameter.
	Qi []byte

	// Protected Key, used with 'Bring Your Own Key'.
	T []byte

	// X component of an EC public key.
	X []byte

	// Y component of an EC public key.
	Y []byte
}

// KeyAttributes - The attributes of a key managed by the key vault service.
type KeyAttributes struct {
	// Determines whether the object is enabled.
	Enabled *bool

	// Expiry date in UTC.
	Expires *time.Time

	// Indicates if the private key can be exported. Release policy must be provided
	// when creating the first version of an exportable key.
	Exportable *bool

	// Not before date in UTC.
	NotBefore *time.Time

	// READ-ONLY; Creation time in UTC.
	Created *time.Time

	// READ-ONLY; The underlying HSM Platform.
	HsmPlatform *string

	// READ-ONLY; softDelete data retention days. Value should be >=7 and <=90 when softDelete
	// enabled, otherwise 0.
	RecoverableDays *int32

	// READ-ONLY; Reflects the deletion recovery level currently in effect for keys in the
	// current vault. If it contains 'Purgeable' the key can be permanently deleted by
	// a privileged user; otherwise, only the system can purge the key, at the end of
	// the retention interval.
	RecoveryLevel *string

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time
}

// KeyBundle - A KeyBundle consisting of a WebKey plus its attributes.
type KeyBundle struct {
	// The key management attributes.
	Attributes *KeyAttributes

	// The Json web key.
	Key *JSONWebKey

	// The policy rules under which the key can be exported.
	ReleasePolicy *KeyReleasePolicy

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string

	// READ-ONLY; True if the key's lifetime is managed by key vault. If this is a key backing a
	// certificate, then managed will be true.
	Managed *bool
}

// KeyCreateParameters - The key create parameters.
type KeyCreateParameters struct {
	// REQUIRED; The type of key to create. For valid values, see JsonWebKeyType.
	Kty *JSONWebKeyType

	// Elliptic curve name. For valid values, see JsonWebKeyCurveName.
	Curve *JSONWebKeyCurveName

	// The attributes of a key managed by the key vault service.
	KeyAttributes *KeyAttributes

	// Json web key operations. For more information on possible key operations, see
	// JsonWebKeyOperation.
	KeyOps []*JSONWebKeyOperation

	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize *int32

	// The public exponent for a RSA key.
	PublicExponent *int32

	// The policy rules under which the key can be exported.
	ReleasePolicy *KeyReleasePolicy

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// KeyImportParameters - The key import parameters.
type KeyImportParameters struct {
	// REQUIRED; The Json web key
	Key *JSONWebKey

	// Whether to import as a hardware key (HSM) or software key.
	Hsm *bool

	// The key management attributes.
	KeyAttributes *KeyAttributes

	// The policy rules under which the key can be exported.
	ReleasePolicy *KeyReleasePolicy

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// KeyItem - The key item containing key metadata.
type KeyItem struct {
	// The key management attributes.
	Attributes *KeyAttributes

	// Key identifier.
	Kid *string

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string

	// READ-ONLY; True if the key's lifetime is managed by key vault. If this is a key backing a
	// certificate, then managed will be true.
	Managed *bool
}

// KeyListResult - The key list result.
type KeyListResult struct {
	// READ-ONLY; The URL to get the next set of keys.
	NextLink *string

	// READ-ONLY; A response message containing a list of keys in the key vault along with a link to the next page of keys.
	Value []*KeyItem
}

// KeyOperationResult - The key operation result.
type KeyOperationResult struct {
	// READ-ONLY; Additional data to authenticate but not encrypt/decrypt when using
	// authenticated crypto algorithms.
	AdditionalAuthenticatedData []byte

	// READ-ONLY; The tag to authenticate when performing decryption with an authenticated
	// algorithm.
	AuthenticationTag []byte

	// READ-ONLY; Cryptographically random, non-repeating initialization vector for symmetric
	// algorithms.
	Iv []byte

	// READ-ONLY; Key identifier
	Kid *string

	// READ-ONLY; The result of the operation.
	Result []byte
}

// KeyOperationsParameters - The key operations parameters.
type KeyOperationsParameters struct {
	// REQUIRED; algorithm identifier
	Algorithm *JSONWebKeyEncryptionAlgorithm

	// REQUIRED; The value to operate on.
	Value []byte

	// Additional data to authenticate but not encrypt/decrypt when using
	// authenticated crypto algorithms.
	AAD []byte

	// Cryptographically random, non-repeating initialization vector for symmetric
	// algorithms.
	Iv []byte

	// The tag to authenticate when performing decryption with an authenticated
	// algorithm.
	Tag []byte
}

// KeyReleaseParameters - The release key parameters.
type KeyReleaseParameters struct {
	// REQUIRED; The attestation assertion for the target of the key release.
	TargetAttestationToken *string

	// The encryption algorithm to use to protected the exported key material
	Enc *KeyEncryptionAlgorithm

	// A client provided nonce for freshness.
	Nonce *string
}

// KeyReleasePolicy - The policy rules under which the key can be exported.
type KeyReleasePolicy struct {
	// Content type and version of key release policy
	ContentType *string

	// Blob encoding the policy rules under which the key can be released. Blob must
	// be base64 URL encoded.
	EncodedPolicy []byte

	// Defines the mutability state of the policy. Once marked immutable, this flag
	// cannot be reset and the policy cannot be changed under any circumstances.
	Immutable *bool
}

// KeyReleaseResult - The release result, containing the released key.
type KeyReleaseResult struct {
	// READ-ONLY; A signed object containing the released key.
	Value *string
}

// KeyRestoreParameters - The key restore parameters.
type KeyRestoreParameters struct {
	// REQUIRED; The backup blob associated with a key bundle.
	KeyBundleBackup []byte
}

// KeyRotationPolicy - Management policy for a key.
type KeyRotationPolicy struct {
	// The key rotation policy attributes.
	Attributes *KeyRotationPolicyAttributes

	// Actions that will be performed by Key Vault over the lifetime of a key. For
	// preview, lifetimeActions can only have two items at maximum: one for rotate,
	// one for notify. Notification time would be default to 30 days before expiry and
	// it is not configurable.
	LifetimeActions []*LifetimeActions

	// READ-ONLY; The key policy id.
	ID *string
}

// KeyRotationPolicyAttributes - The key rotation policy attributes.
type KeyRotationPolicyAttributes struct {
	// The expiryTime will be applied on the new key version. It should be at least 28
	// days. It will be in ISO 8601 Format. Examples: 90 days: P90D, 3 months: P3M, 48
	// hours: PT48H, 1 year and 10 days: P1Y10D
	ExpiryTime *string

	// READ-ONLY; The key rotation policy created time in UTC.
	Created *time.Time

	// READ-ONLY; The key rotation policy's last updated time in UTC.
	Updated *time.Time
}

// KeySignParameters - The key operations parameters.
type KeySignParameters struct {
	// REQUIRED; The signing/verification algorithm identifier. For more information on possible
	// algorithm types, see JsonWebKeySignatureAlgorithm.
	Algorithm *JSONWebKeySignatureAlgorithm

	// REQUIRED; The value to operate on.
	Value []byte
}

// KeyUpdateParameters - The key update parameters.
type KeyUpdateParameters struct {
	// The attributes of a key managed by the key vault service.
	KeyAttributes *KeyAttributes

	// Json web key operations. For more information on possible key operations, see
	// JsonWebKeyOperation.
	KeyOps []*JSONWebKeyOperation

	// The policy rules under which the key can be exported.
	ReleasePolicy *KeyReleasePolicy

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string
}

// KeyVerifyParameters - The key verify parameters.
type KeyVerifyParameters struct {
	// REQUIRED; The signing/verification algorithm. For more information on possible algorithm
	// types, see JsonWebKeySignatureAlgorithm.
	Algorithm *JSONWebKeySignatureAlgorithm

	// REQUIRED; The digest used for signing.
	Digest []byte

	// REQUIRED; The signature to be verified.
	Signature []byte
}

// KeyVerifyResult - The key verify result.
type KeyVerifyResult struct {
	// READ-ONLY; True if the signature is verified, otherwise false.
	Value *bool
}

// LifetimeActions - Action and its trigger that will be performed by Key Vault over the lifetime of
// a key.
type LifetimeActions struct {
	// The action that will be executed.
	Action *LifetimeActionsType

	// The condition that will execute the action.
	Trigger *LifetimeActionsTrigger
}

// LifetimeActionsTrigger - A condition to be satisfied for an action to be executed.
type LifetimeActionsTrigger struct {
	// Time after creation to attempt to rotate. It only applies to rotate. It will be
	// in ISO 8601 duration format. Example: 90 days : "P90D"
	TimeAfterCreate *string

	// Time before expiry to attempt to rotate or notify. It will be in ISO 8601
	// duration format. Example: 90 days : "P90D"
	TimeBeforeExpiry *string
}

// LifetimeActionsType - The action that will be executed.
type LifetimeActionsType struct {
	// The type of the action. The value should be compared case-insensitively.
	Type *KeyRotationPolicyAction
}

// RandomBytes - The get random bytes response object containing the bytes.
type RandomBytes struct {
	// REQUIRED; The bytes encoded as a base64url string.
	Value []byte
}
