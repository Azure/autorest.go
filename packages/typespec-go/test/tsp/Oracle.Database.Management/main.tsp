import "@azure-tools/typespec-autorest";

import "@typespec/rest";
import "@typespec/versioning";
import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";
import "./operations/oraclesubscription/oracleSubscriptionOperations.tsp";

import "./versions.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.Versioning;
using Azure.ResourceManager;
using OpenAPI;
using Autorest;

@armProviderNamespace
@service(#{ title: "Oracle Database Resource Manager" })
@armCommonTypesVersion(Azure.ResourceManager.CommonTypes.Versions.v5)
@versioned(Oracle.Database.Versions)
namespace Oracle.Database;

interface Operations extends Azure.ResourceManager.Operations {}

#suppress "@azure-tools/typespec-azure-core/no-openapi" "TODO FIX - DO NOT COPY - BAD PATTERN"
@@Autorest.example(Operations.list,
  "./examples/operations_list.json",
  "List Operations"
);
