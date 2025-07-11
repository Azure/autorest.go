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

/** body should not be optional */
@patch
op patch(@body body?: Widget): void;

/** body should not be optional */
@put
op put(@body body?: Widget): void;

/** body should be optional */
@post
op post(@body body?: Widget): void;
