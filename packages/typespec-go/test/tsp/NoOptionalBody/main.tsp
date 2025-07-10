import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using TypeSpec.Http;

@service(#{
  title: "No optional body params",
})
namespace NoOptionalBody;

model Widget {
  weight: int32;
}

@put
op put(@body body?: Widget): void;
