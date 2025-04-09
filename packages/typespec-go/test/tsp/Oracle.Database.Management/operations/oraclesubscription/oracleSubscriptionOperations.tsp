import "../commons.tsp";
import "../../models/oraclesubscription/oracleSubscription.tsp";
using Azure.ResourceManager;
using Azure.ResourceManager.Foundations;
using Autorest;
using TypeSpec.Versioning;

namespace Oracle.Database;

@armResourceOperations(OracleSubscription<Plan>)
interface OracleSubscriptions
  extends OracleSubscriptionLevelResource<
      OracleSubscription<Plan>,
      OracleSubscriptionProperties
    > {
  #suppress "@azure-tools/typespec-azure-core/no-openapi" "TODO FIX - DO NOT COPY - BAD PATTERN"
  @doc("List Cloud Account Details")
  @Autorest.example(
    "./examples/oracleSubscriptions_listCloudAccountDetails.json",
    "List Cloud Account details for the Oracle Subscription"
  )
  listCloudAccountDetails is ListActions.ArmResourceGetActionSync<
    OracleSubscription<Plan>,
    CloudAccountDetails
  >;

  #suppress "@azure-tools/typespec-azure-core/no-openapi" "TODO FIX - DO NOT COPY - BAD PATTERN"
  @doc("List Saas Subscription Details")
  @Autorest.example(
    "./examples/oracleSubscriptions_listSaasSubscriptionDetails.json",
    "List Saas Subscription details for the Oracle Subscription"
  )
  listSaasSubscriptionDetails is ListActions.ArmResourceGetActionAsync<
    OracleSubscription<Plan>,
    SaasSubscriptionDetails
  >;
}