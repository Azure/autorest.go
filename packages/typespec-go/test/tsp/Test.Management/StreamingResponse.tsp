import "@typespec/http";
import "@typespec/rest";
import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-azure-resource-manager";

using TypeSpec.Http;
using TypeSpec.Rest;
using Azure.ResourceManager;

namespace Microsoft.Test;

@doc("Configuration resource.")
model Configuration is TrackedResource<ConfigurationProperties> {
  @doc("The configuration name.")
  @pattern("^(?=.{3,24}$)[^0-9][A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$")
  @path
  @key("configurationName")
  @segment("configurations")
  name: string;
}

@doc("Properties of the configuration resource.")
model ConfigurationProperties {
  @doc("Description of the configuration.")
  description?: string;

  @doc("Status of the current operation on configuration.")
  @visibility(Lifecycle.Read)
  provisioningState?: ProvisioningState;
}

@armResourceOperations
interface Configurations {
  @doc("Get the content of the configuration.")
  getContent is ArmResourceRead<Configuration>;

  @doc("Get the content of the configuration identified by configuration name.")
  @action("content")
  getStreamingContent(...ResourceInstanceParameters<Configuration>): {
    @header("Content-Type")
    contentType: "text/powershell";

    @bodyRoot
    @encode("bytes")
    body: bytes;
  };
}
