import "@azure-tools/typespec-client-generator-core";
import "@azure-tools/typespec-azure-core";
import "@typespec/http";

using TypeSpec.Http;

@service(#{
  title: "Raw JSON tests",
})
namespace RawJson;

model ContainsRawJson {
    anything1: unknown;
    anything2: unknown;
}

interface InputOnly {
  @put
  @route("/input")
  op put(@body body: unknown): void;
}

interface OutputOnly {
  @get
  @route("/output")
  op get(): unknown;
}

interface ModelCase {
  @put
  @route("/unknown2")
  op forModel(@body body: ContainsRawJson): ContainsRawJson;
}
