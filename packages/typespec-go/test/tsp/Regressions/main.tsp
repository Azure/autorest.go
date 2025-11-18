import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using TypeSpec.Http;

@service(#{
  title: "Various tests for emitter regressions",
})
@server(
    "{endpoint}/supplemental-path",
    "Contains supplemental path with no params",
    {
        endpoint: url,
    }
)
namespace Emitter.Regressions;

model SpreadParams {
  name: string;
  color?: string;
  inner?: InnerSpreadParam;
}

model InnerSpreadParam {
  name: string;
  code: int32;
}

@route("/spread-with-model")
op spreadWithModel(...SpreadParams): void;
