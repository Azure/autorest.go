import "@azure-tools/typespec-client-generator-core";
import "@typespec/http";

using Azure.ClientGenerator.Core;
using TypeSpec.Http;

@service(#{
  title: "Various tests for clientOption decorator",
})
@server(
    "{endpoint}",
    "Contains supplemental path with no params",
    {
        endpoint: url,
    }
)
namespace ClientOption;

model Person {
  name: string;
  age: int32;
}

@get
@route("/rename-monomorphic-field")
op customFieldName(): Array<Person>;

#suppress "@azure-tools/typespec-client-generator-core/client-option" "response has custom field name Users"
@@clientOption(customFieldName, "monomorphicResponseFieldName", "Users", "go");

#suppress "@azure-tools/typespec-client-generator-core/client-option" "omit serde methods for this type"
@@clientOption(Person, "omitSerdeMethods", "all", "go");
