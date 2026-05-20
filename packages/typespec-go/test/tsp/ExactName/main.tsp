import "@azure-tools/typespec-client-generator-core";
import "@typespec/http";

using Azure.ClientGenerator.Core;
using TypeSpec.Http;

@service(#{
  title: "Tests for the exact() function with @clientName",
})
@server(
    "{endpoint}",
    "Contains supplemental path with no params",
    {
        endpoint: url,
    }
)
namespace ExactName;

model Widget {
  // model property with exact() - field should remain "URL_Path" verbatim (no underscore stripping/recase)
  @clientName(exact("URL_Path"))
  rawField: string;

  // normal property — gets canonicalized
  normalField: int32;
}

// enum with exact() - constant type name should remain "My_HTTP_Status" verbatim (no underscore stripping)
@clientName(exact("My_HTTP_Status"))
enum MyEnum {
  one: "one",
  two: "two",
}

// model with exact() - struct name should remain "My_Exact_Model" verbatim (no underscore stripping)
@clientName(exact("My_Exact_Model"))
model RenamedModel {
  value: string;
  kind: MyEnum;
}

@get
@route("/get-widget")
op getWidget(
  // path parameter with exact()
  @path
  @clientName(exact("Widget_ID"))
  widgetId: string,
): {
  @header
  @clientName(exact("X_Request_ID"))
  xRequestId: string;

  @body
  body: Widget;
};

@get
@route("/get-renamed")
op getRenamed(): RenamedModel;
