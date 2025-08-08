import "@azure-tools/typespec-azure-core";
import "@azure-tools/typespec-client-generator-core";
import "@typespec/http";

using Azure.Core;
using Azure.ClientGenerator.Core;
using TypeSpec.Http;

@service(#{
  title: "DeserializeEmptyString Test"
})
namespace DeserializeEmptyString;

model ResponseModel {
  @deserializeEmptyStringAsNull
  sampleUrl: url;
  
  @doc("Regular string property for comparison")
  normalProperty: string;
}

@route("/test")
@get
op testOperation(): ResponseModel;