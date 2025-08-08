import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using Azure.ClientGenerator.Core;
using Azure.Core;
using TypeSpec.Http;

@service(#{
  title: "Internal Pager",
})
namespace Microsoft.Internal.Pager;

model WidgetList {
  @pageItems
  values: Widget[];
}

model Widget {
  weight: int32;
  color: WidgetColor;
}

union WidgetColor {
  string,
  Blue: "Blue",
  Green: "Green",
  Red: "Red",
}

@route("/widgets")
@tag("Widgets")
interface Widgets {
  @access(Access.internal)
  @get
  @list
  list(): WidgetList;
}
